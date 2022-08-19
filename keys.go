package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	rand2 "crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func publicInit() (*ecdsa.PublicKey, *ecdsa.PrivateKey) {
	var priv2 *ecdsa.PrivateKey
	var pub2 *ecdsa.PublicKey
	dataPrivate, err := ioutil.ReadFile("private.txt")
	if err != nil {
		fmt.Println("This is")
		filePrivate, err := os.Create("private.txt")
		filePublic, err := os.Create("public.txt")
		if err != nil {
			log.Fatal(err)
		}
		privateKey, _ = ecdsa.GenerateKey(elliptic.P384(), rand2.Reader)
		publicKey = &privateKey.PublicKey

		encPriv, encPub := encode1(privateKey, publicKey)
		_, err = filePublic.WriteString(encPub)
		_, err = filePrivate.WriteString(encPriv)
		defer func(filePublic *os.File) {
			err := filePublic.Close()
			if err != nil {

			}
		}(filePublic)
		defer func(filePrivate *os.File) {
			err := filePrivate.Close()
			if err != nil {

			}
		}(filePrivate)

		if !reflect.DeepEqual(privateKey, priv2) {
			fmt.Println("Private keys do not match.")
		}
		if !reflect.DeepEqual(publicKey, pub2) {
			fmt.Println("Public keys do not match.")
		}
	} else {
		dataPublic, err := ioutil.ReadFile("public.txt")

		if err != nil {
			log.Fatalf("failed writing to file: %s", err)
		}

		priv2, pub2 = decode1(string(dataPrivate), string(dataPublic))

		fmt.Println("ok", priv2, pub2)

	}
	return pub2, priv2

}

func encode1(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncoded), string(pemEncodedPub)
}

func decode1(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	block, _ := pem.Decode([]byte(pemEncoded))
	x509Encoded := block.Bytes
	privateKey, _ := x509.ParseECPrivateKey(x509Encoded)

	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return privateKey, publicKey
}

func vid(id string) string {
	/*data := []byte(id)
	hash := sha256.Sum256(data)
	st := strings.Trim(strings.Replace(fmt.Sprint(hash), " ", "", -1), "[]")
	int2, err := strconv.ParseInt(st, 2, 32)
	if err != nil {
		fmt.Println((err))
	}
	rand.Seed(int2)
	fmt.Println(rand.Int63())*/
	sig := digitalSign(privateKey, id)
	hash := sha256.Sum256(sig)
	fmt.Println(hash)
	st := fmt.Sprintf("%x", hash[:])
	return st

}

func digitalSign(privateKey *ecdsa.PrivateKey, message string) []byte {
	hash := sha256.Sum256([]byte(message))
	sig, err := ecdsa.SignASN1(rand2.Reader, privateKey, hash[:])
	if err != nil {
		panic(err)
	}
	print(message, sig)
	return sig
}

func verifySign(key *ecdsa.PublicKey, msg string, sig []byte) bool {
	hash := sha256.Sum256([]byte(msg))
	valid := ecdsa.VerifyASN1(key, hash[:], sig)
	return valid
}

func uuidGen() string {
	id := uuid.New()
	return id.String()
}
