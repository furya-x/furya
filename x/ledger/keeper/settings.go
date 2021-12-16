package keeper

import (
	//gogotypes "github.com/gogo/protobuf/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/stafiprotocol/stafihub/x/ledger/types"
)

func (k Keeper) AddPool(ctx sdk.Context, denom string, addr string) error {
	pool, found := k.TryFindPool(ctx, denom, addr, types.PoolPrefix)
	if found {
		return types.ErrPoolAlreadyAdded
	}

	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(denom), b)
	return nil
}

func (k Keeper) SetPool(ctx sdk.Context, pool *types.Pool, pref []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), pref)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(pool.Denom), b)
}

func (k Keeper) SetEraUnbondLimit(ctx sdk.Context, denom string, limit uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	eul := &types.EraUnbondLimit{
		Denom: denom,
		Limit: limit,
	}
	b := k.cdc.MustMarshal(eul)
	store.Set([]byte(denom), b)
}

func (k Keeper) EraUnbondLimit(ctx sdk.Context, denom string) (val *types.EraUnbondLimit, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.EraUnbondLimitPrefix)
	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetInitBond(ctx sdk.Context, denom, pool string, amount *sdk.Int, receiver sdk.AccAddress) error {
	// todo use cacheContext
	_, found := k.TryFindPool(ctx, denom, pool, types.PoolPrefix)
	if !found {
		return types.ErrPoolNotFound
	}

	bpool, found := k.TryFindPool(ctx, denom, pool, types.BondedPoolPrefix)
	if found {
		return types.ErrRepeatInitBond
	}

	rbalance := k.rateKeeper.TokenToRtoken(ctx, denom, *amount)
	rcoins := sdk.Coins{
		sdk.NewCoin(denom, rbalance),
	}

	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, rcoins); err != nil {
		return err
	}

	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, rcoins)
	if err != nil {
		return err
	}

	if k.rateKeeper.GetRate(ctx, denom) == nil {
		k.rateKeeper.SetRate(ctx, denom, sdk.NewInt(0), sdk.NewInt(0))
	}

	bond := sdk.NewInt(0)
	unbond := sdk.NewInt(0)
	active := sdk.NewInt(0)
	pipe := &types.BondPipeline{
		Denom: denom,
		Pool: pool,
		Chunk: &types.LinkChunk{
			Bond: &bond,
			Unbond: &unbond,
			Active: &active,
		},
	}

	k.AddBondedPool(ctx, bpool)
	k.SetPipeline(ctx, pipe)
	return nil
}

func (k Keeper) AddBondedPool(ctx sdk.Context, pool *types.Pool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)
	b := k.cdc.MustMarshal(pool)
	store.Set([]byte(pool.Denom), b)
}

func (k Keeper) GetBondedPoolByDenom(ctx sdk.Context, denom string) (val *types.Pool, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondedPoolPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetChainBondingDuration(ctx sdk.Context, denom string, era uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)
	cbd := &types.ChainBondingDuration{
		Denom: denom,
		Era: era,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom), b)
}

func (k Keeper) ChainBondingDuration(ctx sdk.Context, denom string) (val *types.ChainBondingDuration, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.ChainBondingDurationPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetPoolDetail(ctx sdk.Context, denom string, pool string, subAccounts []string, threshold uint32) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)
	cbd := &types.PoolDetail{
		Denom: denom,
		Pool: pool,
		SubAccounts: subAccounts,
		Threshold: threshold,
	}
	b := k.cdc.MustMarshal(cbd)
	store.Set([]byte(denom+pool), b)
}

func (k Keeper) PoolDetail(ctx sdk.Context, denom string, pool string) (val *types.PoolDetail, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolDetailPrefix)

	b := store.Get([]byte(denom+pool))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) SetLeastBond(ctx sdk.Context, denom string, amount *sdk.Int) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)
	lb := &types.LeastBond{
		Denom: denom,
		Amount: amount,
	}
	b := k.cdc.MustMarshal(lb)
	store.Set([]byte(denom), b)
}

func (k Keeper) LeastBond(ctx sdk.Context, denom string) (val *types.LeastBond, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.LeastBondPrefix)

	b := store.Get([]byte(denom))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}

func (k Keeper) ClearCurrentEraSnapShots(ctx sdk.Context, denom string) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.CurrentEraSnapShotPrefix)

	cess := &types.CurrentEraSnapShot{
		Denom: denom,
		ShotIds: [][]byte{},
	}
	b := k.cdc.MustMarshal(cess)
	store.Set([]byte(denom), b)
}

//func (k Keeper) SetBondPipeLine(ctx sdk.Context, denom string, )

func (k Keeper) BondPipeLine(ctx sdk.Context, denom string, pool string) (val *types.BondPipeline, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.BondPipelinePrefix)

	b := store.Get([]byte(denom+pool))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, val)
	return val, true
}




func (k Keeper) TryFindPool(ctx sdk.Context, denom, addr string, pref []byte) (val *types.Pool, found bool) {
	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), pref)

	b := store.Get([]byte(denom))
	if b == nil {
		return &types.Pool{
			Denom: denom,
			Addrs: map[string]bool{addr: true},
		}, false
	}

	k.cdc.MustUnmarshal(b, val)
	if _, ok := val.Addrs[addr]; ok {
		return val, true
	}

	val.Addrs[addr] = true
	return val, false
}

func (k Keeper) SetPipeline(ctx sdk.Context, pipe *types.BondPipeline) {

}

//func (k Keeper) GetPoolByDenom(ctx sdk.Context, denom string) (val *types.Pool, found bool) {
//	store :=  prefix.NewStore(ctx.KVStore(k.storeKey), types.PoolPrefix)
//
//	b := store.Get([]byte(denom))
//	if b == nil {
//		return val, false
//	}
//
//	k.cdc.MustUnmarshal(b, val)
//	return val, true
//}






