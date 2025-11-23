package yahoo_finance

import (
	"encoding/json"
	"fmt"
	"github.com/Antonious-Stewart/Aggregator/internal/config"
	"github.com/Antonious-Stewart/Aggregator/internal/db"
	"github.com/Antonious-Stewart/Aggregator/internal/models"
	"log"
	"net/http"
)

type YahooFinance struct {
	SqlRunner db.Runner
	ApiUrl    string
}

func New(runner db.Runner) *YahooFinance {
	apiUrl, err := config.GetVar("API_URL")
	if err != nil {
		log.Fatal(err)
	}

	return &YahooFinance{
		SqlRunner: runner,
		ApiUrl:    apiUrl,
	}
}

func (yh YahooFinance) GetRealTime() error {
	apiKey, err := config.GetVar("X_RAPID_API_KEY")
	if err != nil {
		return err
	}

	apiHost, err := config.GetVar("X_RAPID_API_HOST")
	if err != nil {
		return err
	}

	aaplPath := "https://yahoo-finance15.p.rapidapi.com/api/v1/markets/quote?ticker=AAPL&type=STOCKS"
	req, err := http.NewRequest(http.MethodGet, aaplPath, nil)
	if err != nil {
		return err
	}

	req.Header.Add("x-rapidapi-key", apiKey)
	req.Header.Add("x-rapidapi-host", apiHost)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	var realTimeResponse models.MarketQuoteResponse

	err = json.NewDecoder(res.Body).Decode(&realTimeResponse)
	if err != nil {
		return err
	}

	query := `INSERT INTO quotes ("symbol", "company_name", "stock_type", "exchange", "primary_data", "secondary_data", "market_status", "assetclass", "key_stats") values (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := yh.SqlRunner.Exec(query,
		realTimeResponse.Body.Symbol,
		realTimeResponse.Body.CompanyName,
		realTimeResponse.Body.StockType,
		realTimeResponse.Body.Exchange,
		realTimeResponse.Body.PrimaryData,
		realTimeResponse.Body.SecondaryData,
		realTimeResponse.Body.MarketStatus,
		realTimeResponse.Body.AssetClass,
		realTimeResponse.Body.KeyStats,
	)

	fmt.Println(result)
	return nil
}
