package alipay

type Route struct {
	Open *OpenRoute
	Mini *MiniRoute
}

type OpenRoute struct {
	Member *OpenMemberRouter
	Pay    *OpenPayRouter
}

type MiniRoute struct {
	OpenOrder *MiniOpenOrderRouter
}

func NewRoute(client Client) *Route {
	return &Route{
		Open: &OpenRoute{
			Pay:    NewOpenPayRouter(client),
			Member: NewOpenMemberRouter(client),
		},
		Mini: &MiniRoute{
			OpenOrder: NewMiniOpenOrderRouter(client),
		},
	}
}
