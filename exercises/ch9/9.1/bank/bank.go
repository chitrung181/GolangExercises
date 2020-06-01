package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
//Deposit ...
func Deposit(amount int) { deposits <- amount }

//Balance ...
func Balance() int { return <-balances }

//Withdraw ...
func Withdraw(amount int) bool {
	if Balance() < amount {
		return false
	}
	deposits <- (-amount)
	return true
}
func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func init() {
	go teller() // start the monitor goroutine
}
