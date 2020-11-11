1.package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) 
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}
length of slice 2 capacity 6
After re-slicing length is 6 and capacity is 6

Program exited.
2.package main

import (  
    "fmt"
)

func main() {  
    i := make([]int,8, 8)
    fmt.Println(i)
}
[0 0 0 0 0 0 0 0]

Program exited.
3.package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
[3 5 7]

Program exited.
4.package main

import "fmt"

func main() {
	names := [4]string{
		"sandhya",
		"chinni",
		"raadha",
		"raguvaran",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "chintu"
	fmt.Println(a, b)
	fmt.Println(names)
}
[sandhya chinni raadha raguvaran]
[sandhya chinni] [chinni raadha]
[sandhya chintu] [chintu raadha]
[sandhya chintu raadha raguvaran]
5.package main
import "fmt"
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
[2 3 5 7 11 13]
[true false true true false true]
[{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
