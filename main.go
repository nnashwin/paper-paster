package main

import "fmt"

func convertString(str string) string {

	return ""
}

func main() {
	str := `Like any procedural technique
, a fractal shape is defined
 by an
algorithm
 for generating
 the shape. In the case of fractals
 these
algorithms
 are recursive
 and successive
 recursions
 yield more
detailed
 versions
 of the basic shape. The example of the Koch
snowflake in Figure 2 shows four such recursions.
 Self-similarit
y
is achieved
 by generating
 the same shapes or patterns
 at smaller
and smaller
 scales as the recursion
 progresses,
 a propert
y often
referred
 to as scale invariance.
 There is no theoretical
 limit to the
amount
 of recursion
 that can be done and hence infinite
 levels of
detail exist within the shape.`

	bs := []byte(str)

	for i, r := range bs {
		if r == 10 {
			s := append(bs[:i], byte(32))
			bs = append(s, bs[i+1:]...)
		}
	}

	for i := len(bs) - 1; i >= 0; i-- {
		if bs[i] == 32 && bs[i+1] == 32 {
			bs = append(bs[:i], bs[i+1:]...)
		}
	}

	fmt.Println(bs)

	s := string(bs)
	fmt.Println(s)
}
