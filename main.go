package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

func main() {
	var i string
	for true {
		triangle, err := readInput()
		if err != nil {
			color.Set(color.BgHiRed)
			fmt.Println(err)
			color.Unset()
		} else {

			tType := triangle.GetTriangleType()
			color.Set(color.FgHiGreen)
			fmt.Printf("Your triange type is: %s\n", tType)
			color.Unset()

		}
		color.Set(color.FgRed)
		fmt.Println("Type \"exit\" to quit. Type anything else to retry.")
		color.Unset()
		fmt.Scanln(&i)
		if strings.ToLower(i) == "exit" {
			fmt.Println("Goodbye!")
			break
		}
	}

}

func readInput() (*Triangle, error) {
	color.Set(color.FgCyan)
	defer color.Unset()
	triangle := new(Triangle)
	values := reflect.ValueOf(triangle).Elem()
	triangleType := values.Type()
	for i := 0; i < values.NumField(); i++ {
		fmt.Printf("Please enter an integer for side %s:\n", triangleType.Field(i).Name)
		var x int64
		_, err := fmt.Scanf("%d\n", &x)
		if err != nil {
			fmt.Println(err)
			fmt.Scanln()
			return nil, errors.New("You entered an inccorect value. Value must be an integer")

		}

		if x < 1 || x > 200 {
			return nil, fmt.Errorf("The value of %s must be inbetween 1 and 200", triangleType.Field(i).Name)
		}

		values.Field(i).SetInt(x)
	}

	return triangle, nil
}
