package internal

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// type contextKey string
// func (c contextKey) String() string {
// 	return string(c)
// }

// Claims : declares structure
type Claims struct {
	UUID  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// JSONStandardResponse : structure to classify JSON response
type JSONStandardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// JSONJwtResponse : structure to classify JSON JWT response
type JSONJwtResponse struct {
	Code int    `json:"code"`
	JWT  string `json:"jwt"`
}

// JSONRequestCurrencyQuote : structure responsible for mapping currency quote JSON response
type JSONRequestCurrencyQuote struct {
	Status struct {
		Timestamp    time.Time   `json:"timestamp"`
		ErrorCode    int         `json:"error_code"`
		ErrorMessage interface{} `json:"error_message"`
		Elapsed      int         `json:"elapsed"`
		CreditCount  int         `json:"credit_count"`
		Notice       interface{} `json:"notice"`
	} `json:"status"`
	Data struct {
		Num1 struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			DateAdded         time.Time   `json:"date_added"`
			Tags              []string    `json:"tags"`
			MaxSupply         int         `json:"max_supply"`
			CirculatingSupply int         `json:"circulating_supply"`
			TotalSupply       int         `json:"total_supply"`
			Platform          interface{} `json:"platform"`
			CmcRank           int         `json:"cmc_rank"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				BRL struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"BRL"`
			} `json:"quote"`
		} `json:"1"`
	} `json:"data"`
}
