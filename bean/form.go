package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value
}

// AccountForm AccountForm
type AccountForm struct {
	Code    string `form:"code" json:"code"`       // account code
	Account string `form:"account" json:"account"` // account
	Type    int8   `form:"type" json:"type"`       // account type 0:eth,1:token
	Address string `form:"address" json:"address"` // token address
	Decimal int64  `form:"decimal" json:"decimal"` // token decimal
	Desc    string `form:"desc" json:"desc"`       // desc
}
