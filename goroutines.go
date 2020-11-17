package main

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 5; i++ {
        fmt.Println(from, ":", i)
    }
}

func main() {

    f("direct")

    go f("goroutine")

    go func(msg string) {
        fmt.Println(msg)
    }("going")

    time.Sleep(time.Second)
    fmt.Println("done")
}
$go run main.go
direct : 0
direct : 1
direct : 2
direct : 3
direct : 4
goroutine : 0
goroutine : 1
goroutine : 2
goroutine : 3
goroutine : 4
going
done
2.package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("golang Goroutines")
}
$go run main.go
hello
golang Goroutines
golang Goroutines
hello
golang Goroutines
hello
hello
golang Goroutines
golang Goroutines
3.package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	go responseSize("https://www.golangprograms.com")
	go responseSize("https://codevita.com")
	go responseSize("https://mythra.com")
	time.Sleep(10 * time.Second)
}
$go run main.go
Step1:  https://www.golangprograms.com
Step1:  https://codevita.com
Step1:  https://mythra.com
2020/11/17 08:12:15 Get https://www.golangprograms.com: dial tcp: lookup www.golangprograms.com on 213.133.99.99:53: dial udp 213.133.99.99:53: connect: network is unreachable
exit status 1
4.package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	fmt.Println("Step1: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Step2: ", url)
	defer response.Body.Close()

	fmt.Println("Step3: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step4: ", len(body))
}

func main() {
	// Add a count of three, one for each goroutine.
	wg.Add(3)
	fmt.Println("Start Goroutines")

	go responseSize("https://www.golangprograms.com")
	go responseSize("https://www.google.com")
	go responseSize("https://codevita.com")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Terminating Program")
}
$go run main.go
Start Goroutines
Step1:  https://codevita.com
Step1:  https://www.golangprograms.com
Step1:  https://www.google.com
2020/11/17 08:17:16 Get https://codevita.com: dial tcp: lookup codevita.com on 213.133.99.99:53: dial udp 213.133.99.99:53: connect: network is unreachable
exit status 1
5.package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)
// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func responseSize(url string, nums chan int) {
	// Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// Send value to the unbuffered channel
	nums <- len(body)
}

func main() {
	nums := make(chan int) // Declare a unbuffered channel
	wg.Add(1)
	go responseSize("https://www.golangprograms.com", nums)
	fmt.Println(<-nums) // Read the value from unbuffered channel
	wg.Wait()
	close(nums) // Closes the channel
}
$go run main.go
2020/11/17 08:19:00 Get https://www.golangprograms.com: dial tcp: lookup www.golangprograms.com on 213.133.99.99:53: dial udp 213.133.99.99:53: connect: network is unreachable
exit status 1
6.package main

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
}
$go run main.go
1
2
3
4
Pause
Play
5
6
7
8
Stop
7.package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32          
	wg      sync.WaitGroup 
)

func main() {
	wg.Add(3) 

	go increment("Python")
	go increment("Java")
	go increment("Golang")

	wg.Wait() 
	fmt.Println("Counter:", counter)

}

func increment(name string) {
	defer wg.Done()
	for range name {
		atomic.AddInt32(&counter, 3)
		runtime.Gosched() 
	}
}
$go run main.go
Counter: 48
8.package main

import (
	"fmt"
	"sync"
)

var (
	counter int32          
	wg      sync.WaitGroup 
	mutex   sync.Mutex     
)

func main() {
	wg.Add(3) 

	go increment("Python")
	go increment("Go Programming Language")
	go increment("Java")

	wg.Wait() 
	fmt.Println("Counter:", counter)

}

func increment(lang string) {
	defer wg.Done() 
	for i := 0; i < 6; i++ {
		mutex.Lock()
		{
			fmt.Println(lang)
			counter++
		}
		mutex.Unlock()
	}
}
$go run main.go
Java
Java
Java
Java
Java
Java
Python
Python
Python
Python
Python
Python
Go Programming Language
Go Programming Language
Go Programming Language
Go Programming Language
Go Programming Language
Go Programming Language
Counter: 18