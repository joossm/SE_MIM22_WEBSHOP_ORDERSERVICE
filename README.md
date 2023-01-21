# ORDERSERVICE

2023-01-:12:28.117182467Z About to listen on Port: 8460.

SUPPORTED REQUESTS:

GET:

Get Order By ID: http://127.0.0.1:8460/getOrdersByUserId?id=1 requiers a url parameter id

POST:
Place Order: http://127.0.0.1:8460/placeOrder requiers a Body with following json:
{
"produktId": "1",
"userId": "1",
"amount": "1"
}
