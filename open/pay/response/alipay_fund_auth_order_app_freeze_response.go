package response

type AlipayFundAuthOrderAppFreezeResponse struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	OrderStr string
}
