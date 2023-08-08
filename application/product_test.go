package application_test

import (
	"github.com/edwildson/go-hexagonal/application"
	"testing"
	"github.com/stretchr/testify/require"
)


func TestProduct_Enable(t *testing.T)  {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0

	err = product.Enable()

	require.Equal(t, "T he price must be greater than zero to enable the product", err.Error())
}