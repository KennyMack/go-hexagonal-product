package cli

import (
	"fmt"
	"github.com/kennymack/go-hexagonal-product/application"
)

func Run(service application.ProductServiceInterface, action string,
	productId string, productName string,
	productPrice float64) (string, error) {
	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, productPrice)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("ID %s Name %s has been created", product.GetId(), product.GetName())
	case "enable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Enable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("ID %s Name %s has been enabled", product.GetId(), product.GetName())
	case "disable":
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		product, err = service.Disable(product)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("ID %s Name %s has been disabled", product.GetId(), product.GetName())
	default:
		product, err := service.Get(productId)
		if err != nil {
			return result, err
		}

		result = fmt.Sprintf("ID:%s|NAME:%s|PRICE:%v|STATUS:%s", product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())
	}

	return result, nil
}
