package c2api

import "net/http"

const (
	SubmitSignalUrl string = "submitSignal"
	CancelSignalUrl string = "cancelSignal"
)

type SignalEntryService struct {
	client *Client
}

type submitSignalReq struct {
	ApiKey   string `json:"apikey"`
	SystemId string `json:"systemid"`
	Signal   Signal `json:"signal"`
}

// Action types
const (
	BuyToOpen   string = "BTO"
	SellToOpen  string = "STO"
	SellShort   string = "SSHORT"
	BuyToCover  string = "BTC"
	SellToCover string = "STC"
)

// Types of symbols
const (
	Stock  string = "stock"
	Option string = "option"
	Future string = "future"
	Forex  string = "forex"
)

// Durations
const (
	DayOrder       string = "DAY"
	GoodTillCancel string = "GTC"
)

type Signal struct {
	Action          string  `json:"action"`
	TypeOfSymbol    string  `json:"typeofsymbol"`
	Duration        string  `json:"duration"`
	StopPrice       float64 `json:"stop,omitempty"`
	LimitPrice      float64 `json:"limit,omitempty"`
	Market          string  `json:"market,omitempty"`
	ProfitTarget    float64 `json:"profittarget,omitempty"`
	StopLoss        float64 `json:"stoploss,omitempty"`
	ConditionalUpon string  `json:"conditional,omitempty"`
	Xreplace        string  `json:"xreplace,omitempty"`
}

type SignalConfirmation struct {
}

type submitSignalResp struct {
	Ok          string             `json:"ok"`
	Signal      SignalConfirmation `json:"signal"`
	SignalId    string             `json:"signalid"`
	ElapsedTime string             `json:"elapsed_time"`
}

func (s *SignalEntryService) SubmitSignal(systemId string, signal Signal) (*SignalConfirmation, *http.Response, error) {
	reqBody := submitSignalReq{
		ApiKey:   s.client.apiKey,
		SystemId: systemId,
		Signal:   signal,
	}

	req, err := s.client.NewPostRequest(SubmitSignalUrl, reqBody)
	if err != nil {
		return nil, nil, err
	}

	c := &submitSignalResp{}
	resp, err := s.client.Do(req, c)

	if err != nil {
		return nil, resp, err
	}

	return &c.Signal, resp, err
}
