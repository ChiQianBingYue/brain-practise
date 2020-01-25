package playground

import (
	"errors"
	"fmt"
)

func generateErr(foo interface{}) (uint, error) {
	if str, ok := foo.(uint); ok {
		return str, nil
	}
	return 0, errors.New("hahahaha")
}

func PlayScope() {
	var err error
	var one uint = 1
	var a interface{} = one
	var u uint
	if u, err = generateErr(a); err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("u", u)
	}
	fmt.Println("err", err)
}
