package fetch



import (
	"net/http"
	"io/ioutil"
	"log"
	"stocktick/stockquote"
	"encoding/json"
)


func FetchSymbol(url string) *stockquote.StockQuote {

	resp,err := http.Get(url);

	if err != nil {
		//failed to fetch stream
		log.Fatal(err);
	}

	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);

	stock := &stockquote.StockList{};

	error := json.Unmarshal(body, &stock);
	
	if error != nil {
		log.Fatal(err);
	}

	m := stock.Query.Results.Quote;
	quote := stockquote.NewStockQuote(m);

	return quote;
}

