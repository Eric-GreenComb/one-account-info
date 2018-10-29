package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Eric-GreenComb/one-account-info/bean"
	"github.com/Eric-GreenComb/one-account-info/nsq"
)

// WriteNsq WriteNsq
func WriteNsq(c *gin.Context) {

	var _formParams bean.FormParams
	c.BindJSON(&_formParams)

	nsq.Producer.Publish("topic_string", []byte(_formParams.GoodsID))

	c.JSON(http.StatusOK, gin.H{"errcode": 0, "msg": "OK"})
}
