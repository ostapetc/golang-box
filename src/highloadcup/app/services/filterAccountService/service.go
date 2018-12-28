package filterAccountService

import (
	"fmt"
	"highloadcup/app/slices"
	"highloadcup/app/structs"
	"log"
	"strconv"
	"strings"
	"time"
)

// FilterAccount filters accounts by query params
func Filter(accountsIndex *[]structs.Account, query *structs.ServiceQuery) *[]map[string]interface{} {
	var result []map[string]interface{}
	var size int32

	for _, account := range *accountsIndex {
		if filter(&account, query) {
			result = append(result, makeAccountRow(&account, query))
			size++
		}

		if query.Limit != 0 && size == query.Limit {
			break
		}
	}

	return &result
}

func makeAccountRow(account *structs.Account, query *structs.ServiceQuery) map[string]interface{} {
	obj := map[string]interface{}{}
	obj["id"] = account.ID
	obj["email"] = account.Email

	if _, ok := query.Params["sex"]; ok {
		obj["sex"] = account.Sex
	}

	if _, ok := query.Params["status"]; ok {
		obj["status"] = account.Status
	}

	if _, ok := query.Params["fname"]; ok {
		obj["fname"] = account.Fname
	}

	if _, ok := query.Params["sname"]; ok {
		obj["sname"] = account.Sname
	}

	if _, ok := query.Params["phone"]; ok {
		obj["phone"] = account.Phone
	}

	if _, ok := query.Params["country"]; ok {
		obj["country"] = account.Country
	}

	if _, ok := query.Params["city"]; ok {
		obj["city"] = account.City
	}

	if _, ok := query.Params["birth"]; ok {
		obj["birth"] = account.Birth
	}

	if _, ok := query.Params["interests"]; ok {
		obj["interests"] = account.Interests
	}

	if _, ok := query.Params["likes"]; ok {
		obj["likes"] = account.Likes
	}

	if _, ok := query.Params["premium"]; ok {
		obj["premium"] = account.Premium
	}

	return obj
}

func filter(account *structs.Account, query *structs.ServiceQuery) bool {
	// sex filter
	if val, ok := query.Params["sex"]["eq"]; ok {
		if account.Sex != val {
			return false
		}
	}

	// email filter
	if val, ok := query.Params["email"]["domain"]; ok {
		if !strings.HasSuffix(account.Email, val) {
			return false
		}
	}

	if val, ok := query.Params["email"]["lt"]; ok {
		if account.Email >= val {
			return false
		}
	}

	if val, ok := query.Params["email"]["gt"]; ok {
		if account.Email <= val {
			return false
		}
	}

	// status filter
	if val, ok := query.Params["status"]["eq"]; ok {
		if account.Status != val {
			return false
		}
	}

	if val, ok := query.Params["status"]["neq"]; ok {
		if account.Status == val {
			return false
		}
	}

	// fname filter
	if val, ok := query.Params["fname"]["eq"]; ok {
		if account.Fname != val {
			return false
		}
	}

	if val, ok := query.Params["fname"]["any"]; ok {
		names := strings.Split(val, ",")

		if !slices.Contains(names, account.Fname) {
			return false
		}
	}

	if val, ok := query.Params["fname"]["null"]; ok {
		if val == "1" && account.Fname != "" {
			return false
		}

		if val == "0" && account.Fname == "" {
			return false
		}
	}

	// sname filter
	if val, ok := query.Params["sname"]["eq"]; ok {
		if account.Sname != val {
			return false
		}
	}

	if val, ok := query.Params["sname"]["starts"]; ok {
		if !strings.HasPrefix(account.Sname, val) {
			return false
		}
	}

	if val, ok := query.Params["sname"]["null"]; ok {
		if val == "1" && account.Sname != "" {
			return false
		}

		if val == "0" && account.Sname == "" {
			return false
		}
	}

	// phone filter
	if val, ok := query.Params["phone"]["code"]; ok {
		if !strings.HasPrefix(account.Phone, "8("+val+")") {
			return false
		}
	}

	if val, ok := query.Params["phone"]["null"]; ok {
		if val == "1" && account.Phone != "" {
			return false
		}

		if val == "0" && account.Phone == "" {
			return false
		}
	}

	// country filter
	if val, ok := query.Params["country"]["eq"]; ok {
		if account.Country != val {
			return false
		}
	}

	if val, ok := query.Params["country"]["null"]; ok {
		if val == "1" && account.Country != "" {
			return false
		}

		if val == "0" && account.Country == "" {
			return false
		}
	}

	// city filter
	if val, ok := query.Params["city"]["eq"]; ok {
		if account.City != val {
			return false
		}
	}

	if val, ok := query.Params["city"]["any"]; ok {
		names := strings.Split(val, ",")

		if !slices.Contains(names, account.City) {
			return false
		}
	}

	if val, ok := query.Params["city"]["null"]; ok {
		if val == "1" && account.City != "" {
			return false
		}

		if val == "0" && account.City == "" {
			return false
		}
	}

	// birth filter
	if val, ok := query.Params["birth"]["lt"]; ok {
		birth, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		if account.Birth >= int32(birth) {
			return false
		}
	}

	if val, ok := query.Params["birth"]["gt"]; ok {
		birth, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		if account.Birth <= int32(birth) {
			return false
		}
	}

	if val, ok := query.Params["birth"]["year"]; ok {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		birthYear := time.Unix(int64(account.Birth), 0).Year()

		if birthYear != year {
			return false
		}
	}

	// interests filter
	if val, ok := query.Params["interests"]["contains"]; ok {
		interests := strings.Split(val, ",")

		if !slices.ContainsAll(account.Interests, interests) {
			return false
		}
	}

	if val, ok := query.Params["interests"]["any"]; ok {
		interests := strings.Split(val, ",")

		if !slices.ContainsAny(account.Interests, interests) {
			return false
		}
	}

	// likes filter
	if val, ok := query.Params["likes"]["contains"]; ok {
		var searchIds []int32
		vals := strings.Split(val, ",")

		for _, sid := range vals {
			id, err := strconv.Atoi(sid)
			if err != nil {
				log.Fatal(err)
			}

			searchIds = append(searchIds, int32(id))
		}

		var likeIds []int32

		for _, like := range account.Likes {
			likeIds = append(likeIds, like.ID)
		}

		if !slices.ContainsAllInt32(likeIds, searchIds) {
			return false
		}
	}

	// premium filter
	if _, ok := query.Params["premium"]["now"]; ok {
		if int64(account.Premium.Finish) < time.Now().Unix() {
			return false
		}

		fmt.Println("premium", time.Unix(int64(account.Premium.Finish), 0))
	}

	if val, ok := query.Params["premium"]["null"]; ok {
		if val == "1" && account.Premium != (structs.PremiumPeriod{}) {
			return false
		}

		if val == "0" && account.Premium == (structs.PremiumPeriod{}) {
			return false
		}
	}

	return true
}
