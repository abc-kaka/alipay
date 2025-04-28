package request

import (
	"github.com/abc-kaka/alipay/common/req"
)

// AlipayOpenMiniOrderDeliveryModifyRequest 履约状态变更 - 参数
type AlipayOpenMiniOrderDeliveryModifyRequest struct {
	req.BizContentRequest
	OrderId              string                `json:"order_id,omitempty"`
	OutOrderId           string                `json:"out_order_id,omitempty"`
	OpenId               string                `json:"open_id,omitempty"`
	UserId               string                `json:"user_id,omitempty"`
	Status               string                `json:"status,omitempty"`
	ContactInfo          *ContactInfo          `json:"contact_info,omitempty"`
	BookingInfo          *BookingInfo          `json:"booking_info,omitempty"`
	TicketInfos          []TicketInfo          `json:"ticket_infos,omitempty"`
	ActivityInfos        []ActivityInfos       `json:"activity_infos,omitempty"`
	TourInfo             *TourInfo             `json:"tour_info,omitempty"`
	PriceInfo            *PriceInfo            `json:"price_info,omitempty"`
	AddressInfo          *AddressInfo          `json:"address_info,omitempty"`
	SendOrderContactInfo *SendOrderContactInfo `json:"send_order_contact_info,omitempty"`
	OrderInspectPrice    string                `json:"order_inspect_price,omitempty"`
	UserConfirmPrice     string                `json:"user_confirm_price,omitempty"`
	ItemInfos            []ItemInfo            `json:"item_infos,omitempty"`
	ReasonCode           string                `json:"reason_code,omitempty"`
}

type ContactInfo struct {
	ContactName string `json:"contact_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type BookingInfo struct {
	CustomerServiceMobile string `json:"customer_service_mobile,omitempty"`
	RoomNum               string `json:"room_num,omitempty"`
	RefundRule            string `json:"refund_rule,omitempty"`
	CheckOutDate          string `json:"check_out_date,omitempty"`
	CheckInTime           string `json:"check_in_time,omitempty"`
	ConfirmBookingTime    string `json:"confirm_booking_time,omitempty"`
	HaveStayTime          string `json:"have_stay_time,omitempty"`
	CheckOutTime          string `json:"check_out_time,omitempty"`
	BookingTime           string `json:"booking_time,omitempty"`
	Deadline              string `json:"deadline,omitempty"`
	CheckInDate           string `json:"check_in_date,omitempty"`
}

type TicketInfo struct {
	TicketStatus     string `json:"ticket_status,omitempty"`
	EventEndTime     string `json:"event_end_time,omitempty"`
	EventId          string `json:"event_id,omitempty"`
	TicketLink       string `json:"ticket_link,omitempty"`
	EventStartTime   string `json:"event_start_time,omitempty"`
	TicketId         string `json:"ticket_id,omitempty"`
	PerformanceSeats string `json:"performance_seats,omitempty"`
	EventStatus      string `json:"event_status,omitempty"`
}

type ActivityInfos struct {
	ActivityName string `json:"activity_name,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	ActivityId   string `json:"activity_id,omitempty"`
	EndTime      string `json:"end_time,omitempty"`
	Link         string `json:"link,omitempty"`
	Status       string `json:"status,omitempty"`
}

type TourInfo struct {
	ToLocation   string       `json:"to_location,omitempty"`
	FromLocation string       `json:"from_location,omitempty"`
	PackageInfo  *PackageInfo `json:"package_info,omitempty"`
}

type PackageInfo struct {
	ReturnTime    string `json:"return_time,omitempty"`
	Name          string `json:"name,omitempty"`
	Mobile        string `json:"mobile,omitempty"`
	DepartureTime string `json:"departure_time,omitempty"`
	IdentityCard  string `json:"identity_card,omitempty"`
	Population    string `json:"population,omitempty"`
}

type PriceInfo struct {
	RealOrderPrice string `json:"real_order_price,omitempty"`
	OrderPrice     string `json:"order_price,omitempty"`
}

type AddressInfo struct {
	DetailedAddress      string `json:"detailed_address,omitempty"`
	TelNumber            string `json:"tel_number,omitempty"`
	ReceiverZip          string `json:"receiver_zip,omitempty"`
	ReceiverName         string `json:"receiver_name,omitempty"`
	ReceiverDivisionCode string `json:"receiver_division_code,omitempty"`
}

type SendOrderContactInfo struct {
	ContactName string `json:"contact_name,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type ItemInfo struct {
	OutItemId     string    `json:"out_item_id,omitempty"`
	OutSkuId      string    `json:"out_sku_id,omitempty"`
	RecycleStatus string    `json:"recycle_status,omitempty"`
	InspectPrice  string    `json:"inspect_price,omitempty"`
	RentInfo      *RentInfo `json:"rent_info,omitempty"`
}

type RentInfo struct {
	BuyoutPrice      string `json:"buyout_price,omitempty"`
	DepositPrice     string `json:"deposit_price,omitempty"`
	RentStartTime    string `json:"rent_start_time,omitempty"`
	RentEndTime      string `json:"rent_end_time,omitempty"`
	InitialRentPrice string `json:"initial_rent_price,omitempty"`
}
