package main

import (
	"fmt"
	"strings"

	"github.com/mingyeungAA/learngo/something"
)

//main.go 파일은 컴파일을 위한 파일이다.
//컴파일을 하지 않으려면 main.go파일이 필요 없음

//argument의 타입, return값의 타입 모두 적어줘야 함!!
func multiply(a int, b int) int {
	return a * b
}

//GO는 여러개의 return 값을 가질 수 있다!
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//nacked return
func lenAndUpper02(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func repeatMe(words ...string) {
	fmt.Println(words)
}

//for
func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return total
}

//컴파일러는 main package와 그 안의 main function을
//먼저 찾고 싫행시킨다.
func main() {
	fmt.Println("Hello world!")
	fmt.Println("Hello World!!")
	//다른 패키지에서 export된 함수!
	something.SayHello()

	//타입이 없는 상수
	//const name = "nico"
	//Go는 타입언어라서 type을 지정해줘야함!
	const name string = "nico"

	fmt.Println(name)

	//변수
	var nn string = "nico"
	nn = "Lynn"
	fmt.Println(nn)

	//축약형
	//func안에서만 가능
	//상수가 아닌 변수에서만 적용 가능함
	//func밖에서는
	//var nna string = "ming"
	//이렇게 다 적어줘야함
	nna := "ming"
	fmt.Println(nna)

	fmt.Println(multiply(2, 2))

	totalLegth, upperName := lenAndUpper("nico")
	fmt.Println(totalLegth, upperName)

	totalLength22, _ := lenAndUpper("ming")
	fmt.Println(totalLength22)

	repeatMe("nico", "lynn", "dal", "marl", "fllynn")

	totalLegth, len := lenAndUpper02("nico")
	fmt.Println(totalLegth, len)

	result := superAdd(1, 2, 3, 4, 5, 6)
	fmt.Println(result)
}
