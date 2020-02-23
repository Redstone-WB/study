### 1강

1. 'main.go' 는 compile을 하고 싶다면, not optional.
   1. complie 할 필요가 없다면, (Ex. making library) main.go가 없어도 됨.
2. 'main.go' 파일 안에서는 func main(){} 이 정의되어야 함. (entry point)
3. package 밖으로 Export 하고 싶은 function이나 변수명은 대문자로 시작함.
   1. ex. fmt.Println
4. function 안에서는, name := "redstone" 과 같이 변수 선언 가능. (shorthand functionality)
5. naked return : function이 return하는 value를 앞에서 선언할 수 있음. return하는 값들의 type을 지정해 줄 때, 변수 이름도 같이 알려주면 됨.
6. defer : function이 끝난 뒤 무언가를 하게 함.
7. Variable Expiration : if 문 바로 다음에 변수를 만들면, if/else 구문이 끝나고 나면 변수가 사라짐. (switch문도 마찬가지)
8. array 선언 예제
   1. names = [5]string{"Hongseok", "Jaeeun", "Cholong"}
   2. names[3] = "sth"
   3. names[4] = "sth"
9. array 중, length 지정이 되지 않은 것을 slice라 함.
   1.  slice에 value를 추가하기 위해서는, append() 를 씀.
   2.  names := []string{"Redstone", "Jenny"}
   3.  names = append(names, "Hongseok")   // append는 new slice를 return하기 때문에.


