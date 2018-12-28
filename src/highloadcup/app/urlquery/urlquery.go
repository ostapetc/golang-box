package urlquery

import (
	"errors"
	"fmt"
	"highloadcup/app/slices"
	"highloadcup/app/structs"
	"strconv"
	"strings"
)

// Parse url query
func Parse(form map[string][]string, allowedParams []string, requiredParams []string) (*structs.ServiceQuery, error) {
	query := structs.NewServiceQuery()

	for _, param := range requiredParams {
		if value, ok := form[param]; !ok || value[0] == "" {
			return nil, errors.New("missing required param")
		}
	}

	for key, value := range form {
		if key == "query_id" || key == "limit" || key == "order" {
			continue
		}

		if !slices.Contains(allowedParams, key) && !slices.Contains(requiredParams, key) {
			return nil, errors.New("param " + key + " is not allowed")
		}

		if strings.Contains(key, "_") {
			tmp := strings.Split(key, "_")
			attr := tmp[0]
			oper := tmp[1]

			if _, ok := query.Params[attr]; !ok {
				query.Params[attr] = make(map[string]string)
			}

			query.Params[attr][oper] = value[0]
		} else {
			query.Params[key] = make(map[string]string)
			query.Params[key]["eq"] = value[0]
		}
	}

	if slices.Contains(requiredParams, "limit") {
		limit, err := getInt32ParamValue(form, "limit")
		if err != nil {
			return nil, err
		}

		query.Limit = limit
	}

	if slices.Contains(requiredParams, "order") {
		order, err := getInt32ParamValue(form, "order")
		if err != nil {
			return nil, err
		}

		query.Order = order
	}

	fmt.Println("----------------------")
	print(query)

	return query, nil
}

func print(query *structs.ServiceQuery) {
	for attr, data := range query.Params {
		for oper, value := range data {
			fmt.Println(attr, oper, value)
		}
	}

	if query.Limit != 0 {
		fmt.Println("Limit", query.Limit)
	}

	if query.Order != 0 {
		fmt.Println("Order", query.Order)
	}
}

func getInt32ParamValue(form map[string][]string, param string) (int32, error) {
	value, ok := form[param]

	if !ok {
		return 0, nil
	}

	intValue, err := strconv.Atoi(value[0])
	if err != nil {
		return 0, err
	}

	return int32(intValue), nil
}
