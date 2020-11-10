1.package main

import (  
    "fmt"
)

type student struct {  
    firstName string
    lastName  string
    age       int
    qualification    string
}

func main() {  
    var student4 student //zero valued struct
    fmt.Println("First Name:", student4.firstName)
    fmt.Println("Last Name:", student4.lastName)
    fmt.Println("Age:", student4.age)
    fmt.Println("qualification:", student4.qualification)
}
First Name: 
Last Name: 
Age: 0
qualification: 
2.package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp8 := &Employee{
        firstName: "sanju",
        lastName:  "kaali",
        age:       55,
        salary:    70000,
    }
    fmt.Println("First Name:", (*emp8).firstName)
    fmt.Println("Age:", (*emp8).age)
}
First Name: sanju
Age: 55

Program exited.
3.package main

import (  
    "fmt"
)

type Employee struct {  
    firstName string
    lastName  string
    age       int
    salary    int
}

func main() {  
    emp8 := &Employee{
        firstName: "ram",
        lastName:  "chandhran",
        age:       25,
        salary:    10000,
    }
    fmt.Println("First Name:", emp8.firstName)
    fmt.Println("Age:", emp8.age)
	
	}
First Name: ram
Age: 25

Program exited.
4.package main

import (  
    "fmt"
)

type Address struct {  
    city  string
    state string
}

type Person struct {  
    name    string
    age     int
    address Address
}

func main() {  
    p := Person{
        name: "sandhya",
        age:  50,
        address: Address{
            city:  "kadapa",
            state: "andhrapradesh",
        },
    }

    fmt.Println("Name:", p.name)
    fmt.Println("Age:", p.age)
    fmt.Println("City:", p.address.city)
    fmt.Println("State:", p.address.state)
}
Name: sandhya
Age: 50
City: kadapa
State: andhrapradesh
5.package main

import (  
    "fmt"
)

type name struct {  
    firstName string
    lastName  string
}

func main() {  
    name1 := name{
        firstName: "sandhya",
        lastName:  "kaali",
    }
    name2 := name{
        firstName: "Sasikala",
        lastName:  "kaali",
    }
    if name1 == name2 {
        fmt.Println("name1 and name2 are equal")
    } else {
        fmt.Println("name1 and name2 are not equal")
    }

    name3 := name{
        firstName: "akash",
        lastName:  "kanda",
    }
    name4 := name{
        firstName: "Steve",
    }

    if name3 == name4 {
        fmt.Println("name3 and name4 are equal")
    } else {
        fmt.Println("name3 and name4 are not equal")
    }
}
name1 and name2 are not equal
name3 and name4 are not equal

Program exited.