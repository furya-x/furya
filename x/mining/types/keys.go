package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "mining"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_mining"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var RewardFactor = sdk.NewInt(1e12)

var (
	StakePoolStoreKeyPrefix            = []byte{0x01}
	StakeItemStoreKeyPrefix            = []byte{0x02}
	UserStakeRecordStoreKeyPrefix      = []byte{0x03}
	UserStakeRecordIndexStoreKeyPrefix = []byte{0x04}
	RewardPoolIndexStoreKeyPrefix      = []byte{0x05}
	StakeItemIndexStoreKey             = []byte{0x06}
	RewarderStoreKeyPrefix             = []byte{0x07}
	StakePoolIndexStoreKey             = []byte{0x08}
	RewardTokenStoreKeyPrefix          = []byte{0x09}
	MaxRewardPoolNumberStoreKey        = []byte{0x0a}
	MiningProviderSwitchStoreKey       = []byte{0x0b}
)

var (
	SwitchStateClose = []byte{0x00}
	SwitchStateOpen  = []byte{0x01}
)

func StakePoolStoreKey(index uint32) []byte {
	bts := make([]byte, 4)
	binary.LittleEndian.PutUint32(bts, index)
	return append(StakePoolStoreKeyPrefix, bts...)
}

func StakeItemStoreKey(index uint32) []byte {
	indexBts := make([]byte, 4)
	binary.LittleEndian.PutUint32(indexBts, index)
	return append(StakeItemStoreKeyPrefix, indexBts...)
}

// prefix + len(userAddress) + userAddress + stakePoolIndex + index
func UserStakeRecordStoreKey(userAddress string, stakePoolIndex, index uint32) []byte {
	userAddressLen := len(userAddress)

	key := make([]byte, 1+1+userAddressLen+4+4)
	key[0] = UserStakeRecordStoreKeyPrefix[0]
	key[1] = byte(len(userAddress))
	copy(key[2:2+userAddressLen], userAddress)

	binary.LittleEndian.PutUint32(key[2+userAddressLen:], stakePoolIndex)
	binary.LittleEndian.PutUint32(key[2+userAddressLen+4:], index)
	return key
}

func RewarderStoreKey(addr sdk.AccAddress) []byte {
	return append(RewarderStoreKeyPrefix, addr.Bytes()...)
}

func RewardTokenStoreKey(denom string) []byte {
	return append(RewardTokenStoreKeyPrefix, []byte(denom)...)
}
