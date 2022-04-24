package c2api

import (
	"net/http"
)

const (
	GetSystemRoasterUrl        string = "getSystemRoaster"
	GetListSubscribersUrl      string = "getListSubscribers"
	UnsubscribeFromStrategyUrl string = "unsubscribeFromStrategy"
	GetSystemDetailsUrl        string = "getSystemDetails"
	GetDesiredPositionsUrl     string = "getDesiredPositions"
	RetrieveSignalsWorkingUrl  string = "retrieveSignalsWorking"
	RetrieveSignalsAllUrl      string = "retrieveSignalsAll"
	RequestMargingEquityUrl    string = "requestMarginEquity"
	CreateNewSystemUrl         string = "createNewSystem"
	ChangeSystemAttributeUrl   string = "changeSystemAttribute"
	RequestAllTradesUrl        string = "requestAllTrades_overview"
	SubscribeToStrategyUrl     string = "subscribeToStrategy"
	SetDesiredPositionUrl      string = "setDesiredPositions"
	RetrieveSystemEquityUrl    string = "retrieveSystemEquity"
)

type StrategyAccessService struct {
	client *Client
}

type SystemRoaster struct {
	SystemId                     string `json:"system_id"`
	OwnerPersonId                string `json:"ownerpersonid"`
	TradesStocks                 string `json:"trades_stocks"`
	TradesStocksShort            string `json:"trades_stocks_short"`
	TradesOptionsShort           string `json:"trades_options_short"`
	MinimumPortfolioSizeRequired string `json:"minimum_portfolio_size_required"`
	FreeTrialPeriodDays          string `json:"freeTrialPeriodDays"`
	OwnerScreenName              string `json:"owner_screenname"`
	TradesOptions                string `json:"trades_options"`
	CreatedWhen                  string `json:"created_when"`
	SystemName                   string `json:"system_name"`
	RecentlyInactiveSince        string `json:"recently_inactive_since"`
	EquityCurveStartingCapital   string `json:"equitycurve_startingcapital"`
	MonthlyFee                   string `json:"monthlyFee"`
	TradesForex                  string `json:"trades_forex"`
	IsAlive                      string `json:"isAlive"`
	TradesFutures                string `json:"trades_futures"`
}

type systemRoasterReq struct {
	ApiKey string `json:"apikey"`
	Filter string `json:"filter"`
}

type systemRoasterResp struct {
	Ok       string          `json:"ok"`
	Response []SystemRoaster `json:"response"`
}

func (s *StrategyAccessService) GetSystemRoaster(filter string) ([]SystemRoaster, *http.Response, error) {
	reqBody := systemRoasterReq{
		ApiKey: s.client.apiKey,
		Filter: filter,
	}

	req, err := s.client.NewPostRequest(GetSystemRoasterUrl, reqBody)

	if err != nil {
		return nil, nil, err
	}

	c := &systemRoasterResp{}

	resp, err := s.client.Do(req, c)

	if err != nil {
		return nil, resp, err
	}

	return c.Response, resp, err
}
