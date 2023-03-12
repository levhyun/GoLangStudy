// Title : 1.0 Main Package (04:02)
// https://nomadcoders.co/go-for-beginners/lectures/1500

package main // 패키지 선언

import "fmt"

func main() { // 함수 선언
	// 컴파일러는 소스코드를 실행하게 되면 main패키지에 main함수를 가장 먼저 찾아 실행한다.
	fmt.Println("Hello World!") // 출력
}
