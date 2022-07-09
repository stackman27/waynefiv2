package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

// GetBorrowAccuredCount get the total number of borrowAccured
func (k Keeper) GetBorrowAccuredCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BorrowAccuredCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetBorrowAccuredCount set the total number of borrowAccured
func (k Keeper) SetBorrowAccuredCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.BorrowAccuredCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendBorrowAccured appends a borrowAccured in the store with a new id and update the count
func (k Keeper) AppendBorrowAccured(
	ctx sdk.Context,
	borrowAccured types.BorrowAccured,
) uint64 {
	// Create the borrowAccured
	count := k.GetBorrowAccuredCount(ctx)

	// Set the ID of the appended value
	borrowAccured.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccuredKey))
	appendedValue := k.cdc.MustMarshal(&borrowAccured)
	store.Set(GetBorrowAccuredIDBytes(borrowAccured.Id), appendedValue)

	// Update borrowAccured count
	k.SetBorrowAccuredCount(ctx, count+1)

	return count
}

// SetBorrowAccured set a specific borrowAccured in the store
func (k Keeper) SetBorrowAccured(ctx sdk.Context, borrowAccured types.BorrowAccured) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccuredKey))
	b := k.cdc.MustMarshal(&borrowAccured)
	store.Set(GetBorrowAccuredIDBytes(borrowAccured.Id), b)
}

// GetBorrowAccured returns a borrowAccured from its id
func (k Keeper) GetBorrowAccured(ctx sdk.Context, id uint64) (val types.BorrowAccured, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccuredKey))
	b := store.Get(GetBorrowAccuredIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveBorrowAccured removes a borrowAccured from the store
func (k Keeper) RemoveBorrowAccured(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccuredKey))
	store.Delete(GetBorrowAccuredIDBytes(id))
}

// GetAllBorrowAccured returns all borrowAccured
func (k Keeper) GetAllBorrowAccured(ctx sdk.Context) (list []types.BorrowAccured) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BorrowAccuredKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.BorrowAccured
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBorrowAccuredIDBytes returns the byte representation of the ID
func GetBorrowAccuredIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBorrowAccuredIDFromBytes returns ID in uint64 format from a byte array
func GetBorrowAccuredIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
