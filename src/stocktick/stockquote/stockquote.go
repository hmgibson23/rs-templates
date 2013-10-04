package stockquote

import (
	"fmt"
	"strconv"
)


type StockList struct {
	Query struct {
		Count uint
		Lang string
		Created string
		Results struct {
			Quote map[string] string 
		}
	}
}


 

/*
 * Struct to represent stock price information
 */
type StockQuote struct {
	symbol string
	bookValue float64
	EPS float64
	EBITDA float64
	PEratio float64
}

func NewStockQuote(m map[string] string) *StockQuote {
	sq := new(StockQuote)
	sq.symbol = m["symbol"]
	sq.bookValue, _ = strconv.ParseFloat(m["BookValue"], 32)
	sq.EPS, _ = strconv.ParseFloat(m["EarningsShare"], 32)
	sq.EBITDA, _ = strconv.ParseFloat(m["EBITDA"], 32)
	sq.PEratio, _ = strconv.ParseFloat(m["PERatio"], 32)
	return sq
}


func PrintStock(s *StockQuote) {
	fmt.Printf("Symbol: %s\n", s.symbol);
	fmt.Printf("Book Value: %f\n", s.bookValue);
	fmt.Printf("EPS: %f\n", s.EPS);
	fmt.Printf("EBITDA: %f\n", s.EBITDA);
	fmt.Printf("P/E Ratio: %f\n", s.PEratio);
}
