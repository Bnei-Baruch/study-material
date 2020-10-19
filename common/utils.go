package common

import "log"

func FatalIfNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
