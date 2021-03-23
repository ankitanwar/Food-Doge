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

func (service *StoreService) CreateStore(ctx context.Context, req *storespb.CreateStoreRequest) (*storespb.CreateStoreResponse, error) {
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
	response := &storespb.CreateStoreResponse{
		StoreID:    storeID,
		AddedStore: req,
	}
	return response, nil

}
func (service *StoreService) Explore(req *storespb.ExploreOutletsRequest, stream storespb.StoresService_ExploreServer) error {
	return nil
}

func (service *StoreService) UpdateStoreDetails(ctx context.Context, req *storespb.UpdateStoreRequest) (*storespb.UpdateStoreResponse, error) {
	return nil, nil
}

func (service *StoreService) DeleteStore(ctx context.Context, req *storespb.DeleteStoreRequest) (*storespb.DeleteStoreResponse, error) {
	givenStoreID := req.GetStoreID()
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
