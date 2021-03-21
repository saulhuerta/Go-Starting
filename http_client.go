package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main(){

	response, err := http.Get("https://gobyexample.com")

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Response status:", response.Status)

	scanner := bufio.NewScanner(response.Body)

	for i := 0; scanner.Scan() && i < 5; i++{
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}