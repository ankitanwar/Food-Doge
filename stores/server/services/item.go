package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	storespb "github.com/ankitanwar/Food-Doge/stores/proto"
	"github.com/ankitanwar/Food-Doge/stores/server/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (service *StoreService) ViewItemOfStore(req *storespb.ViewParticularStoreRequest, stream storespb.StoresService_ViewItemOfStoreServer) error {
	storeID := req.GetStoreID()
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("Please Enter The valid StoreID")
	}
	result, err := itemDB.FetchAllItems(storeKey)
	if err != nil {
		return errors.New("Unable To Fetch The Details")
	}
	items := []domain.Items{}
	fmt.Println("The value of result is")
	result.All(context.Background(), items)
	fmt.Println("The value of iems is", items)

	return nil
}

func (service *StoreService) FilterDish(req *storespb.FilterFoodRequest, stream storespb.StoresService_FilterDishServer) error {
	storeID := req.GetStoreID()
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("Please Enter The valid StoreID")
	}
	filterByCusine := "*"
	// filterByVegetarian := false
	filterByName := "*"
	var filterByPrice int64 = 0
	if req.Price > 0 {
		filterByPrice = req.Price
	}
	if req.Cuisine != "" {
		filterByCusine = req.Cuisine
	}
	// if req.Vegetarian == true {
	// 	filterByVegetarian = true
	// }
	if req.Name != "" {
		filterByName = req.Name
	}
	details, err := itemDB.FilterItems(storeKey, filterByPrice, filterByCusine, "", filterByName)
	if err != nil {
		log.Println("Unable To Filter The details", err)
		return errors.New("Unable To Filter The Products")
	}
	items := []domain.Items{}
	details.All(context.Background(), items)
	for i := 0; i < len(items); i++ {
		current := items[i]
		fmt.Println("the value of current is", current)
	}
	return nil
}

func (service *StoreService) DeleteItem(ctx context.Context, req *storespb.DeleteItemRequest) (*storespb.DeleteItemResponse, error) {
	storeID := req.GetStoreID()
	fmt.Println("The value of store id is", storeID)
	itemID := req.GetItemID()
	fmt.Println("The value of item ID is", itemID)
	userID := req.GetUserID()
	fmt.Println("the value of user Id is ", userID)
	fmt.Println("The value of item id is ", itemID)
	StoreKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Please Enter The Valid Store ID")
	}
	err = itemDB.DeleteItem(StoreKey, userID, itemID)
	if err != nil {
		return nil, errors.New("Error While Deleting The item")
	}
	response := &storespb.DeleteItemResponse{
		Message: "Item Has Been Deleted Successfully",
	}
	return response, nil

}

func (service *StoreService) UpdateItemDetail(ctx context.Context, req *storespb.UpdateItemRequest) (*storespb.UpdateItemResponse, error) {
	return nil, nil
}
func (service *StoreService) NewItem(ctx context.Context, req *storespb.AddItemsInStoreRequest) (*storespb.AddItemsInStoreResponse, error) {
	userID := req.GetUserID()
	storeID := req.GetStoreID()
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Unable To Get StoreID")
	}
	itemDetails := &domain.Items{
		ItemName:    req.Details.GetItemName(),
		Description: req.Details.GetDescription(),
		Vegetarain:  req.Details.GetVegetarian(),
		Price:       req.Details.GetPrice(),
		Cuisine:     req.Details.GetCuisine(),
	}
	err = itemDetails.CheckItemDetails()
	if err != nil {
		return nil, err
	}
	itemDetails.ItemID, err = itemDetails.GenerateUniqueID()
	if err != nil {
		return nil, errors.New("Unable To Add Item Into The StoreS")
	}
	err = itemDB.AddItemToStore(storeKey, userID, itemDetails)
	if err != nil {
		return nil, errors.New("Unable To Add Item Into The store")
	}
	response := &storespb.AddItemsInStoreResponse{
		Message: "Item Has Been Added Successfully",
	}
	return response, nil

}
func (service *StoreService) GetItemDetail(ctx context.Context, req *storespb.GetItemDetailsRequest) (*storespb.GetItemDetailsResponse, error) {
	storeID := req.GetStoreID()
	itemID := req.GetItemID()
	if storeID == "" {
		return nil, errors.New("Invalid Store ID")
	}
	if req.ItemID == "" {
		return nil, errors.New("Invalid Item ID")
	}
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Unable To Get StoreID")
	}
	details := itemDB.GetItemDetail(storeKey, itemID)
	if details == nil {
		return nil, errors.New("Unable To Fetch The Item Details")
	}
	result := &domain.Items{}
	err = details.Decode(result)
	if err != nil {
		return nil, errors.New("Error While Decoding The Result")
	}
	response := &storespb.GetItemDetailsResponse{}
	response.Details.ItemName = result.ItemName
	response.Details.Description = result.Description
	response.Details.Vegetarian = result.Vegetarain
	response.Details.Price = result.Price
	return response, nil
}

func (service *StoreService) OrderFood(ctx context.Context, req *storespb.OrderFoodRequest) (*storespb.OrderFoodRespomse, error) {
	storeID := req.GetStoreID()
	itemID := req.GetItemID()
	if storeID == "" {
		return nil, errors.New("Invalid Store ID")
	}
	if itemID == "" {
		return nil, errors.New("Please Enter The valid Item ID")
	}
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Invalid Store ID")
	}
	details := itemDB.GetItemDetail(storeKey, itemID)
	if details == nil {
		return nil, errors.New("Unable To Fetch The Given Item")
	}
	item := &domain.Items{}
	details.Decode(item)
	deliveryTime := time.Now().Local().Add(time.Minute * time.Duration(40)).String()
	response := &storespb.OrderFoodRespomse{
		DeliveryTime: deliveryTime,
		Price:        item.Price,
		ItemName:     item.ItemName,
	}
	return response, nil
}
