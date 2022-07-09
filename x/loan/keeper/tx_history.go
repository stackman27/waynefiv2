package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"loan/x/loan/types"
)

// GetTxHistoryCount get the total number of txHistory
func (k Keeper) GetTxHistoryCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TxHistoryCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTxHistoryCount set the total number of txHistory
func (k Keeper) SetTxHistoryCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TxHistoryCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTxHistory appends a txHistory in the store with a new id and update the count
func (k Keeper) AppendTxHistory(
	ctx sdk.Context,
	txHistory types.TxHistory,
) uint64 {
	// Create the txHistory
	count := k.GetTxHistoryCount(ctx)

	// Set the ID of the appended value
	txHistory.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	appendedValue := k.cdc.MustMarshal(&txHistory)
	store.Set(GetTxHistoryIDBytes(txHistory.Id), appendedValue)

	// Update txHistory count
	k.SetTxHistoryCount(ctx, count+1)

	return count
}

// SetTxHistory set a specific txHistory in the store
func (k Keeper) SetTxHistory(ctx sdk.Context, txHistory types.TxHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	b := k.cdc.MustMarshal(&txHistory)
	store.Set(GetTxHistoryIDBytes(txHistory.Id), b)
}

// GetTxHistory returns a txHistory from its id
func (k Keeper) GetTxHistory(ctx sdk.Context, id uint64) (val types.TxHistory, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	b := store.Get(GetTxHistoryIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTxHistory removes a txHistory from the store
func (k Keeper) RemoveTxHistory(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	store.Delete(GetTxHistoryIDBytes(id))
}

// GetAllTxHistory returns all txHistory
func (k Keeper) GetAllTxHistory(ctx sdk.Context) (list []types.TxHistory) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TxHistoryKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TxHistory
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTxHistoryIDBytes returns the byte representation of the ID
func GetTxHistoryIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTxHistoryIDFromBytes returns ID in uint64 format from a byte array
func GetTxHistoryIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
