package main

import "fmt"

func enum() {

	const (
		a = 1
		b = 2
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	const (
		q = iota
		w
		e
		r
		t
		y
		u
	)
	fmt.Println(q, w, e, r, t, y, u)

	const (
		z = iota + 1
		x
		v
		g
	)
	fmt.Println(z, x, v, g)

}

func main() {

	enum()
}
