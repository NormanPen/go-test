package composite

import "fmt"

func DemoCollections() {
	primesArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(primesArray)

	// primeSlice := []int{1,2,3,4,5,6,7,8}
	primesSlice := make([]int, 5, 7)
	primesSlice = append(primesSlice, 2)
	primesSlice = append(primesSlice, 3)
	primesSlice = append(primesSlice, 5)
	fmt.Println(primesSlice)
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))

	for index, value := range primesSlice {
		fmt.Println(index, value)
	}

	primeSlice2 := []string{"Hallo", "Welt"}

	for _, value := range primeSlice2 { //index kann durch Unterstrich weggelassen werden
		fmt.Print(" " + value)
		fmt.Println(" ")
	}

	points := map[string]Point{
		"A": *NewPoint(3, 7),
		"B": *NewPoint(4, 8),
	}
	fmt.Println(points)
	somePoint, ok := points["D"]
	fmt.Println(somePoint, ok)

	points["E"] = *NewPoint(1, 1)
	delete(points, "A")
	fmt.Println(points)

}
