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
6.package main

import (  
    "fmt"
)

func main() {  
    a := [3]int{12} 
    fmt.Println(a)
}
7.package main

import (  
    "fmt"
)

func main() {  
    a := [...]int{87, 10, 67} // ... makes the compiler determine the length
    fmt.Println(a)
}
[87 10 67]

Program exited.
8.package main

func main() {  
    a := [3]int{5, 78, 8}
    var b [5]int
    b = a //not possible since [3]int and [5]int are distinct types
}
./prog.go:6:7: cannot use a (type [3]int) as type [5]int in assignment

Go build failed.
9.package main

import "fmt"

func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[0] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}
a is  [USA China India Germany France]
b is  [Singapore China India Germany France]

Program exited.
10.package main

import "fmt"

func changeLocal(num [5]int) {  
    num[0] = 55
    fmt.Println("inside function ", num)

}
func main() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}
before passing to function  [5 6 7 8 8]
inside function  [55 6 7 8 8]
after passing to function  [5 6 7 8 8]
11.package main

import "fmt"

func main() {  
    a := [...]float64{68.7, 84.8, 56, 90}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}
0 th element of a is 68.70
1 th element of a is 84.80
2 th element of a is 56.00
3 th element of a is 90.00
12.package main

import (  
    "fmt"
)

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func main() {  
    a := [3][2]string{
        {"lion", "dog"},
        {"mutton", "chicken"},
        {"rabbit", "tiger"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "nikitha"
    b[0][1] = "sasikala"
    b[1][0] = "teju"
    b[1][1] = "navya"
    b[2][0] = "akash"
    b[2][1] = "chandhuhasini"
    fmt.Printf("\n")
    printarray(b)
}
lion dog 
mutton chicken 
rabbit tiger 

nikitha sasikala 
teju navya 
akash chandhuhasini