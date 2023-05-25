package main

import (
	// "errors"
	"fmt"
	"nt"
)

func main(){
	t1:=nt.Create()
	t1.Watch("h1","bye")
	t1.Join(func(c *nt.Ctx) {
		world:="hello world"
		c.Set("world",world)
	})
	t1.Join(func(c *nt.Ctx) {
		fmt.Println("hello sb")
	})
	t1.Call(func(c *nt.Ctx) {
		c.Set("h2","you got a nil")
		fmt.Println("hello nt")
		fmt.Println(c.Get("world"))
		fmt.Println(c.Get("h1"))
		fmt.Println(c.Get("h2"))
	})
}

// func main(){
// 	nt.Create("example")
// 	nt.Find("example").Watch("h1","bye")
// 	nt.Find("example").Join(func(c *nt.Ctx) {
// 		c.Set("world","hello world")
// 	})
// 	nt.Find("example").Join(func(c *nt.Ctx) {
// 		fmt.Println("hello sb")
// 	})
// 	nt.Find("example").Call(func(c *nt.Ctx) {
// 		fmt.Println("hello nt")
// 		fmt.Println(c.Get("world"))
// 		fmt.Println(c.Get("h1"))
// 	})
// }

// func main() {
// 	t1 := nt.Create()
// 	t1.Join(func(c *nt.Ctx) {
// 		c.Set("test", "you may not see this")
// 	})
// 	if err := t1.Call(func(c *nt.Ctx) {
// 		shouldBeNilValue := c.Get("nothing")
// 		if shouldBeNilValue == nil {
// 			c.Error(errors.New("key not found"))
// 			return
// 		}
// 		fmt.Println("Do something here")
// 	}); err != nil {
// 		fmt.Println(err)
// 	}
// }
