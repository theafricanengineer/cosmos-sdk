package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// StoreKey is string representation of the store key for auth
	StoreKey = "acc"

	// FeeStoreKey is a string representation of the store key for fees
	FeeStoreKey = "fee"

	// QuerierRoute is the querier route for acc
	QuerierRoute = StoreKey
)

var (
	// AddressStoreKeyPrefix prefix for account-by-address store
	AddressStoreKeyPrefix = []byte{0x01}

	globalAccountNumberKey = []byte("globalAccountNumber")
)

// AddressStoreKey turn an address to key used to get it from the account store
func AddressStoreKey(addr sdk.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}