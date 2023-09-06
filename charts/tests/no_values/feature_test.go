package test

import (
	htIstio "github.com/krogebry/golibs/helm_tests/istio"

	"github.com/stretchr/testify/assert"
	htUtils "github.com/krogebry/golibs/helm_tests/utils"
	v1 "k8s.io/api/core/v1"
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

func TestGatewayDoesNotExist(t *testing.T) {
	htIstio.TestGatewayDoesNotExist(t, config, htIstio.GatewayOptions{})
}
