package test

import (
	htIstio "github.com/krogebry/golibs/helm_tests/istio"
	htOAuth2Proxy "github.com/krogebry/golibs/helm_tests/oauth2_proxy"

	htTelegraf "github.com/krogebry/golibs/helm_tests/telegraf"
	"testing"
)

var telegrafContainerOptions htTelegraf.TelegrafContainerOptions
var oAuth2ProxyContainerOptions htOAuth2Proxy.OAuth2ProxyContainerOptions

// OAuth Proxy tests
func TestOAuth2ProxyReadinessProbe(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyContainerSecurityContext(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyLivenessProbe(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyLivenessProbe(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyContainerSecurityContext(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyContainerSecurityContext(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyPorts(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyPorts(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyArgs(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyArgs(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyImagePullPolicy(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyImagePullPolicy(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyEnvironmentVars(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyEnvironmentVars(t, config, oAuth2ProxyContainerOptions)
}

func TestOAuth2ProxyVolumeMounts(t *testing.T) {
	htOAuth2Proxy.TestOAuth2ProxyVolumeMounts(t, config, oAuth2ProxyContainerOptions)
}

// Telegraf tests
func TestTelegrafReadinessProbe(t *testing.T) {
	htTelegraf.TestTelegrafReadinessProbe(t, config, telegrafContainerOptions)
}

func TestTelegrafLivenessProbe(t *testing.T) {
	htTelegraf.TestTelegrafLivenessProbe(t, config, telegrafContainerOptions)
}

func TestTelegrafPorts(t *testing.T) {
	htTelegraf.TestTelegrafPorts(t, config, telegrafContainerOptions)
}

func TestTelegrafVolumeMounts(t *testing.T) {
	htTelegraf.TestTelegrafVolumeMounts(t, config, telegrafContainerOptions)
}

func TestGatewayDoesNotExist(t *testing.T) {
	htIstio.TestGatewayDoesNotExist(t, config, htIstio.GatewayOptions{})
}
