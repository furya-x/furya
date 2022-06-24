package types

import (
// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:               DefaultParams(),
		StakePoolList:        []*StakePool{},
		StakeItemList:        []*StakeItem{},
		UserStakeRecordList:  []*UserStakeRecord{},
		MiningProviderList:   []string{},
		RewardTokenList:      []*RewardToken{},
		MaxRewardPoolNumber:  DefaultMaxRewardPoolNumber,
		MiningProviderSwitch: true,
		MaxStakeItemNumber:   DefaultMaxStakeItemNumber,
		StakeTokenList:       []string{},
		StakeItemLimit: &StakeItemLimit{
			MaxPowerRewardRate: DefaultMaxPowerRewardRate,
			MaxLockSecond:      DefaultMaxLockSecond,
		},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
