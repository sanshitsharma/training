package main

import "testing"

func TestDeposit(t *testing.T) {
    balance := float32(100)
    deposit := float32(200)
    if Deposit(balance, deposit) != 300 {
        t.Error("Unit test for Deposit failed")
    }
}

func TestWithdrawal(t *testing.T) {
    balance := float32(400)
    withdrawal := float32(200)
    result, _ := Withdrawal(balance, withdrawal)
    if result != 200 {
        t.Error("Unit test for Withdrawal failed")
    }
}

