package myquote

import "fmt"
import "rsc.io/quote"

func PrintQuotes1() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
}

func PrintQuotes2() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
