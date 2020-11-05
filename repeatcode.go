//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var width, height, area float64
//	width = 4.2
//	height = 3.0
//	area = width * height
//	//fmt.Println(area/10.0, "liters needed") до Printf
//	fmt.Printf("%.2f liters need\n", area/10.0)
//
//	width = 5.2
//	height = 3.5
//	area = width * height
//	//fmt.Println(area/10.0, "liters needed") до Printf
//	fmt.Printf("%.2f liters need\n", area/10.0)
//
//}

//package main
//
//import (
//	"fmt"
//)
//
//func painNeeded(width float64, heigth float64) {
//	area := width * heigth
//	fmt.Printf("%.2f liters need\n", area/10.0)
//}
//
//func main() {
//	painNeeded(4.2, 3.0)
//	painNeeded(5.2, 3.5)
//	painNeeded(5.0, 3.3)
//}

package main

import (
	"fmt"
)

func painNeeded(width float64, height float64) float64 {
	area := width * height
	return area / 10.0
}

func main() {
	var amount, total float64
	amount = painNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = painNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	fmt.Printf("Total: %0.2f liters\n", total)
}
