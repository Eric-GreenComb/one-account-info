package handler

import (
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"

	"github.com/Eric-GreenComb/one-account-info/ethereum"
	"github.com/Eric-GreenComb/one-account-info/persist"
)

// GetEtherBalance GetEtherBalance
func GetEtherBalance(c *gin.Context) {

	_code := c.Params.ByName("code")

	_account, err := persist.GetPersist().AccountInfo(_code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": "get account error"})
		return
	}

	_ethCoin, err := ethereum.GetBalance(_account.Account)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_wei, _ := StringToWei(_ethCoin)

	_eth, _ := Wei2Eth(_wei, _account.Decimal)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "wei": _wei.String(), "eth": _eth.String()})
}

// GetTokenBalance GetTokenBalance
func GetTokenBalance(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": 0})
}

// StringToWei StringToWei
func StringToWei(hex string) (*big.Int, bool) {
	n := new(big.Int)
	_str := Remove0x(hex)
	return n.SetString(_str, 16)
}

// Wei2Eth Wei2Eth
func Wei2Eth(wei *big.Int, quantity int64) (decimal.Decimal, error) {
	_wei, err := decimal.NewFromString(wei.String())
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	_float, err := strconv.ParseFloat(strconv.FormatInt(quantity, 10), 64)
	if err != nil {
		return decimal.NewFromFloat(0), err
	}

	_quantity := decimal.NewFromFloat(_float)
	_ether := _wei.Div(_quantity)
	return _ether, nil
}

// Remove0x Remove0x
func Remove0x(s string) string {
	var _ret string
	if len(s) > 1 {
		if s[0:2] == "0x" || s[0:2] == "0X" {
			_ret = s[2:]
		}
	}
	return _ret
}