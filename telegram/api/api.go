package telegram

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"

	"encoding/json"
	"celebot/telegram/models"
)


func getUpdates(url string) telegram.UpdateResponse {
	res := telegram.UpdateResponse {}

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
		json.Unmarshal([]byte(bodyString), &res)
	}

	return res
}


func LongPolling() {
	url_head := os.Getenv("TELEGRAM_BOT_URL")
	token := os.Getenv("BOTTOKEN")

	for {
		url := url_head + token + "/" +
			"getUpdates" +
			"?timeout=5" +
			"&limit=1"

		updates := getUpdates(url)

		fmt.Print(updates.Result[0].MessageInfo.Text)
	}
}
