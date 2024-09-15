package main

//##########################################
//##################ERRORS##################
//##########################################
//
//import (
//	"fmt"
//	"os"
//)
//
//func readFile(filename string) (string, error) {
//	data, err := os.ReadFile(filename)
//	if err != nil {
//		return "", err
//	}
//
//	if len(data) == 0 {
//		//return "", errors.New("empty content") // erreur simple
//		return "", fmt.Errorf("empty content (filename=\"%v\")", filename) // erreur avec formatage pour ajouter des informations supplémentaires tel que le nom du fichier
//	}
//
//	return string(data), nil
//}
//
//func main() {
//	dat, err := readFile("./test.txt")
//	if err != nil {
//		fmt.Printf("Error while reading file: %v\n", err)
//		return
//	}
//	fmt.Println("Content:")
//	fmt.Println(dat)
//}

// ##########################################
// ##################DEFER###################
// ##########################################
// permet de reporter l'exécution d'une fonction jusqu'à la fin de la fonction courante

import (
	"fmt"
)

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

func main() {
	//start()
	//defer finish()
	//
	//names := []string{"Bob", "Alice", "Bobette", "John"}
	//for _, n := range names {
	//	fmt.Printf("Hello %v\n", n)
	//  fmt.Printf("Goodbye %v\n", n)
	//}

	start()
	defer finish() // defer fonctionne en LIFO (Last In First Out)
	// ordre d'ajout dans la pile du defer
	// 1. finish()          5
	// 2. Goodbye Bob     ^ 4
	// 3. Goodbye Alice   ^ 3
	// 4. Goodbye Bobette ^ 2
	// 5. Goodbye John    ^ 1
	//                    ordre d'éxécution de la pile du defer

	names := []string{"Bob", "Alice", "Bobette", "John"}
	for _, n := range names {
		fmt.Printf("Hello %v\n", n)
		defer fmt.Printf("Goodbye %v\n", n) // ici on ajoute chaque Printf dans la pile du defer de la fonction courante
	}

}
