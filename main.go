package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	serverkey string
	token     string
	namespace string
)

var Usage = func() {
	fmt.Println("USAGE: ./token -serverkey=serverkey,-token=token,-namespace=namespace")
	fmt.Println("\nThe commond are:\n\t-serverkey server key \n\t-token client token 。\n\t-namespac service namespace。")
}

func main() {

	flag.StringVar(&serverkey, "serverkey", "", "server key")
	flag.StringVar(&token, "token", "", "client token")
	flag.StringVar(&namespace, "namespace", "", "server namespace")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}
	if serverkey == "" {
		fmt.Println("serverkey is empty,Please enter serverkey on the command line")
		return
	}
	if token != "" && namespace == "" {
		//namespace is resolved from token
		ns, err := Verify([]byte(token), []byte(serverkey))
		if err != nil {
			fmt.Printf("Parse failed :%s\n", err)
			return
		}
		fmt.Println("Prase sucess")
		fmt.Printf("namespace: %s\n", ns)
		return
	} else if token == "" && namespace != "" {
		//Token is generated based on serverkey and namespace
		token, err := Token([]byte(serverkey), []byte(namespace), time.Now().Unix())
		if err != nil {
			fmt.Printf("Create token failed %s\n", err)
			return
		}
		fmt.Printf("token : %s\n", token)
		return
	} else if token != "" && namespace != "" {
		//Verify that the token passed in matches the namespace
		ns, err := Verify([]byte(token), []byte(serverkey))
		if err != nil {
			fmt.Printf("Parse failed :%s\n", err)
			return
		}
		tk, err := Token([]byte(serverkey), []byte(namespace), time.Now().Unix())
		if err != nil {
			fmt.Printf("Create token failed: %s\n", err)
			return
		}
		if string(ns) == namespace && string(tk) == token {
			fmt.Println("Auth success")
		} else {
			fmt.Println("Auth failed")
		}
		return
	} else {
		Usage()
	}
	return
}
