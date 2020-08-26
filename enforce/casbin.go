package enforce

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
