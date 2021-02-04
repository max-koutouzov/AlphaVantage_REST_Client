package src

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func flags() {
	flag.StringVar(&URL, "url", "",
		"Default URL")
	flag.StringVar(&Symbol, "symbol", "",
		"Stock symbol")
	flag.StringVar(&QueryFunction, "query", "",
		"Query to call for stock")
	flag.StringVar(&Key, "key", os.Getenv("KEY"),
		"API Key used for authentication")
	flag.StringVar(&FileName, "file", "output.json",
		"Add filename to export output of call into")
	flag.StringVar(&DataType, "datatype", "json",
		"Default is JSON or you can choose CSV.")
	flag.Parse()
}

func RestCall() {
	flags()
	response, err := http.Get(URL + "/query?function=" + QueryFunction + "&symbol=" +
		Symbol + "&apikey=" + Key + "&datatype=" + DataType)
	if err != nil {
		fmt.Printf("The HTTP request failed %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		if FileName != "" {
			f, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Printf("Could not create file %s\n", err)
			}
			defer f.Close()
			if _, err := f.Write(data); err != nil {
				log.Fatalln(err)
			}
		}
	}
}
