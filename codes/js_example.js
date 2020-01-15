var pizza = True
pizza = false
console.log(pizza) //false 출력

const cheeze = True //상수는 변경이 불가능.

//////////////////

// if/else 혹은 for loop 중괄호(블록) 안에서 변수를 만들거나 수정하면, global로 만들어지거나 수정됨. (local 변수 x)
// local 변수를 사용하고자 하면, 'let' 을 사용해야 함.
var topic = "javascript"

if (topic) {
  let topic = "React"
  console.log('블록', topic)
  
}
console.log('글로벌', topic)

////////////////// ${}를 사용하여 문자열 안에 변수 삽입. ₩로 전체 문자열을 묶어 줘야함.
var lastName = 'HA';
const firstName = 'Hongseok';
console.log(`Hello : ${lastName}, ${firstName}`);
