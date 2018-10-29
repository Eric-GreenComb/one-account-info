package handler

import (
	"bytes"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-account-info/badger"
	"github.com/Eric-GreenComb/one-account-info/bean"
	"github.com/Eric-GreenComb/one-account-info/common"
	"github.com/Eric-GreenComb/one-account-info/config"
	"github.com/Eric-GreenComb/one-account-info/ethereum"
	"github.com/Eric-GreenComb/one-account-info/persist"
)

// WriteBlock WriteBlock
func WriteBlock(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	var _buf bytes.Buffer
	_buf.WriteString("订单号:")
	_buf.WriteString(_formParams.OrderCode)
	_buf.WriteString(";商品ID:")
	_buf.WriteString(_formParams.GoodsID)
	_buf.WriteString(";商品名称:")
	_buf.WriteString(_formParams.GoodName)
	_buf.WriteString(";金额:")
	_buf.WriteString(_formParams.Amount)
	_buf.WriteString(";时间:")
	_buf.WriteString(_formParams.BuyTime)
	_buf.WriteString(";用户:")
	_buf.WriteString(_formParams.UserName)
	_buf.WriteString(";购买编码:")
	_buf.WriteString(_formParams.Desc)

	_desc := _buf.String()

	_txID, err := SendEthereumCoin(_desc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _order bean.Order
	_order.OrderCode = _formParams.OrderCode
	_order.GoodsID = _formParams.GoodsID
	_order.GoodName = _formParams.GoodName
	_order.Amount = _formParams.Amount
	_order.BuyTime = _formParams.BuyTime
	_order.UserName = _formParams.UserName
	_order.Type = 0
	_order.Desc = _formParams.Desc
	_order.TxID = _txID

	err = persist.GetPersist().CreateOrder(_order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txID})
}

// ReadBlock ReadBlock
func ReadBlock(c *gin.Context) {

	_orderid := c.Params.ByName("orderid")

	_order, err := persist.GetPersist().OrderInfo(_orderid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _order})
}

// PutWinerTxID PutWinerTxID
func PutWinerTxID(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	var _buf bytes.Buffer
	_buf.WriteString("商品ID:")
	_buf.WriteString(_formParams.GoodsID)
	_buf.WriteString(";商品名称:")
	_buf.WriteString(_formParams.GoodName)
	_buf.WriteString(";时间:")
	_buf.WriteString(_formParams.WinTime)
	_buf.WriteString(";备注:")
	_buf.WriteString(_formParams.Desc)

	_desc := _buf.String()

	_txID, err := SendEthereumCoin(_desc)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	var _order bean.Order
	_order.GoodsID = _formParams.GoodsID
	_order.GoodName = _formParams.GoodName
	_order.BuyTime = _formParams.WinTime
	_order.Type = 1
	_order.Desc = _formParams.Desc
	_order.TxID = _txID

	err = persist.GetPersist().CreateOrder(_order)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _txID})
}

// GetAllOrders GetAllOrders
func GetAllOrders(c *gin.Context) {

	_catid := c.Params.ByName("catid")
	_patchid := c.Params.ByName("patchid")

	_orders, err := persist.GetPersist().GetAllOrders(_catid, _patchid)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"errcode": 1, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": _orders})
}

// SendEthereumCoin SendEthereumCoin
func SendEthereumCoin(desc string) (string, error) {
	txID := &ethcommon.Hash{}

	_from := config.Ethereum.Address
	_to := _from
	_pwd := string(common.FromHex(config.Ethereum.Passphrase))

	_value, err := badger.NewRead().Get(_from)
	if err != nil {
		return txID.String(), err
	}

	var _keystore string
	_keystore = strings.Replace(string(_value), "\\\"", "\"", -1)
	_key, err := keystore.DecryptKey([]byte(_keystore), _pwd)
	if err != nil {
		return txID.String(), err
	}

	_amountBigInt := ethereum.StringToWei("0.001", 18)
	fmt.Println(_amountBigInt)
	_chainIDBigInt := big.NewInt(config.Ethereum.ChainID)

	// _nonce, err := ethereum.PendingNonce(_from)
	// if err != nil {
	// 	return txID.String(), err
	// }

	_inputData := []byte(desc)
	_txID, err := ethereum.SendEthCoins(_to, config.PendingNonce, _amountBigInt, _key.PrivateKey, _chainIDBigInt, _inputData)
	if err != nil {
		return txID.String(), err
	}
	config.PendingNonce++
	return _txID, err
}
