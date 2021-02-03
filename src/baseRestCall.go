package src

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func RestCall() {

	flag.StringVar(&URL, "url", "", "Default URL")
	flag.StringVar(&Symbol, "symbol", "", "Stock symbol")
	flag.StringVar(&QueryFunction, "query", "", "Query to call for stock")
	flag.StringVar(&Key, "key", os.Getenv("KEY"), "API Key used for authentication")
	flag.StringVar(&DataType, "datatype", "json", "Default is JSON or you can"+
		"choose CSV.")
	flag.Parse()

	response, err := http.Get(URL + "/query?function=" + QueryFunction + "&symbol=" +
		Symbol + "&apikey=" + Key + "&datatype=" + DataType)
	if err != nil {
		fmt.Printf("The HTTP request failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
