package models

type Request struct {
	Type string
}

type Work struct {
	Resource   []byte
	Difficulty int
}

type Proof struct {
	Hash  string
	Nonce string
}
