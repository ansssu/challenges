package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	//create http request
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, "https://image-placeholder.com/images/actual-size/75x75.png", nil)
	if err != nil {
		panic(err)
	}

	//perform http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	imageData, err := ioutil.ReadAll(res.Body)
	fmt.Printf("downloaded image of size %d\n", len(imageData))

}
