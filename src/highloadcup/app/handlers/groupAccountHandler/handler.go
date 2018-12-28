package groupAccountHandler

import (
	"encoding/json"
	"highloadcup/app/services/groupAccountService"
	"highloadcup/app/slices"
	"highloadcup/app/structs"
	"highloadcup/app/urlquery"
	"net/http"
	"strings"
)

var AccountList *[]structs.Account

var allowedParams = []string{
	"query_id",
	"sex",
	"email",
	"status",
	"fname",
	"sname",
	"phone",
	"country",
	"city",
	"birth",
	"interests",
	"likes",
	"premium",
	"limit",
	"order",
	"joined",
}

var requiredParams = []string{
	"keys",
	"limit",
	"order",
}

var allowedKeys = []string{
	"sex",
	"status",
	"interests",
	"country",
	"city",
}

func Handle(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err := urlquery.Parse(request.Form, allowedParams, requiredParams)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	keys := strings.Split(query.Params["keys"]["eq"], ",")
	if !slices.ContainsAll(allowedKeys, keys) {
		http.Error(writer, "Keys has now allowed values", http.StatusBadRequest)
		return
	}

	groups := groupAccountService.Group(AccountList, query)

	data, err := json.Marshal(struct {
		Groups interface{} `json:"groups"`
	}{groups})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Write(data)
}
