package main

import (
	"flag"
	"fmt"

	"github.com/meitu/token"
)

var (
	key     string
	auth    string
	payload string
)

var Usage = func() {
	fmt.Println("USAGE: ./token -key=key,-auth=auth,-payload=payload")
	fmt.Println("\nThe commond are:\n\t-key server key \n\t-auth client auth 。\n\t-payload service payload。")
}

func main() {

	flag.StringVar(&key, "key", "", "server key")
	flag.StringVar(&auth, "auth", "", "client auth")
	flag.StringVar(&payload, "payload", "", "server payload")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	if key == "" {
		fmt.Println("serverkey is empty,Please enter serverkey on the command line")
		return
	}
	tt := token.New([]byte(key))
	if auth != "" && payload == "" {
		res, err := tt.Verify([]byte(auth))
		if err != nil {
			fmt.Printf("Parse failed :%s\n", err)
			return
		}
		if res {
			fmt.Println("Prase sucess")
		} else {
			fmt.Printf("Prase faild")

		}
		return
	} else if auth == "" && payload != "" {
		//Token is generated based on serverkey and namespace
		auth, err := tt.Sign([]byte(payload))
		if err != nil {
			fmt.Printf("Create auth failed %s\n", err)
			return
		}
		fmt.Printf("auth : %s\n", auth)
		return
	} else {
		Usage()
	}
	return
}
