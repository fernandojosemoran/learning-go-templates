package enums

import "fmt"

type Method int

const (
	get Method = iota + 1
	post
	put
	delete
	patch
)

var MethodOptions = map[Method]string{
	get:    "GET",
	post:   "POST",
	put:    "PUT",
	delete: "DELETE",
	patch:  "PATCH",
}

func (mo Method) String() string {
	if val, ok := MethodOptions[mo]; ok {
		return val
	}
	return "UNKNOWN"
}

func GetMethod(m Method) Method {
	switch m {
	case get, post, put, delete, patch:
		return m
	default:
		panic(fmt.Errorf("unknown method option: %d", m))
	}
}
