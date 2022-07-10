package types

const (
	// ModuleName defines the module name
	ModuleName = "loan"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_loan"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	LoanKey      = "Loan-value-"
	LoanCountKey = "Loan-count-"
)

const (
	TxHistoryKey      = "TxHistory-value-"
	TxHistoryCountKey = "TxHistory-count-"
)

const (
	DepositEarnedKey      = "DepositEarned-value-"
	DepositEarnedCountKey = "DepositEarned-count-"
)

const (
	BorrowAccuredKey      = "BorrowAccured-value-"
	BorrowAccuredCountKey = "BorrowAccured-count-"
)

const (
	DepositKey      = "Deposit-value-"
	DepositCountKey = "Deposit-count-"
)

const (
	BorrowKey      = "Borrow-value-"
	BorrowCountKey = "Borrow-count-"
)

const (
	AprKey      = "Apr-value-"
	AprCountKey = "Apr-count-"
)

const (
	UserKey      = "User-value-"
	UserCountKey = "User-count-"
)

const (
	PoolKey      = "Pool-value-"
	PoolCountKey = "Pool-count-"
)
