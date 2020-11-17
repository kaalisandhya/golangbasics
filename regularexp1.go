1.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := `http://www.suon.co.uk/product/1/7/3/`

	re := regexp.MustCompile(`^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)`)
	fmt.Printf("Pattern Type: %v\n", re.String()) 
	fmt.Println(re.MatchString(str1)) 

	submatchall := re.FindAllString(str1,-1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
$go run main.go
Pattern Type: ^(?:https?:\/\/)?(?:[^@\/\n]+@)?(?:www\.)?([^:\/\n]+)
true
http://www.suon.co.uk
2.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "sandhyakaali@gmail.com"
	str2 := "kaalisandhya24@gmail_yahoo.com"
	str3 := "suji@gmail-yahoo.com"
	str4 := "rani@gmailyahoo"
	str5 := "suraj56@gmail.yahoo"

	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern	
	fmt.Printf("\nEmail: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Email: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Email: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Email: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Email: %v :%v\n", str5, re.MatchString(str5))
}
$go run main.go
Pattern: ^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$

Email: sandhyakaali@gmail.com :true
Email: kaalisandhya24@gmail_yahoo.com :false
Email: suji@gmail-yahoo.com :true
Email: rani@gmailyahoo :true
Email: suraj56@gmail.yahoo :true
3.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "5678901x1234"
	str2 := "(+351) 282 43 50 50"
	str3 := "90191919908"
	str4 := "7689056789"
	str5 := "001 6867684"
	str6 := "00106867684x1"
	str7 := "1 (234) 567-8901"
	str8 := "1-234-567-8901 ext1234"

	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nPhone: %v\t:%v\n", str1, re.MatchString(str1))
	fmt.Printf("Phone: %v\t:%v\n", str2, re.MatchString(str2))
	fmt.Printf("Phone: %v\t\t:%v\n", str3, re.MatchString(str3))
	fmt.Printf("Phone: %v\t\t\t:%v\n", str4, re.MatchString(str4))
	fmt.Printf("Phone: %v\t\t:%v\n", str5, re.MatchString(str5))
	fmt.Printf("Phone: %v\t\t:%v\n", str6, re.MatchString(str6))
	fmt.Printf("Phone: %v\t\t:%v\n", str7, re.MatchString(str7))
	fmt.Printf("Phone: %v\t:%v\n", str8, re.MatchString(str8))
}
$go run main.go
Pattern: ^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$

Phone: 5678901x1234	        :true
Phone: (+351) 282 43 50 50	:true
Phone: 90191919908		:true
Phone: 7689056789		:true
Phone: 001 6867684		:true
Phone: 00106867684x1		:true
Phone: 1 (234) 567-8901		:true
Phone: 1-234-567-8901 ext1234	:true
4.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "17/1q/2020"
	str2 := "16/10/2020"
	str3 := "29/2/2007"
	str4 := "31/08/2010"
	str5 := "29/02/200a"
	str6 := "29/0a/200a"
	str7 := "55/02/200a"
	str8 := "2_/02/2009"

	re := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nDate: %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("Date: %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("Date: %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("Date: %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("Date: %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("Date: %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("Date: %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("Date: %v :%v\n", str8, re.MatchString(str8))
}
$go run main.go
Pattern: (0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\d\d)

Date: 17/1q/2020 :false
Date: 16/10/2020 :true
Date: 29/2/2007 :true
Date: 31/08/2010 :true
Date: 29/02/200a :false
Date: 29/0a/200a :false
Date: 55/02/200a :false
Date: 2_/02/2009 :false
5.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "778899787888989"
	str2 := "346823285239073"
	str3 := "3707909-7718351"
	str4 := "455622983649586"
	str5 := "501971570101037"
	str6 := "7600924456189087"
	str7 := "4111-1111-1111-1111"
	str8 := "5610591081018250"
	str9 := "30569309025904000"
	str10 := "6011111111111117"

	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	fmt.Printf("Pattern: %v\n", re.String()) // print pattern
	fmt.Printf("\nCC : %v :%v\n", str1, re.MatchString(str1))
	fmt.Printf("CC : %v :%v\n", str2, re.MatchString(str2))
	fmt.Printf("CC : %v :%v\n", str3, re.MatchString(str3))
	fmt.Printf("CC : %v :%v\n", str4, re.MatchString(str4))
	fmt.Printf("CC : %v :%v\n", str5, re.MatchString(str5))
	fmt.Printf("CC : %v :%v\n", str6, re.MatchString(str6))
	fmt.Printf("CC : %v :%v\n", str7, re.MatchString(str7))
	fmt.Printf("CC : %v :%v\n", str8, re.MatchString(str8))
	fmt.Printf("CC : %v :%v\n", str9, re.MatchString(str9))
	fmt.Printf("CC : %v :%v\n", str10, re.MatchString(str10))
}
$go run main.go
Pattern: ^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$

CC : 778899787888989 :false
CC : 346823285239073 :true
CC : 3707909-7718351 :false
CC : 455622983649586 :false
CC : 501971570101037 :false
CC : 7600924456189087 :false
CC : 4111-1111-1111-1111 :false
CC : 5610591081018250 :true
CC : 30569309025904000 :false
CC : 6011111111111117 :true
6.package main

import (
	"fmt"
	"log"
	"regexp"
)

func main() {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	newStr := reg.ReplaceAllString("#Golang#Python$kaali&sandhya@@", "-")
	fmt.Println(newStr)
}
$go run main.go
-Golang-Python-kaali-sandhya-
11.package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "kaali sandhya X42 I'm a Y-32.35 string Z30"

	re := regexp.MustCompile(`[A-Z][^A-Z]*`)

	fmt.Printf("Pattern Type: %v\n", re.String()) 

	submatchall := re.FindAllString(str1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
$go run main.go
Pattern Type: [A-Z][^A-Z]*
X42 
I'm a 
Y-32.35 string 
Z30