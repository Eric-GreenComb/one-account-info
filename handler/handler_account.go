package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-account-info/bean"
	"github.com/Eric-GreenComb/one-account-info/persist"
)

// CreateAccount CreateAccount
func CreateAccount(c *gin.Context) {

	var _account bean.Account
	c.BindJSON(&_account)

	err := persist.GetPersist().CreateAccount(_account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "create account error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _account.ID})
}

// UpdateAccount UpdateAccount
func UpdateAccount(c *gin.Context) {

	var _account bean.AccountForm
	c.BindJSON(&_account)

	err := persist.GetPersist().UpdateAccount(_account)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "update account error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}

// GetAccountInfo GetAccountInfo
func GetAccountInfo(c *gin.Context) {

	_code := c.Params.ByName("code")

	_account, err := persist.GetPersist().AccountInfo(_code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get account error"})
		return
	}

	c.JSON(http.StatusOK, _account)
}

// ListAccount ListAccount
func ListAccount(c *gin.Context) {

	_accounts, err := persist.GetPersist().GetAllAccounts()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get account error"})
		return
	}

	c.JSON(http.StatusOK, _accounts)
}
