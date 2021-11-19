package application_test

import (
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/wesleysaraujo/arquitetura-hexagonal-golang/application"
	"testing"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Redmi Note 10"
	product.Status = application.DISABLED
	product.Price = 1290.99

	err := product.Enable()

	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enable the product.", err.Error())
}

func TestProduct_Disabled(t *testing.T) {
	product := application.Product{}
	product.Name = "Redmi Note 10"
	product.Status = application.DISABLED
	product.Price = 0

	err := product.Disabled()

	require.Nil(t, err)

	product.Price = 1299.90
	err = product.Disabled()
	require.Equal(t, "Price must be zero in order to have the product disabled!", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := application.Product{}
	product.ID = uuid.New().String()
	product.Name = "Redmi Note 10"
	product.Status = application.DISABLED
	product.Price = 1000

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"

	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED

	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -1200

	_, err =  product.IsValid()
	require.Equal(t, "the price must be greater or equal zero", err.Error())
}

func TestProduct_GetID(t *testing.T) {
	id := uuid.New().String()

	product := application.Product{}
	product.ID = id

	require.Equal(t, id, product.GetID())
}

func TestProduct_GetName(t *testing.T) {
	name := "Redmi Note 10"

	product := application.Product{}
	product.Name = name

	require.Equal(t, name, product.GetName())
}

func TestProduct_GetStatus(t *testing.T) {
	product := application.Product{}
	product.Status = application.ENABLED

	require.Equal(t, application.ENABLED, product.GetStatus())
}

func TestProduct_GetPrice(t *testing.T) {
	product := application.Product{}
	product.Price = 1299.90

	require.Equal(t, float32(1299.90), product.GetPrice())
}


