package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/ekowdd89/go-crud-mysql/config"
	"github.com/ekowdd89/go-crud-mysql/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	connection := &config.Database{}
	con, err := connection.Db()
	if err != nil {
		log.Println(err)
	}
	router := mux.NewRouter()
	user := handler.NewUser(*con)
	category := handler.NewCategory(*con)
	article := handler.NewArticle(*con)
	/*
	* User routes
	 */
	router.HandleFunc("/users", user.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", user.FindUser).Methods("GET")
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", user.DeleteUser).Methods("DELETE")

	/*
	* category router handler
	 */

	router.HandleFunc("/categorys", category.GetCategory).Methods("GET")
	router.HandleFunc("/categorys/{id}", category.FindCategory).Methods("GET")
	router.HandleFunc("/categorys", category.CreateCategory).Methods("POST")
	router.HandleFunc("/categorys/{id}", category.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categorys/{id}", category.DeleteCategory).Methods("DELETE")

	router.HandleFunc("/articles", article.GetArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", article.FindArticle).Methods("GET")
	router.HandleFunc("/articles", category.CreateCategory).Methods("POST")
	router.HandleFunc("/articles/{id}", category.UpdateCategory).Methods("PUT")
	router.HandleFunc("/articles/{id}", category.DeleteCategory).Methods("DELETE")

	var address = "localhost:8081"
	log.Println("Server running on :", strings.Join([]string{"http://", address}, ""))

	server := &http.Server{
		Handler: router,
		Addr:    address,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
