package main

import "fmt"

func main() {

	//声明一个变量,保存接收用户输入的选项
	key := ""
	//声明一个变量,通知是否退出for
	loop := true

	//账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支说明
	note := ""
	//收支详情,当有收支时,只需要对details进行拼接处理
	details := "收支\t账户金额\t收支金额\t说明"
	//定义变量是否有收支明细
	flag := false

	//显示主菜单
	for {
		fmt.Println("\n------------家庭收入记账软件------------")
		fmt.Println("              1收支明细")
		fmt.Println("              2登记收入")
		fmt.Println("              3登记支出")
		fmt.Println("              4退出软件")
		fmt.Print("请选择(1-4)")
		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("------------当前收支明细记录------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支,先来一笔吧")
			}
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			//修改账户余额
			balance += money
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			//将收入情况记录到details中
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足不能支出")
				break
			}
			//修改账户余额
			balance -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			//将收入情况记录到details中
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "4":
			fmt.Println("您确定要退出吗?y/n:")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("您的输入有误,您确定要退出吗?y/n:")
			}
			if choice == "y" {
				loop = false
			}

		default:
			fmt.Println("请输入正确的选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出家庭记账软件的使用...")
}
