var pizza = True
pizza = false
console.log(pizza) //false 출력

const cheeze = True //상수는 변경이 불가능.

//////////////////

// if/else 중괄호(블록) 안에서 변수를 만들거나 수정하면, global로 만들어지거나 수정됨. (local 변수 x)
// local 변수를 사용하고자 하면, 'let' 을 사용해야 함.
var topic = "javascript"

if (topic) {
  let topic = "React"
  console.log('블록', topic)
  
}
console.log('글로벌', topic)

//////////////////
