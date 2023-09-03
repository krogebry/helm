package test

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"testing"
)

func TestApplicationVolumeMounts(t *testing.T) {
	container := getContainer(t, config.name, config.name)

	assert.Len(t, container.VolumeMounts, 1, "Should only have 1 mount")
	assert.Contains(t, container.VolumeMounts, v1.VolumeMount{
		Name:      "tmp",
		MountPath: "/tmp",
	}, "/tmp should be mounted")

}

func TestApplicationPorts(t *testing.T) {
	container := getContainer(t, config.name, config.name)

	assert.Len(t, container.Ports, 1)

	assert.Contains(t, container.Ports, v1.ContainerPort{
		Name:          "http",
		ContainerPort: 8080,
		Protocol:      "TCP",
	})

}

func TestApplicationEnvironmentVars(t *testing.T) {
	container := getContainer(t, config.containerName, config.name)
	assert.Len(t, container.Env, 0)
}

func TestApplicationImagePullPolicy(t *testing.T) {
	container := getContainer(t, config.name, config.name)
	assert.Equal(t, container.ImagePullPolicy, v1.PullPolicy("Always"))
}

func TestDeploymentApplicationReadinessProbe(t *testing.T) {
	container := getContainer(t, config.name, config.name)

	assert.Equal(t, container.ReadinessProbe.InitialDelaySeconds, int32(0))
	assert.Equal(t, container.ReadinessProbe.SuccessThreshold, int32(1))
	assert.Equal(t, container.ReadinessProbe.PeriodSeconds, int32(10))
	assert.Equal(t, container.ReadinessProbe.FailureThreshold, int32(3))

	assert.Equal(t, container.ReadinessProbe.ProbeHandler, v1.ProbeHandler{
		HTTPGet: &v1.HTTPGetAction{
			Path:   "/healthz",
			Port:   intstr.IntOrString{Type: 1, StrVal: "http"},
			Scheme: v1.URISchemeHTTP,
		},
	})
}

func TestDeploymentApplicationLivenessProbe(t *testing.T) {
	container := getContainer(t, config.name, config.name)

	assert.Equal(t, container.LivenessProbe.InitialDelaySeconds, int32(0))
	assert.Equal(t, container.LivenessProbe.SuccessThreshold, int32(1))
	assert.Equal(t, container.LivenessProbe.PeriodSeconds, int32(10))
	assert.Equal(t, container.LivenessProbe.FailureThreshold, int32(3))

	assert.Equal(t, container.LivenessProbe.ProbeHandler, v1.ProbeHandler{
		HTTPGet: &v1.HTTPGetAction{
			Path:   "/healthz",
			Port:   intstr.IntOrString{Type: 1, StrVal: "http"},
			Scheme: v1.URISchemeHTTP,
		},
	})
}
