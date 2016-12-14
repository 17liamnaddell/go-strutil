package strutil_test

import (
	"fmt"
	s "github.com/skilstak/go-strutil"
)

func ExampleToSlug() {
	fmt.Println(s.ToSlug("Hello"))
	fmt.Println(s.ToSlug("Hello World!"))
	fmt.Println(s.ToSlug(" Hello  there  "))
	fmt.Println(s.ToSlug(" Hello  there  日本語! "))
	// Output:
	// hello
	// hello-world
	// hello-there
	// hello-there-日本語
}
