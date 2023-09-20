package entity

type SavingAccount struct {
	Holder                     Holder
	Agency, Account, Operation int
	balance                    float64
}

func (s *SavingAccount) Withdraw(value float64) string {
	isAbleToWithdraw := value > 0 && value <= s.balance
	if !isAbleToWithdraw {
		return "Not enough balance."
	}
	s.balance -= value
	return "Withdraw executed successfully!"
}

func (s *SavingAccount) Deposit(value float64) (string, float64) {
	if value < 0 {
		return "Deposit not executed. Negative value", s.balance
	}
	s.balance += value
	return "Deposit executed successfully!", s.balance
}

func (s *SavingAccount) GetBalance() float64 {
	return s.balance
}
