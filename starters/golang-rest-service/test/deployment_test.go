package test

import (
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"testing"
)

func TestImagePullSecrets(t *testing.T) {
	deployment := getDeployment(t, config.name)
	assert.Contains(t, deployment.Spec.Template.Spec.ImagePullSecrets, v1.LocalObjectReference{Name: "ext-docker-regcred"})
}

func TestDeploymentVolumes(t *testing.T) {
	deployment := getDeployment(t, config.name)

	assert.Len(t, deployment.Spec.Template.Spec.Volumes, 1, "There should be exactly 1 volume setup")

	assert.Contains(t, deployment.Spec.Template.Spec.Volumes, v1.Volume{
		Name: "tmp",
		VolumeSource: v1.VolumeSource{
			EmptyDir: &v1.EmptyDirVolumeSource{},
		},
	})

}

func TestDeploymentSecurity(t *testing.T) {
	deployment := getDeployment(t, config.name)
	assert.True(t, *deployment.Spec.Template.Spec.AutomountServiceAccountToken)
	assert.Equal(t, deployment.Spec.Template.Spec.ServiceAccountName, config.name)
}
