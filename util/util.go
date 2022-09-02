package util

import (
	"errors"
	"log"
)

func HandleErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func ValidateID(src, dest int) error {
	if dest < src {
		return errors.New("invalid route")
	} else if src > 14 || src < 1 || dest > 14 || dest < 1 {
		return errors.New("invalid source/desination")
	}
	return nil
}
