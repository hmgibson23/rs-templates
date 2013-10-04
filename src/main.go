package main

import (
	"stocktick/fetch"
	"stocktick/stockquote"
)




func main() {
	stock := fetch.FetchSymbol("http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20%28%22MSFT%22%29&env=store://datatables.org/alltableswithkeys&format=json");
	stockquote.PrintStock(stock);
}
