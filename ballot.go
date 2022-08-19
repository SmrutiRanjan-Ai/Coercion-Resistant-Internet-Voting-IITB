package main

type candidate1 struct {
	option        string
	candidateName string
	encryptedCode string
}

type candidate2 struct {
	candidateName string
	encryptedCode string
}

type choice1 struct {
	index  int
	choice bool
}

type choice2 struct {
	index  int
	option string
	choice bool
}

type ballot1 struct {
	candidateList []candidate1
	options       []choice1
	serial        string
	signature     []byte
}

type ballot2 struct {
	candidateList []candidate2
	options       []choice2
	serial        string
	signature     []byte
}
