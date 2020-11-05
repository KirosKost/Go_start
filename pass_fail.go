// pass_fail сообщает, сдал ли пользователь экзамен.

package main

import 	(
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//func main() {
//	fmt.Print("Inter a grade: ")
//	reader := bufio.NewReader(os.Stdin)
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	input = strings.TrimSpace(input)
//	grade, err := strconv.ParseFloat(input, 64)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var status string
//	if grade >= 60 {
//		status = "passing"
//
//	} else {
//		status = "failing"
//	}
//
//	fmt.Println("A grade of", grade, "is", status)
//}
func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, nil
	}
	return number, nil
}
func main() {
	fmt.Print("Enter a grade: ")
	grade, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string

	if grade >= 60 {
		status = "passing"

	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
