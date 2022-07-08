/* eslint-disable array-callback-return */
import React, {useState, useEffect} from 'react'
import useAxios from 'axios-hooks'
import { Keplr } from "@keplr-wallet/types";
import { OfflineSigner } from "@cosmjs/launchpad";
import loanLoan from '../generated/loan/loan.loan';

interface rootGetters {
    apiNode: string; 
    rpcNode: string; 
    wsNode: string;
    chainId: string; 
    addrPrefix: string; 
    sdkVersion: string;
    getTXApi: string; 
    offlineSigner: OfflineSigner | null;
}

function Loans(props: {userAddr: string, keplr: Keplr | null, rootGetters: rootGetters | null}) {

    const [state, setState] = React.useState({
        amount: 0,
        collateral: 0,
      })

    const [{ data: userData}] = useAxios(
        `http://0.0.0.0:1317/cosmos/auth/v1beta1/accounts/${props.userAddr}`
    )

    const [{ data: userBankData}] = useAxios(
        `http://0.0.0.0:1317/cosmos/bank/v1beta1/balances/${props.userAddr}/by_denom?denom=uloan`
    )
    
    const [{data: loanData}] = useAxios(
        `http://0.0.0.0:1317/loan/loan/loan`
    )

    
    const accountNumber =
       (userData && userData.account.account_number) || "0";
     const sequence = (userData && userData.account.sequence) || "0";


    const requestLoan = async() => {
        if(!props.keplr) {
            throw Error("Keplr isn't connected")
        }       

        var data = {
            creator: props.userAddr,
            amount: state.amount + "uloan",
            fee: 200 + "uloan",
            collateral: state.collateral + "uloan", 
            deadline: "1000",
        }


  
        const newTx = await loanLoan.actions
        .sendMsgRequestLoan({rootGetters: props.rootGetters},
                            {value: data})

        console.log(newTx)
    }

    const approveLoan = async (loanId) => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }

        // start creating msg for approving loan
        var data = {
            creator: props.userAddr,
            id: loanId,
        }
 
        const newTx = await loanLoan.actions
                .sendMsgApproveLoan({rootGetters: props.rootGetters},
                {value: data})

        console.log(newTx)
    }

    const repayLoan = async (loanId) => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }
        
        var data = {
            creator: props.userAddr,
            id: loanId
        }

        const newTx = await loanLoan.actions
        .sendMsgRepayLoan({rootGetters: props.rootGetters},
                            {value: data})

console.log(newTx)
    }

    const liquidateLoan = async (loanId) => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }
        
        var data = {
            creator: props.userAddr,
            id: loanId
        }

        const newTx = await loanLoan.actions
        .sendMsgLiquidateLoan({rootGetters: props.rootGetters},
        {value: data})

        console.log(newTx)
    }
    
    const cancelLoan = async(loanId) => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }
        
        var data = {
            creator: props.userAddr,
            id: loanId
        }

        const newTx = await loanLoan.actions
        .sendMsgCancelLoan({rootGetters: props.rootGetters},
                            {value: data})

        console.log(newTx)
    }

    function handleChange(evt) {
        const value = evt.target.value;
        setState({
          ...state,
          [evt.target.name]: value
        });
      }

    return (
    <div style = {{textAlign: 'left', width: '80%', marginLeft: 'auto', marginRight: 'auto',marginTop: 0}}>
        <div style={{float: 'right',}}> 
            <h3> User Balance: {userBankData && userBankData.balance.amount.concat(userBankData.balance.denom) } </h3>
        </div>
        <br/> <br/>
        <div>
            <h3 style = {{color: 'red'}}> Request Loan </h3>
       
                <label style = {{float: 'left'}}> 
                    Amount: <br/>
                    <input type="text" name="amount" value={state.amount} onChange = {handleChange}/>
                </label>

                <label  style = {{float: 'left', marginLeft: 20}}>
                    Fixed Fee: <br/>
                    <input type="text" name="fixedFee" disabled= {true} placeholder = {"200uloan"}/>
                </label>

                <label  style = {{float: 'left', marginLeft: 20}}>
                    Collateral:<br/>
                    <input type="text" name="collateral" value={state.collateral} onChange={handleChange}/>
                </label>

                <label  style = {{float: 'left', marginLeft: 20}}> 
                Deadline: <br/>
                <input type="text" name="fixedFee" disabled= {true} placeholder = {"1000 block height"}/> 
                </label>
               
                <button onClick={() => requestLoan()} style = {{marginTop: 20, marginLeft: 25}}> Request Loan </button> 
                
        
        </div>
        <br/>
        <hr/>
        <h2 style = {{color: 'red'}}> All Loans </h2>
        <div style = {{}}> 
            {loanData && loanData.Loan.map((loan: { amount: string, id: string , borrower: string, state: string, collateral: string, deadline:string}) => (
                <div key={loan.id} style = {{height: '300px', width: '420px', background: '#eee', textAlign: 'left', padding: 10, float: 'left', margin: 20, }}>
                        <h2>Loan Id: {loan.id}</h2>
                        <p> Loan Amount: {loan.amount}</p>
                        <p> Loan Borrower: {loan.borrower} </p>
                        <p> Loan State: {loan.state} </p>
                        <p> Loan Collateral: {loan.collateral}</p>
                        <p> Loan Deadline: {loan.deadline} </p>
                    <button onClick={() => approveLoan(loan.id)} disabled = {props.userAddr !== loan.borrower ? false : true}> Approve Loan </button> &nbsp;
                    <button onClick={() => repayLoan(loan.id)} disabled = {loan.state === "approved" ? false : true}> Repay Loan </button> &nbsp;
                    <button onClick={() => liquidateLoan(loan.id)} disabled = {loan.state === "approved" ? false : true}> Liquidate Loan </button> &nbsp;
                    <button onClick={() => cancelLoan(loan.id)} disabled = {props.userAddr === loan.borrower ? false : true}> Cancel Loan </button>
                </div>
            ))}
       </div> 
    </div>
  )
}

export default Loans 

