/*
*
@description  实现IScriptPlugin接口的tcp插件
*/
package tcpPlugin

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"com.chinatelecom.oneops.client/pkg/utils"
	"go.uber.org/zap"
)

type TLSRoundTripperSettings struct {
	CA, CAFile     string
	Cert, CertFile string
	Key, KeyFile   string
}

type TLSVersion uint16
type Secret string

type TLSConfig struct {
	// Text of the CA cert to use for the targets.
	CA string `yaml:"ca,omitempty" json:"ca,omitempty"`
	// Text of the client cert file for the targets.
	Cert string `yaml:"cert,omitempty" json:"cert,omitempty"`
	// Text of the client key file for the targets.
	Key Secret `yaml:"key,omitempty" json:"key,omitempty"`
	// The CA cert to use for the targets.
	CAFile string `yaml:"ca_file,omitempty" json:"ca_file,omitempty"`
	// The client cert file for the targets.
	CertFile string `yaml:"cert_file,omitempty" json:"cert_file,omitempty"`
	// The client key file for the targets.
	KeyFile string `yaml:"key_file,omitempty" json:"key_file,omitempty"`
	// Used to verify the hostname for the targets.
	ServerName string `yaml:"server_name,omitempty" json:"server_name,omitempty"`
	// Disable target certificate validation.
	InsecureSkipVerify bool `yaml:"insecure_skip_verify" json:"insecure_skip_verify"`
	// Minimum TLS version.
	MinVersion TLSVersion `yaml:"min_version,omitempty" json:"min_version,omitempty"`
	// Maximum TLS version.
	MaxVersion TLSVersion `yaml:"max_version,omitempty" json:"max_version,omitempty"`
}

func (c *TLSConfig) usingClientCert() bool {
	return len(c.Cert) > 0 || len(c.CertFile) > 0
}

func (c *TLSConfig) usingClientKey() bool {
	return len(c.Key) > 0 || len(c.KeyFile) > 0
}

func (c *TLSConfig) roundTripperSettings() TLSRoundTripperSettings {
	return TLSRoundTripperSettings{
		CA:       c.CA,
		CAFile:   c.CAFile,
		Cert:     c.Cert,
		CertFile: c.CertFile,
		Key:      string(c.Key),
		KeyFile:  c.KeyFile,
	}
}

// getClientCertificate reads the pair of client cert and key from disk and returns a tls.Certificate.
func (c *TLSConfig) getClientCertificate(_ *tls.CertificateRequestInfo) (*tls.Certificate, error) {
	var (
		certData, keyData []byte
		err               error
	)

	if c.CertFile != "" {
		certData, err = os.ReadFile(c.CertFile)
		if err != nil {
			return nil, fmt.Errorf("unable to read specified client cert (%s): %w", c.CertFile, err)
		}
	} else {
		certData = []byte(c.Cert)
	}

	if c.KeyFile != "" {
		keyData, err = os.ReadFile(c.KeyFile)
		if err != nil {
			return nil, fmt.Errorf("unable to read specified client key (%s): %w", c.KeyFile, err)
		}
	} else {
		keyData = []byte(c.Key)
	}

	cert, err := tls.X509KeyPair(certData, keyData)
	if err != nil {
		return nil, fmt.Errorf("unable to use specified client cert (%s) & key (%s): %w", c.CertFile, c.KeyFile, err)
	}

	return &cert, nil
}

// Validate validates the TLSConfig to check that only one of the inlined or
// file-based fields for the TLS CA, client certificate, and client key are
// used.
func (c *TLSConfig) Validate() error {
	if len(c.CA) > 0 && len(c.CAFile) > 0 {
		return fmt.Errorf("at most one of ca and ca_file must be configured")
	}
	if len(c.Cert) > 0 && len(c.CertFile) > 0 {
		return fmt.Errorf("at most one of cert and cert_file must be configured")
	}
	if len(c.Key) > 0 && len(c.KeyFile) > 0 {
		return fmt.Errorf("at most one of key and key_file must be configured")
	}

	if c.usingClientCert() && !c.usingClientKey() {
		return fmt.Errorf("exactly one of key or key_file must be configured when a client certificate is configured")
	} else if c.usingClientKey() && !c.usingClientCert() {
		return fmt.Errorf("exactly one of cert or cert_file must be configured when a client key is configured")
	}

	return nil
}

// TcpScriptPlugin 实现了IScriptPlugin接口的tcp插件
type TcpScriptPlugin struct {
	logger             *zap.Logger
	Deadline           int
	IPProtocol         int
	IPProtocolFallback bool
	SourceIPAddress    string
	Expect             string
	Send               string
	StartTLS           bool
	TLS                bool
	TLSConfig          TLSConfig
}

func NewTcpScriptPlugin(logger *zap.Logger) *TcpScriptPlugin {
	return &TcpScriptPlugin{
		logger:   logger,
		Deadline: 5,
	}
}

// @return string tcp插件
func (thisPlugin TcpScriptPlugin) GetCode() string {
	return "tcp"
}

// CallPluginsMethod 实现IScriptPlugin接口的CallPluginsMethod方法
func (thisPlugin TcpScriptPlugin) CallPluginsMethod(method string, params interface{}) interface{} {
	switch method {
	case "probe":
		{
			thisPlugin.probe(params.(string))
		}
	}
	return nil
}

func (thisPlugin *TcpScriptPlugin) SetLogger(logger *zap.Logger) {
	thisPlugin.logger = logger
}

func (thisPlugin *TcpScriptPlugin) probe(target string) bool {
	logger := thisPlugin.logger
	ctx, _ := context.WithDeadline(context.Background(),
		time.Now().Add(time.Duration(thisPlugin.Deadline)*time.Second))

	deadline, _ := ctx.Deadline()

	conn, err := thisPlugin.dialTCP(ctx, target)
	if err != nil {
		logger.Info(fmt.Sprint("Error dialing TCP", "err", err))
		return false
	}
	defer conn.Close()
	logger.Info(fmt.Sprint("Successfully dialed"))

	// Set a deadline to prevent the following code from blocking forever.
	// If a deadline cannot be set, better fail the probe by returning an error
	// now rather than blocking forever.
	if err := conn.SetDeadline(deadline); err != nil {
		logger.Info(fmt.Sprint("Error setting deadline", "err", err))
		return false
	}
	if thisPlugin.TLS {
		state := conn.(*tls.Conn).ConnectionState()
		fmt.Println("cert expire ", getEarliestCertExpiry(&state),
			"tls version ", getTLSVersion(&state),
			"last chain expire ", getLastChainExpiry(&state),
			"fingerprint", getFingerprint(&state),
			"subject", getSubject(&state),
			"issuer", getIssuer(&state),
			"dnsnames", getDNSNames(&state),
			"ciphers", getTLSCipher(&state))
	}
	return true
}

func (thisPlugin *TcpScriptPlugin) dialTCP(ctx context.Context, target string) (net.Conn, error) {
	var dialProtocol, dialTarget string
	logger := thisPlugin.logger
	dialer := &net.Dialer{}
	targetAddress, port, err := net.SplitHostPort(target)
	if err != nil {
		logger.Info(fmt.Sprint("Error splitting target address and port", "err", err))
		return nil, err
	}

	ip, _, _, err := utils.ChooseProtocol(ctx, thisPlugin.IPProtocol, thisPlugin.IPProtocolFallback, targetAddress, thisPlugin.logger)
	if err != nil {
		logger.Info(fmt.Sprint("Error resolving address", "err", err))
		return nil, err
	}

	if ip.IP.To4() == nil {
		dialProtocol = "tcp6"
	} else {
		dialProtocol = "tcp4"
	}

	if len(thisPlugin.SourceIPAddress) > 0 {
		srcIP := net.ParseIP(thisPlugin.SourceIPAddress)
		if srcIP == nil {
			logger.Info(fmt.Sprint("Error parsing source ip address", "srcIP", thisPlugin.SourceIPAddress))
			return nil, fmt.Errorf("error parsing source ip address: %s", thisPlugin.SourceIPAddress)
		}
		logger.Info(fmt.Sprint("Using local address", "srcIP", srcIP))
		dialer.LocalAddr = &net.TCPAddr{IP: srcIP}
	}

	dialTarget = net.JoinHostPort(ip.String(), port)

	if !thisPlugin.TLS {
		logger.Info(fmt.Sprint("Dialing TCP without TLS"))
		return dialer.DialContext(ctx, dialProtocol, dialTarget)
	}
	tlsConfig, err := newTLSConfig(&thisPlugin.TLSConfig)
	if err != nil {
		logger.Info(fmt.Sprint("Error creating TLS configuration", "err", err))
		return nil, err
	}

	if len(tlsConfig.ServerName) == 0 {
		// If there is no `server_name` in tls_config, use
		// targetAddress as TLS-servername. Normally tls.DialWithDialer
		// would do this for us, but we pre-resolved the name by
		// `chooseProtocol` and pass the IP-address for dialing (prevents
		// resolving twice).
		// For this reason we need to specify the original targetAddress
		// via tlsConfig to enable hostname verification.
		tlsConfig.ServerName = targetAddress
	}
	timeoutDeadline, _ := ctx.Deadline()
	dialer.Deadline = timeoutDeadline

	logger.Info(fmt.Sprint("Dialing TCP with TLS"))
	return tls.DialWithDialer(dialer, dialProtocol, dialTarget, tlsConfig)
}
func getEarliestCertExpiry(state *tls.ConnectionState) time.Time {
	earliest := time.Time{}
	for _, cert := range state.PeerCertificates {
		if (earliest.IsZero() || cert.NotAfter.Before(earliest)) && !cert.NotAfter.IsZero() {
			earliest = cert.NotAfter
		}
	}
	return earliest
}

func getFingerprint(state *tls.ConnectionState) string {
	cert := state.PeerCertificates[0]
	fingerprint := sha256.Sum256(cert.Raw)
	return hex.EncodeToString(fingerprint[:])
}

func getSubject(state *tls.ConnectionState) string {
	cert := state.PeerCertificates[0]
	return cert.Subject.String()
}

func getIssuer(state *tls.ConnectionState) string {
	cert := state.PeerCertificates[0]
	return cert.Issuer.String()
}

func getDNSNames(state *tls.ConnectionState) string {
	cert := state.PeerCertificates[0]
	return strings.Join(cert.DNSNames, ",")
}

func getLastChainExpiry(state *tls.ConnectionState) time.Time {
	lastChainExpiry := time.Time{}
	for _, chain := range state.VerifiedChains {
		earliestCertExpiry := time.Time{}
		for _, cert := range chain {
			if (earliestCertExpiry.IsZero() || cert.NotAfter.Before(earliestCertExpiry)) && !cert.NotAfter.IsZero() {
				earliestCertExpiry = cert.NotAfter
			}
		}
		if lastChainExpiry.IsZero() || lastChainExpiry.Before(earliestCertExpiry) {
			lastChainExpiry = earliestCertExpiry
		}

	}
	return lastChainExpiry
}

func getTLSVersion(state *tls.ConnectionState) string {
	switch state.Version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS13:
		return "TLS 1.3"
	default:
		return "unknown"
	}
}

func getTLSCipher(state *tls.ConnectionState) string {
	return tls.CipherSuiteName(state.CipherSuite)
}

func newTLSConfig(cfg *TLSConfig) (*tls.Config, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: cfg.InsecureSkipVerify,
		MinVersion:         uint16(cfg.MinVersion),
		MaxVersion:         uint16(cfg.MaxVersion),
	}

	if cfg.MaxVersion != 0 && cfg.MinVersion != 0 {
		if cfg.MaxVersion < cfg.MinVersion {
			return nil, fmt.Errorf("tls_config.max_version must be greater than or equal to tls_config.min_version if both are specified")
		}
	}

	// If a CA cert is provided then let's read it in so we can validate the
	// scrape target's certificate properly.
	if len(cfg.CA) > 0 {
		if !updateRootCA(tlsConfig, []byte(cfg.CA)) {
			return nil, fmt.Errorf("unable to use inline CA cert")
		}
	} else if len(cfg.CAFile) > 0 {
		b, err := readCAFile(cfg.CAFile)
		if err != nil {
			return nil, err
		}
		if !updateRootCA(tlsConfig, b) {
			return nil, fmt.Errorf("unable to use specified CA cert %s", cfg.CAFile)
		}
	}

	if len(cfg.ServerName) > 0 {
		tlsConfig.ServerName = cfg.ServerName
	}

	// If a client cert & key is provided then configure TLS config accordingly.
	if cfg.usingClientCert() && cfg.usingClientKey() {
		// Verify that client cert and key are valid.
		if _, err := cfg.getClientCertificate(nil); err != nil {
			return nil, err
		}
		tlsConfig.GetClientCertificate = cfg.getClientCertificate
	}

	return tlsConfig, nil
}

// readCAFile reads the CA cert file from disk.
func readCAFile(f string) ([]byte, error) {
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, fmt.Errorf("unable to load specified CA cert %s: %w", f, err)
	}
	return data, nil
}

// updateRootCA parses the given byte slice as a series of PEM encoded certificates and updates tls.Config.RootCAs.
func updateRootCA(cfg *tls.Config, b []byte) bool {
	caCertPool := x509.NewCertPool()
	if !caCertPool.AppendCertsFromPEM(b) {
		return false
	}
	cfg.RootCAs = caCertPool
	return true
}
