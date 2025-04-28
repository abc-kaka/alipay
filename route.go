package alipay

type Route struct {
	Open *OpenRoute
	Mini *MiniRoute
}

type OpenRoute struct {
	Member *OpenMemberRouter
	Pay    *MiniPayRouter
}

type MiniRoute struct {
	OpenOrder *MiniOpenOrderRouter
}

func NewRoute(client Client) *Route {
	return &Route{
		Open: &OpenRoute{
			Pay:    NewMiniPayRouter(client),
			Member: NewH5AppMemberRouter(client),
		},
		Mini: &MiniRoute{
			OpenOrder: NewMiniOpenOrderRouter(client),
		},
	}
}
