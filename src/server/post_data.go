package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//PostData POST数据到服务器
func PostData(url string, jsondata []byte) (err error) {
	var req *http.Request
	var resp *http.Response
	if req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsondata)); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	if resp, err = client.Do(req); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return
}
