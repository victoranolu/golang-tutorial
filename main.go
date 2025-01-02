// package main

// import "fmt"

// func main() {
// 	fmt.Println("Hello Go")
// 	var agency string = "Go Language Programming"
// 	fmt.Println("Welcome to", agency)
// 	name := "Victor"
// 	fmt.Println("Hello,", name, "!")
// 	TotalCars := 50
// 	fmt.Println("The total numbers of cars are", TotalCars, "I am rich and wealthy")

// 	startingPrice := 70.42

// 	fmt.Printf("Our price starts at %v\n", startingPrice)
// 	fmt.Println("Choose from our catalogue")

// 	str := "Supercalafragilisticespialidious"
// 	length := len(str)
// 	fmt.Println(length)
// }

// Variables

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str1 := "Come outside and play"

// 	wIndex := strings.Index(str1, "o")
// 	substr := str1[wIndex:21]
// 	fmt.Println(substr)
// }                                string slicing  with the index on the starting with an number and an ending number

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str1 := "come and be going"
// 	str2 := "do not eat"

// 	strings.EqualFold(str1, str2)
// 	fmt.Println(strings.EqualFold(str1, str2))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str5 := "Hello, Golang"
// 	fmt.Println(strings.Replace(str5, "Golang", "Yummylack", 1))
// }

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str7 := "British Bull Dog"
// 	fmt.Println(strings.ToUpper(str7))
// 	fmt.Println(strings.ToLower(str7))
// }                                                to change strings from uppercae to lowercase

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	str8 := "Towel is on the floor"
// 	fmt.Println(strings.Contains(str8, "police"))
// 	fmt.Println(strings.Contains(str8, "floor"))
// }

package main

import "fmt"

func main() {
	temperatureC := 32.0
	temperatureK := 0.4

	temperatureK = temperatureC + 273.16

	fmt.Println("Temperature in Celcius", temperatureC)
	fmt.Println("Temperature in Kelvin", temperatureK)
}
