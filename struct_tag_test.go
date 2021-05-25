package struct_tag

import (
	"fmt"
)

func ExampleGet() {
	fmt.Println(Get(`a:"av"`, `a`))
	// Output:
	// av
}

func ExampleLookup_false() {
	fmt.Println(Lookup(``, ``))
	fmt.Println(Lookup(` `, ``))
	fmt.Println(Lookup(`a`, ``))
	fmt.Println(Lookup(`a:""`, ``))
	fmt.Println(Lookup(`a:"av"`, `b`))
	// Output:
	//  false
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
		b:"b v" `, `b`))
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

func ExampleParse_true() {
	fmt.Println(Parse(`a:""`))
	fmt.Println(Parse(`a:"av"`))
	fmt.Println(Parse(`a:"av" b:"b\" v"`))
	fmt.Println(Parse(`a:"av"
		b:"b v" `))
	fmt.Println(Parse(`名称:"值
\"def"`))
	// Output:
	// map[a:]
	// map[a:av]
	// map[a:av b:b" v]
	// map[a:av b:b v]
	// map[名称:值
	// "def]
}

func ExampleTrimLeadingSpace() {
	fmt.Printf("%#v\n", trimLeadingSpace("a"))
	fmt.Printf("%#v\n", trimLeadingSpace("  "))
	fmt.Printf("%#v\n", trimLeadingSpace("  b"))
	// Output:
	// "a"
	// ""
	// "b"
}

func ExampleStripName() {
	fmt.Println(stripName("a:xx"))
	name, tag := stripName("a :")
	fmt.Printf("%s%s\n", name, tag)
	fmt.Println(stripName("a:"))
	// Output:
	// a xx
	//
	// a
}

func ExampleStripValue() {
	fmt.Println(stripValue(`"xx"yy`))
	name, tag := stripValue(`a`)
	fmt.Printf("%s%s\n", name, tag)
	name, tag = stripValue(`"a:`)
	fmt.Printf("%s%send\n", name, tag)
	// Output:
	// "xx" yy
	//
	// end
}
