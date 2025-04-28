package request

import (
	"github.com/abc-kaka/alipay/common/req"
)

// AlipayOpenMiniOrderModifyRequest 订单修改 - 参数
type AlipayOpenMiniOrderModifyRequest struct {
	req.BizContentRequest
	OutOrderId      string           `json:"out_order_id,omitempty"`
	OrderId         string           `json:"order_id,omitempty"`
	UserId          string           `json:"user_id,omitempty"`
	OpenId          string           `json:"open_id,omitempty"`
	PriceInfo       *PriceInfo       `json:"price_info,omitempty"`
	ItemInfos       []ItemInfo       `json:"item_infos,omitempty"`
	StagePayPlans   []StagePayPlan   `json:"stage_pay_plans,omitempty"`
	AllocAmountInfo *AllocAmountInfo `json:"alloc_amount_info,omitempty"`
	SubMerchant     map[string]any   `json:"sub_merchant,omitempty"`
	SendAddressInfo *SendAddressInfo `json:"send_address_info,omitempty"`
	DeliveryDetail  *DeliveryDetail  `json:"delivery_detail,omitempty"`
	AddressInfo     *AddressInfo     `json:"address_info,omitempty"`
}

type StagePayPlan struct {
	StageNo           int                `json:"stage_no,omitempty"`
	StagePayPlanInfos []StagePayPlanInfo `json:"stage_pay_plan_infos,omitempty"`
}

type StagePayPlanInfo struct {
	PlanPayTime     string `json:"plan_pay_time,omitempty"`
	PlanBuyoutPrice string `json:"plan_buyout_price,omitempty"`
	PlanPayNo       int    `json:"plan_pay_no,omitempty"`
	PlanPayPrice    string `json:"plan_pay_price,omitempty"`
}

type AllocAmountInfo struct {
	InvestAppId        string              `json:"invest_app_id"`
	InvestId           string              `json:"invest_id"`
	BuyOutRoyalty      *BuyOutRoyalty      `json:"buy_out_royalty"`
	RentRoyaltyDetails []RentRoyaltyDetail `json:"rent_royalty_details"`
}

type BuyOutRoyalty struct {
	ExpectRoyaltyTime string `json:"expect_royalty_time"`
	RoyaltyPeriod     int    `json:"royalty_period"`
	RoyaltyPrice      string `json:"royalty_price"`
	BuyOutPrice       string `json:"buy_out_price"`
	RoyaltyType       string `json:"royalty_type"`
}

type RentRoyaltyDetail struct {
	Royalties []Royalties `json:"royalties"`
	StageNo   int         `json:"stage_no"`
}

type Royalties struct {
	ExpectRoyaltyTime string `json:"expect_royalty_time"`
	RoyaltyPeriod     int    `json:"royalty_period"`
	RoyaltyPrice      string `json:"royalty_price"`
	BuyOutPrice       string `json:"buy_out_price"`
	RoyaltyType       string `json:"royalty_type"`
}

type SendAddressInfo struct {
	DetailedAddress string `json:"detailed_address"`
	TelNumber       string `json:"tel_number"`
	Name            string `json:"name"`
	DivisionCode    string `json:"division_code"`
}

type DeliveryDetail struct {
	DeliveryStartTime string `json:"delivery_start_time"`
	DeliveryEndTime   string `json:"delivery_end_time"`
}
