package list

type (
	Model struct {
		ID          *string
		Name        *string
		Description *string
		Website     *string
		Staking     Staking
		Payout      Payout
		Status      Status
	}

	Staking struct {
		FreeSpace         int
		MinDelegation     int
		OpenForDelegation bool
	}

	Payout struct {
		Commission   float64 // In percents (%).
		PayoutDelay  int     // In cycles.
		PayoutPeriod int
	}

	Status struct {
		Disabled bool
		Note     string
	}
)
