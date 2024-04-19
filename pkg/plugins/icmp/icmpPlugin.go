/*
*
@description  实现IScriptPlugin接口的http插件
*/
package icmpPlugin

import (
	"math/rand"
	"os"
	"sync"
	"time"

	"com.chinatelecom.oneops.client/pkg/utils"
	"go.uber.org/zap"
)

var (
	icmpID            int
	icmpSequence      uint16
	icmpSequenceMutex sync.Mutex
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
	logger *zap.Logger
}

func NewICMPScriptPlugin(logger *zap.Logger) *ICMPScriptPlugin {
	return &ICMPScriptPlugin{
		logger: logger,
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

func (this *ICMPScriptPlugin) ping(target string) {
	utils.ChooseProtocol(nil, 4, true, target, this.logger)
}
