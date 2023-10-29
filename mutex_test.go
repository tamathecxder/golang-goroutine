package golanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type UserBalance struct {
	sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(userSender *UserBalance, userReceiver *UserBalance, amount int) {
	userSender.Lock()
	fmt.Println("Lock Sender:", userSender.Name)
	userSender.Change(-amount)

	time.Sleep(1 * time.Second)

	userReceiver.Lock()
	fmt.Println("Lock Receiver:", userReceiver.Name)
	userReceiver.Change(amount)

	time.Sleep(1 * time.Second)

	userSender.Unlock()
	userReceiver.Unlock()
}

func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Karina",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Winter",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000) // karina = 90.000, winter = 110.000
	go Transfer(&user2, &user1, 200000)

	time.Sleep(5 * time.Second)

	fmt.Println("User 1", user1.Name, ", Balance:", user1.Balance)
	fmt.Println("User 2", user2.Name, ", Balance:", user2.Balance)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()

	return balance
}

func TestReadWriteMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Total Balance:", account.GetBalance())
}

func TestMutex(t *testing.T) {
	x := 0

	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)

	fmt.Println("X value is", x)
}
