package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func CasbinEnforce(e *casbin.Enforcer, sub, obj, act string) {
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}

	if ok == true {
		// permit alice to read data1
		fmt.Printf("Allow %s %s %s\n", sub, act, obj)
	} else {
		// deny the request, show an error
		fmt.Printf("Deny %s %s %s\n", sub, act, obj)
	}
}

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		fmt.Println("Error occurred: ", err.Error())
	}

	requests := map[string][]string{
		"ti":  {"admin", "/users", "POST"},
		"suu": {"user", "/products", "POST"},
		"dan": {"user", "/products", "GET"},
	}

	for key, value := range requests {
		fmt.Printf("Check permission for %s\n", key)
		CasbinEnforce(e, value[0], value[1], value[2])
	}
}
