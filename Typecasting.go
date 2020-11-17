1.package main
 import (
    "fmt"
)
 func main() {
    var a int = 42
    f := float64(a)
    fmt.Println(f)       
}
$go run main.go
67
2.package main 
  import "fmt"
  func main() { 
          var totalsum int = 840
    var number int = 14
    var avg float32 
          avg = float32(totalsum) / float32(number)   
    fmt.Printf("Average = %f\n", avg) 
} 
$go run main.go
Average = 60.000000
3.package main
import "fmt"
func main() {
    var x int = 18
    var y int = 23
    var mul float32
    mul = float32(x) * float32(y)
    fmt.Printf("Multiplication = %f\n", mul)
}
$go run main.go
Multiplication = 414.000000
4.package main
 import (
    "fmt"
    "strconv"
)
 func main() {
    var s string = "28"
    v, _ := strconv.Atoi(s)       
        fmt.Println(v)    
    var i int = 28
    str := strconv.Itoa(i)         
    fmt.Println(str) 
}
$go run main.go
28
28
5.package main
 import (
    "fmt"
)
 func main() {
    f := 19.34567
    i := int(f)  
    fmt.Println(i)           
    ii := 37
    ff := float64(ii)
    fmt.Println(ff)     
}
$go run main.go
19
37
