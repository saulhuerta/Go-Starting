package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main(){
	s := "Digest this string"

	hash_sha1 := sha1.New()
	hash_sha1.Write([]byte(s))

	hash_sha256 := sha256.New()
	hash_sha256.Write([]byte(s))

	hash_sha512 := sha512.New()
	hash_sha512.Write([]byte(s))

	fmt.Println(s)
	fmt.Printf("SHA1:   %x\n",  hash_sha1.Sum(nil))
	fmt.Printf("SHA256: %x\n",  hash_sha256.Sum(nil))
	fmt.Printf("SHA512: %x\n",  hash_sha512.Sum(nil))

}
