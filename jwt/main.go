package main

import (
	"fmt"
	"github.com/newbmiao/go-staffs/jwt/hmac"
	"github.com/newbmiao/go-staffs/jwt/rsa"
)

func main() {
	fmt.Println("======hmac example=======")
	hmac.ExampleNew_hmac()
	hmac.ExampleParse_hmac()

	fmt.Println("=======rsa example=======")
	rsa.Example_getTokenViaHTTP()
	rsa.Example_useTokenViaHTTP()
}
