package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format bill
func (b *bill) format() string {
	fs := "Bill BreakDown: \n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v...$%v\n", k+":", v)
		total += v
	}
	fs += fmt.Sprintf("%-15v...$%v\n", "tip", b.tip)
	//total
	fs += fmt.Sprintf("%-15v...$%0.2f", "total:", total+b.tip)
	return fs
}

// add item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// save bill
func (b *bill) save() {
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved")
}
