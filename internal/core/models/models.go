package models

import "time"

type PowTask struct {
	Text      string
	Salt      []byte
	Iteration uint32
	Memory    uint32
	Hash      []byte
	CreatedAt time.Time
}

type PowResult struct {
	Hash []byte
}

type Passage struct {
	Text string
}
