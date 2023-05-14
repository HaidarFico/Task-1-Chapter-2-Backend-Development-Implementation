package main

import (
	"log"
	"minio/controllers"
	"minio/database"
	"minio/entity"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/users", controllers.RegisterUser).Methods("POST")
	router.HandleFunc("/allusers", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	router.HandleFunc("/product/{id}", controllers.GetProductById).Methods("GET")
	router.HandleFunc("/products/search", controllers.SearchProducts).Methods("GET")
	router.HandleFunc("/products/addproducts", controllers.AddProducts).Methods("POST")
	router.HandleFunc("/products/deleteproduct", controllers.DeleteProduct).Methods("POST")
	// router.HandleFunc("/shoppingcart", controllers.GetShoppingCart).Methods("GET")
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "minio",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Connector.AutoMigrate(&entity.User{}, &entity.Product{})

}
