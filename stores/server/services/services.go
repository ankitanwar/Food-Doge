package services

import (
	"context"
	"errors"

	storespb "github.com/ankitanwar/Food-Doge/stores/proto"
	db "github.com/ankitanwar/Food-Doge/stores/server/database"
	"github.com/ankitanwar/Food-Doge/stores/server/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	AvailableServices storespb.StoresServiceServer = &StoreService{}
	storeDB                                        = &db.StoreDB{}
	itemDB                                         = &db.ItemsDB{}
)

type StoreService struct {
}

func (s *StoreService) CreateStore(ctx context.Context, req *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
	storeDetails := &domain.Store{
		StoreName:   req.StoreName,
		State:       req.State,
		Street:      req.Street,
		Country:     req.Country,
		PhoneNumber: req.PhoneNumber,
		Description: req.Description,
		Pincode:     req.Pincode,
		StoreOwner:  req.UserID,
	}
	err := storeDetails.CheckForError()
	if err != nil {
		return nil, err
	}
	res, err := storeDB.SaveStore(storeDetails)
	if err != nil {
		return nil, errors.New("Unable To Add The Store Into The Database")
	}
	storeID := res.InsertedID.(primitive.ObjectID).Hex()
	response := &storespb.CreateStoreResponse{
		StoreID:    storeID,
		AddedStore: req,
	}
	return response, nil

}
func (s *StoreService) Explore(req *storespb.ExploreOutletsRequest, stream storespb.StoresService_ExploreServer) error {
	return nil
}
func (s *StoreService) OrderFood(ctx context.Context, req *storespb.OrderFoodRequest) (*storespb.OrderFoodRespomse, error) {
	return nil, nil
}
func (s *StoreService) UpdateStoreDetails(ctx context.Context, req *storespb.UpdateStoreRequest) (*storespb.UpdateStoreResponse, error) {
	return nil, nil
}
func (s *StoreService) UpdateItemDetail(ctx context.Context, req *storespb.UpdateItemRequest) (*storespb.UpdateItemResponse, error) {
	return nil, nil
}
func (s *StoreService) Filter(req *storespb.FilterOutletsRequest, stream storespb.StoresService_FilterServer) error {
	return nil
}
func (s *StoreService) NewItem(ctx context.Context, req *storespb.AddItemsInStoreRequest) (*storespb.AddItemsInStoreResponse, error) {
	userID := req.UserID
	id := req.StoreID
	storeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("Unable To Get StoreID")
	}
	itemDetails := &domain.Items{
		ItemName:    req.ItemName,
		Description: req.Description,
		Vegetarain:  req.Vegetarian,
		Price:       req.Price,
	}
	err = itemDetails.CheckItemDetails()
	if err != nil {
		return nil, err
	}
	itemDetails.ItemID, err = itemDetails.GenerateUniqueID()
	if err != nil {
		return nil, errors.New("Unable To Add Item Into The StoreS")
	}
	err = itemDB.AddItemToStore(storeID, userID, itemDetails)
	if err != nil {
		return nil, errors.New("Unable To Add Item Into The store")
	}
	response := &storespb.AddItemsInStoreResponse{
		Message: "Item Has Been Added Successfully",
	}
	return response, nil

}
func (s *StoreService) DeleteItem(ctx context.Context, req *storespb.DeleteItemRequest) (*storespb.DeleteItemResponse, error) {
	userID := req.UserID
	itemID := req.ItemID
	storeID := req.StoreID
	PstoreID, err := primitive.ObjectIDFromHex(storeID)
	if err != nil {
		return nil, errors.New("Please Enter The Valid Store ID")
	}
	err = itemDB.DeleteItem(PstoreID, userID, itemID)
	if err != nil {
		return nil, errors.New("Error While Deleting The item")
	}
	response := &storespb.DeleteItemResponse{
		Message: "Item Has Been Deleted Successfully",
	}
	return response, nil

}
func (s *StoreService) DeleteStore(ctx context.Context, req *storespb.DeleteStoreRequest) (*storespb.DeleteStoreResponse, error) {
	givenStoreID := req.StoreID
	storeID, err := primitive.ObjectIDFromHex(givenStoreID)
	if err != nil {
		return nil, errors.New("Invalid storeID")
	}
	err = storeDB.DeleteStore(storeID)
	if err != nil {
		return nil, errors.New("Unable To Delete The Store")
	}
	response := &storespb.DeleteStoreResponse{
		Message: "Store Has Been Removed Successfullly",
	}
	return response, nil
}
