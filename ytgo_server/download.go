package main

import (
	"log"
	"net/http"
	"errors"
	"io"
	"bytes"
)

 
func downloadMultipleFiles(urls []string) ([][]byte, error) {
	done := make(chan []byte, len(urls))
	errch := make(chan error, len(urls))
	for _, URL := range urls {
		go func(URL string) {
			b, err := downloadFile(URL)
			if err != nil {
				errch <- err
				done <- nil
				return
			}
			done <- b
			errch <- nil
		}(URL)
	}
	bytesArray := make([][]byte, 0)
	var errStr string
	for i := 0; i < len(urls); i++ {
		bytesArray = append(bytesArray, <-done)
		if err := <-errch; err != nil {
			errStr = errStr + " " + err.Error()
		}
	}
	var err error
	if errStr!=""{
		err = errors.New(errStr)
	}
	return bytesArray, err
}

func downloadFile(URL string) ([]byte, error) {


	if URL== ""{
		return nil, errors.New("bad url")
	}
	const (
		vi     = "https://i.ytimg.com/vi/"
		resMax = "/maxresdefault.jpg"
		resHQ  = "/hqdefault.jpg"
	)
	resp, err := http.Get(vi+URL+resHQ)
	if err != nil  || resp.StatusCode != http.StatusOK{
		resp, err = http.Get(vi + URL + resMax)
		if err != nil || resp.StatusCode != 200 {
			log.Printf("Response Status Code: %v\n", resp.StatusCode)
			return nil,err
		}
		return nil, err
	}
        defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	var data bytes.Buffer
	_, err = io.Copy(&data, resp.Body)
	if err != nil {
		return nil, err
	}
	return data.Bytes(), nil
}
