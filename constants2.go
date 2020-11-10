1.package main

import "fmt"

func main() {
   const LENGTH int = 10
   const WIDTH int = 5   
   var area int

   area = LENGTH * WIDTH
   fmt.Printf("value of area : %d", area)   
}
value of area : 60
Program exited.
2.package main

import "fmt"

func main() {
  const x string = "Hello sandhya"
  fmt.Println(x)
}
Hello sandhya

Program exited.
3.package main

import "fmt"

func main() {
  const x string = "golang example"
  fmt.Println(x)
}
golang example

Program exited.
4.package main

import (  
    "fmt"
)

func main() {  
    const (
        name = "John"
        age = 50
        country = "Canada"
    )
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(country)

}
John
50
Canada
5.package main

import (  
    "fmt"
)

func main() {  
    const (
        name = "sandhya"
        age = 50
        country = "india"
    )
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(country)

}
sandhya
50
india
6.package main

import (  
    "fmt"
)

func main() {  
    const (
        name = "sandhya"
        age = 50
	qualification="MCA"
	nativeplace="rajampet"
        country = "india"
	phonenum=8877766600
		    )
    fmt.Println(name)
    fmt.Println(age)
    fmt.Println(qualification)
    fmt.Println(nativeplace)
    fmt.Println(country)
    fmt.Println(phonenum)

}
sandhya
50
MCA
rajampet
india
8877766600