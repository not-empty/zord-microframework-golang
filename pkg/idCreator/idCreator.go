package idCreator

import "github.com/oklog/ulid/v2"

type IdCreator struct {
}

func NewIdCreator() *IdCreator {
	return &IdCreator{}
}

func (*IdCreator) Create() string {
	return ulid.Make().String()
}
