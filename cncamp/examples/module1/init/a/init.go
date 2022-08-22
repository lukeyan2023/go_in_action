package a

import "fmt"
import _ "github.com/cncamp/golang/examples/module1/init/b"

func init() {
	fmt.Println("init from a")
}
