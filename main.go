package main
import "C"
import "fmt"

func main(){
	fmt.Println("I am in Go code know")
	C.callFromC()
}

