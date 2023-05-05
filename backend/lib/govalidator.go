package lib

import (
	"log"
	"reflect"

	"github.com/asaskevich/govalidator"
)

func SetupCustomValidators() {
	govalidator.CustomTypeTagMap.Set("passwordConfirmation", govalidator.CustomTypeValidator(func(i interface{}, context interface{}) bool {
		var password string
		var ok bool

		if password, ok = i.(string); !ok {
			log.Printf("Error asserting %v to string\n", i)
			return false
		}

		r := reflect.ValueOf(context)
		f := reflect.Indirect(r).FieldByName("PasswordConfirmation")
		passwordConfirmation := f.String()

		return password == passwordConfirmation
	}))
}
