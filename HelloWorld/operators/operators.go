package operators

import "fmt"

func doOperator() {
	var a float32 = 12.1
	var b float32 = 13.2

	sum := a / b

	fmt.Printf("%v\n", sum)
}
