package database

func EmailAlreadyExists(email string, model interface{}) bool {
	result := db.Model(model).Where("email = ?", email).First(model)
	if result.RowsAffected == 1 {
		return true
	}
	return false
}

func Update(model interface{}) (err error) {
	result := db.Model(model).Updates(model)
	err = result.Error
	return
}
