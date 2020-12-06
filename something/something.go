package something

import "fmt"

//대문자로 시작하기 때문에 export가 가능!
func SayHello() {
	fmt.Println("Hello")
}

//private function
func sayBye() {
	fmt.Println("Bye")
}
