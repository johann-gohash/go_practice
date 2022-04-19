package main


import (
    "fmt"
    "errors"
)

var (
    ErrInsufficientBalance = errors.New("Insufficient balance.")
    ErrOverdraftIncurred= errors.New("Overdraft Incurred")
)

type BankAccount struct {
    Owner string // Consider capitalized variables to be public attributes.
    balance float64 // Consider lowercase variables to be private attributes.
}

func (ba BankAccount) Balance() float64 {
    return ba.balance
}

// When writing methods that change a struct's data ( as opposed to just read )
// we pass in pointers to the struct. That way we don't pass in a copy through
// pass-by-value.
func (ba *BankAccount) Deposit(amount float64) {
    //fmt.Println("In Deposit method")
    //fmt.Println(*ba)

    // Why don't we have to dereference the struct to use its method?
    // This seems like an important question.
    ba.balance += amount
}

func (ba *BankAccount) Withdraw(amount float64) (float64, error) {
    if ba.balance < amount {
        return 0, ErrInsufficientBalance
    }
    ba.balance -= amount
    return ba.balance, nil
}

type OverdraftableBankAccount struct {
    BankAccount
    Fee float64
}

func (oba * OverdraftableBankAccount) Withdraw(amount float64) (float64, error) {
    var err error
    if oba.balance < amount {
        oba.balance -= oba.Fee
        err = ErrOverdraftIncurred
    }
    oba.balance -= amount
    return oba.balance, err

}

func main() {
    a1 := BankAccount{"Bob", 50}
    a1.Deposit(20)
    fmt.Println("Balance:", a1.Balance())
    a2 := OverdraftableBankAccount{BankAccount{"Jill", 100}, 20}
    a2.Deposit(20)
    fmt.Println("Balance for Jill:", a2.Balance())
}
