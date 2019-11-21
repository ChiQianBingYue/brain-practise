package main

import (
	"fmt"
	"runtime"
)

type Err struct {
	msg string
}

func (e *Err) Error() string {
	return e.msg
}
func main() {
	// playground.PlayGorm()
	// u64, err := strconv.ParseUint("0", 10, 64)
	// fmt.Println(u64, err)
	_, file, line, _ := runtime.Caller(3)

	fmt.Println(file, line)
}
