package struct_tag

import (
	"fmt"
)

func ExampleLookup_false() {
	fmt.Println(Lookup(``, ``))
	fmt.Println(Lookup(`a`, ``))
	fmt.Println(Lookup(`a:""`, ``))
	fmt.Println(Lookup(`a:"av"`, `b`))
	// Output:
	//  false
	//  false
	//  false
	//  false
}

func ExampleLookup_true() {
	fmt.Println(Lookup(`a:""`, `a`))
	fmt.Println(Lookup(`a:"av"`, `a`))
	fmt.Println(Lookup(`a:"av" b:"b\" v"`, `b`))
	fmt.Println(Lookup(`a:"av"
		b:"b v"`, `b`))
	fmt.Println(Lookup(`名称:"值
\"def"`, `名称`))
	// Output:
	//  true
	// av true
	// b" v true
	// b v true
	// 值
	// "def true

}
