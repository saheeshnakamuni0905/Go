package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 23, 45, 89
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum ", intSum)

	f1, f2, f3 := 23.99, 45.777, 89.909
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum ", floatSum)

	floatSum = math.Round(floatSum)
	fmt.Println("The sum is now", floatSum)

	floatSum = math.Round(floatSum*100) / 100 //safest way to round digits
	fmt.Println("The sum is now", floatSum)

	circleRadius := 5.9
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

}
