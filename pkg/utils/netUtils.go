package utils

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.uber.org/zap"
)

// Returns the IP for the IPProtocol and lookup time.
/**
* @param ctx context.Context
* @param IPProtocol int 4或6
* @param fallbackIPProtocol bool
* @param target string
* @param logger *zap.Logger
* @return ip *net.IPAddr
* @return lookupTime float64
* @return probePrototocol int
* @return err error
 */
func ChooseProtocol(ctx context.Context, IPProtocol int, fallbackIPProtocol bool, target string, logger *zap.Logger) (ip *net.IPAddr, lookupTime float64, probePrototocol int, err error) {
	var fallbackProtocol int

	if IPProtocol == 6 {
		IPProtocol = 6
		fallbackProtocol = 4
	} else {
		IPProtocol = 4
		fallbackProtocol = 6
	}
	logger.Info(fmt.Sprint("Resolving target address target", target, " ip_protocol ", IPProtocol))
	resolveStart := time.Now()

	defer func() {
		lookupTime = time.Since(resolveStart).Seconds()
	}()

	resolver := &net.Resolver{}
	if !fallbackIPProtocol {
		ips, err := resolver.LookupIP(ctx, fmt.Sprintf("ip%d", IPProtocol), target)
		if err == nil {
			for _, ip := range ips {
				logger.Info(fmt.Sprint("Resolved target address", "target", target, "ip", ip.String()))
				return &net.IPAddr{IP: ip}, lookupTime, IPProtocol, nil
			}
		}
		logger.Info(fmt.Sprint("Resolution with IP protocol failed", "target", target, "ip_protocol", IPProtocol, "err", err))
		return nil, 0.0, IPProtocol, err
	}

	ips, err := resolver.LookupIPAddr(ctx, target)
	if err != nil {
		logger.Info(fmt.Sprint("Resolution with IP protocol failed", "target", target, "err", err))
		return nil, 0.0, IPProtocol, err
	}

	// Return the IP in the requested protocol.
	var fallback *net.IPAddr
	for _, ip := range ips {
		switch IPProtocol {
		case 4:
			if ip.IP.To4() != nil {
				logger.Info(fmt.Sprint("Resolved target address", "target", target, "ip", ip.String()))
				return &ip, lookupTime, IPProtocol, nil
			}

			// ip4 as fallback
			fallback = &ip

		case 6:
			if ip.IP.To4() == nil {
				logger.Info(fmt.Sprint("Resolved target address", "target", target, "ip", ip.String()))
				return &ip, lookupTime, IPProtocol, nil
			}

			// ip6 as fallback
			fallback = &ip
		}
	}

	// Unable to find ip and no fallback set.
	if fallback == nil || !fallbackIPProtocol {
		return nil, 0.0, IPProtocol, fmt.Errorf("unable to find ip; no fallback")
	}

	// Use fallback ip protocol.
	if fallbackProtocol == 4 {
		IPProtocol = 6
	} else {
		IPProtocol = 4
	}
	logger.Info(fmt.Sprint("Resolved target address", "target", target, "ip", IPProtocol))
	return fallback, lookupTime, IPProtocol, nil
}
