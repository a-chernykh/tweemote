package errors

import (
	"log"

	"github.com/stvp/rollbar"
)

func Check(e error) {
	if e != nil {
		log.Printf("Error: %s", e)
		rollbar.Error(rollbar.ERR, e)
		panic(e)
	}
}
