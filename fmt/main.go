package testfmt

import "fmt"

func Run() {
	s := "GET 200"
	var m string
	var sc int64
	fmt.Scanf(s, "%d %s", &m, &sc)
	fmt.Println(m, sc)
}
