package main

import (
	"fmt"
	"net/http"
	"os"
)

func fabnacci(sum chan int, quit chan int) (ret int) {
	x, y := 0, 1
	for {
		select {
		case sum <- x:
			x, y = y, x+y
			/* is rquivalant with follow
			z := x + y
			x = y
			y = z
			*/
		case <-quit:
			return
		}
	}
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")

	fmt.Fprintln(res, "Init concurrency")
	//fabnacci
	fab := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprintln(res, "fab =", <-fab)
		}
		quit <- 0
	}()

	fabnacci(fab, quit)
}
