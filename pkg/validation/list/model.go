package list

type (
	Model struct {
		ID          *string
		Name        *string
		Description *string
		Website     *string
		Staking     Staking
		Payout      Payout
		Status      ValidatorStatus
	}

	Staking struct {
		FreeSpace         int
		MinDelegation     int
		OpenForDelegation bool
	}

	Payout struct {
		Commission   float64 // in %
		PayoutDelay  int     // in cycles
		PayoutPeriod int
	}

	ValidatorStatus struct {
		Disabled bool
		Note     string
	}
)
