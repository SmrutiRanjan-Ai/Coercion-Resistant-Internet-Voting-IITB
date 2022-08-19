package main

import "crypto/ecdsa"

const (
	host     = "localhost"
	port     = 5432
	user1    = "postgres"
	password = "213050077"
	dbname   = "MyDB"
)

var voters []voter
var publicKey *ecdsa.PublicKey
var privateKey *ecdsa.PrivateKey

type candidate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var candidates = []candidate{
	{ID: uuidGen(), Name: "Narendra Modi 1"},
	{ID: uuidGen(), Name: "Narendra Modi 2"},
}

type voter struct {
	ID                string `json:"id"`
	VoterID           string
	BallotIndex       int
	ConfirmationToken string
}
