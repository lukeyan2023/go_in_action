package main

import "fmt"
import _ "github.com/cncamp/golang/examples/module1/init/a"
import _ "github.com/cncamp/golang/examples/module1/init/b"

func init() {
	fmt.Println("main init")
}

func main() {

}
