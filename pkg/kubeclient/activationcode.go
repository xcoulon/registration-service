package kubeclient

import (
	"context"

	crtapi "github.com/codeready-toolchain/api/api/v1alpha1"
)

const (
	activationCodeResourcePlural = "activationcodes"
)

type ActivationCodeInterface interface {
	GetByCode(code string) (*crtapi.ActivationCode, error)
}

type activationCodeClient struct {
	crtClient
}

func (c *activationCodeClient) GetByCode(code string) (*crtapi.ActivationCode, error) {
	result := &crtapi.ActivationCode{}
	err := c.client.Get().
		Namespace(c.ns).
		Resource(activationCodeResourcePlural).
		Name(code).
		Do(context.TODO()).
		Into(result)
	return result, err
}
