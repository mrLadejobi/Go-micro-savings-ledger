package main

import (
	"fmt"
	"errors"
)

type SavingGoal struct {
	ItemName string
	TargetAmount int
	CurrentSavings int
	IsAchieved bool
}
 
func (s *SavingGoal) DepositMoney(amount int) error {
	if amount <= 0 {
		return errors.New("Deposited Amount cannot be 0, please try again starting from 5000")
	}

	s.CurrentSavings = s.CurrentSavings + amount

	

	if s.CurrentSavings >= s.TargetAmount {
		s.IsAchieved = true

	}
	return nil
}

func (s SavingGoal) ViewStatus() {
	fmt.Println("You have saved:", s.CurrentSavings, "out of:", s.TargetAmount)
}

func main() {
    myItem := SavingGoal{
        ItemName:       "Macbook Pro",
        TargetAmount:   500000,
        CurrentSavings: 0,
    }

    // TEST CASE 1: Valid Deposit
    fmt.Println("--- Test 1: Depositing ₦200,000 ---")
    myItem.DepositMoney(200000)
    myItem.ViewStatus()

    // TEST CASE 2: The Attack (Negative Deposit)
    fmt.Println("\n--- Test 2: Trying to deposit -5,000 ---")
    err := myItem.DepositMoney(-5000)
    if err != nil {
        fmt.Println("❌ ALERT:", err) // This should catch your custom error!
    }
    myItem.ViewStatus() // Balance should STILL be 200,000

    // TEST CASE 3: Smashing the Goal
    fmt.Println("\n--- Test 3: Depositing ₦300,000 to finish goal ---")
    myItem.DepositMoney(300000)
    myItem.ViewStatus()
    
    // Check if the flag flipped to true!
    fmt.Println("Is Goal Achieved?:", myItem.IsAchieved) 
}