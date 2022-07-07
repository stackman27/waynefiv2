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
            amount: 1000 + "uloan",
            fee: 100 + "uloan",
            collateral: 1500 + "uloan", 
            deadline: 1000,
        }

        const newTx = await loanLoan.actions
                            .sendMsgRequestLoan({rootGetters: props.rootGetters},
                                                {value: data})

        console.log(newTx)
    }

    const approveLoan = async () => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }

        // start creating msg for approving loan
        var data = {
            creator: props.userAddr,
            id: 0,
        }
 
        const newTx = await loanLoan.actions
                .sendMsgApproveLoan({rootGetters: props.rootGetters},
                {value: data})

        console.log(newTx)
    }
    
    const cancelLoan = () => {
        console.log("CANCEL THIS LOAN")
    }

    return (
    <div style = {{textAlign: 'left', width: '80%', marginLeft: 'auto', marginRight: 'auto'}}>
        <h2> User Balance: {userBankData && userBankData.balance.amount.concat(userBankData.balance.denom) } </h2>
        <div>
            <button onClick={() => requestLoan()}>RequestLoan</button>
        </div>

        <h3 style = {{color: 'red'}}> Loans </h3>
        <div style = {{}}> 
            {loanData && loanData.Loan.map((loan: { amount: string, id: string , borrower: string, state: string, collateral: string, deadline:string}) => (
                <div id={loan.id} style = {{height: '300px', width: '400px', background: '#eee', textAlign: 'left', padding: 10, float: 'left', margin: 20}}>
                        <h2>Loan Id: {loan.id}</h2>
                        <p> Loan Amount: {loan.amount}</p>
                        <p> Loan Borrower: {loan.borrower} </p>
                        <p> Loan State: {loan.state} </p>
                        <p> Loan Collateral: {loan.collateral}</p>
                        <p> Loan Deadline: {loan.deadline} </p>
                    <button onClick={() => approveLoan()}> Approve Loan </button>
                    <button onClick={() => cancelLoan()} disabled = {props.userAddr === loan.borrower ? false : true}> Cancel Loan </button>
                </div>
            ))}
       </div> 
    </div>
  )
}

export default Loans 
