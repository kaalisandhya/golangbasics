1.package main

import (
    "bytes"
    "fmt"
    "regexp"
)

func main() {

    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    r, _ := regexp.Compile("p([a-z]+)ch")

    fmt.Println(r.MatchString("peach"))

    fmt.Println(r.FindString("peach punch"))

    fmt.Println(r.FindStringIndex("peach punch"))

    fmt.Println(r.FindStringSubmatch("peach punch"))

    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    fmt.Println(r.FindAllString("peach punch pinch", -1))

    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    fmt.Println(r.FindAllString("peach punch pinch", 2))

    fmt.Println(r.Match([]byte("peach")))

    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
$go run main.go
true
true
peach
[0 5]
[peach ea]
[0 5 1 3]
[peach punch pinch]
[[0 5 1 3] [6 11 7 9] [12 17 13 15]]
[peach punch]
true
p([a-z]+)ch
a <fruit>
a PEACH
2.package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str1 := "this is a [sample] [[string]] with [SOME] special words"

	re := regexp.MustCompile(`\[([^\[\]]*)\]`)
	fmt.Printf("Pattern type: %v\n", re.String())     
	fmt.Println("Matched:", re.MatchString(str1)) 

	fmt.Println("\nText between square brackets:")
	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "[")
		element = strings.Trim(element, "]")
		fmt.Println(element)
	}
}
$go run main.go
Pattern type: \[([^\[\]]*)\]
Matched: true

Text between square brackets:
sample
string
SOME
3.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "I @@@ LIKE @@@@ #Golang!$! ****Programming****Language^^^"

	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)

	fmt.Printf("Pattern Type: %v\n", re.String()) 
	fmt.Println(re.MatchString(str1))        

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
$go run main.go
Pattern Type: [^a-zA-Z0-9]+
true
@@@ 
@@@@ #
!$! ****
****
^^^
4.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "If I am 20 years 10 months and 14 days old as of August 17,2016 then my DOB would be 1995-10-03"

	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)

	fmt.Printf("Pattern: %v\n", re.String()) 

	fmt.Println(re.MatchString(str1))

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
$go run main.go
Pattern: \d{4}-\d{2}-\d{2}
true
1995-10-03
5.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `Proxy Port Last Check Proxy Speed Proxy Country Anonymity 118.99.81.204
	118.99.81.204 8080 34 sec Indonesia - Tangerang Transparent 2.184.31.2 8080 58 sec 
	Iran Transparent 93.126.11.189 8080 1 min Iran - Esfahan Transparent 202.118.236.130 
	7777 1 min China - Harbin Transparent 62.201.207.9 8080 1 min Iraq Transparent`

	re := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)

	fmt.Printf("Pattern: %v\n", re.String()) 
	fmt.Println(re.MatchString(str1)) 

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
$go run main.go
Pattern: (25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}
true
118.99.81.204
118.99.81.204
2.184.31.2
93.126.11.189
202.118.236.130
62.201.207.9
