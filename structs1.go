1.package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
Hello World
[Hello World]
[2 3 5 7 11 13]
2.package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "sandhya"
	a[1] = "rani"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{3, 4, 5, 7, 11, 15}
	fmt.Println(primes)
}
sandhya rani
[sandhya rani]
[3 4 5 7 11 15]
3.package main

import "fmt"

func main() {
   var n [10]int /* n is an array of 10 integers */
   var i,j int

   /* initialize elements of array n to 0 */         
   for i = 0; i < 10; i++ {
      n[i] = i + 110 /* set element at location i to i + 100 */
   }
   
   /* output each array element's value */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
Element[0] = 110
Element[1] = 111
Element[2] = 112
Element[3] = 113
Element[4] = 114
Element[5] = 115
Element[6] = 116
Element[7] = 117
Element[8] = 118
Element[9] = 119
4.package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    fmt.Println(a)
}
[0 0 0]

Program exited.
5.package main

import (  
    "fmt"
)


func main() {  
    var a [3]int //int array with length 3
    a[0] = 23 // array index starts at 0
    a[1] = 89
    a[2] = 55
    fmt.Println(a)
}
[23 89 55]

Program exited.