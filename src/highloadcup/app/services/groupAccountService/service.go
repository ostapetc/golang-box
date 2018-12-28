package groupAccountService

import (
	"highloadcup/app/slices"
	"highloadcup/app/structs"
	"log"
	"strconv"
	"strings"
	"time"
)

// ServiceQuery is a parsed http query
type AggGroup struct {
	Count     int32
	Sex       string
	City      string
	Country   string
	Status    string
	Interests string
}

type RowGroup struct {
	Key   string
	Count int32
}

// Group filters accounts by query params
func Group(accountsIndex *[]structs.Account, query *structs.ServiceQuery) *[]map[string]interface{} {
	var result []map[string]interface{}

	keys := strings.Split(query.Params["keys"]["eq"], ",")

	countMap := make(map[string]int32)

	for _, account := range *accountsIndex {
		if !filter(&account, query, keys) {
			continue
		}

		keyValues := getKeyValues(account, keys)

		for _, keyValue := range keyValues {
			countMap[keyValue]++
		}
	}

	groups := sortCountMap(countMap, query.Limit, query.Order)

	for _, group := range groups {
		result = append(result, makeAccountRow(group.Key, group.Count))
	}

	return &result
}

func sortCountMap(countMap map[string]int32, limit int32, order int32) []RowGroup {
	var result []RowGroup

	sorted := make(map[string]bool)

	for i := int32(0); i < limit; i++ {
		key, count := findNextLessOrEqual(countMap, sorted, order)

		sorted[key] = true

		result = append(result, RowGroup{key, count})
	}

	return result
}

func getMaxCount(countMap map[string]int32) int32 {
	var max int32

	for _, count := range countMap {
		if count > max {
			max = count
		}
	}

	return max
}

func findNextLessOrEqual(countMap map[string]int32, sortedKeys map[string]bool, order int32) (string, int32) {
	var maxCount int32
	var maxKey string

	for key, count := range countMap {
		if _, ok := sortedKeys[key]; ok {
			continue
		}

		if order == -1 {
			if count > maxCount {
				maxCount = count
				maxKey = key
			}
		} else {
			if count < maxCount || maxKey == "" {
				maxCount = count
				maxKey = key
			}
		}

		if count == maxCount {
			currCompareKey := strings.Split(key, ":")[1]
			maxCompareKey := strings.Split(maxKey, ":")[1]

			if currCompareKey < maxCompareKey {
				maxCount = count
				maxKey = key
			}
		}
	}

	return maxKey, maxCount
}

func getKeyValues(account structs.Account, keys []string) []string {
	var result []string
	var values []string

	for _, key := range keys {
		if key != "interests" {
			values = append(values, key+":"+account.GetFieldValue(strings.Title(key)))
		}
	}

	primValue := strings.Join(values, ";")

	if slices.Contains(keys, "interests") {
		if len(primValue) > 0 {
			primValue += ";"
		}

		for _, interest := range account.Interests {
			result = append(result, primValue+"interest:"+interest)
		}
	} else {
		result = append(result, primValue)
	}

	return result
}

func makeAccountRow(groupKey string, count int32) map[string]interface{} {
	obj := map[string]interface{}{}
	obj["count"] = count

	keyValues := strings.Split(groupKey, ";")

	for _, keyValue := range keyValues {
		value := strings.Split(keyValue, ":")
		obj[value[0]] = value[1]
	}

	// obj["key"] = groupKey

	return obj
}

func filter(account *structs.Account, query *structs.ServiceQuery, groupKeys []string) bool {
	for _, key := range groupKeys {
		value := account.GetFieldValue(strings.Title(key))

		if len(value) == 0 {
			return false
		}
	}

	// sex filter
	if val, ok := query.Params["sex"]["eq"]; ok {
		if account.Sex != val {
			return false
		}
	}

	// email filter
	if val, ok := query.Params["email"]["eq"]; ok {
		if account.Email != val {
			return false
		}
	}

	// status filter
	if val, ok := query.Params["status"]["eq"]; ok {
		if account.Status != val {
			return false
		}
	}

	// fname filter
	if val, ok := query.Params["fname"]["eq"]; ok {
		if account.Fname != val {
			return false
		}
	}

	// sname filter
	if val, ok := query.Params["sname"]["eq"]; ok {
		if account.Sname != val {
			return false
		}
	}

	// phone filter
	if val, ok := query.Params["phone"]["eq"]; ok {
		if account.Phone != val {
			return false
		}
	}

	// country filter
	if val, ok := query.Params["country"]["eq"]; ok {
		if account.Country != val {
			return false
		}
	}

	// city filter
	if val, ok := query.Params["city"]["eq"]; ok {
		if account.City != val {
			return false
		}
	}

	// birth filter
	if val, ok := query.Params["birth"]["eq"]; ok {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		birthYear := time.Unix(int64(account.Birth), 0).Year()

		if birthYear != year {
			return false
		}
	}

	// joined filter
	if val, ok := query.Params["joined"]["eq"]; ok {
		year, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		joinYear := time.Unix(int64(account.Joined), 0).Year()

		if joinYear != year {
			return false
		}
	}

	// interests filter
	if val, ok := query.Params["interests"]["eq"]; ok {
		if !slices.Contains(account.Interests, val) {
			return false
		}
	}

	// likes filter
	if val, ok := query.Params["likes"]["eq"]; ok {
		id, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}

		var likeIds []int32

		for _, like := range account.Likes {
			likeIds = append(likeIds, like.ID)
		}

		if !slices.ContainsInt32(likeIds, int32(id)) {
			return false
		}
	}

	return true
}
