package main

import (
	"fmt"
	"slices"
)

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s)
	
	s = append(s, "d")
	s = append(s, "e")
	s = append(s, "f")
	s = append(s, "g")
	s = append(s, "h")
	s = append(s, "i")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("copy", c)

	fmt.Println(s[2:5])
	fmt.Println(s[2:])
	fmt.Println(s[:5])

	t1 := []string{"a", "b", "c"}
	fmt.Println("t1", t1)

	t2 := []string{"a", "b", "c"}

	if slices.Equal(t1, t2) {
		fmt.Println("very equal guys")
	}

	twoD := make([][]int, 3)
	for i := 0; i< 3; i++ {
		innerLen := i + 1;
		twoD[i] = make([]int, innerLen)
		for j :=0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
