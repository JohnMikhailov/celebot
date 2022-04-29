package main

import (
	//"os"
	//"io/ioutil"
	//"net/http"

	"fmt"

	"github.com/meehighlov/celebot/telegram"
	"github.com/meehighlov/celebot/app"
)


func main() {
	fmt.Printf("start polling")
	token := app.GetConfig().BOTTOKEN
	telegram.StartPolling(token)
}
