package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	v1 "k8s.io/api/core/v1"

	"os"
	"testing"
)

var config struct {
	name           string
	stackId        string
	kubectlOptions k8s.KubectlOptions
	namespaceName  string
	containerName  string
}

var cache struct {
	deployments []appsv1.Deployment
}

func TestMain(m *testing.M) {
	config.stackId = os.Getenv("STACK_ID")
	config.namespaceName = os.Getenv("NAMESPACE")

	config.name = config.stackId

	config.kubectlOptions = k8s.KubectlOptions{
		Namespace: config.namespaceName,
	}

	m.Run()
}

func getContainer(t *testing.T, containerName string, deploymentName string) v1.Container {
	deployment := getDeployment(t, deploymentName)
	assert.NotNil(t, deployment)
	container := getContainerf(containerName, &deployment)
	assert.NotNil(t, container)
	return container
}

func getContainerf(containerName string, deployment *appsv1.Deployment) v1.Container {
	for _, container := range deployment.Spec.Template.Spec.Containers {
		if container.Name == containerName {
			return container
		}
	}
	return v1.Container{}
}

func getDeployment(t *testing.T, instanceName string) appsv1.Deployment {
	logger.Log(t, "Getting cache")
	for _, deployment := range cache.deployments {
		for k, v := range deployment.Labels {
			logger.Log(t, fmt.Sprintf("Label: %s - %s", k, v))
			if k == "app.kubernetes.io/instance" && v == instanceName {
				logger.Log(t, fmt.Sprintf("Found deployment: %s", deployment.Name))
				return deployment
			}
		}
	}
	logger.Log(t, "No cache found, searching")
	deployments := k8s.ListDeployments(t, &config.kubectlOptions, metav1.ListOptions{LabelSelector: fmt.Sprintf("app.kubernetes.io/instance=%s", instanceName)})
	cache.deployments = append(cache.deployments, deployments[0])
	return deployments[0]
}
