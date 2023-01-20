package main

import (
	"SE_MIM22_WEBSHOP_ORDERSERVICE/handler"
	"log"
	"net/http"
	"time"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/placeOrder", handler.PlaceOrder)
	serveMux.HandleFunc("/getOrdersByUserId", handler.GetOrdersByUserID)
	log.Printf("\n\n\tORDERSERVICE\n\nAbout to listen on Port: 8460." +
		"\n\nSUPPORTED REQUESTS:" +
		"\nGET:" +
		"\nGet Order By ID: http://127.0.0.1:8460/getOrdersByUserId?id=1 requiers a url parameter id" +
		"\nPOST:" +
		"\nPlace Order: http://127.0.0.1:8460/placeOrder requiers a Body with following json:\n{\n    \"produktId\": \"1\",\n    \"userId\": \"1\",\n    \"amount\": \"1\"\n}")
	server := &http.Server{
		Addr:              ":8460",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
