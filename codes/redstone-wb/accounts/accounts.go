package accounts

import (
	"errors"
	"fmt"
)

// Account struct
type Account struct {
	owner   string // 대문자로 안해주면 main.go에서 참조할 수 없음. (소문자는 private variable)
	balance int
}

// Fuction for Creating Struct : owner를 받아서 Account의 copy를 반환함
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	// fmt.Println(account)

	return &account // 복사본이 아니라, object를 return 하는 것, memory address를 return함.

}

// Method 작성 : func 과 이름 사이에 (input type)을 넣어 줌. => 괄호 안의 변수 input을 Receiver라 함.
// Receiver 규칙 : type의 첫 lowercase를 따서 지어야 함.
// (a Account) 와 같이 선언하면, Account의 copy를 가져오는 것임. => 복사본을 만들고 싶은 경우
// 복사본을 만들지 않고, 실제로 Account의 balance를 변경 시키려면? (a *Account) 와 같이 써줘야 함.
// 의미는? accounts.Deposit() 을 어딘가에서 호출한다면, Account를 복사하지 말고, Deposit()을 call 하고 있는 account라는 변수를 활용하라는 것임! (main.go에 있는...)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance를 보는 method
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't Withdraw, you are poor!")
	}
	a.balance -= amount
	return nil // error type도 nil과 Error의 value가 있음.
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Show owner of an account
func (a *Account) Owner() string {
	return a.owner
}

// GO가 내부적으로 제공하는 메쏘드(여기서는 String() )를 사용하는 방법
func (a *Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n Has : ", a.Balance()) // 값을 그대로 문자열로 만드는 함수, fmt.Sprint()
}
