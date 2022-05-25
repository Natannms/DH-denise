package cases

import (
	"fmt"
)

func CaseOne() {
	idade := 45

	switch {
	case idade < 10:
		fmt.Println(" é menor que 10")
	case idade < 50:
		fmt.Println(" é menor que 50")
	case idade < 100:
		fmt.Println(" é menor que 100")
	}
}

func CaseTow() {
	idade := 45

	switch {
	case idade < 10:
		fmt.Println(" é menor que 10")
		fallthrough
	case idade < 50:
		fmt.Println(" é menor que 50")
		fallthrough
	case idade < 100:
		fmt.Println(" é menor que 100")
	}
}

func ErrorCase() {
	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
		// fmt.Println("Not allowed")
	case i < 100:
		fmt.Println("i is less than 100")
	}
}
