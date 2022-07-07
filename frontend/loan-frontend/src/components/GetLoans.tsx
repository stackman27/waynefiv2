/* eslint-disable array-callback-return */
import React, {useState, useEffect} from 'react'
import useAxios from 'axios-hooks'
import { Keplr } from "@keplr-wallet/types";
import { chainInfo } from '../config/chain';
import { txClient } from '../generated/loan/loan.loan/module';
import loanLoan from '../generated/loan/loan.loan';



 function GetLoans(props: {userAddr: string, keplr: Keplr | null}) {

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


    const approveLoan = async () => {
        if (!props.keplr) {
            throw Error("Keplr isn't connected");
        }

        const offlineSigner = props.keplr.getOfflineSigner(chainInfo.chainId)

        console.log("APPROVE THIS LOAN")
        // start creating msg for approving loan

        var data = {
            creator: props.userAddr,
            id: 0,
        }

        var rootGetters =  {
            apiNode: 'http://localhost:1317',
            rpcNode: 'http://localhost:26657',
            wsNode: 'ws://localhost:26657/websocket',
            chainId: 'loan-1',
            addrPrefix: 'loan',
            sdkVersion: 'Stargate',
            getTXApi: 'http://localhost:26657/tx?hash=0x',       
            offlineSigner: props.keplr.getOfflineSigner(chainInfo.chainId),
        }
 
        const newTx = await loanLoan.actions
                .sendMsgApproveLoan({rootGetters: rootGetters},
                {value: data})

        console.log(newTx)

    //    const aminoMsgs: Msg[] = [
    //         {
    //           type: "approve-loan",
    //           value: {
    //             id: "0" 
    //           },
    //         },
    //       ];

    //     const protoMsgs = [
    //         {
    //             typeUrl: "approve-loan",
    //             value: {
    //                 id: "0",
    //             }  
    //         }
    //     ]

    //       const fee: StdFee = {
    //         gas: "200000",
    //         amount: [
    //           {
    //             amount: "12",
    //             denom: "uatom",
    //           },
    //         ],
    //       };
      
    //       const broadCastMode = "async";
      
    //       const signDoc = makeSignDoc(
    //         aminoMsgs,
    //         fee,
    //         chainInfo.chainId,
    //         "",
    //         accountNumber,
    //         sequence
    //       );
          
         
    //       const signResponse = await props.keplr.signAmino(
    //         chainInfo.chainId,
    //         props.userAddr,
    //         signDoc,
    //         undefined
    //       );

    //       const signedTx = makeStdTx(signResponse.signed, signResponse.signature);

    //       const client = await SigningStargateClient.connectWithSigner(
    //         "https://rpc-osmosis.blockapsis.com",
    //         offlineSigner
    //     )

        
    //       await props.keplr.sendTx(
    //         chainInfo.chainId,
    //         null,
    //         broadCastMode as BroadcastMode
    //       );

    }

    const cancelLoan = () => {
        console.log("CANCEL THIS LOAN")
    }

    return (
    <div style = {{textAlign: 'left', width: '80%', marginLeft: 'auto', marginRight: 'auto'}}>
        <h2> User Balance: {userBankData && userBankData.balance.amount.concat(userBankData.balance.denom) } </h2>
        <h3 style = {{color: 'red'}}> Loans </h3>
            {loanData && loanData.Loan.map((loan: { amount: string, id: string , borrower: string, state: string, collateral: string, deadline:string}) => (
                <div id={loan.id} style = {{height: '300px', width: '400px', background: '#eee', textAlign: 'left', padding: 10,}}>
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
  )
}

export default GetLoans
