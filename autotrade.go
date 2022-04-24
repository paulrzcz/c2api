package c2api

const (
	GetListAvailableBrokersUrl string = "getListAvailableBrokers"
	ChangeAutoTradeUrl         string = "changeAutoTrade"
	StopAutoTradeUrl           string = "stopAutoTrade"
	StartAutoTradeUrl          string = "startAutoTrade"
)

type AutoTradeManagementService struct {
	client *Client
}
