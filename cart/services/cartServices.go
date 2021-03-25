package services

import (
	"fmt"
	"net/http"
	"time"

	cartdatabase "github.com/ankitanwar/Food-Doge/cart/database"
	domain "github.com/ankitanwar/Food-Doge/cart/domain"
	product "github.com/ankitanwar/Food-Doge/middleware/Products"
	"github.com/ankitanwar/GoAPIUtils/errors"
)

//AddToCart : To add the given product details into the cart of the given user
func AddToCart(userID, storeID, itemID string) *errors.RestError {
	getItemDetails, err := product.ItemSerivce.GetItemDetails(storeID, itemID)
	if err != nil {
		return err
	}
	fmt.Println(getItemDetails)
	details := getItemDetails.Detail
	details.ItemID = itemID
	details.StoreID = storeID
	err = cartdatabase.AddToCart(userID, details)
	if err != nil {
		return err
	}
	return nil

}

//RemoveFromCart : To remove the given item from the cart
func RemoveFromCart(userID, itemID string) *errors.RestError {
	removeErr := cartdatabase.RemoveFromCart(userID, itemID)
	if removeErr != nil {
		return errors.NewInternalServerError("Error while removing the item from the cart")
	}
	return nil
}

//ViewCart : To view all the items in the cart
func ViewCart(userID string) (*domain.Details, *errors.RestError) {
	userCart, err := cartdatabase.ViewCart(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Error while fetching the cart")
	}
	return userCart, nil
}

//Checkout : To checkout all the given items in the cart
func Checkout(req *http.Request, userID string) (*domain.CheckoutCart, *errors.RestError) {
	userCart, err := cartdatabase.ViewCart(userID)
	if err != nil {
		return nil, errors.NewInternalServerError("Error while fetching the cart")
	}
	response := &domain.CheckoutCart{}
	deliveryTime := time.Now()
	deliveryTime.Format("01-02-2006")
	deliveryTime = deliveryTime.AddDate(0, 0, 10)
	response.DeliveryTime = deliveryTime.String()
	for i := 0; i < len(userCart.Detail); i++ {
		currentItem := userCart.Detail[i]
		err := product.ItemSerivce.BuyItem(req, currentItem.StoreID, currentItem.ItemID)
		if err != nil {
			return nil, err
		}
		response.TotalPrice += currentItem.Price
		response.Items = append(response.Items, currentItem)
		cartdatabase.RemoveFromCart(userID, currentItem.ItemID)
	}
	return response, nil
}
