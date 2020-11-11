1.package main

import "fmt"

func plus(a int, b int) int {

    return a + b
}

func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {

    res := plus(6, 8)
    fmt.Println("6+8 =", res)

    res = plusPlus(6, 8, 9)
    fmt.Println("6+8+9 =", res)
}
6+8 = 14
6+8+9 = 23

Program exited.
2.package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("b-c-u-o-l", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("howareyou"))
    p("Char:", "howareyou"[1])
}
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [b c u o l]
ToLower:    test
ToUpper:    TEST

Len:  9
Char: 111  

3.package main
  
import "fmt"
  
// function which swap values
func swap(a, b int)int{
 
    var o int
    o= a
    a=b
    b=o
    
   return o 
}
  
// Main function
func main() {
 var p int = 45
 var q int = 67
  fmt.Printf("p = %d and q = %d", p, q)
  
 // call by values
 swap(p, q)
   fmt.Printf("\np = %d and q = %d",p, q)
}
p = 45 and q = 67
p = 45 and q = 67
Program exited.
4.package main 
  
import( 
    "fmt"
    "strings"
) 
  
// Variadic function to join strings 
func joinstr(element...string)string{ 
    return strings.Join(element, "-") 
} 
  
func main() { 
    
  // zero argument 
   fmt.Println(joinstr()) 
     
   // multiple arguments 
   fmt.Println(joinstr("SANDHYA", "KAALI")) 
   fmt.Println(joinstr("KAALI", "for", "GOOD")) 
   fmt.Println(joinstr("S", "A", "N", "D", "Y")) 
     
} 
SANDHYA-KAALI
KAALI-for-GOOD
S-A-N-D-Y
5.package main 
  
// Importing packages 
import ( 
    "fmt"
    "sort"
    "strings"
    "time"
) 
  
// Main function 
func main() { 
  
    // Sorting the given slice 
    s := []int{7, 0, 123, 10, 796, 122, 57, 05,78,89,70,455} 
    sort.Ints(s) 
    fmt.Println("Sorted slice: ", s) 
  
    // Finding the index 
    fmt.Println("Index value: ", strings.Index("sandhyakaali", "ks")) 
  
    // Finding the time 
    fmt.Println("Time: ", time.Now().Unix()) 
  
} 
Sorted slice:  [0 5 7 10 57 70 78 89 122 123 455 796]
Index value:  -1
Time:  1257894000
6.package main 
  
import "fmt"
  
// myfunc return 3 values of int type 
func func1(s, k int)(int, int, int ){ 
    return s - k, s * k, s + k  
} 
  
// Main Method 
func main() { 
      
    // The return values are assigned into  
    // three different variables 
   var var1, var2, var3 = func1(4, 2) 
     
   // Display the values 
   fmt.Printf("Result is: %d", var1 ) 
   fmt.Printf("\nResult is: %d", var2) 
   fmt.Printf("\nResult is: %d", var3) 
} 
Result is: 2
Result is: 8
Result is: 6
Program exited.
