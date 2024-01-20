# BillBook

- An app to create, store and manage bills.
- This app is built using Go.
  
## Quick Start

- Clone the repo: ```bash git clone https://github.com/Justtejas/Bills-App.git ```
- run ```bash go run main.go bill.go``` in the root directory.

## Example

- ```bash go run main.go bill.go```
- ```bash
    Enter bill name: John\`s Bill
    Created the bill -  Created the bill -  Rent
    Choose Options (a - add item, s - save bill, t - add tip): a
    Whats the item name: Coffee
    Whats the price: 10
    Item added Coffee 10
    Choose Options (a - add item, s - save bill, t - add tip): t
    Enter the tip amount: 2
    Choose Options (a - add item, s - save bill, t - add tip): s
    Bill was saved

    In bills directory a file was created with the name John\`s Bill.txt
    Bill BreakDown: 
    Coffee:        ...$10
    tip            ...$2
    total:         ...$12.00
    ```