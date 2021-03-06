package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"

	"casbin-example/enforce"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}

	requests := map[string][]string{
		"dang":  {"admin", "/users", "POST"},
		"phuoc": {"user", "/products", "POST"},
		"quang": {"user", "/products", "GET"},
	}

	for key, value := range requests {
		fmt.Printf("Check permission for %s\n", key)
		enforce.CasbinEnforce(e, value[0], value[1], value[2])
	}
}
