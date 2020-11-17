1.package main
import (
    "fmt"
)
func main() {
    var s string = "Hello golang"
    var b []byte = []byte(s)     
    fmt.Println(b)  
    ss := string(b)              
    fmt.Println(ss)    
}
$go run main.go
[72 101 108 108 111 32 103 111 108 97 110 103]
Hello golang
2.package main
 import (
    "fmt"
)
 func main() {
    a := 7/3        
    f := 6.3/3       
    fmt.Println(a, f)     
}
$go run main.go
2 2.1
3.package main
import (
	"fmt"
	"strconv"
)
func main() {
	s := "3.1415926535"
	f, _ := strconv.ParseFloat(s, 8)
	fmt.Printf("%T, %v\n", f, f)
	s1 := "-3.14189"
	f1, _ := strconv.ParseFloat(s1, 8)
	fmt.Printf("%T, %v\n", f1, f1)
	s2 := "-3.14165"
	f2, _ := strconv.ParseFloat(s2, 32)
	fmt.Printf("%T, %v\n", f2, f2)
}
$go run main.go
float64, 3.1415926535
float64, -3.14189
float64, -3.1416499614715576
4.package main
import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	var b bool = true
	fmt.Println(reflect.TypeOf(b))
	var s string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(s))
}
$go run main.go
bool
string
5.package main
import (
	"fmt"
	"reflect"
	"strconv"
)
func main() {
	var i int64 = -655
	fmt.Println(reflect.TypeOf(i))
	fmt.Println(i)
	var s string = strconv.FormatInt(i, 11)
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(s)
}
$go run main.go
int64
-655
string
-546