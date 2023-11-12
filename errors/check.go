package errors

import (
	"fmt"
	"os"
)

func Check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(0)
	}
}
