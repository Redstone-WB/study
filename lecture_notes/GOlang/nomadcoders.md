### 1강

1. 'main.go' 는 compile을 하고 싶다면, not optional.
   1. complie 할 필요가 없다면, (Ex. making library) main.go가 없어도 됨.
2. 'main.go' 파일 안에서는 func main(){} 이 정의되어야 함. (entry point)
3. package 밖으로 Export 하고 싶은 function이나 변수명은 대문자로 시작함.
   1. ex. fmt.Println
4. function 안에서는, name := "redstone" 과 같이 변수 선언 가능. (shorthand functionality)
5.
