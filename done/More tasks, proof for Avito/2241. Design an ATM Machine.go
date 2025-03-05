package main

var emptyWithdraw = []int{-1}

type ATM struct {
	nominals       []int // 20, 50, 100, 200, 500
	banknotesCount []int
}

func Constructor5() ATM {
	nominals := []int{20, 50, 100, 200, 500}
	banknotesCount := make([]int, len(nominals))
	return ATM{nominals, banknotesCount}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i, cnt := range banknotesCount {
		this.banknotesCount[i] += cnt
	}
}

func (this *ATM) Withdraw(amount int) []int {
	result := make([]int, len(this.nominals))
	for i := len(this.nominals) - 1; i >= 0; i-- {
		cnt := amount / this.nominals[i]
		takenCnt := min(this.banknotesCount[i], cnt)
		result[i] = takenCnt
		amount -= takenCnt * this.nominals[i]
	}

	if amount != 0 {
		return emptyWithdraw
	}

	for i, cnt := range result { // важно вычитать банкноты, если транзакция успешна
		this.banknotesCount[i] -= cnt
	}

	return result
}

/**
 * Your ATM object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Deposit(banknotesCount);
 * param_2 := obj.Withdraw(amount);
 */
