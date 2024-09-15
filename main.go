package main

import (
	"fmt"
	"os"
)

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		//return "", errors.New("empty content") // erreur simple
		return "", fmt.Errorf("empty content (filename=\"%v\")", filename) // erreur avec formatage pour ajouter des informations supplémentaires tel que le nom du fichier
	}

	return string(data), nil
}

func main() {
	dat, err := readFile("./test.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("Content:")
	fmt.Println(dat)
}
