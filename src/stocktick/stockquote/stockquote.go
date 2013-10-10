package stockquote

import (
	"fmt"
	"strconv"
	//"labix.org/v2/mgo"
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
	bookValue float32
	ask float32
	EPS float32
	EBITDA string
	PEratio float32
	EPSEstimateCurrentYear float32
	
}

func floatParse(i float64, err error) float32 {
	return float32(i);
}

func NewStockQuote(m map[string] string) *StockQuote {
	sq := new(StockQuote)
	sq.symbol = m["symbol"]
	sq.ask = floatParse(strconv.ParseFloat(m["Ask"], 32))
	sq.bookValue = floatParse(strconv.ParseFloat(m["BookValue"], 32))
	sq.EPS = floatParse(strconv.ParseFloat(m["EarningsShare"], 32))
	sq.EBITDA = m["EBITDA"]
	sq.EPSEstimateCurrentYear = floatParse(strconv.ParseFloat(m["EPSEstimateCurrentYear"], 32))
	sq.PEratio = floatParse(strconv.ParseFloat(m["PERatio"], 32))
	return sq
}

func PrintStock(s *StockQuote) {
	fmt.Printf("Symbol: %s\n", s.symbol);
	fmt.Printf("Ask: %v\n", s.ask);
	fmt.Printf("Book Value: %v\n", s.bookValue);
	fmt.Printf("EPS: %v\n", s.EPS);
	fmt.Printf("EPS Estimate this year: %v\n", s.EPSEstimateCurrentYear)
	fmt.Printf("EBITDA: %v\n", s.EBITDA);
	fmt.Printf("P/E Ratio: %v\n", s.PEratio);
}
/*
//can't use on windows
func SaveToDB(s *StockQuote) {
	//connect to the db
	session, err := mgo.Dial("localhost");

	if err != nil {
		panic(err);
	}
	defer session.Close();
	
	c := session.DB("stocktick").C("people")
	err = c.Insert(s);

	if err != nil {
		panic(err);
	}
}
*/
func FetchFromDB(symbol string) {
}

