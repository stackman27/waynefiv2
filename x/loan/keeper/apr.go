package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

// GetAprCount get the total number of apr
func (k Keeper) GetAprCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AprCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetAprCount set the total number of apr
func (k Keeper) SetAprCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.AprCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendApr appends a apr in the store with a new id and update the count
func (k Keeper) AppendApr(
	ctx sdk.Context,
	apr types.Apr,
) uint64 {
	// Create the apr
	count := k.GetAprCount(ctx)

	// Set the ID of the appended value
	apr.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	appendedValue := k.cdc.MustMarshal(&apr)
	store.Set(GetAprIDBytes(apr.Id), appendedValue)

	// Update apr count
	k.SetAprCount(ctx, count+1)

	return count
}

// SetApr set a specific apr in the store
func (k Keeper) SetApr(ctx sdk.Context, apr types.Apr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	b := k.cdc.MustMarshal(&apr)
	store.Set(GetAprIDBytes(apr.Id), b)
}

// GetApr returns a apr from its id
func (k Keeper) GetApr(ctx sdk.Context, id uint64) (val types.Apr, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	b := store.Get(GetAprIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveApr removes a apr from the store
func (k Keeper) RemoveApr(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	store.Delete(GetAprIDBytes(id))
}

// GetAllApr returns all apr
func (k Keeper) GetAllApr(ctx sdk.Context) (list []types.Apr) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AprKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Apr
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetAprIDBytes returns the byte representation of the ID
func GetAprIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetAprIDFromBytes returns ID in uint64 format from a byte array
func GetAprIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
