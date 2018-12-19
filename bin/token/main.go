package main

import (
	"flag"
	"fmt"

	"github.com/meitu/token"
)

func main() {
	var (
		key     string
		auth    string
		payload string
	)

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
	if auth == "" && payload == "" {
		fmt.Println("Error parameter. Please enter auth or payload")
		return
	}
	if auth != "" && payload != "" {
		fmt.Println("Parameter error, auth and payload cannot exist simultaneously")
		return
	}
	tt := token.New([]byte(key))
	if auth != "" {
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
	} else if payload != "" {
		auth, err := tt.Sign([]byte(payload))
		if err != nil {
			fmt.Printf("Create auth failed %s\n", err)
			return
		}
		fmt.Printf("auth : %s\n", auth)
		return
	}

	return
}
