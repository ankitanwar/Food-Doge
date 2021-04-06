## Store Service

This Services Is To Help The Store To Manage Orders

<b>Store Services</b>

| Feature  | Description  |
|----------|:-------------|
| View Orders| To View All The Placed Orders |
| Receive Orders | To Receive The Orders By The Customers |
| Order Completed | To Delete The Order Once Completed  |
<br></br>

<b>End Points</b>

| Request  | Description  | Url |
|----------|:-------------|:-------------|
| Post | To Place The Order With The Store | host:9080/orders/:storeID |
| Get | To View All The Place Orders |host:9080/orders/store/:storeID |
| Delete | To Delete The Completed Order |host:9080/orders/store/:storeID/:itemID |

<br></br>