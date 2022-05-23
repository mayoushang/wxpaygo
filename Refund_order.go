package wxpay

import (
	
)

// UnifiedOrderParam 参数
type RefundOrderParam struct {
	Param `xml:",innerXml"`

	Transaction_id string `xml:"transaction_id"`

	Out_trade_no string `xml:"out_trade_no"`

	Out_refund_no string `xml:"out_refund_no"`

	Total_fee int `xml:"total_fee"`

	Refund_fee int `xml:"refund_fee"`

	Refund_fee_type string `xml:"refund_fee_type"`

	Refund_desc string `xml:"refund_desc"`

	Refund_account string `xml:"refund_account"`

	Notify_url string `xml:"notify_url"`
}

// UnifiedOrderResponse 统一下单返回
type RefundOrderResponse struct {
	Response `xml:",innerXml"`

	ResultCode string `xml:"result_code"`
	ErrCode  string `xml:"err_code"`
	ErrCodeDes   string `xml:"err_code_des"`
	Appid   string `xml:"appid"`
}