package businesscustomer_test

import (
	"github.com/golang/mock/gomock"
	"github.com/kennymack/go-hexagonal-product/application"
	businesscustomer "github.com/kennymack/go-hexagonal-product/application/business/customer"
	mock_bucustomer "github.com/kennymack/go-hexagonal-product/application/business/customer/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCustomerService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	customer := mock_bucustomer.NewMockCustomerInterface(ctrl)
	persistence := mock_bucustomer.NewMockCustomerPersistenceInterface(ctrl)
	persistence.EXPECT().Get(gomock.Any()).Return(customer, nil).AnyTimes()

	service := businesscustomer.CustomerService{
		Persistence: persistence,
	}

	result, err := service.Get("abc")

	require.Nil(t, err)
	require.Equal(t, customer, result)
}
