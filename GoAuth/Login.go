package GoAuth

import "fmt"

func Login[T comparable,S string](data T,key S) bool {
	fmt.Println(data,key)
	return false
}
