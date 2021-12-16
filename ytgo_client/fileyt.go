package main

import (
	"os"
	"errors"
	"log"
	"math/rand"
	"time"
)
func CreateFolder(pic string) error {

	// create folder if already exist do nothing
	err := os.MkdirAll(pic, os.ModePerm)
	if err != nil {
		return errors.New("Can't create folder")
	}
	return nil
}

func t_rand() string {
	var letter string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 10)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}

func CreateFile(Name string) (*os.File, error) {

	createdFile, err := os.Create(Name)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Can't create file")
	}

	return createdFile, nil
}


