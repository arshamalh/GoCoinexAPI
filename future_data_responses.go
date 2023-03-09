package gocoinex

import (
	"encoding/json"
	"net/http"
)

type Ping struct {
	Data string `json:"data"` // Pong

}

func (r *Ping) Parse(raw_response *http.Response) (*Ping, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type SystemTime struct {
	Data int `json:"data"` // Server time, milliseconds

}

func (r *SystemTime) Parse(raw_response *http.Response) (*SystemTime, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketList struct {
	Name        string   `json:"name"`        // Market name
	Type        int      `json:"type"`        // Contract type, 1: Linear contract, 2: Inverse contract
	Stock       string   `json:"stock"`       // Base coin
	Money       string   `json:"money"`       // Price coin
	Fee_prec    int      `json:"fee_prec"`    // Rate decimal
	Stock_prec  int      `json:"stock_prec"`  // Base coin decimal
	Money_prec  int      `json:"money_prec"`  // Price coin decimal
	Multiplier  int      `json:"multiplier"`  // Multiplier
	Amount_prec int      `json:"amount_prec"` // Quantity decimal
	Amount_min  string   `json:"amount_min"`  // Minimum amount
	Tick_size   string   `json:"tick_size"`   // Price granularity
	Leverages   []string `json:"leverages"`   // Margin list
	Available   bool     `json:"available"`   // Whether the market is open

}

func (r *MarketList) Parse(raw_response *http.Response) (*MarketList, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type PositionLevel struct {
	Params0 string `json:"params0"` // amount, amount
	Params1 string `json:"params1"` // leverage, leverage
	Params2 string `json:"params2"` // mainten margin, maintenance margin rate

}

func (r *PositionLevel) Parse(raw_response *http.Response) (*PositionLevel, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketStatus struct {
	Period               int    `json:"period"`               // Period
	Funding_time         int    `json:"funding_time"`         // The remaining time for the next collection of the funding rate, unit: minute
	Position_amount      string `json:"position_amount"`      // Amount
	Funding_rate_last    string `json:"funding_rate_last"`    // Last funding rate
	Funding_rate_next    string `json:"funding_rate_next"`    // Next funding rate
	Funding_rate_predict string `json:"funding_rate_predict"` // Predicted funding rate
	Last                 string `json:"last"`                 // Latest Price
	Sign_price           string `json:"sign_price"`           // Mark Price
	Index_price          string `json:"index_price"`          // Index Price
	Sell_total           string `json:"sell_total"`           // The number of ask orders in the last 1,000 transactions
	Buy_total            string `json:"buy_total"`            // The number of bid orders in the last 1,000 transactions
	Open                 string `json:"open"`                 // Opening price
	Close                string `json:"close "`               // Closing price
	High                 string `json:"high"`                 // Highest price
	Low                  string `json:"low"`                  // Lowest price
	Vol                  string `json:"vol"`                  // Amount
	Buy                  string `json:"buy"`                  // Bid1 price
	Buy_amount           string `json:"buy_amount"`           // Bid1 amount
	Sell                 string `json:"sell"`                 // Ask1 price
	Sell_amount          string `json:"sell_amount"`          // Ask1 amount
	Date                 int    `json:"date"`                 // Date timestamp
}

func (r *MarketStatus) Parse(raw_response *http.Response) (*MarketStatus, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type AllMarketStatus struct {
	Period               int    `json:"period"`               // Period
	Funding_time         int    `json:"funding_time"`         // The remaining time for the next collection of the funding rate, unit: minute
	Position_amount      string `json:"position_amount"`      // Amount
	Funding_rate_last    string `json:"funding_rate_last"`    // Last funding rate
	Funding_rate_next    string `json:"funding_rate_next"`    // Next funding rate
	Funding_rate_predict string `json:"funding_rate_predict"` // Predicted funding rate
	Last                 string `json:"last"`                 // Latest Price
	Sign_price           string `json:"sign_price"`           // Mark Price
	Index_price          string `json:"index_price"`          // Index Price
	Sell_total           string `json:"sell_total"`           // The number of ask orders in the last 1,000 transactions
	Buy_total            string `json:"buy_total"`            // The number of bid orders in the last 1,000 transactions
	Open                 string `json:"open"`                 // Opening price
	Close                string `json:"close "`               // Closing price
	High                 string `json:"high"`                 // Highest price
	Low                  string `json:"low"`                  // Lowest price
	Vol                  string `json:"vol"`                  // Amount
	Buy                  string `json:"buy"`                  // Bid1 price
	Buy_amount           string `json:"buy_amount"`           // Bid1 amount
	Sell                 string `json:"sell"`                 // Ask1 price
	Sell_amount          string `json:"sell_amount"`          // Ask1 amount
	Date                 int    `json:"date"`                 // Date timestamp
}

func (r *AllMarketStatus) Parse(raw_response *http.Response) (*AllMarketStatus, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketDepthFuture struct {
	Asks0       string `json:"asks00"`      // Ask1 price
	Asks1       string `json:"asks01"`      // Ask1 amount
	Bids00      string `json:"bids00"`      // Bid1 price
	Bids01      string `json:"bids01"`      // Bid1 amount
	Last        string `json:"last"`        // Price
	Sign_price  string `json:"sign_price"`  // Mark Price
	Index_price string `json:"index_price"` // Index Price
}

func (r *MarketDepthFuture) Parse(raw_response *http.Response) (*MarketDepthFuture, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketLatestTransaction struct {
	Id      int    `json:"id"`      // Txid
	Type    string `json:"type"`    // Type, “buy”: buy, “sell”: sell
	Price   string `json:"price"`   // Executed Price
	Amount  string `json:"amount"`  // Amount
	Date    string `json:"date"`    // Transaction time, unit: second
	Data_ms string `json:"data_ms"` // Transaction time, unit: milliseconds
}

func (r *MarketLatestTransaction) Parse(raw_response *http.Response) (*MarketLatestTransaction, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type MarketKLine struct {
	Data0 int    `json:"data0"` // Timestamp
	Data1 string `json:"data1"` // Opening price
	Data2 string `json:"data2"` // Closing price
	Data3 string `json:"data3"` // Highest price
	Data4 string `json:"data4"` // Lowest price
	Data5 string `json:"data5"` // Amount
	Data6 string `json:"data6"` // Value

}

func (r *MarketKLine) Parse(raw_response *http.Response) (*MarketKLine, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}

type UserTransaction struct {
	Offset          int    `json:"offset"`          // Offset
	Limit           int    `json:"limit"`           // Number of query
	Id              int    `json:"id"`              // Transaction ID
	Time            int    `json:"time"`            // Timestamp
	Deal_type       int    `json:"deal_type"`       // Transaction type, 1: open position, 2: add margin, 3: reduce margin, 4: close position, 5: reduce by system, 6: liquidate position, 7: position adl
	Market          string `json:"market"`          //  Market name
	User_id         int    `json:"user_id"`         // User ID
	Deal_user_id    int    `json:"deal_user_id"`    // Counter-party user id
	Order_id        int    `json:"order_id"`        // Order id
	Deal_order_id   int    `json:"deal_order_id"`   // Counter-party order id
	Position_id     int    `json:"position_id"`     // Position id
	Side            int    `json:"side"`            // 1: sell, 2: buy
	Role            int    `json:"role"`            // 1: maker, 2: taker
	Position_type   int    `json:"position_type"`   // Margin type, 1: Isolated, 2: Cross
	Price           string `json:"price"`           // Price
	Open_price      string `json:"open_price"`      // Opening Price
	Settle_price    string `json:"settle_price"`    // Settlement Price
	Amount          string `json:"amount"`          // Amount
	Position_amount string `json:"position_amount"` // Amount
	Margin_amount   string `json:"margin_amount"`   // Position margin after transaction
	Leverage        string `json:"leverage"`        // Margin
	Deal_stock      string `json:"deal_stock"`      // Number of base coin transactions
	Deal_fee        string `json:"deal_fee"`        // Fee
	Deal_margin     string `json:"deal_margin"`     // Trading margin
	Fee_rate        string `json:"fee_rate"`        // Fee rate
	Deal_profit     string `json:"deal_profit"`     // Realized PNL
	Deal_insurance  string `json:"deal_insurance "` // Consumed or increased insurance fund
}

func (r *UserTransaction) Parse(raw_response *http.Response) (*UserTransaction, error) {
	err := json.NewDecoder(raw_response.Body).Decode(r)
	defer raw_response.Body.Close()
	return r, err
}
