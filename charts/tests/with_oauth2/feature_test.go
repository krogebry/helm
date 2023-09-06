package test

import (
	"github.com/stretchr/testify/assert"
	htIstio "github.com/krogebry/golibs/helm_tests/istio"
	htOAuth2Proxy "github.com/krogebry/golibs/helm_tests/oauth2_proxy"
	htUtils "github.com/krogebry/golibs/helm_tests/utils"
	v1 "k8s.io/api/core/v1"
	"testing"
)

var oAuth2ProxyContainerOptions htOAuth2Proxy.OAuth2ProxyContainerOptions

func TestTelegrafContainerShouldNotExist(t *testing.T) {
	deployment := htUtils.GetDeployment(t, config, config.DeploymentName)
	assert.NotContains(t, deployment.Spec.Template.Spec.Containers, v1.Container{Name: "telegraf"})
}

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

func TestGatewayDoesNotExist(t *testing.T) {
	htIstio.TestGatewayDoesNotExist(t, config, htIstio.GatewayOptions{})
}
