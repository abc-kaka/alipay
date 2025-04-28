package response

type AlipayUserAgreementQueryResponse struct {
	Code                string          `json:"code"`
	Msg                 string          `json:"msg"`
	PrincipalId         string          `json:"principal_id"`
	PrincipalOpenId     string          `json:"principal_open_id"`
	ValidTime           string          `json:"valid_time"`
	AlipayLogonId       string          `json:"alipay_logon_id"`
	InvalidTime         string          `json:"invalid_time"`
	PricipalType        string          `json:"pricipal_type"`
	DeviceId            string          `json:"device_id"`
	SignScene           string          `json:"sign_scene"`
	AgreementNo         string          `json:"agreement_no"`
	ThirdPartyType      string          `json:"third_party_type"`
	Status              string          `json:"status"`
	SignTime            string          `json:"sign_time"`
	PersonalProductCode string          `json:"personal_product_code"`
	ExternalAgreementNo string          `json:"external_agreement_no"`
	ZmOpenId            string          `json:"zm_open_id"`
	ExternalLogonId     string          `json:"external_logon_id"`
	CreditAuthMode      string          `json:"credit_auth_mode"`
	SingleQuota         string          `json:"single_quota"`
	LastDeductTime      string          `json:"last_deduct_time"`
	NextDeductTime      string          `json:"next_deduct_time"`
	ExecutionPlans      []ExecutionPlan `json:"execution_plans"`
}

type ExecutionPlan struct {
	SingleAmount      string `json:"single_amount"`
	PeriodId          string `json:"period_id"`
	ExecuteTime       string `json:"execute_time"`
	LatestExecuteTime string `json:"latest_execute_time"`
}
