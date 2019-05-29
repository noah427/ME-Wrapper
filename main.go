package main

import memeClient "github.com/noah427/ME-Wrapper/wrapper"
import "fmt"

func main() {
	me := memeClient.GetUser("creativityisdying")
	fmt.Println(me.ID)
	fmt.Println(me.Name)
	fmt.Println(me.Networth)
	myFirm := me.GetFirm()
	fmt.Println(myFirm.Name)
	myInvestments := me.GetInvestments()
	for _,investment := range myInvestments{
		fmt.Println(investment.Profit)
		fmt.Println(investment.CoinsInvested)
	}
}
