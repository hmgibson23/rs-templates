package crawler

import (
	"net/http"
	"fmt"
	"log"
	"io/ioutil"
	"code.google.com/p/go.net/html"
	"bytes"
)

type rawHTML struct {
	body []byte
}

func urlFetch(url string) *rawHTML {
	
	resp, err := http.Get(url);
	
	if err != nil {
		log.Fatal(err);
	}
	
	defer resp.Body.Close();
	body, err := ioutil.ReadAll(resp.Body);
	
	if err != nil {
		log.Fatal(err);
	}

	return &rawHTML {body};
}


func navigateTree(node *html.Node) {
	if node.FirstChild == nil && node.NextSibling == nil {
		fmt.Println("End of tree...");
		return;
	} else {
		navigateTree(node.FirstChild);
	}
}

func parseHTML(raw *rawHTML)  {
	fmt.Println("Got given some html...");
	doc, err := html.Parse(bytes.NewReader(raw.body));

	if err != nil {
		log.Fatal(err);
	}
	

	//parse all the nodes
	fmt.Printf("%+v", doc.LastChild.LastChild.FirstChild);
	bodyNode := doc.LastChild.LastChild.FirstChild;
	navigateTree(bodyNode);
	//this needs to be done recursively
}



func loopURLS(url string, count int) {
	//loops through the URLS, assigning the parsing to goroutines

	//written to when each routine has read the HTML
	quit := make(chan bool)
	//routine for parsing HTML
	rawC := make(chan *rawHTML)
	
	//channel for the final parsed strings
	//symbols := make(chan []string)

	for i := 0; i < count; i++ {
		fetchURL := url //+ string(i)
		
		go func() {
			rawH := urlFetch(fetchURL);
			rawC <- rawH
		}()
			
		go func() {
			//read from the channel
			raw := <-rawC
			parseHTML(raw);
			quit <- true
		}()
	}
	

	//wait for all routines to finish before returning
	for i := 0; i < count; i++ {
		<-quit;
	}

	//once it's all done collect all the string results and merge them
	
}


//return nothing while testing
func FetchAIMSymbols()  {
	//Fetch Every Symbol on the London AIM Stock Exchange and return them in an array of strings
	//eseentially works by using a concurrent producers consumers queue

	url := "http://www.londonstockexchange.com/exchange/prices-and-markets/stocks/indices/summary/summary-indices-constituents.html?index=AXX&page=1"

	loopURLS(url, 1)
	
}
