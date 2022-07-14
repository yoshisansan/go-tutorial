package main

func f(xp *int) {
	*xp++
}
func main() {
	var x int
	f(&x)
	f(&x)
	println(x)
}
