package test

import (
	htIstio "github.com/krogebry/golibs/helm_tests/istio"

	"fmt"
	"github.com/stretchr/testify/assert"
	htSecrets "github.com/krogebry/golibs/helm_tests/secrets"
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

// Quay secret should exist
func TestExternalSecretDockerRegcredQuay(t *testing.T) {
	htSecrets.TestExternalSecret(t, config, htSecrets.ExternalSecretOptions{
		Name:               fmt.Sprintf("%s-stargate-artifactory-docker-quay", config.ContainerName),
		RemoteKeyRef:       "stargate-artifactory-regcred-quay",
		SecretStoreRefName: config.ContainerName,
	})
}

func TestArtifactoryRegistrySecretQuay(t *testing.T) {
	htSecrets.TestArtifactoryRegistrySecret(t, config, htSecrets.ArtifactoryRegistrySecretOptions{
		Name: fmt.Sprintf("%s-stargate-artifactory-docker-quay", config.ContainerName),
	})
}
