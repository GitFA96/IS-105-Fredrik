package myquote

import "fmt"
import "rsc.io/quote"

func PrintQuotesone() {
	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
}

func PrintQuotestwo() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
