package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	serverkey string //server key
	token     string // appkey
	namespace string //namespace
)

var Usage = func() {
	fmt.Println("USAGE: ./main -s=serverkey,-t=token,-n=namespace")
	fmt.Println("\nThe commond are:\n\t-s   server key \n\t-t client token 。\n\t-n service namespace。")
}

func main() {

	flag.StringVar(&serverkey, "serverkey", "", "server key")
	flag.StringVar(&token, "token", "", "client token")
	flag.StringVar(&namespace, "namespace", "", "server namespace")
	flag.Parse()
	args := os.Args[1:]

	if len(args) == 0 {
		Usage()
		return
	}
	if serverkey == "" {
		fmt.Println("serverkey is empty,Please enter serverkey on the command line")
		return
	}
	//如果username不为空
	if token != "" && namespace == "" {
		ns, err := Verify([]byte(token), []byte(serverkey))
		if err != nil {
			fmt.Printf("Parse failed :%s\n", err)
			return
		}
		fmt.Println("Prase sucess")
		fmt.Printf("namespace: %s\n", ns)
		return
	} else if token == "" && namespace != "" {
		token, err := Token([]byte(serverkey), []byte(namespace), time.Now().Unix())
		if err != nil {
			fmt.Printf("Create token failed %s\n", err)
			return
		}
		fmt.Printf("token : %s\n", token)
		return
	} else if token != "" && namespace != "" {
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
