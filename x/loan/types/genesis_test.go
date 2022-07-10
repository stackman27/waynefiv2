package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"loan/x/loan/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				LoanList: []types.Loan{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LoanCount: 2,
				TxHistoryList: []types.TxHistory{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TxHistoryCount: 2,
				DepositEarnedList: []types.DepositEarned{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				DepositEarnedCount: 2,
				BorrowAccuredList: []types.BorrowAccured{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				BorrowAccuredCount: 2,
				DepositList: []types.Deposit{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				DepositCount: 2,
				BorrowList: []types.Borrow{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				BorrowCount: 2,
				AprList: []types.Apr{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				AprCount: 2,
				UserList: []types.User{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				UserCount: 2,
				PoolList: []types.Pool{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				PoolCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated loan",
			genState: &types.GenesisState{
				LoanList: []types.Loan{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid loan count",
			genState: &types.GenesisState{
				LoanList: []types.Loan{
					{
						Id: 1,
					},
				},
				LoanCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated txHistory",
			genState: &types.GenesisState{
				TxHistoryList: []types.TxHistory{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid txHistory count",
			genState: &types.GenesisState{
				TxHistoryList: []types.TxHistory{
					{
						Id: 1,
					},
				},
				TxHistoryCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated depositEarned",
			genState: &types.GenesisState{
				DepositEarnedList: []types.DepositEarned{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid depositEarned count",
			genState: &types.GenesisState{
				DepositEarnedList: []types.DepositEarned{
					{
						Id: 1,
					},
				},
				DepositEarnedCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated borrowAccured",
			genState: &types.GenesisState{
				BorrowAccuredList: []types.BorrowAccured{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid borrowAccured count",
			genState: &types.GenesisState{
				BorrowAccuredList: []types.BorrowAccured{
					{
						Id: 1,
					},
				},
				BorrowAccuredCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated deposit",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid deposit count",
			genState: &types.GenesisState{
				DepositList: []types.Deposit{
					{
						Id: 1,
					},
				},
				DepositCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated borrow",
			genState: &types.GenesisState{
				BorrowList: []types.Borrow{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid borrow count",
			genState: &types.GenesisState{
				BorrowList: []types.Borrow{
					{
						Id: 1,
					},
				},
				BorrowCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated apr",
			genState: &types.GenesisState{
				AprList: []types.Apr{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid apr count",
			genState: &types.GenesisState{
				AprList: []types.Apr{
					{
						Id: 1,
					},
				},
				AprCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated user",
			genState: &types.GenesisState{
				UserList: []types.User{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid user count",
			genState: &types.GenesisState{
				UserList: []types.User{
					{
						Id: 1,
					},
				},
				UserCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated pool",
			genState: &types.GenesisState{
				PoolList: []types.Pool{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid pool count",
			genState: &types.GenesisState{
				PoolList: []types.Pool{
					{
						Id: 1,
					},
				},
				PoolCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
