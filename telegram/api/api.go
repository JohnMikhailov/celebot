package telegram

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)


func getModel() {
	
}


func LongPolling() {
	url_head := os.Getenv("TELEGRAM_BOT_URL")
	token := os.Getenv("BOTTOKEN")

	for {
		url := url_head + token + "/" +
			"getUpdates" +
			"?timeout=5" +
			"&limit=1"

		fmt.Print(url)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Printf(bodyString)
		}
	}
}
