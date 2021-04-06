## Food Service
<br> </br>

ALERT : Since the frontend is not there I have integrated grpc with the REST API To Take The Input which is not the correct way of doing so

<br> </br>

<b>FOOD Service</b>

| Feature  | Description  |
|----------|:-------------|
| Register Store | To Register Your Store With Us |
| Get Nearby Restaurants | To Explore The Near By Restaurants  |
| Filter Food Item | Filter Food Items According To Your Choices |
| Order Food Item | To Order Food Item  |
| Add Item Into The Store | Allowing Restaurants Owner Can Add More Food Item Into The Store |
| Delete Food Item From The Store | Allowing Restaurants Owner To Delete Food Item From The Store |
| Get Every Item Of The Store | To Explore All The Food Items Of The Particular Store |
| Get Specific Food Item Of The Store | To View The Specific Food Item Of the Store|

<br></br>

<b>End Points</b>

| Request  | Description  | Url |
|----------|:-------------|:-------------|
| Post | To Register The New Store | host:8070/stores/newstore |
| Get | To Get All The Stores In The Given Location |host:8070/stores/:location |
| Get | To Filter Food Items |host:8070/store/:storeID/filter |
| Post | To Order The Food Item |host:8070/buy/food/:storeID/:itemID/:addressID |
| Post | To Add The Food Item Into The Store |host:8070/food/:storeID |
| Delete | To Delete The Food Item From The Store |host:8070/:storeID/:itemID |
| Get | To Get The Every Food Item In The Store |host:8070/food/all/:storeID |
| Get | To Get The Specific Food Item Detail Of The Store |host:8070/itemDetail/:storeID/:itemID |






<br></br>
