package cli_test

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/kennymack/go-hexagonal-product/adapters/cli"
	"github.com/kennymack/go-hexagonal-product/application"
	mock_application "github.com/kennymack/go-hexagonal-product/application/mocks"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	productName := "Product 3"
	productPrice := 40.0
	productStatus := application.ENABLED
	productID := "abc3"

	productMock := mock_application.NewMockProductInterface(ctrl)

	productMock.EXPECT().GetId().Return(productID).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)

	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productID).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	resultCreateExpected := fmt.Sprintf("ID %s Name %s has been created", productID, productName)
	resultEnableExpected := fmt.Sprintf("ID %s Name %s has been enabled", productID, productName)
	resultDisableExpected := fmt.Sprintf("ID %s Name %s has been disabled", productID, productName)
	resultGetExpected := fmt.Sprintf("ID:%s|NAME:%s|PRICE:%v|STATUS:%s", productID, productName, productPrice, productStatus)

	result, err := cli.Run(service, "create", productID, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, result, resultCreateExpected)

	result, err = cli.Run(service, "enable", productID, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, result, resultEnableExpected)

	result, err = cli.Run(service, "disable", productID, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, result, resultDisableExpected)

	result, err = cli.Run(service, "get", productID, productName, productPrice)

	require.Nil(t, err)
	require.Equal(t, result, resultGetExpected)
}
