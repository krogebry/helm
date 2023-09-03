package test

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"testing"
)

func TestServicePorts(t *testing.T) {
	service := k8s.GetService(t, &config.kubectlOptions, config.name)

	assert.Len(t, service.Spec.Ports, 1)
	assert.Contains(t, service.Spec.Ports, v1.ServicePort{
		Name:       "http",
		Protocol:   "TCP",
		Port:       8080,
		TargetPort: intstr.IntOrString{Type: 1, StrVal: "http"},
	})
}
