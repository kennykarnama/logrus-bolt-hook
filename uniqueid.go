package boltlogrus

import uuid "github.com/satori/go.uuid"

//UniqueID is an interface to provide
//ID generation
//It is used as key in boltdb
type UniqueID interface {
	GenerateID() (string, error)
}

type satoru struct{}

//NewSatoru will return new empty satoru struct
func NewSatoru() UniqueID {
	return &satoru{}
}

func (s *satoru) GenerateID() (string, error) {
	id := uuid.NewV4()
	return id.String(), nil
}
