package api

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func RapidAPIGet(word string, print bool) ([]byte, error) {
	url := fmt.Sprintf("https://wordsapiv1.p.rapidapi.com/words/%s", word)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-rapidapi-host", "your_host")
	req.Header.Add("x-rapidapi-key", "your_key")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("unable to fetch data from rapidapi, error: ", err)
		return nil, err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}
