package syntax

import "fmt"

func Map() {
	myMap := map[string]string{
		"France":  "Paris",
		"Belgium": "Brussels",
		"Spain":   "Madrid",
	}
	fmt.Println("Capitals are: ", myMap)

	myMapFromMake := make(map[string]string, 0)
	fmt.Println("An empty map: ", myMapFromMake)
}

func MapOperations() {
	myMap := map[string]string{
		"France": "Paris",
	}

	myMap["Belgium"] = "Brussels" // Assign a new element
	myMap["France"] = "Lille"     // Replace an element's value

	franceCapital := myMap["France"] // Accessing to an element
	fmt.Println("The new capital in France is: ", franceCapital)

	switzerlandCapital, found := myMap["Switzerland"] // Access to an element and check existence
	if found {
		fmt.Println("The capital of Switzerland is: ", switzerlandCapital)
	} else {
		fmt.Println("Looks like we don't know the capital of Switzerland")
	}

	_, found = myMap["Belgium"] // Check existence
	if found {
		fmt.Println("Looks like we know the capital of Belgium!")
	}
}

func MapIteration() {
	myMap := map[string]string{
		"France":  "Paris",
		"Belgium": "Brussels",
	}

	for country := range myMap {
		fmt.Printf("The capital of %s is %s\n", country, myMap[country])
	}

	for country, capital := range myMap {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	for _, capital := range myMap {
		fmt.Println("Oh! I know this capital: ", capital)
	}
}
