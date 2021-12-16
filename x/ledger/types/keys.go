package types

const (
	// ModuleName defines the module name
	ModuleName = "ledger"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

    // QuerierRoute defines the module's query routing key
    QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_ledger"


)

var (
	PoolPrefix      = []byte{0x00}
	BondedPoolPrefix = []byte{0x01}
	EraUnbondLimitPrefix = []byte{0x02}
	ChainBondingDurationPrefix = []byte{0x03}
	PoolDetailPrefix = []byte{0x04}
	LeastBondPrefix = []byte{0x05}
	SnapShotPrefix = []byte{0x06}
	CurrentEraSnapShotPrefix = []byte{0x07}
	BondPipelinePrefix = []byte{0x08}
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
