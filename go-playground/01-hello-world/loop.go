package main

import "fmt"

func lesson3() {
	i := 0
	c := 0
	s := 0

	for {
		i++

		if i == 10 {
			c++
			fmt.Println(c)
			i = 0
		}

		if c == 3 {
			fmt.Println("bow")
			c = 0
			s++
		}

		if s == 3 {
			break
		}
	}
}
