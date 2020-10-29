package common

import "log"

func PanicIfNotNil(err error) {
	if err != nil {
		log.Panic(err)
	}
}
