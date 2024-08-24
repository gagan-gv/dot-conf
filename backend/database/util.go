package database

import "gorm.io/gorm"

func EmailAlreadyExists(email string, model interface{}) bool {
	if FindByEmail(email, model).RowsAffected == 1 {
		return true
	}
	return false
}

func Update(model interface{}) (err error) {
	result := db.Model(model).Updates(model)
	err = result.Error
	return
}

func FindByEmail(email string, model interface{}) *gorm.DB {
	return db.Model(model).Where("email = ?", email).First(model)
}
