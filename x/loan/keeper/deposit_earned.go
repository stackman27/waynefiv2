package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

// GetDepositEarnedCount get the total number of depositEarned
func (k Keeper) GetDepositEarnedCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DepositEarnedCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDepositEarnedCount set the total number of depositEarned
func (k Keeper) SetDepositEarnedCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DepositEarnedCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDepositEarned appends a depositEarned in the store with a new id and update the count
func (k Keeper) AppendDepositEarned(
	ctx sdk.Context,
	depositEarned types.DepositEarned,
) uint64 {
	// Create the depositEarned
	count := k.GetDepositEarnedCount(ctx)

	// Set the ID of the appended value
	depositEarned.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	appendedValue := k.cdc.MustMarshal(&depositEarned)
	store.Set(GetDepositEarnedIDBytes(depositEarned.Id), appendedValue)

	// Update depositEarned count
	k.SetDepositEarnedCount(ctx, count+1)

	return count
}

// SetDepositEarned set a specific depositEarned in the store
func (k Keeper) SetDepositEarned(ctx sdk.Context, depositEarned types.DepositEarned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	b := k.cdc.MustMarshal(&depositEarned)
	store.Set(GetDepositEarnedIDBytes(depositEarned.Id), b)
}

// GetDepositEarned returns a depositEarned from its id
func (k Keeper) GetDepositEarned(ctx sdk.Context, id uint64) (val types.DepositEarned, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	b := store.Get(GetDepositEarnedIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDepositEarned removes a depositEarned from the store
func (k Keeper) RemoveDepositEarned(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	store.Delete(GetDepositEarnedIDBytes(id))
}

// GetAllDepositEarned returns all depositEarned
func (k Keeper) GetAllDepositEarned(ctx sdk.Context) (list []types.DepositEarned) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DepositEarnedKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.DepositEarned
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDepositEarnedIDBytes returns the byte representation of the ID
func GetDepositEarnedIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDepositEarnedIDFromBytes returns ID in uint64 format from a byte array
func GetDepositEarnedIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
