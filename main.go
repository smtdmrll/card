package main

import "fmt"

func main() {


cards := deste{newCard(), newCard(), "karo 5 "}


fmt.Println(cards)
cards = append(cards, "6 lı karo ")
fmt.Println(cards)


cards.samet()

}

func newCard () string { 
	return "karo beşli"
}