### 1강

1. 'main.go' 는 compile을 하고 싶다면, not optional.

   1. complie 할 필요가 없다면, (Ex. making library) main.go가 없어도 됨.

2. 'main.go' 파일 안에서는 func main(){} 이 정의되어야 함. (entry point)

3. package 밖으로 Export 하고 싶은 function이나 변수명은 대문자로 시작함.

   1. ex. fmt.Println

4. function 안에서는, name := "redstone" 과 같이 변수 선언 가능. (shorthand functionality)

5. 변수를 n개 받는 함수를 만들고 싶을 때, ...[type] 와 같이 사용
    func repeatMe (words …string) {
      fmt.Println(words)
    }

6. naked return : function이 return하는 value를 앞에서 선언할 수 있음. return하는 값들의 type을 지정해 줄 때, 변수 이름도 같이 알려주면 됨.
   ex> func lenAndUpper (name string) (length int, uppercase string){

   ​    length = len(name)

   ​    uppercase = strings.ToUpper(name)

   ​    return
   }

7. defer : function이 끝난 뒤 무언가를 하게 함.

8. Variable Expiration : if 문 바로 다음에 변수를 만들면, if/else 구문이 끝나고 나면 변수가 사라짐. (switch문도 마찬가지)

9. Pointer

    1. b := &a // a의 주소값을 b에 저장
        fmt.Println(*b)
    2. *는 'see through' 를 의마함
    3. &는 주솟값을 의미함

10. array 선언 예제

   11. names = [5]string{"Hongseok", "Jaeeun", "Cholong"}

   12. names[3] = "sth"

   3. names[4] = "sth"

11. array 중, length 지정이 되지 않은 것을 slice라 함.

     1.  slice에 value를 추가하기 위해서는, append() 를 씀.
     2.  names := []string{"Redstone", "Jenny"}
     3.  names = append(names, "Hongseok")   // append는 new slice를 return하기 때문에.

15. map

      1.  rdst := map[string]string{"name" : "Redstone", "age" : "29"} // [key의 type] value의 type
      2.  map에 element 추가, map에서 element 검색 등 가능
       3.  python에서 dictionary와 유사한 것.

16. Struct

     1. map과 다르게, 서로 다른 type의 element들을 넣을 수 있음.

17. channels & go routines

18. GO는 constructor가 없다.  ==> 이를 대신할 함수를 사용함

     1. struct 안의 변수들을 대문자로 지정하여 main.go 에서 사용하면, struct 안의 변수들을 임의로 바꿀 수 있다는 문제가 생김

     2. 따라서, struct 안의 변수들은 소문자로 쓰되, struct를 만들어 주는 function을 정의해서 사용함

     3. 기 정의된 Account struct에 대해...
         func NewAccount(owner string) *Account {

         ​	account := Account{owner : owner, balance : 0}
         }

19. Method : func와 [Method_NAME] 사이에 (a *Account) 와 같은 Receiver를 넣어 주면 Method가 됨.

     1. Pointer Receiver : (a Account) 를 넣으면 복사본을 만들고,  (a *Account)를 넣으면 Method를 호출한 변수의 값을 실제로
         바꿀 수 있음.

20. Handling Errors

     1. Golang 에서는 error 를 직접 handling해 주어야 함. (if 조건 넣고, return errors.New("Can't Withdaraw!"))
     2. error type에 해당하는 값은 nil, error 가 있음. (nil은 문제 없는 경우)

21. empty map (ex> var mapname map[string]string ) 을 만들 경우, nil 이 되고, 여기에 새로운 값을 write할 수가 없음

     1. panic 에러가 나게 됨.
     2. map을 초기화할 때 empty map을 만들면 안 되고, make() 함수를 사용하거나, var mapname map[string]string{} 과 같이 써야함.

22. GO ROUTINE

     1. 프로그램이 동작하는 동안만 유효함 (main 함수가 run 하는 동안만 유효)
     2. main function은 go routine을 기다리지 않음.. 

23. CHANNEL 

     1. goroutine과 main func. 사이에 정보를 전달하기 위한 방법.
     2. fmt.Println(<-c) 와 같은 함수를 마주하면, main()은 c로부터 데이터를 받을 때까지 멈춘다. (blocking operation)

     

     



