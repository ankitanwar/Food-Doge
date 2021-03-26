package services

import (
	"context"
	"errors"
	"fmt"

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
	storeID := res.InsertedID.(primitive.ObjectID)
	err = AddToLocation(req.GetState(), storeID.Hex(), req.GetStoreName(), req.GetDescription(), req.GetStreet())
	if err != nil {
		storeDB.DeleteStore(storeID)
		return nil, err
	}
	response := &foodpb.CreateStoreResponse{
		StoreID:    storeID.Hex(),
		AddedStore: req,
	}
	return response, nil

}
func (service *FoodService) Explore(req *foodpb.ExploreOutletsRequest, stream foodpb.StoresService_ExploreServer) error {
	state := req.GetAddress()
	location := getLocationKey(state)
	details := storeDB.ExploreStore(location)
	stores := &domain.StoreExplore{}
	err := details.Decode(stores)
	if err != nil {
		return errors.New("Unable To Search For The stores")
	}
	if err != nil {
		fmt.Println("error occured here")
		return errors.New("Error While Unmarshalling the store Details")
	}
	for i := 0; i < len(stores.StoreLocationInformation); i++ {
		currentStore := stores.StoreLocationInformation[i]
		response := &foodpb.ExporeOutletsResponse{
			StoreName:   currentStore.Name,
			Street:      currentStore.Street,
			State:       currentStore.State,
			Description: currentStore.Description,
			StoreID:     currentStore.StoreID,
		}
		stream.Send(response)
	}
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

func AddToLocation(state, storeID, name, description, street string) error {
	location := getLocationKey(state)
	details := domain.StoreLocationInformation{
		StoreID:     storeID,
		Name:        name,
		Description: description,
		State:       state,
		Street:      street,
	}
	locationDetails := &domain.StoreLocation{
		Location:                 location,
		StoreLocationInformation: details,
	}
	err := storeDB.AddStoreInLocationSearch(locationDetails)
	if err != nil {
		return errors.New("Unable To Add Store In The Given Location")
	}
	return nil
}

func getLocationKey(state string) string {
	lat, long := GetGeoLocation(state)
	location := ""
	location += fmt.Sprintf("%f", lat)
	location += "-"
	location += fmt.Sprintf("%f", long)
	return location
}
