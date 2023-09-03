package test

import (
	v1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"
	htCommon "github.com/krogebry/golibs/helm_tests/common"
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

func TestGatewayDoesNotExist(t *testing.T) {
	htIstio.TestGatewayDoesNotExist(t, config, htIstio.GatewayOptions{})
}

var asgOptions = htCommon.HorizontalPodAutoscalerOptions{}

func TestHorizontalPodAutoscaler(t *testing.T) {
	htCommon.TestHorizontalPodAutoscaler(t, config, asgOptions)
}
