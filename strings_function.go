package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	p("Contains:", strings.Contains("test", "t"))
	p("count:", strings.Count("tests", "s"))
	p("hasPrefix:", strings.HasPrefix("tests", "te"))
	p("hasSuffix:", strings.HasSuffix("tests", "ts"))
	p("index:", strings.Index("tests", "s"))
	p("Join:", strings.Join([]string{"a", "b", "c"}, ","))
	p("repeact:", strings.Repeat("fk", 4))
	p("replact:", strings.Replace("fkkster", "ks", "x", 4))
	p("split:", strings.Split("abckdesf", ""))
	p("tolower", strings.ToLower("ADLKFJASDFJ"))
	p("toupper", strings.ToUpper("adsflajdsfkajsf"))

	p("len:", len("fjkasdflkajsdf"))
	p("char:", "hello"[2])
}
