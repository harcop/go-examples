package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("this is some base description=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	co := container {
		base: base {
			num: 1,
		},
		str: "this is a container baby",
	}

	fmt.Println(co.base, co.num, co.str, co.base.num)

	type describer interface{
		describe() string
	}

	var d describer = co
	fmt.Println(d.describe())

}
