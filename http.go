package common

import "io/ioutil"
import "net/http"

func SimpleHttpGet(url string) ([]byte, error) {
    resp, err := http.Get(url)
    if err != nil {
	return nil, err
    }

    defer resp.Body.Close()
    return ioutil.ReadAll(resp.Body)
}
