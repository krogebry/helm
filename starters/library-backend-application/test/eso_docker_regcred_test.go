package test

import (
	"github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExternalSecretStore(t *testing.T) {
	output, _ := k8s.RunKubectlAndGetOutputE(t, &config.kubectlOptions, "get", "externalsecret",
		config.stackId, "-o", "yaml")
	externalSecret := v1beta1.ExternalSecret{}
	helm.UnmarshalK8SYaml(t, output, &externalSecret)

	assert.Equal(t, externalSecret.Spec.SecretStoreRef.Kind, "SecretStore")
	assert.Equal(t, externalSecret.Spec.SecretStoreRef.Name, config.stackId)

	assert.Equal(t, externalSecret.Spec.Data[0].RemoteRef.Key, "stargate-artifactory-regcred-docker")
}
