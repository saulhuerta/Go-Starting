package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "abc123!?$*&()'-=@~"
	strEnc := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(strEnc)

	strDec, _ := base64.StdEncoding.DecodeString(strEnc)
	fmt.Print(string(strDec))
}