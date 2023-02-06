package api

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"../utils"
)

func ApiBulk(data string) {
	req, errHttpR := http.NewRequest("POST", "http://localhost:4080/api/_bulk", strings.NewReader(data))
	utils.PrintError(errHttpR)
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, errHttpD := http.DefaultClient.Do(req)
	utils.PrintError(errHttpD)
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, errRead := io.ReadAll(resp.Body)
	utils.PrintError(errRead)
	fmt.Println(string(body))
}
