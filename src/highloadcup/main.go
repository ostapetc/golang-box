package main

import (
	"fmt"
	"highloadcup/app/handlers/filterAccountHandler"
	"highloadcup/app/handlers/groupAccountHandler"
	"highloadcup/app/importer"
	"highloadcup/app/structs"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var AccountList []structs.Account

func groupAccountsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("GroupAccountsHandler!"))
}

func accountsRecommendationsHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("AccountsRecommendationsHandler!"))
}

func accountsSuggestHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("AccountsSuggestHandler!"))
}

func createAccountHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Println(vars)
	writer.Write([]byte("CreateAccountHandler!"))
}

func updateAccountHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Println(vars)
	writer.Write([]byte("UpdateAccountHandler!"))
}

func createLikeHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Println(vars)
	writer.Write([]byte("CreateLikeHandler!"))
}

func main() {
	AccountList = importer.ImportAccounts()

	filterAccountHandler.AccountList = &AccountList
	groupAccountHandler.AccountList = &AccountList

	router := mux.NewRouter()

	router.HandleFunc("/accounts/filter/", filterAccountHandler.Handle)
	router.HandleFunc("/accounts/group/", groupAccountHandler.Handle)
	// router.HandleFunc("/accounts/{id:[0-9]+}/recommend/", handlers.AccountRecommendations)
	// router.HandleFunc("/accounts/{id:[0-9]+}/suggest/", handlers.AccountSuggest)
	// router.HandleFunc("/accounts/new/", handlers.CreateAccount).Methods("post")
	// router.HandleFunc("/accounts/{id:[0-9]+}/", handlers.UpdateAccount).Methods("post")
	// router.HandleFunc("/accounts/likes/", handlers.CreateLike).Methods("post")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
