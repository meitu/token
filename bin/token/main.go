package main

import (
	"flag"
	"fmt"

	"github.com/meitu/token"
)

func main() {
	var (
		key     string
		sign    string
		payload string
	)

	flag.StringVar(&key, "key", "", "server key")
	flag.StringVar(&sign, "token", "", "client token")
	flag.StringVar(&payload, "payload", "", "server payload")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	if key == "" {
		fmt.Println("key is empty,Please enter key on the command line")
		return
	}
	if sign == "" && payload == "" {
		fmt.Println("Error parameter. Please enter token or payload")
		return
	}
	if sign != "" && payload != "" {
		fmt.Println("Parameter error, token and payload cannot exist simultaneously")
		return
	}
	tt := token.New([]byte(key))
	if sign != "" {
		err := tt.Verify([]byte(sign))
		if err != nil {
			fmt.Printf("Parse failed :%s\n", err)
		} else {
			fmt.Println("Prase sucess")
		}
	} else if payload != "" {
		token, err := tt.Sign([]byte(payload))
		if err != nil {
			fmt.Printf("Create token failed %s\n", err)
			return
		}
		fmt.Printf("token : %s\n", token)
		return
	}
	return
}
