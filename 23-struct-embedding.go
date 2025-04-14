package main

import "fmt"

/* Go supports EMBEDDING of structs and interface for a more seamless composition of types */

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct { // container embeds base
	base // embedding looks like a field without a name
	str  string
}

func main() {
	co := container{ // when creating structs with literals, initialize embedding explicitly
		base: base{ // embedded type serves as field name
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str) // can access base's fields directly from container

	fmt.Println("also num:", co.base.num) // can spell out full path using embedded type name

	fmt.Println("describe:", co.describe()) // methods of base become methods of container due to embeddment

	type describer interface {
		describe() string
	}
	var d describer = co
	fmt.Println("describer:", d.describe()) // automatically uses describe from base because it is part of container
}
