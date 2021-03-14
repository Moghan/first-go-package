package hello

import "rsc.io/quote/v3"
import "fmt"

func Hello() string {
	fmt.Println("Hello World")
    return quote.HelloV3()
}

func Proverb() string {
    return quote.Concurrency()
}