package test

import (
	"github.com/external-secrets/external-secrets/apis/externalsecrets/v1beta1"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestSecretStore(t *testing.T) {
	output, _ := k8s.RunKubectlAndGetOutputE(t, &config.kubectlOptions, "get", "secretstore",
		config.stackId, "-o", "yaml")
	secretStoreVault := v1beta1.SecretStore{}
	helm.UnmarshalK8SYaml(t, output, &secretStoreVault)

	assert.Equal(t, *secretStoreVault.Spec.Provider.Vault.Auth.Jwt.KubernetesServiceAccountToken.ExpirationSeconds, int64(600))
	assert.Contains(t, *secretStoreVault.Spec.Provider.Vault.Auth.Jwt.KubernetesServiceAccountToken.Audiences, "vault")
	assert.Equal(t, secretStoreVault.Spec.Provider.Vault.Auth.Jwt.KubernetesServiceAccountToken.ServiceAccountRef.Name, config.stackId)
	assert.Equal(t, secretStoreVault.Spec.Provider.Vault.Auth.Jwt.Path, "jwt_mtfuji_gc_0_apps_us_east_1")
	assert.Equal(t, secretStoreVault.Spec.Provider.Vault.Auth.Jwt.Role, "jwt_mtfuji_gc_0_apps_us_east_1_cloud_services")

	assert.Equal(t, *secretStoreVault.Spec.Provider.Vault.Namespace, "ns_stargate/ns_dev_arenepe_cloudservicesdevops")
	assert.Equal(t, *secretStoreVault.Spec.Provider.Vault.Path, "devops")
	assert.Equal(t, secretStoreVault.Spec.Provider.Vault.Server, "https://dev.vault.tmc-stargate.com")
	assert.Equal(t, secretStoreVault.Spec.Provider.Vault.Version, v1beta1.VaultKVStoreV2)
}
