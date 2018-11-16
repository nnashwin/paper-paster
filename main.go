package main

import "fmt"

func main() {
	str := `The basic concept of fractals
 is that they contain
 a large degree of
self similarity
. This means
 that they
 usually
 contain little copi
es of
themselves
  buried   deep   within   the   original
  like   the   stars
embedded
 in the Koch Snowflake
[17]
 shown in 
 Figure 1
. Also,
fractals
 possess infinite
 detail, so for any given fractal the closer
we look at it the more detail it 
can reveal. 
[20]`

	bs := []byte(str)

	for i, r := range bs {
		if r == 10 {
			bs = append(bs[:i], bs[i+1:]...)
		}
	}
	fmt.Println(bs)

	s := string(bs)
	fmt.Println(s)
}
