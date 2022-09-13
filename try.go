package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

var v = vote{"1", "hello", 5, "ok"}

func main1() {
	/*data := []byte("hellojkkoyutyutyut")
	hash := sha256.Sum256(data)
	st := strings.Trim(strings.Replace(fmt.Sprint(hash), " ", "", -1), "[]")
	fmt.Println(st)*/

	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	publicKey := &privateKey.PublicKey

	msg := "hello, world"
	hash := sha256.Sum256([]byte(msg))

	sig, err := ecdsa.SignASN1(rand.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	fmt.Printf("signature: %x\n", sig)
	fmt.Println(hash[:])
	valid := ecdsa.VerifyASN1(publicKey, hash[:], sig)
	fmt.Println("signature verified:", valid)
	st := fmt.Sprintf("%x", hash[:])
	fmt.Println(st)

}
