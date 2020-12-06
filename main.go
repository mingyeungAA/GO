package main

import (
	"fmt"

	"github.com/mingyeungAA/learngo/something"
)

//main.go 파일은 컴파일을 위한 파일이다.
//컴파일을 하지 않으려면 main.go파일이 필요 없음

//컴파일러는 main package와 그 안의 main function을
//먼저 찾고 싫행시킨다.
func main() {
	fmt.Println("Hello world!")
	fmt.Println("Hello World!!")
	//다른 패키지에서 export된 함수!
	something.SayHello()
}
