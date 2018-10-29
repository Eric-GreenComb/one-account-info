package handler

import (
	"math/big"
	"net/http"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-account-info/badger"
	"github.com/Eric-GreenComb/one-account-info/bean"
	"github.com/Eric-GreenComb/one-account-info/config"
	"github.com/Eric-GreenComb/one-account-info/ethereum"
)

// PendingNonce PendingNonce
func PendingNonce(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_nonce, err := ethereum.PendingNonce(_formParams.Params)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _nonce})
}

// SendEthCoin SendEthCoin
func SendEthCoin(c *gin.Context) {
	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	_from := _formParams.From
	_to := _formParams.To
	_amount := _formParams.Amount
	_pwd := _formParams.Pwd
	_decimals := _formParams.Decimals
	_desc := _formParams.Desc

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _keystore string
	_keystore = strings.Replace(string(_value), "\\\"", "\"", -1)
	_key, err := keystore.DecryptKey([]byte(_keystore), _pwd)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_amountBigInt := ethereum.StringToWei(_amount, _int)
	_chainIDBigInt := big.NewInt(config.Ethereum.ChainID)

	_nonce, err := ethereum.PendingNonce(_from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	_inputData := []byte(_desc)
	_txid, err := ethereum.SendEthCoins(_to, _nonce, _amountBigInt, _key.PrivateKey, _chainIDBigInt, _inputData)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txid})
}

// GetBalance GetBalance
func GetBalance(c *gin.Context) {

	_addr := c.Params.ByName("addr")

	_ethCoin, err := ethereum.GetBalance(_addr)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _ethCoin})
}

// StringToWei StringToWei
func StringToWei(c *gin.Context) {

	_val := c.Params.ByName("val")
	_decimals := c.Params.ByName("decimals")

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_wei := ethereum.StringToWei(_val, _int)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _wei.String()})
}

// FloatToWei FloatToWei
func FloatToWei(c *gin.Context) {

	_val := c.Params.ByName("val")
	_decimals := c.Params.ByName("decimals")

	_float64, err := strconv.ParseFloat(_val, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
	}

	_int, err := strconv.Atoi(_decimals)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 2, "msg": err.Error()})
	}

	_wei := ethereum.FloatToWei(_float64, _int)

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _wei})
}
