package crawler

import (
	"net/http"
	"log"
	"io/ioutil"
	"strconv"
	"code.google.com/p/go.net/html"
	"bytes"
	"regexp"
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


func parseHTML(raw *rawHTML) []string  {
	doc, err := html.Parse(bytes.NewReader(raw.body));
	
	if err != nil {
		log.Fatal(err);
	}

	//initial capacity of ten
	symbols := make([]string, 10);

	var nav func(*html.Node)

	
        nav = func(n *html.Node)  {
		if n.Type == html.ElementNode && n.Data == "td"  {
			func() {
				s := make([]html.Attribute, 2)
				s = n.Attr
				if len(s) == 2 {
					//in this very specific case a td element 
					//with 2 will contain our symbol
					symbols = append(symbols, n.FirstChild.Data)
				}
			} ()
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			nav(c);
		}
	}

	nav(doc)
	return symbols
}



func loopURLS(url string, count int) []string {
	//loops through the URLS, assigning the parsing to goroutines

	//routine for parsing HTML
	rawC := make(chan *rawHTML)
	
	symChan := make(chan []string)

	for i := 0; i < count; i++ {
		fetchURL := url + strconv.Itoa(i)
		go func() {
			rawH := urlFetch(fetchURL);
			rawC <- rawH
		}()
			
		go func() {
			//read from the channel
			raw := <-rawC
			str := parseHTML(raw);
			symChan <- str
		}()
	}
	

	//wait for all routines to finish before returning
	//then collect the results and return the array
	final := make([]string, 10)
	
	//make sure the symbols match
	var symRegex = regexp.MustCompile("^[a-zA-Z]+$")

	for i := 0; i < count; i++ {
		s := <-symChan;
		for _, elem := range s {
			if symRegex.MatchString(elem) {
				final = append(final, elem + ".L") //the .L is added becuase that's 
				//what YQL understands
			}
		}
		
	}

	return final;
}


func FetchAIMSymbols()  []string {
	//Fetch Every Symbol on the London AIM  and return them in an array of strings
	//eseentially works by using a concurrent producers consumers queue

	url := "http://www.londonstockexchange.com/exchange/prices-and-markets/stocks/indices/summary/summary-indices-constituents.html?index=AXX&page="


	return loopURLS(url, 43)
	
}
