package fetch



import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
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
		//panic(err);
	}

	stock := &stockquote.StockList{};

	error := json.Unmarshal(body, &stock);
	
	if error != nil {
		//panic(error);
	}

	m := stock.Query.Results.Quote;
	quote := stockquote.NewStockQuote(m);

	return quote;
}

func FetchSymbols(symbols []string) {
	//make it concurrent
	length := len(symbols)
	sem := make(chan bool, length)

	for _, elem := range symbols {
		go func(elem string) {
			fmt.Println(elem);
			q := FetchSymbol(elem)
			stockquote.PrintStock(q)
//			stockquote.SaveToDB(q);
			sem <- true;
		}(elem)
	}
	
	for i := 0; i < length; i++ { <-sem }
}
