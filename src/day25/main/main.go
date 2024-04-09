package main

import "day25/account"

/**
家庭收支记账软件
收入
支出
输出收入支出表

1 收支明细
2 登记收入
3 登记支出
4 退出

*/
// 改成面向对象的

func main() {
	acc := account.Account{
		Balance: 10000.0,
		Money:   0.0,
		Note:    "",
		Details: "收支\t账户金额\t收支金额\t说明",
		Key:     "",
		Flag:    true,
	}
	acc.GetMenu()
}
