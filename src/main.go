package main

import (
	"stocktick/fetch"
	"stocktick/stockquote"
	"crawler"
	"fmt"
	"os"
	"bufio"
)




func main() {
	
	fmt.Println("Enter a symbol to get quote:");
	bio := bufio.NewReader(os.Stdin);
	line, _, _ := bio.ReadLine();
	stock := fetch.FetchSymbol(string(line));
	stockquote.PrintStock(stock);
	
	
	fmt.Println("Getting AIM Symbols...");
	crawler.FetchAIMSymbols();
}
