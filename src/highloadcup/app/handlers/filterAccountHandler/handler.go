package filterAccountHandler

import (
	"encoding/json"
	"highloadcup/app/services/filterAccountService"
	"highloadcup/app/structs"
	"highloadcup/app/urlquery"
	"net/http"
)

var AccountList *[]structs.Account

var allowedParams = []string{
	"query_id",
	"sex_eq",
	"email_domain",
	"email_lt",
	"email_gt",
	"status_eq",
	"status_neq",
	"fname_eq",
	"fname_any",
	"fname_null",
	"sname_eq",
	"sname_starts",
	"sname_null",
	"phone_code",
	"phone_null",
	"country_eq",
	"country_null",
	"city_eq",
	"city_any",
	"city_null",
	"birth_lt",
	"birth_gt",
	"birth_year",
	"interests_contains",
	"interests_any",
	"likes_contains",
	"premium_now",
	"premium_null",
}

var requiredParams = []string{"limit"}

func Handle(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err := urlquery.Parse(request.Form, allowedParams, requiredParams)
	if err != nil {
		http.Error(writer, "", http.StatusBadRequest)
		return
	}

	accounts := filterAccountService.Filter(AccountList, query)

	data, err := json.Marshal(struct {
		Accounts interface{} `json:"accounts"`
	}{accounts})

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Write(data)
}
