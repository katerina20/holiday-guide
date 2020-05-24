package http

import (
	"fmt"
	"net/http"
)

func RequestApi() {
	resp, err := http.Get("https://date.nager.at/api/v2/publicholidays/2020/UA")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
}
