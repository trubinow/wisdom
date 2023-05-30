package models

type Request struct {
	Type string
}

type Work struct {
	Resource   []byte
	Difficulty int
}
