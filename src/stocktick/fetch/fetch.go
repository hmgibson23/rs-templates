package fetch



import (
	"net/http"
	"io/ioutil"
	"log"
	"stocktick/stockquote"
	"encoding/json"
	"yql/yqlbuilder"
)


func FetchSymbol(symbol string) *stockquote.StockQuote {

	url := yqlbuilder.GetSymbol(symbol);
	resp,err := http.Get(url);
	
	if err != nil {
		//failed to fetch stream
		log.Fatal(err);
	}

	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	
	if err != nil {
		log.Fatal(err);
	}

	stock := &stockquote.StockList{};

	error := json.Unmarshal(body, &stock);
	
	if error != nil {
		log.Fatal(error);
	}

	m := stock.Query.Results.Quote;
	quote := stockquote.NewStockQuote(m);

	return quote;
}

