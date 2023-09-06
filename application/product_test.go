package application_test

import (
	"github.com/c4st3ll4n/go-hexagon/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 100

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "the price must be bigger than 0 to enable the product", err.Error())
}

func TestProduct_Disable(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 100
	err = product.Disable()
	require.Equal(t, "the price must be zero to disable the product", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.Name = "Product 1"
	product.Status = application.DISABLED
	product.Price = 10
	product.ID = uuid.NewV4().String()

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "price must be zero or bigger", err.Error())
}
