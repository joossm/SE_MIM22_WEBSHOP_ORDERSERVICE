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

func PlaceOrder(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			order := model.Order{}
			jsonErr := json.Unmarshal(body, &order)
			if jsonErr != nil {
				js, err := json.Marshal("Error")
				errorHandler(err)
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
			db := openDB()
			defer closeDB(db)
			_, insertErr := db.Query("INSERT INTO orders (produktId, userId, Amount) VALUES (?, ?, ?)",
				order.ProduktId, order.UserId, order.Amount)
			errorHandler(insertErr)
			js, err := json.Marshal("true")
			errorHandler(err)
			_, responseErr := responseWriter.Write(js)
			errorHandler(responseErr)
			return
		}
	default:
		js, err := json.Marshal("THIS IS A POST REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func GetOrdersByUserId(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT * FROM orders WHERE userId = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var orders []model.Order
		if result != nil {
			for result.Next() {
				var order model.Order
				err = result.Scan(&order.Id, &order.ProduktId, &order.UserId, &order.Amount)
				errorHandler(err)
				orders = append(orders, order)
			}
		}
		var orderResults []model.OrderResult
		for _, order := range orders {
			var orderResultItem model.OrderResult
			var bookAndAmount model.BookAndAmount
			bookAndAmount.Amount = order.Amount
			result, err := db.Query("SELECT * FROM books WHERE Id = ?", order.ProduktId)
			errorHandler(err)
			if result != nil {
				for result.Next() {
					var book model.Book
					err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
					errorHandler(err)
					bookAndAmount.Book = book
					bookAndAmount.Amount = order.Amount
				}
			}
			orderResultItem.BasketID = order.Id
			orderResultItem.UserId = order.UserId
			orderResultItem.Books = append(orderResultItem.Books, bookAndAmount)
			orderResults = append(orderResults, orderResultItem)
		}
		orderResultJson, jsonErr := json.Marshal(orderResults)
		errorHandler(jsonErr)
		_, responseErr := responseWriter.Write(orderResultJson)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
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
