package utils

import (
	"fmt"
	"reflect"
)

var MultiValueFilters = []string{"specialization_id", "writer_id"}

func QueryConditions(filters interface{}) []string {

	var where []string

	v := reflect.ValueOf(filters)

	filterValues := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		filterValues[i] = v.Field(i).Interface()
	}

	for name, key := range []string{"writer_id", "specialization_id"} {
		if filterValues[name] != "" {
			if StringSlicecontains(MultiValueFilters, key) {
				where = append(where, fmt.Sprintf("%s IN (%s)", key, filterValues[name]))
			} else {
				where = append(where, fmt.Sprintf("%s = %s", key, fmt.Sprintf("%v", filterValues[name])))
			}

		}
	}
	return where
}
