package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	htCommon "github.com/krogebry/golibs/helm_tests/common"
	htIstio "github.com/krogebry/golibs/helm_tests/istio"
	htSecrets "github.com/krogebry/golibs/helm_tests/secrets"
	htUtils "github.com/krogebry/golibs/helm_tests/utils"
	"os"
	"testing"
)

var config htUtils.Config
var applicationContainerOptions htCommon.ApplicationContainerOptions
var deploymentOptions htCommon.DeploymentOptions
var serviceOptions htCommon.ServiceOptions
var secretStoreVaultOptions htSecrets.SecretStoreVaultOptions

var istioVirtualServiceOptions htIstio.VirtualServiceOptions

func TestMain(m *testing.M) {
	//config := htUtils.Config{}
	config.StackId = os.Getenv("STACK_ID")
	config.NamespaceName = os.Getenv("NAMESPACE")

	debug := os.Getenv("DEBUG")
	if debug == "true" {
		config.Debug = true
	}

	deploymentOptions.NumberOfVolumes = 2

	config.ContainerName = config.StackId
	config.DeploymentName = config.StackId

	config.KubectlOptions = k8s.KubectlOptions{
		Namespace: config.NamespaceName,
	}

	deploymentOptions.ServiceAccountName = config.StackId
	deploymentOptions.ImagePullSecretName = fmt.Sprintf("%s-stargate-artifactory-docker", config.ContainerName)

	m.Run()
}

// Application Container Tests
func TestApplicationVolumeMounts(t *testing.T) {
	htCommon.TestApplicationVolumeMounts(t, config, applicationContainerOptions)
}

func TestApplicationPorts(t *testing.T) {
	htCommon.TestApplicationPorts(t, config, applicationContainerOptions)
}

func TestApplicationImagePullPolicy(t *testing.T) {
	htCommon.TestApplicationImagePullPolicy(t, config, applicationContainerOptions)
}

func TestDeploymentApplicationReadinessProbe(t *testing.T) {
	htCommon.TestDeploymentApplicationReadinessProbe(t, config, applicationContainerOptions)
}

func TestDeploymentApplicationLivenessProbe(t *testing.T) {
	htCommon.TestDeploymentApplicationLivenessProbe(t, config, applicationContainerOptions)
}

func TestApplicationContainerSecurityContext(t *testing.T) {
	htCommon.TestApplicationContainerSecurityContext(t, config, applicationContainerOptions)
}

func TestApplicationContainerResources(t *testing.T) {
	htCommon.TestApplicationContainerResources(t, config, applicationContainerOptions)
}

// Deployment tests
func TestDeploymentImagePullSecrets(t *testing.T) {
	htCommon.TestDeploymentImagePullSecrets(t, config, deploymentOptions)
}

func TestDeploymentVolumes(t *testing.T) {
	htCommon.TestDeploymentVolumes(t, config, deploymentOptions)
}

func TestDeploymentSecurity(t *testing.T) {
	htCommon.TestDeploymentSecurity(t, config, deploymentOptions)
}

// Secret tests
func TestSecretStore(t *testing.T) {
	secretStoreVaultOptions.JWTRole = "jwt_mtfuji_gc_0_apps_us_east_1_cloud_services_dev"
	secretStoreVaultOptions.ServiceAccountName = "cloud-services-grpc"
	htSecrets.TestSecretStoreForVault(t, config, secretStoreVaultOptions)
}

func TestExternalSecretDockerRegcred(t *testing.T) {
	htSecrets.TestExternalSecret(t, config, htSecrets.ExternalSecretOptions{
		Name:               fmt.Sprintf("%s-stargate-artifactory-docker", config.ContainerName),
		RemoteKeyRef:       "stargate-artifactory-regcred-docker",
		SecretStoreRefName: config.ContainerName,
	})
}

func TestArtifactoryRegistrySecret(t *testing.T) {
	htSecrets.TestArtifactoryRegistrySecret(t, config, htSecrets.ArtifactoryRegistrySecretOptions{
		Name: "ext-docker-regcred",
	})
}

func TestDefaultVirtualService(t *testing.T) {
	istioVirtualServiceOptions.Name = config.ContainerName
	istioVirtualServiceOptions.Region = "us-east-1"
	istioVirtualServiceOptions.Hostname = config.ContainerName
	htIstio.TestDefaultVirtualService(t, config, istioVirtualServiceOptions)
}

// Service tests
func TestDefaultService(t *testing.T) {
	htCommon.TestDefaultService(t, config, serviceOptions)
}
