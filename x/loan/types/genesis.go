package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LoanList:          []Loan{},
		TxHistoryList:     []TxHistory{},
		DepositEarnedList: []DepositEarned{},
		BorrowAccuredList: []BorrowAccured{},
		DepositList:       []Deposit{},
		BorrowList:        []Borrow{},
		AprList:           []Apr{},
		UserList:          []User{},
		PoolList:          []Pool{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in loan
	loanIdMap := make(map[uint64]bool)
	loanCount := gs.GetLoanCount()
	for _, elem := range gs.LoanList {
		if _, ok := loanIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for loan")
		}
		if elem.Id >= loanCount {
			return fmt.Errorf("loan id should be lower or equal than the last id")
		}
		loanIdMap[elem.Id] = true
	}
	// Check for duplicated ID in txHistory
	txHistoryIdMap := make(map[uint64]bool)
	txHistoryCount := gs.GetTxHistoryCount()
	for _, elem := range gs.TxHistoryList {
		if _, ok := txHistoryIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for txHistory")
		}
		if elem.Id >= txHistoryCount {
			return fmt.Errorf("txHistory id should be lower or equal than the last id")
		}
		txHistoryIdMap[elem.Id] = true
	}
	// Check for duplicated ID in depositEarned
	depositEarnedIdMap := make(map[uint64]bool)
	depositEarnedCount := gs.GetDepositEarnedCount()
	for _, elem := range gs.DepositEarnedList {
		if _, ok := depositEarnedIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for depositEarned")
		}
		if elem.Id >= depositEarnedCount {
			return fmt.Errorf("depositEarned id should be lower or equal than the last id")
		}
		depositEarnedIdMap[elem.Id] = true
	}
	// Check for duplicated ID in borrowAccured
	borrowAccuredIdMap := make(map[uint64]bool)
	borrowAccuredCount := gs.GetBorrowAccuredCount()
	for _, elem := range gs.BorrowAccuredList {
		if _, ok := borrowAccuredIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for borrowAccured")
		}
		if elem.Id >= borrowAccuredCount {
			return fmt.Errorf("borrowAccured id should be lower or equal than the last id")
		}
		borrowAccuredIdMap[elem.Id] = true
	}
	// Check for duplicated ID in deposit
	depositIdMap := make(map[uint64]bool)
	depositCount := gs.GetDepositCount()
	for _, elem := range gs.DepositList {
		if _, ok := depositIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for deposit")
		}
		if elem.Id >= depositCount {
			return fmt.Errorf("deposit id should be lower or equal than the last id")
		}
		depositIdMap[elem.Id] = true
	}
	// Check for duplicated ID in borrow
	borrowIdMap := make(map[uint64]bool)
	borrowCount := gs.GetBorrowCount()
	for _, elem := range gs.BorrowList {
		if _, ok := borrowIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for borrow")
		}
		if elem.Id >= borrowCount {
			return fmt.Errorf("borrow id should be lower or equal than the last id")
		}
		borrowIdMap[elem.Id] = true
	}
	// Check for duplicated ID in apr
	aprIdMap := make(map[uint64]bool)
	aprCount := gs.GetAprCount()
	for _, elem := range gs.AprList {
		if _, ok := aprIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for apr")
		}
		if elem.Id >= aprCount {
			return fmt.Errorf("apr id should be lower or equal than the last id")
		}
		aprIdMap[elem.Id] = true
	}
	// Check for duplicated ID in user
	userIdMap := make(map[uint64]bool)
	userCount := gs.GetUserCount()
	for _, elem := range gs.UserList {
		if _, ok := userIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for user")
		}
		if elem.Id >= userCount {
			return fmt.Errorf("user id should be lower or equal than the last id")
		}
		userIdMap[elem.Id] = true
	}
	// Check for duplicated ID in pool
	poolIdMap := make(map[uint64]bool)
	poolCount := gs.GetPoolCount()
	for _, elem := range gs.PoolList {
		if _, ok := poolIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for pool")
		}
		if elem.Id >= poolCount {
			return fmt.Errorf("pool id should be lower or equal than the last id")
		}
		poolIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
