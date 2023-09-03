package test

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtRegcredSecret(t *testing.T) {
	secret := k8s.GetSecret(t, &config.kubectlOptions, "ext-docker-regcred")
	assert.NotEmpty(t, secret.Data[".dockerconfigjson"])
}
