package persist

import (
	"github.com/Eric-GreenComb/one-account-info/bean"
)

// CreateAccount CreateAccount Persist
func (persist *Persist) CreateAccount(account bean.Account) error {
	err := persist.db.Create(&account).Error
	return err
}

// AccountInfo AccountInfo Persist
func (persist *Persist) AccountInfo(code string) (bean.AccountView, error) {

	var account bean.AccountView

	err := persist.db.Table("accounts").Where("code = ?", code).First(&account).Error

	return account, err
}

// UpdateAccount UpdateAccount Persist
func (persist *Persist) UpdateAccount(account bean.AccountForm) error {

	_map := make(map[string]interface{})
	_map["account"] = account.Account
	_map["type"] = account.Type
	_map["address"] = account.Address
	_map["desc"] = account.Desc

	return persist.db.Table("accounts").Where("code = ?", account.Code).Update(_map).Error
}

// GetAllAccounts GetAllAccounts Persist
func (persist *Persist) GetAllAccounts() ([]bean.AccountView, error) {

	var accounts []bean.AccountView

	err := persist.db.Table("accounts").Find(&accounts).Error

	return accounts, err
}
