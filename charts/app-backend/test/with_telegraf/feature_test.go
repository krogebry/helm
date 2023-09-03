package test

import (
	"github.com/stretchr/testify/assert"
	htIstio "github.com/krogebry/golibs/helm_tests/istio"
	htTelegraf "github.com/krogebry/golibs/helm_tests/telegraf"
	v1 "k8s.io/api/core/v1"
	"testing"

	htUtils "github.com/krogebry/golibs/helm_tests/utils"
)

var telegrafContainerOptions htTelegraf.TelegrafContainerOptions

func TestOAuth2ContainerShouldNotExist(t *testing.T) {
	deployment := htUtils.GetDeployment(t, config, config.DeploymentName)
	assert.NotContains(t, deployment.Spec.Template.Spec.Containers, v1.Container{Name: "oauth2_proxy"})
}

func TestTelegrafContainerShouldNotExist(t *testing.T) {
	deployment := htUtils.GetDeployment(t, config, config.DeploymentName)
	assert.NotContains(t, deployment.Spec.Template.Spec.Containers, v1.Container{Name: "telegraf"})
}

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
