package order

import (
	"fmt"
	"net/http"
	"time"

	foodpb "github.com/ankitanwar/Food-Doge/food/proto"
	"github.com/ankitanwar/GoAPIUtils/errors"
	"github.com/mercadolibre/golang-restclient/rest"
)

const (
	headerXCallerID      = "X-Caller-ID"
	headerXAccessTokenID = "X-Token-ID"
)

var (
	paramAccessToken = headerXAccessTokenID
	headers          = make(http.Header)
	oauthRestClient  = rest.RequestBuilder{
		BaseURL: "http://localhost:9080",
		Timeout: 200 * time.Millisecond,
		Headers: headers,
	}
)

//GetCallerID : To get the caller id from the url
func GetCallerID(request *http.Request) string {
	if request == nil {
		return ""
	}
	callerID := request.Header.Get(headerXCallerID)
	return callerID
}

//GetAccessID: To get the caller id from the url
func GetAccessID(request *http.Request) string {
	if request == nil {
		return ""
	}
	accessID := request.Header.Get(headerXAccessTokenID)
	return accessID
}

type PlaceOrder struct {
	HouseNumber string `json:"houseNo" json:"houseNumber"`
	Street      string `json:"street" json:"street"`
	State       string `json:"state" json:"state"`
	Phone       string `json:"phone" json:"phone"`
	Items       Order  `json:"order" bson:"order"`
}

type Order struct {
	ItemName string `json:"itemName" bson:"itemName"`
	Price    int64  `json:"price" bson:"price"`
}

func PlaceOrders(req *http.Request, storeID string, order *foodpb.OrderFoodRespomse) *errors.RestError {
	orderedItem := Order{
		ItemName: order.ItemName,
		Price:    order.Price,
	}
	details := &PlaceOrder{
		HouseNumber: order.CustomerHouseNo,
		Street:      order.CustomerStreet,
		State:       order.CustomerStreet,
		Phone:       order.CustomerPhoneNumber,
		Items:       orderedItem,
	}
	err := PlaceOrderWithStore(req, details, storeID)
	if err != nil {
		return err
	}
	return nil
}

func PlaceOrderWithStore(req *http.Request, order *PlaceOrder, storeID string) *errors.RestError {
	userID := GetCallerID(req)
	tokenID := GetAccessID(req)
	headers.Set(headerXCallerID, userID)
	headers.Set(headerXAccessTokenID, tokenID)
	response := oauthRestClient.Post(fmt.Sprintf("/orders/%s", storeID), order)
	fmt.Println("The value of response is", response.String())
	if response == nil || response.Response == nil {
		errors.NewNotFound("Not found")
	}

	if response.StatusCode > 299 {
		return errors.NewInternalServerError("Unable To Place The Order")
	}
	return nil

}
