package cli_test

import (
	"fmt"
	"github.com/c4st3ll4n/go-hexagon/adapters/cli"
	mock_application "github.com/c4st3ll4n/go-hexagon/application/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productId := "any_id"
	productName := "Product Test"
	productPrice := 66.66
	productStatus := "enabled"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetID().Return(productId).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()

	productService := mock_application.NewMockProductServiceInterface(ctrl)
	productService.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	productService.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productService.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	productService.EXPECT().FindOne(productId).Return(productMock, nil).AnyTimes()

	createdResultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())

	enabledResultExpected := fmt.Sprintf("Product ID %s has been enabled",
		productMock.GetID())

	disabledResultExpected := fmt.Sprintf("Product ID %s has been disabled",
		productMock.GetID())

	defaultResultExpected := fmt.Sprintf("Product ID: %s * name: %s * price: %f * status: %s",
		productMock.GetID(), productMock.GetName(), productMock.GetPrice(), productMock.GetStatus())

	resultCreate, err := cli.Run(productService, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, createdResultExpected, resultCreate)

	resultDisable, err := cli.Run(productService, "disable", "any_id", "", 0)
	require.Nil(t, err)
	require.Equal(t, disabledResultExpected, resultDisable)

	resultEnable, err := cli.Run(productService, "enable", "any_id", "", 0)
	require.Nil(t, err)
	require.Equal(t, enabledResultExpected, resultEnable)

	resultDefault, err := cli.Run(productService, "", "any_id", "", 0)
	require.Nil(t, err)
	require.Equal(t, defaultResultExpected, resultDefault)
}
