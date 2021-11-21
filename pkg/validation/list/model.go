package list

type Model struct {
	ID          *string
	Name        *string
	Description *string
	Website *string
	Staking Staking
	Payout  Payout
	Status  ValidatorStatus
}

type Staking struct {
	FreeSpace         int
	MinDelegation     int
	OpenForDelegation bool
}

type Payout struct {
	Commission   float64 // in %
	PayoutDelay  int     // in cycles
	PayoutPeriod int
}

type ValidatorStatus struct {
	Disabled bool
	Note     string
}
