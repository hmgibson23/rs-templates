package yqlbuilder


func GetSymbol(symbol string) string {
	//returns a url to Fetch the quote from
	url := "http://query.yahooapis.com/v1/public/yql?q=select%20*%20from%20yahoo.finance.quotes%20where%20symbol%20in%20%28%22" + symbol  + "%22%29&env=store://datatables.org/alltableswithkeys&format=json";
	
	return url;

}
