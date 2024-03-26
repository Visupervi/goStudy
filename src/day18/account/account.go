package account

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

func (a *Account) SaveMoney(n float64, pwd string) {

	if pwd != a.Pwd {
		fmt.Println("密码不正确...")
		return
	}
	if n > 0 {
		a.Balance += n
	} else {
		fmt.Println("请输入正确金额...")
		return
	}
	fmt.Println("您的当前余额是", a.Balance)
	//return a.Balance
}

func (a *Account) GetMoney(money float64, pwd string) {

	if pwd != a.Pwd {
		fmt.Println("密码不正确")
		return
	}
	if money > a.Balance {
		fmt.Println("余额不足")
		return
	} else if money < 0 {
		fmt.Println("请输入正确金额")
		return
	} else {
		a.Balance -= money
	}
	fmt.Println("您的当前余额是", a.Balance)
	//return a.Balance
}

func (a *Account) LookMoney(pwd string) {
	if pwd != a.Pwd {
		fmt.Println("密码不正确")
	}
	fmt.Println("您的当前余额是", a.Balance)
}
