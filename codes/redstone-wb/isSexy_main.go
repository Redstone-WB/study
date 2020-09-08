package main

import (
	"fmt"
	"time"
)

// GO ROUTINE을 활용한 병렬처리

// func main() {
// 	c := make(chan string) // isSexy 함수가 끝나면,, channel을 통해 보내는 내용의 type; isSexy 함수가 main function과 communication 할 수 있도록 하는 channel임.
// 	people := [5]string{"nico", "flynn", "Redstone", "jenny", "five"}
// 	for _, person := range people {
// 		go isSexy(person, c) //만들어 둔 channel을 두 개의 함수(isSexy("nico"), isSexy("flynn"))로 보내고 있음.
// 	}
// 	// result := <-c // c로부터 메시지를 받고 있음. main은 channel로부터 reply를 받을 때까지 기다림.
// 	// fmt.Println(result)
// 	// 위 두 줄을 fmt.Println(<-c) 로 대체할 수 있음.
// 	for i := 0; i < len(people); i++ {
// 		fmt.Print("Waiting for ", i)
// 		fmt.Println(<-c)
// 	}

	// fmt.Println("Received this message: ", <-c)
	// fmt.Println("Received this message: ", <-c) // GO는 2개의 함수가 돌아가고 있는 것을 알고 있으므로, 이렇게 써도 2개를 받을 수 있음.
	// 하나 더 쓰면,, deadlock 에러가 발생!

}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	c <- person + " is Sexy" //아까 만들어 둔 chnnel로 다시 값을 보내는 것.
}
