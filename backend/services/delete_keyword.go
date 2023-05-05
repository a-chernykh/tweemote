package services

import (
	"errors"

	"bitbucket.org/andreychernih/tweemote/models"
)

func DeleteKeyword(id uint) error {
	db := models.Connect()
	deleted := db.Where("id = ?", id).Delete(models.Keyword{}).RowsAffected
	if deleted == 0 {
		return errors.New("Twitter account was not found")
	} else {
		return nil
	}
}
