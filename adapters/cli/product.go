package cli

import (
	"fmt"
	"github.com/c4st3ll4n/go-hexagon/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		resProduct, err := service.Enable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been enabled",
			resProduct.GetID())

	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		resProduct, err := service.Disable(product)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID %s has been disabled",
			resProduct.GetID())

	default:
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("Product ID: %s * name: %s * price: %f * status: %s",
			res.GetID(), res.GetName(), res.GetPrice(), res.GetStatus())

	}

	return result, nil
}
