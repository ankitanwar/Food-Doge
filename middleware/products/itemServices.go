package product

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

var (
	restClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8070",
		Timeout: 100 * time.Millisecond,
	}
	ItemSerivce itemServiceInterface = &itemServicesStruct{}
)

type itemServiceInterface interface {
	GetItemDetails(string, string) (*Details, *errors.RestError)
	BuyItem(*http.Request, string, string) *errors.RestError
}

type itemServicesStruct struct {
}

func (item *itemServicesStruct) GetItemDetails(storeID, itemID string) (*Details, *errors.RestError) {
	res := restClient.Get(fmt.Sprintf("/itemDetail/%s/%s", storeID, itemID))
	if res.Response == nil || res == nil {
		return nil, errors.NewInternalServerError("Error while fetching the item details")
	}
	product := &Details{}
	if res.StatusCode < 299 {
		err := json.Unmarshal(res.Bytes(), &product)
		if err != nil {
			fmt.Println(err)
			return nil, errors.NewInternalServerError("Error while unmarshalling the data")
		}
		return product, nil
	}
	return nil, errors.NewInternalServerError("Error while getting the items details")
}

func (item *itemServicesStruct) BuyItem(req *http.Request, storeID, itemID string) *errors.RestError {
	res := restClient.Post(fmt.Sprintf("/checkout/%s/%s", storeID, itemID), nil)
	if res.StatusCode < 299 {
		details := &Order{}
		fmt.Println("The value of details is", res.String())
		err := json.Unmarshal(res.Bytes(), details)
		fmt.Println("The value of err is", err)
		if err != nil {
			return errors.NewInternalServerError("Error While Unmarshalling The Data")
		}
		return nil
	}
	return errors.NewBadRequest("Unable to purchase items")
}
