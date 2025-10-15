package util

import "fmt"

func ThrowError(desc string, err error) {
	panic(fmt.Sprintf(desc+": %v", err))
}
