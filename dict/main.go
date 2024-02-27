package dict

import "fmt"

func DicExample() {
	m := make(map[string]int64, 4)

	m["dog"] = 1
	m["hen"] = 2
	m["cat"] = 3

	fmt.Println(m)

	which := "hen"

	v, ok := m[which]
	if ok {
		// find
		fmt.Println("finn", which, "value:", v)
	} else {
		// not find
		fmt.Println("not find", which)
	}
}