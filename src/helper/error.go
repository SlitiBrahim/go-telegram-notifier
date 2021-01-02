package helper

import "log"

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}