package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const timoutSec = 5

var clientConf = &http.Client{Timeout: timoutSec * time.Second}

func RequestApi(url string) []byte {
	r, err := clientConf.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	content, _ := ioutil.ReadAll(r.Body)
	return content
}
