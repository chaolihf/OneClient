/*
*
@description  实现IScriptPlugin接口的http插件
*/
package icmpPlugin

import (
	"bytes"
	"context"
	"fmt"
	"math/rand"
	"net"
	"os"
	"runtime"
	"sync"
	"time"

	"com.chinatelecom.oneops.client/pkg/utils"
	"go.uber.org/zap"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

var (
	icmpID            int
	icmpSequence      uint16
	icmpSequenceMutex sync.Mutex
	DefaultICMPTTL    = 64
)

func init() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// PID is typically 1 when running in a container; in that case, set
	// the ICMP echo ID to a random value to avoid potential clashes with
	// other blackbox_exporter instances. See #411.
	if pid := os.Getpid(); pid == 1 {
		icmpID = r.Intn(1 << 16)
	} else {
		icmpID = pid & 0xffff
	}

	// Start the ICMP echo sequence at a random offset to prevent them from
	// being in sync when several blackbox_exporter instances are restarted
	// at the same time. See #411.
	icmpSequence = uint16(r.Intn(1 << 16))
}

func getICMPSequence() uint16 {
	icmpSequenceMutex.Lock()
	defer icmpSequenceMutex.Unlock()
	icmpSequence++
	return icmpSequence
}

// ICMPScriptPlugin 实现了IScriptPlugin接口的ICMP插件
type ICMPScriptPlugin struct {
	logger          *zap.Logger
	DontFragment    bool
	sourceIpAddress string
	PayloadSize     int
	TTL             int
	Deadline        int
}

func NewICMPScriptPlugin(logger *zap.Logger) *ICMPScriptPlugin {
	return &ICMPScriptPlugin{
		logger:       logger,
		DontFragment: false,
		TTL:          DefaultICMPTTL,
		Deadline:     5,
	}
}

// @return string icmp插件
func (this ICMPScriptPlugin) GetCode() string {
	return "icmp"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (this ICMPScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "ping":
		{
			this.ping(params.(string))
		}
	}
	return nil
}

func (this *ICMPScriptPlugin) SetLogger(logger *zap.Logger) {
	this.logger = logger
}

func (this *ICMPScriptPlugin) ping(target string) bool {
	var (
		requestType     icmp.Type
		replyType       icmp.Type
		icmpConn        *icmp.PacketConn
		v4RawConn       *ipv4.RawConn
		hopLimitFlagSet bool = true
	)
	ctx, _ := context.WithDeadline(context.Background(),
		time.Now().Add(time.Duration(this.Deadline)*time.Second))

	logger := this.logger
	dstIPAddr, lookupTime, IPProtocol, err := utils.ChooseProtocol(nil, 4, true, target, logger)
	logger.Info(fmt.Sprint(lookupTime, IPProtocol))
	var srcIP net.IP
	if len(this.sourceIpAddress) > 0 {
		if srcIP = net.ParseIP(this.sourceIpAddress); srcIP == nil {
			logger.Info(fmt.Sprint("Error parsing source ip address", "srcIP", this.sourceIpAddress))
			return false
		}
		logger.Info(fmt.Sprint("Using source address", "srcIP", srcIP))
	}
	setupStart := time.Now()
	logger.Info(fmt.Sprint("Creating socket"))

	privileged := true
	// Unprivileged sockets are supported on Darwin and Linux only.
	tryUnprivileged := runtime.GOOS == "darwin" || runtime.GOOS == "linux"

	if dstIPAddr.IP.To4() == nil {
		requestType = ipv6.ICMPTypeEchoRequest
		replyType = ipv6.ICMPTypeEchoReply

		if srcIP == nil {
			srcIP = net.ParseIP("::")
		}

		if tryUnprivileged {
			// "udp" here means unprivileged -- not the protocol "udp".
			icmpConn, err = icmp.ListenPacket("udp6", srcIP.String())
			if err != nil {
				logger.Debug(fmt.Sprint("Unable to do unprivileged listen on socket, will attempt privileged", "err", err))
			} else {
				privileged = false
			}
		}

		if privileged {
			icmpConn, err = icmp.ListenPacket("ip6:ipv6-icmp", srcIP.String())
			if err != nil {
				logger.Error(fmt.Sprint("Error listening to socket", "err", err))
				return false
			}
		}
		defer icmpConn.Close()

		if err := icmpConn.IPv6PacketConn().SetControlMessage(ipv6.FlagHopLimit, true); err != nil {
			logger.Debug(fmt.Sprint("Failed to set Control Message for retrieving Hop Limit", "err", err))
			hopLimitFlagSet = false
		}
	} else {
		requestType = ipv4.ICMPTypeEcho
		replyType = ipv4.ICMPTypeEchoReply

		if srcIP == nil {
			srcIP = net.ParseIP("0.0.0.0")
		}

		if this.DontFragment {
			// If the user has set the don't fragment option we cannot use unprivileged
			// sockets as it is not possible to set IP header level options.
			netConn, err := net.ListenPacket("ip4:icmp", srcIP.String())
			if err != nil {
				logger.Error(fmt.Sprint("Error listening to socket", "err", err))
				return false
			}
			defer netConn.Close()

			v4RawConn, err = ipv4.NewRawConn(netConn)
			if err != nil {
				logger.Error(fmt.Sprint("Error creating raw connection", "err", err))
				return false
			}
			defer v4RawConn.Close()

			if err := v4RawConn.SetControlMessage(ipv4.FlagTTL, true); err != nil {
				logger.Debug(fmt.Sprint("Failed to set Control Message for retrieving TTL", "err", err))
				hopLimitFlagSet = false
			}
		} else {
			if tryUnprivileged {
				icmpConn, err = icmp.ListenPacket("udp4", srcIP.String())
				if err != nil {
					logger.Debug(fmt.Sprint("Unable to do unprivileged listen on socket, will attempt privileged", "err", err))
				} else {
					privileged = false
				}
			}

			if privileged {
				icmpConn, err = icmp.ListenPacket("ip4:icmp", srcIP.String())
				if err != nil {
					logger.Error(fmt.Sprint("Error listening to socket", "err", err))
					return false
				}
			}
			defer icmpConn.Close()

			if err := icmpConn.IPv4PacketConn().SetControlMessage(ipv4.FlagTTL, true); err != nil {
				logger.Debug(fmt.Sprint("Failed to set Control Message for retrieving TTL", "err", err))
				hopLimitFlagSet = false
			}
		}
	}

	var dst net.Addr = dstIPAddr
	if !privileged {
		dst = &net.UDPAddr{IP: dstIPAddr.IP, Zone: dstIPAddr.Zone}
	}

	var data []byte
	if this.PayloadSize != 0 {
		data = make([]byte, this.PayloadSize)
		copy(data, "Prometheus Blackbox Exporter")
	} else {
		data = []byte("Prometheus Blackbox Exporter")
	}

	body := &icmp.Echo{
		ID:   icmpID,
		Seq:  int(getICMPSequence()),
		Data: data,
	}
	logger.Info(fmt.Sprint("Creating ICMP packet", "seq", body.Seq, "id", body.ID))
	wm := icmp.Message{
		Type: requestType,
		Code: 0,
		Body: body,
	}

	wb, err := wm.Marshal(nil)
	if err != nil {
		logger.Error(fmt.Sprint("Error marshalling packet", "err", err))
		return false
	}

	setupTime := time.Since(setupStart).Seconds()
	logger.Info(fmt.Sprint("Sending packet", "setupTime", setupTime))
	logger.Info(fmt.Sprint("Writing out packet"))
	rttStart := time.Now()

	if icmpConn != nil {
		ttl := this.TTL
		if ttl > 0 {
			if c4 := icmpConn.IPv4PacketConn(); c4 != nil {
				logger.Debug(fmt.Sprint("Setting TTL (IPv4 unprivileged)", "ttl", ttl))
				c4.SetTTL(ttl)
			}
			if c6 := icmpConn.IPv6PacketConn(); c6 != nil {
				logger.Debug(fmt.Sprint("Setting TTL (IPv6 unprivileged)", "ttl", ttl))
				c6.SetHopLimit(ttl)
			}
		}
		_, err = icmpConn.WriteTo(wb, dst)
	} else {
		ttl := DefaultICMPTTL
		if this.TTL > 0 {
			logger.Debug(fmt.Sprint("Overriding TTL (raw IPv4)", "ttl", ttl))
			ttl = this.TTL
		}
		// Only for IPv4 raw. Needed for setting DontFragment flag.
		header := &ipv4.Header{
			Version:  ipv4.Version,
			Len:      ipv4.HeaderLen,
			Protocol: 1,
			TotalLen: ipv4.HeaderLen + len(wb),
			TTL:      ttl,
			Dst:      dstIPAddr.IP,
			Src:      srcIP,
		}

		header.Flags |= ipv4.DontFragment

		err = v4RawConn.WriteTo(header, wb, nil)
	}
	if err != nil {
		logger.Warn(fmt.Sprint("Error writing to socket", "err", err))
		return false
	}

	// Reply should be the same except for the message type and ID if
	// unprivileged sockets were used and the kernel used its own.
	wm.Type = replyType
	// Unprivileged cannot set IDs on Linux.
	idUnknown := !privileged && runtime.GOOS == "linux"
	if idUnknown {
		body.ID = 0
	}
	wb, err = wm.Marshal(nil)
	if err != nil {
		logger.Error(fmt.Sprint("Error marshalling packet", "err", err))
		return false
	}

	if idUnknown {
		// If the ID is unknown (due to unprivileged sockets) we also cannot know
		// the checksum in userspace.
		wb[2] = 0
		wb[3] = 0
	}

	rb := make([]byte, 65536)
	deadline, _ := ctx.Deadline()
	if icmpConn != nil {
		err = icmpConn.SetReadDeadline(deadline)
	} else {
		err = v4RawConn.SetReadDeadline(deadline)
	}
	if err != nil {
		logger.Error(fmt.Sprint("Error setting socket deadline", "err", err))
		return false
	}
	logger.Info(fmt.Sprint("Waiting for reply packets"))
	for {
		var n int
		var peer net.Addr
		var err error
		var hopLimit float64 = -1

		if dstIPAddr.IP.To4() == nil {
			var cm *ipv6.ControlMessage
			n, cm, peer, err = icmpConn.IPv6PacketConn().ReadFrom(rb)
			// HopLimit == 0 is valid for IPv6, although go initialize it as 0.
			if cm != nil && hopLimitFlagSet {
				hopLimit = float64(cm.HopLimit)
			} else {
				logger.Debug(fmt.Sprint("Cannot get Hop Limit from the received packet. 'probe_icmp_reply_hop_limit' will be missing."))
			}
		} else {
			var cm *ipv4.ControlMessage
			if icmpConn != nil {
				n, cm, peer, err = icmpConn.IPv4PacketConn().ReadFrom(rb)
			} else {
				var h *ipv4.Header
				var p []byte
				h, p, cm, err = v4RawConn.ReadFrom(rb)
				if err == nil {
					copy(rb, p)
					n = len(p)
					peer = &net.IPAddr{IP: h.Src}
				}
			}
			if cm != nil && hopLimitFlagSet {
				// Not really Hop Limit, but it is in practice.
				hopLimit = float64(cm.TTL)
			} else {
				logger.Debug(fmt.Sprint("Cannot get TTL from the received packet. 'probe_icmp_reply_hop_limit' will be missing."))
			}
		}
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				logger.Warn(fmt.Sprint("Timeout reading from socket", "err", err))
				return false
			}
			logger.Error(fmt.Sprint("Error reading from socket", "err", err))
			continue
		}
		if peer.String() != dst.String() {
			continue
		}
		if idUnknown {
			// Clear the ID from the packet, as the kernel will have replaced it (and
			// kept track of our packet for us, hence clearing is safe).
			rb[4] = 0
			rb[5] = 0
		}
		if idUnknown || replyType == ipv6.ICMPTypeEchoReply {
			// Clear checksum to make comparison succeed.
			rb[2] = 0
			rb[3] = 0
		}
		if bytes.Equal(rb[:n], wb) {
			rttTime := time.Since(rttStart).Seconds()
			logger.Info(fmt.Sprint("rttTime", rttTime))
			if hopLimit >= 0 {
				logger.Info(fmt.Sprint("hopLimit", hopLimit))
			}
			logger.Info(fmt.Sprint("Found matching reply packet"))
			return true
		}
	}
}
