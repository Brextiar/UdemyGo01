package main

import (
	"UdemyGo01/data"
	"fmt"
	"math/rand"
	"strings"
)

func printInfoNoParam() {
	fmt.Println("Hello Go")
}
func printInfoParams(name string, age int, email string) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Email:", email)
}

func avg(x int, y float64) float64 {
	return (float64(x) + y) / 2
}

func sumNamedReturn(x int, y int) (result int) {
	result = x + y
	return result
}

func TolowerStr(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

// On peu déclarer des variable en dehors de la fonction main mais on ne peu plus utiliser la syntaxe raccourcie :=

func main() {
	var age int   // age = 0
	var age2 = 10 // age2 = 10
	fmt.Println("Age:", age)
	fmt.Println("Age2:", age2)
	/*
			Déclaration de plusieurs variables
			var weight, height int
			print(weight, height) -> 0 0
		------------------------------------------------
			var weight, height int = 70, 180 -> weight = 70, height = 180
			print(weight, height) -> 70 180
	*/
	// inférence de type
	var name = "John"
	fmt.Println("Name:", name) // Name: John

	// := -> déclaration et affectation
	email := "bob@go.org"
	fmt.Println("Email:", email) // Email: bob@go.org
	// email := "" error: no new variables on left side of :=

	/*
		visibilité des variable en Go
		première lettre de la variable en minuscule -> variable privée
		première lettre de la variable en majuscule -> variable publique
	*/
	fmt.Println(data.Name, data.Age, data.Password) // Fix this line
	//fmt.Println(data.pasword) -> error: data.pasword undefined (cannot refer to unexported field or method pasword)

	// day of a mounth [1-31]
	day := rand.Int() % 31

	if day < 15 {
		fmt.Printf("we're in the first half of the month (day = %d)", day)
	} else if day == 18 {
		fmt.Printf("We're in the special day (day = %d)", day)
	} else {
		fmt.Printf("We're in the second half of the month (day = %d)", day)
	}

	printInfoNoParam()
	printInfoParams("John", 30, "test@test.fr")
	fmt.Println(avg(10, 5.5))
	fmt.Println(sumNamedReturn(10, 5))

	//utilisation d'un fonction à retour multiple
	lowerName, length := TolowerStr("ALICE")
	fmt.Println("ALICE -> ", lowerName, length)
	// '_' permet d'ignorer un retour
	name2, _ := TolowerStr("BOB")
	fmt.Println("BOB -> ", name2)
	// println(nam3, _ := TolowerStr("CHARLIE")) // error: nam3 undefined

	// ARRAY
	var nameArr [3]string
	fmt.Printf("name: %v (len=%v)\n", nameArr, len(name))
	nameArr[0] = "Alice"
	nameArr[2] = "Bob"
	fmt.Printf("name: %v (len=%v)\n", nameArr, len(name))
	fmt.Println("nameArr[0] -> ", nameArr[0], "\nnameArr[2] -> ", nameArr[2])

	//Slice

	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -1
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, 64)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, -8)
	fmt.Printf("%v (len=%d, cap=%d)\n", nums, len(nums), cap(nums))
	fmt.Println("------------------------------")
	letters := []string{"a", "b", "c", "d", "e"}
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	letters = append(letters, "f")
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Println("-----------------------------")
	sub1 := letters[0:2]
	sub2 := letters[2:]
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	// Attention, les slices sub1 et sub2 pointent vers le même tableau sous-jacent
	// Si on modifie sub2, on modifie aussi letters
	// Si subs2 commence avant la fin de sub1, alors sub1 est aussi modifié si la valeur modifié est dans la zone de sub1
	fmt.Println("-----------------------------")
	sub2[0] = "z"
	fmt.Printf("%v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("%v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("-----------------------------")
	//copier un slice
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[0] = "DOWN"
	fmt.Printf("COPY %v (len=%d, cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	fmt.Printf("letters %v (len=%d, cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("sub1 %v (len=%d, cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2 %v (len=%d, cap=%d)\n", sub2, len(sub2), cap(sub2))
	fmt.Println("-----------------------------")
	fmt.Println("----------BOUCLE FOR---------")
	//BOUCLE FOR
	sum := 0
	for i := 0; i <= 10; i++ {
		fmt.Println("Loop:", i)
		sum += i
	}
	fmt.Printf("Sum result: %v\n", sum)
	fmt.Println("-----------------------------")
	fmt.Println("----BOUCLE FOR-(WHILE)-------")
	eventCnt := 0
	for eventCnt < 3 {
		fmt.Println("Retrievings events...")
		eventCnt++
		if eventCnt == 3 {
			fmt.Printf("Got %d notifications\n", eventCnt)
		}
	}
	fmt.Println("-----------------------------")
	fmt.Println("----BOUCLE FOR-(INFINITE)----")
	j := 0
	for {
		j++
		if j%2 != 0 {
			fmt.Println(j, "Skipping...")
			continue // passe à l'itération suivante
		}
		fmt.Println(j, "Looping...")
		if j >= 10 {
			fmt.Println("Breaking...")
			break // sort de la boucle
		}
	}
	fmt.Println("-----------------------------")
	fmt.Println("------------RANGE------------") // Permet d'itérer sur une structure de données
	surname := []string{"Dupont", "Durand", "Duchemin", "Duchamp"}
	for i, s := range surname {
		fmt.Printf("Surnname=%s (index=%d)\n", s, i)
	}

	for _, c := range "Hello" {
		fmt.Printf("%v\n", string(c)) // %v -> valeur en byte du character on cast pour afficher la lettre / %c -> affiche le character
	}
}
