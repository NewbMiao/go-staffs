package main

import (
	"fmt"
	"github.com/newbmiao/go-staffs/jwk/utils"
	"io/ioutil"
	"net/http"
)


func main() {
	var url = "http://localhost:3010/api/private"
	token := utils.GetToken()
	requestPrivate(token, url)
	url = "http://localhost:3010/api/private-scoped"
	requestPrivate(token, url)

	//go-auth0
	utils.GetPublicKey()

}


func requestPrivate(token, url string) {
	fmt.Println("request ", url)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", "Bearer "+token)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	//fmt.Println(res)
	fmt.Println(string(body))
}

