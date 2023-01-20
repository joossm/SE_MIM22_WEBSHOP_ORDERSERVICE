package handler

import (
	"SE_MIM22_WEBSHOP_ORDERSERVICE/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

const post = "POST"

func PlaceOrder(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case post:
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			order := model.Order{}
			jsonErr := json.Unmarshal(body, &order)
			if jsonErr != nil {
				_, responserErr := responseWriter.Write([]byte("{ERROR}"))
				errorHandler(responserErr)
				return
			}
			db := openDB()
			defer closeDB(db)
			_, err := db.Query("INSERT INTO orders (produktId, userId, Amount) VALUES (?, ?, ?)",
				order.ProduktId, order.UserId, order.Amount)
			errorHandler(err)
			_, responserErr := responseWriter.Write([]byte("{true}"))
			errorHandler(responserErr)

			return
		}
	default:
		_, responserErr := responseWriter.Write([]byte("THIS IS A POST REQUEST"))
		errorHandler(responserErr)
		return
	}
}
func GetOrdersByUserID(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT produktId, userId, amount FROM orders WHERE userId = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var orders []model.Order
		for result.Next() {
			var order model.Order
			err = result.Scan(&order.ProduktId, &order.UserId, &order.Amount)
			errorHandler(err)
			orders = append(orders, order)
		}
		jsonOrders, err := json.Marshal(orders)
		errorHandler(err)
		_, responserErr := responseWriter.Write(jsonOrders)
		errorHandler(responserErr)
		return
	default:
		_, responserErr := responseWriter.Write([]byte("THIS IS A GET REQUEST"))
		errorHandler(responserErr)
		return
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		print(err)
	}
}
