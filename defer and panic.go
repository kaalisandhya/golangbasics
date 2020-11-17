1.package main
import (
	"fmt"
)
func test() {
	fmt.Println("Testing Golang defer function")
}
func main() {

	fmt.Println("Start")
	defer test()
	fmt.Println("second")

}
$go run main.go
Start
second
Testing Golang defer function
2.package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("staring")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Golang Panic Occurred!")
	fmt.Println("end")

}
$go run main.go
staring
2020/11/17 10:32:40 Error: Golang Panic Occurred!
3.package main 
  
import "fmt"
func add(a1, a2 int) int { 
    res := a1 + a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
 func main() { 
  
    fmt.Println("Start") 
    defer fmt.Println("End") 
    defer add(35, 55) 
    defer add(17, 18) 
} 
$go run main.go
Start
Result:  35
Result:  90
End
4.package main 
  
import "fmt"
 func mul(a1, a2 int) int { 
  
    res := a1 * a2 
    fmt.Println("Result: ", res) 
    return 0 
} 
  
func show() { 
    fmt.Println("Welcome to golang course ") 
} 

func main() { 
  
    mul(28, 44) 
  
    defer mul(28, 44) 
      show() 
} 
$go run main.go
Result:  1232
Welcome to golang course 
Result:  1232
5.package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("Golang Panic example")
	fmt.Println("end")

}
$go run main.go
Start
2020/11/17 10:42:07 Error: Golang Panic example
6.package main 
  
import "fmt"

func main() { 
      var myarr [3]string 
      myarr[0] = "GFG"
    myarr[1] = "GeeksforGeeks"
    myarr[2] = "Geek"
  fmt.Println("Elements of Array:") 
    fmt.Println("Element 1: ", myarr[0]) 
      fmt.Println("Element 2: ", myarr[5]) 
  
} 
7.package main 
  
import "fmt"
 
func entry(lang *string, aname *string) { 
      if lang == nil { 
        panic("Error: Language cannot be nil") 
    } 
       if aname == nil { 
        panic("Error: Author name cannot be nil") 
    } 
      fmt.Printf("Author Language: %s \n Author Name: %s\n", *lang, *aname) 
} 
 func main() { 
  
    A_lang := "Golang advanced topics"
      entry(&A_lang, nil) 
	  }
$go run main.go
panic: Error: Author name cannot be nil
