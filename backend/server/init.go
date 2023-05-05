package server

import (
	"bitbucket.org/andreychernih/tweemote/lib"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	lib.SetupCustomValidators()
}
