package utils

import "fmt"

func StringSlicecontains(s []string, e interface{}) bool {
	for _, a := range s {
		fmt.Println(a)
		fmt.Println(e)

		if a == e {
			return true
		}
	}
	return false
}
