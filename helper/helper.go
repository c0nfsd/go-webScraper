package helper

import (
	"fmt"
	"os"
)

func ErrCheck(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
