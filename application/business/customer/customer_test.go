package businesscustomer_test

import (
	businesscustomer "github.com/kennymack/go-hexagonal-product/application/business/customer"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCustomer_Enable(t *testing.T) {
	customer := businesscustomer.Customer{}
	customer.Name = "Marcia"
	customer.Status = businesscustomer.DISABLED
	customer.Email = "marcia@email.com"
	customer.DeactivationReason = ""

	err := customer.Enable()

	require.Nil(t, err)

	customer.Email = ""

	err = customer.Enable()

	require.Equal(t, "the name and e-mail must be informed", err.Error())

	customer.Email = "marcia@email.com"
	customer.Name = ""

	err = customer.Enable()

	require.Equal(t, "the name and e-mail must be informed", err.Error())
}

func TestCustomer_Disable(t *testing.T) {
	customer := businesscustomer.Customer{}
	customer.Name = "Marcia"
	customer.Status = businesscustomer.DISABLED
	customer.Email = "marcia@email.com"
	customer.DeactivationReason = "Old customer"

	err := customer.Disable()

	require.Nil(t, err)

	customer.DeactivationReason = ""

	err = customer.Disable()

	require.Equal(t, "the reason for deactivation must be informed", err.Error())
}

func TestCustomer_IsValid(t *testing.T) {
	customer := businesscustomer.Customer{}
	customer.ID = uuid.NewV4().String()
	customer.Name = "Marcia"
	customer.Status = businesscustomer.ENABLED
	customer.Email = "marcia@email.com"
	customer.DeactivationReason = ""

	_, err := customer.IsValid()

	require.Nil(t, err)

	customer.Status = "INVALID"

	_, err = customer.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	customer.DeactivationReason = ""
	customer.Status = businesscustomer.DISABLED

	_, err = customer.IsValid()

	require.Equal(t, "the reason for deactivation must be informed", err.Error())

	customer.DeactivationReason = "Old"
	customer.Status = businesscustomer.ENABLED

	_, err = customer.IsValid()

	require.Equal(t, "the reason for deactivation should not be informed", err.Error())
}
