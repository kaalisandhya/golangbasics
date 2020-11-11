1.package main
import "fmt"
func main() {
    var s,k,r int
    s = 50
    k = 10
    r = s & k        
    fmt.Println("bitwise AND integers",r)
    r = s | k       
    fmt.Println("bitwise OR  integers",r)
    r = s ^ k       
    fmt.Println("bitwise aOR integers",r)
    r = s &^ k      
    fmt.Println("bit clear (AND NOT) integers",r)
}
bitwise AND integers 2
bitwise OR  integers 58
bitwise aOR integers 56
bit clear (AND NOT) integers 48
2.package main

import "fmt"

func main() {
    var p int = 20
    var q int = 60
    fmt.Println(p + q)
}
80

Program exited.
3.import "fmt"

func main() {
    var k int = 23
    var b int = 78

    if k != b && k <= b {
        fmt.Println("first True")
    }

    if k != b || k <= b {
        fmt.Println("second True")
    }

    if !(k == b) {
        fmt.Println("third True")
    }
	}
first True
second True
third True
package main

4.import "fmt"

func main() {
    var x int = 11
    var y int = 21

    if x != y && x >= y {
        fmt.Println("first True")
    }

    if x != y || x <= y {
        fmt.Println("second True")
    }

}
second True

Program exited.
5.package main

import "fmt"

func main() {
    p := 19
    k := 21

    
    option1 := p & k
    fmt.Printf("Result of p & k = %d", option1)

    
    option2 := p | k
    fmt.Printf("\nResult of p | k = %d", option2)

    
    option3 := p ^ k
    fmt.Printf("\nResult of p ^ k = %d", option3)

    
    option4 := p << 1
    fmt.Printf("\nResult of p << 1 = %d", option4)

    
    option5 := p >> 1
    fmt.Printf("\nResult of p >> 1 = %d", option5)

    
    option6 := p &^ k
    fmt.Printf("\nResult of p &^ k = %d", option6)

}

Result of p & k = 17
Result of p | k = 23
Result of p ^ k = 6
Result of p << 1 = 38
Result of p >> 1 = 9
Result of p &^ k = 2
6. package main

import "fmt"

func main() {
    s := 10
    k := &s

    fmt.Println(*k)

    *k = 12
    fmt.Println(s)
}
10
12

Program exited.