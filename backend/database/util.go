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

func FindAppByKey(key string, model interface{}) *gorm.DB {
	return db.Model(model).Where("app_key = ?", key).First(model)
}

func FindAppByCompanyIdAndName(companyId int64, name string, model interface{}) *gorm.DB {
	return db.Model(model).Where("company_id = ? AND name = ?", companyId, name).First(model)
}

func AppAlreadyExists(companyId int64, name string, model interface{}) bool {
	if FindAppByCompanyIdAndName(companyId, name, model).RowsAffected == 1 {
		return true
	}
	return false
}

func FindConfigByName(name, appId string, model interface{}) *gorm.DB {
	return db.Model(model).Where("name = ? AND app_id = ?", name, appId).First(model)
}

func ConfigAlreadyExists(name, appId string, model interface{}) bool {
	if FindConfigByName(name, appId, model).RowsAffected == 1 {
		return true
	}
	return false
}

func FindConfigById(configId string, model interface{}) *gorm.DB {
	return db.Model(model).Where("id = ?", configId).First(model)
}
