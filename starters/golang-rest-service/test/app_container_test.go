package test

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"

	"testing"
)

func TestApplicationVolumeMounts(t *testing.T) {
	container := getContainer(t, config.applicationName, config.name)

	assert.Len(t, container.VolumeMounts, 1, "Should only have 1 mount")
	assert.Contains(t, container.VolumeMounts, v1.VolumeMount{
		Name:      "tmp",
		MountPath: "/tmp",
	}, "/tmp should be mounted")

}

func TestApplicationPorts(t *testing.T) {
	container := getContainer(t, config.applicationName, config.name)

	assert.Len(t, container.Ports, 2)

	assert.Contains(t, container.Ports, v1.ContainerPort{
		Name:          "http",
		ContainerPort: 8080,
		Protocol:      "TCP",
	})

	assert.Contains(t, container.Ports, v1.ContainerPort{
		Name:          "metrics",
		ContainerPort: 2112,
		Protocol:      "TCP",
	})

}

func TestApplicationEnvironmentVars(t *testing.T) {
	container := getContainer(t, config.applicationName, config.name)

	assert.Len(t, container.Env, 0)
}

func TestApplicationImagePullPolicy(t *testing.T) {
	container := getContainer(t, config.applicationName, config.name)
	assert.Equal(t, container.ImagePullPolicy, v1.PullPolicy("Always"))
}
