package services

import (
	"context"
	"errors"

	foodpb "github.com/ankitanwar/Food-Doge/food/proto"
	db "github.com/ankitanwar/Food-Doge/food/server/database"
	"github.com/ankitanwar/Food-Doge/food/server/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	AvailableServices foodpb.StoresServiceServer = &FoodService{}
	storeDB                                      = &db.StoreDB{}
	itemDB                                       = &db.ItemsDB{}
)

type FoodService struct {
}

func (service *FoodService) CreateStore(ctx context.Context, req *foodpb.CreateStoreRequest) (*foodpb.CreateStoreResponse, error) {
	products := []domain.Items{}
	storeDetails := &domain.Store{
		StoreName:   req.GetStoreName(),
		State:       req.GetState(),
		Street:      req.GetStreet(),
		Country:     req.GetCountry(),
		PhoneNumber: req.GetPhoneNumber(),
		Description: req.GetDescription(),
		Pincode:     req.GetPincode(),
		StoreOwner:  req.GetUserID(),
		Products:    products,
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
	response := &foodpb.CreateStoreResponse{
		StoreID:    storeID,
		AddedStore: req,
	}
	return response, nil

}
func (service *FoodService) Explore(req *foodpb.ExploreOutletsRequest, stream foodpb.StoresService_ExploreServer) error {
	return nil
}

func (service *FoodService) UpdateStoreDetails(ctx context.Context, req *foodpb.UpdateStoreRequest) (*foodpb.UpdateStoreResponse, error) {
	return nil, nil
}

func (service *FoodService) DeleteStore(ctx context.Context, req *foodpb.DeleteStoreRequest) (*foodpb.DeleteStoreResponse, error) {
	givenStoreID := req.GetStoreID()
	storeID, err := primitive.ObjectIDFromHex(givenStoreID)
	if err != nil {
		return nil, errors.New("Invalid storeID")
	}
	err = storeDB.DeleteStore(storeID)
	if err != nil {
		return nil, errors.New("Unable To Delete The Store")
	}
	response := &foodpb.DeleteStoreResponse{
		Message: "Store Has Been Removed Successfullly",
	}
	return response, nil
}
