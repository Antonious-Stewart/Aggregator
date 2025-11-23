package models

type primaryData struct {
	LastSalePrice      string `json:"lastSalePrice"`
	NetChange          string `json:"netChange"`
	PercentageChange   string
	DeltaIndicator     string `json:"deltaIndicator"`
	LastTradeTimestamp string `json:"lastTradeTimestamp"`
	IsRealTime         bool   `json:"isRealTime"`
	BidPrice           string `json:"bidPrice"`
	AskPrice           string `json:"askPrice"`
	BidSize            string `json:"bidSize"`
	AskSize            string `json:"askSize"`
	Volume             string `json:"volume"`
	Currency           string `json:"currency"`
}

type labelValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type keyStats struct {
	FiftyTwoWeekHighLow labelValue `json:"fiftyTwoWeekHighLow"`
	DayRange            labelValue `json:"dayRange"`
}

type RealTimeResponse struct {
	Symbol        string      `json:"symbol"`
	CompanyName   string      `json:"companyName"`
	StockType     string      `json:"stockType"`
	Exchange      string      `json:"exchange"`
	PrimaryData   primaryData `json:"primaryData"`
	SecondaryData struct{}    `json:"secondary"`
	MarketStatus  string      `json:"marketStatus"`
	AssetClass    string      `json:"assetClass"`
	KeyStats      keyStats    `json:"keyStats"`
}

type MarketQuoteResponse struct {
	Body RealTimeResponse `json:"body"`
}
