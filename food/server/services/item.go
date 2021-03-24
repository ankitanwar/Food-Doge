package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	storespb "github.com/ankitanwar/Food-Doge/food/proto"
	"github.com/ankitanwar/Food-Doge/food/server/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func orderFood(storeID, itemID string) (*domain.Items, error) {
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Unable To Get StoreID")
	}
	details := itemDB.GetItemDetail(storeKey, itemID)
	if details == nil {
		return nil, errors.New("Unable To Fetch The Item Details")
	}
	storeDetails := &domain.ListALlProducts{}
	details.Decode(storeDetails)
	for i := 0; i < len(storeDetails.AllProducts); i++ {
		currentProduct := storeDetails.AllProducts[i]
		if currentProduct.ItemID == itemID {
			result := &domain.Items{
				ItemID:   currentProduct.ItemID,
				ItemName: currentProduct.ItemName,
				Cuisine:  currentProduct.Cuisine,
				Price:    currentProduct.Price,
			}
			return result, nil
		}
	}
	return nil, errors.New("Unable To Find The Prodcut To Be Ordered")

}

func (service *FoodService) ViewItemOfStore(req *storespb.ViewParticularStoreRequest, stream storespb.StoresService_ViewItemOfStoreServer) error {
	storeID := req.GetStoreID()
	storeKey, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return errors.New("Please Enter The valid StoreID")
	}
	result := itemDB.FetchAllItems(storeKey)
	if err != nil {
		return errors.New("Unable To Fetch The Details")
	}
	results := &domain.ListALlProducts{}
	result.Decode(results)
	for i := 0; i < len(results.AllProducts); i++ {
		currentItem := results.AllProducts[i]
		details := &storespb.Food{
			ItemName:    currentItem.ItemName,
			Description: currentItem.Description,
			Vegetarian:  currentItem.Vegetarain,
			Price:       currentItem.Price,
			Cuisine:     currentItem.Cuisine,
		}
		response := &storespb.ViewParticularStoreResponse{
			ItemID: currentItem.ItemID,
			Foods:  details,
		}
		stream.Send(response)

	}

	return nil
}

func (service *FoodService) FilterDish(req *storespb.FilterFoodRequest, stream storespb.StoresService_FilterDishServer) error {
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

func (service *FoodService) DeleteItem(ctx context.Context, req *storespb.DeleteItemRequest) (*storespb.DeleteItemResponse, error) {
	storeID := req.GetStoreID()
	itemID := req.GetItemID()
	userID := req.GetUserID()
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

func (service *FoodService) UpdateItemDetail(ctx context.Context, req *storespb.UpdateItemRequest) (*storespb.UpdateItemResponse, error) {
	return nil, nil
}
func (service *FoodService) NewItem(ctx context.Context, req *storespb.AddItemsInStoreRequest) (*storespb.AddItemsInStoreResponse, error) {
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
func (service *FoodService) GetItemDetail(ctx context.Context, req *storespb.GetItemDetailsRequest) (*storespb.GetItemDetailsResponse, error) {
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
	storeDetails := &domain.ListALlProducts{}
	details.Decode(storeDetails)
	for i := 0; i < len(storeDetails.AllProducts); i++ {
		currentItem := storeDetails.AllProducts[i]
		if currentItem.ItemID == itemID {
			dish := &storespb.Food{
				ItemName:    currentItem.ItemName,
				Description: currentItem.Description,
				Vegetarian:  currentItem.Vegetarain,
				Price:       currentItem.Price,
				Cuisine:     currentItem.Cuisine,
			}
			response := &storespb.GetItemDetailsResponse{
				Details: dish,
			}
			return response, nil
		}
	}
	return nil, errors.New("Unable To Fetch The Item Details ")
}

func (service *FoodService) OrderFood(ctx context.Context, req *storespb.OrderFoodRequest) (*storespb.OrderFoodRespomse, error) {
	itemID := req.GetItemID()
	storeID := req.GetStoreID()
	if storeID == "" {
		return nil, errors.New("Invalid Store ID")
	}
	if itemID == "" {
		return nil, errors.New("Please Enter The valid Item ID")
	}
	result, err := orderFood(storeID, itemID)
	if err != nil {
		return nil, err
	}
	deliveryTime := time.Now().Local().Add(time.Minute * time.Duration(40)).String()
	response := storespb.OrderFoodRespomse{
		ItemName:     result.ItemName,
		Price:        result.Price,
		DeliveryTime: deliveryTime,
	}
	return &response, nil
}

func (service *FoodService) Checkout(ctx context.Context, req *storespb.CheckOutRequest) (*storespb.CheckOutResponse, error) {
	itemID := req.GetItemID()
	storeID := req.GetStoreID()
	if storeID == "" {
		return nil, errors.New("Invalid Store ID")
	}
	if itemID == "" {
		return nil, errors.New("Please Enter The valid Item ID")
	}
	result, err := orderFood(storeID, itemID)
	if err != nil {
		return nil, err
	}
	deliveryTime := time.Now().Local().Add(time.Minute * time.Duration(40)).String()
	response := storespb.CheckOutResponse{
		ItemName:     result.ItemName,
		Price:        result.Price,
		DeliveryTime: deliveryTime,
	}
	return &response, nil
}
