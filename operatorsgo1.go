1.package main
import "fmt"
func main() {
    var s int=35
    var k int=36
    fmt.Println("addition is",s+k)
    fmt.Println("sub is",s-k)
    fmt.Println("multiplications is",s*k)
    fmt.Println("division is ",s/k)
    fmt.Println("modulo division",s%k)
   
}
addition is 71
sub is -1
multiplications is 1260
division is  0
modulo division 35
2.package main
import "fmt"
func main() {
    var a int =45
    var b int =8

    fmt.Printf("Assignment Operators")
    a+=b
    fmt.Printf("a+=b :%d\n", a)
    a-=b
    fmt.Printf("a-=b :%d\n", a)
    a*=b
    fmt.Printf("a*=b :%d\n", a)
    a/=b
    fmt.Printf("a/=b :%d\n", a)
    a%=b
    fmt.Printf("a%%=b :%d\n", a)
     
}
Assignment Operatorsa+=b :53
a-=b :45
a*=b :360
a/=b :45
a%=b :5
3.package main
import "fmt"
func main() {
    var a int =33
    var b int =11

    if (a > b) {
        fmt.Println("Relational Operators")
        fmt.Println("a is greater than b.")
    } else{
        fmt.Println("Relational Operators")
        fmt.Println("b is greater than a.")
    }

    
}
Relational Operators
a is greater than b.

Program exited.
4.package main
import "fmt"
func main() {
    var a int =23
    var b int =33

    if (a > b) {
       
        fmt.Println("a is greater than b.")
    } else{
        
        fmt.Println("b is greater than a.")
    }

    
}
b is greater than a.

Program exited.
5.package main
import "fmt"
func main() {
    var a int =20
    var b int =10
    var c int =25
    var flag bool = false
    var result bool
    fmt.Println("Logical Operators")
    result = (a > b) && (a > c)
    fmt.Printf("(a>b) && (a>c) : %t\n",result)
    result = (a > b) || (a > c)
    fmt.Printf("(a>b) || (a>c) :%t\n",result)
    result = !flag
    fmt.Printf("!flag :%t\n",result)
     
}
Logical Operators
(a>b) && (a>c) : false
(a>b) || (a>c) :true
!flag :true
