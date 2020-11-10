1.package main

import (  
    "math"
)

func main() {  
    var a = math.Sqrt(4)   //allowed
    const b = math.Sqrt(4) //not allowed
}
Go build failed.
2.package main

import (  
    "fmt"
)

func main() {  
    const n = "Sandhya"
    var name = n
    fmt.Printf("type %T value %v", name, name)

}
type string value Sandhya
Program exited.
3.package main

func main() {  
    const trueConst = true
    type myBool bool
    var defaultBool = trueConst //allowed
    var customBool myBool = trueConst //allowed
    defaultBool = customBool //not allowed
}
./prog.go:8:17: cannot use customBool (type myBool) as type bool in assignment
Go build failed.
4.package main

import (  
    "fmt"
)

func main() {  
    const a = 10
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
}
intVar 10 
int32Var 10 
float64Var 10 
complex64Var (10+0i)