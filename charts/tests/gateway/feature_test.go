package test

import (
	v1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"
	htIstio "github.com/krogebry/golibs/helm_tests/istio"
	htUtils "github.com/krogebry/golibs/helm_tests/utils"
	"testing"
)

func TestTelegrafContainerShouldNotExist(t *testing.T) {
	deployment := htUtils.GetDeployment(t, config, config.DeploymentName)
	assert.NotContains(t, deployment.Spec.Template.Spec.Containers, v1.Container{Name: "telegraf"})
}

func TestOAuth2ContainerShouldNotExist(t *testing.T) {
	deployment := htUtils.GetDeployment(t, config, config.DeploymentName)
	assert.NotContains(t, deployment.Spec.Template.Spec.Containers, v1.Container{Name: "oauth2_proxy"})
}

var gatewayOptions = htIstio.GatewayOptions{}

func TestGateway(t *testing.T) {
	gatewayOptions.Region = "us-east-1"
	gatewayOptions.Hostname = config.ContainerName
	htIstio.TestGateway(t, config, gatewayOptions)
}

func TestVirtualService(t *testing.T) {
	istioVirtualServiceOptions.Name = config.ContainerName
	istioVirtualServiceOptions.Region = "us-east-1"
	istioVirtualServiceOptions.Hostname = config.ContainerName
	istioVirtualServiceOptions.GatewayName = config.ContainerName
	htIstio.TestDefaultVirtualService(t, config, istioVirtualServiceOptions)
}
