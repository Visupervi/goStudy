package account

import (
	"fmt"
)

type Account struct {
	Key     string  `json:"key"`
	Money   float64 `json:"money"`   // 每次收支的金额
	Balance float64 `json:"balance"` // 余额
	Note    string  `json:"note"`    // 收支的说明
	Details string  `json:"details"`
	Flag    bool    `json:"flag"`
}

func (a *Account) saveMoney() {
	fmt.Println("请输入收入")
	fmt.Scanln(&a.Money)
	a.Balance += a.Money
	fmt.Println("请输入说明")
	fmt.Scanln(&a.Note)
	a.Details += fmt.Sprintf("\n收入\t%v\t%v\t%v", a.Balance, a.Money, a.Note)
}

func (a *Account) getMoney() {
	fmt.Println("输入支出")
	//fmt.Println("请输入收入")
	fmt.Scanln(&a.Money)
	if a.Money > a.Balance {
		fmt.Println("余额不足")
	}
	a.Balance -= a.Money
	fmt.Println("请输入说明")
	fmt.Scanln(&a.Note)
	a.Details += fmt.Sprintf("\n支出\t%v\t%v\t%v", a.Balance, a.Money, a.Note)

}

func (a *Account) showDetails() {
	fmt.Println("----------------------收支明细----------------------")
	fmt.Println(a.Details)
}

func (a *Account) exit() {
	fmt.Println("你确定要退出么？y/n")
	choice := ""

	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}

		fmt.Println("输入有误")
	}

	if choice == "y" {
		a.Flag = false
	}
}

func (a *Account) GetMenu() {
	for {
		fmt.Println("\n---------------------------收支软件已启动---------------------------")
		fmt.Println("                             输入1显示收支明细")
		fmt.Println("                             输入2登记收入")
		fmt.Println("                             输入3登记支出")
		fmt.Println("                             输入4退出软件")
		fmt.Println("-----------------------------收支软件已启动---------------------------")
		fmt.Println("请输入<1-4>：")
		fmt.Scanln(&a.Key)
		switch a.Key {
		case "1":
			a.showDetails()
		case "2":
			a.saveMoney()
		case "3":
			a.getMoney()
		case "4":
			a.exit()
		}
		if !a.Flag {
			fmt.Println("退出收支软件")
			break
		}
	}

}
