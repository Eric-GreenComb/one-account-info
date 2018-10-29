package bean

import ()

// FormParams FormParams
type FormParams struct {
	Params string `form:"params" json:"params"` // params
	Key    string `form:"key" json:"key"`       // key
	Value  string `form:"value" json:"value"`   // value

	TTL         string `form:"ttl" json:"ttl"`                   // ttl
	Address     string `form:"address" json:"address"`           //
	From        string `form:"from" json:"from"`                 //
	To          string `form:"to" json:"to"`                     //
	Amount      string `form:"amount" json:"amount"`             //
	Decimals    string `form:"decimals" json:"decimals"`         //
	Pwd         string `form:"pwd" json:"pwd"`                   //
	Name        string `form:"name" json:"name"`                 //
	Symbol      string `form:"symbol" json:"symbol"`             //
	Total       string `form:"total" json:"total"`               //
	Desc        string `form:"desc" json:"desc"`                 //
	IsWait      string `form:"iswait" json:"iswait"`             //
	Number      string `form:"number" json:"number"`             //
	CallAddress string `form:"call_address" json:"call_address"` //
	Mnemonic    string `form:"mnemonic" json:"mnemonic"`         //
	Path        string `form:"path" json:"path"`                 //

	OrderCode string `form:"ordercode" json:"ordercode"` // 订单编码
	GoodsID   string `form:"goodsid" json:"goodsid"`     // 货物id
	GoodName  string `form:"goodname" json:"goodname"`   // Iphone(第三期）
	BuyTime   string `form:"buytime" json:"buytime"`     // 购买时间
	UserName  string `form:"username" json:"username"`   // 购买用户名称
	Type      int8   `form:"type" json:"type"`           // 类型  0为下单，1为抽奖

	WinTime string `form:"wintime" json:"wintime"` // 开奖时间
}
