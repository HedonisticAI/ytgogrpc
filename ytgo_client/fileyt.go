package main

import (
	"os"
	"errors"
)
func CreateFolder(pic string) error {

	// create folder if already exist do nothing
	err := os.MkdirAll(pic, os.ModePerm)
	if err != nil {
		return errors.New("Can't create folder")
	}
	return nil
}

func createFile(thumbnailsName string) (*os.File, error) {

	// create file with auto set in the name's last number
	createdFile, err := os.Create(thumbnailsName)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Can't create file")
	}

	return createdFile, nil
}