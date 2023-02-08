package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Fetch(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("wrong status code %d",
			res.StatusCode)
	}
	html, ok := ioutil.ReadAll(res.Body)

	return string(html), ok
}
