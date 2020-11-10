1.package main

import (  
    "fmt"
)

func main() {  
    a := [5]int{70, 77, 76, 79, 80}
    var b []int = a[1:4] 
    fmt.Println(b)
}
[77 76 79]

Program exited.
2.package main

import (  
    "fmt"
)

func main() {  
    c := []int{9, 8, 5} 
    fmt.Println(c)
}
[9 8 5]

Program exited.
3.package main

import (  
    "fmt"
)

func main() {  
    darr := [...]int{57, 66, 99, 34,67,112,88,69,43}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 
}
array before [57 66 99 34 67 112 88 69 43]
array after [57 66 100 35 68 112 88 69 43]

Program exited.
4.package main

import (  
    "fmt"
)

func main() {  
    num := [3]int{78, 79 ,80}
    nums1 := num[:] 
    nums2 := num[:]
    fmt.Println("array before change 1",num)
    nums1[0] = 101
    fmt.Println("array after modification to slice nums1", num)
    nums2[1] = 102
    fmt.Println("array after modification to slice nums2", num)
}
array before change 1 [78 79 80]
array after modification to slice nums1 [101 79 80]
array after modification to slice nums2 [101 102 80]
5.package main

import (  
    "fmt"
)

func main() {  
    fruitarray := [...]string{"orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) 
}
length of slice 2 capacity 5
Program exited.