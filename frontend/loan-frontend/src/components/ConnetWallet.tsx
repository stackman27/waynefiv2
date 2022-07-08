import React, {useState, useEffect} from 'react';
import { getKeplrFromWindow } from "@keplr-wallet/stores"
import { Keplr } from "@keplr-wallet/types"
import { chainInfo } from '../config/chain';

function ConnectWallet(props: { connectWallet: () => void; signOut: () => void; userAddr: string; }) {
  return (
    <div style = {{width: '80%', marginLeft: 'auto', marginRight: 'auto', justifyContent: 'center', height: '60px', borderBottom: '1px solid #ccc'}}>
      <div style = {{float: 'left',  height: '100%', width: '15%', display: 'flex', alignItems: 'center'}}> 
        <button onClick={() => props.connectWallet()}> Connect Wallet </button> &nbsp; &nbsp;
        <button onClick={() => props.signOut()}> SignOut </button>
      </div>
      <div style = {{float: 'right'}}> 
      <h3> Address: {props.userAddr || "Address Not found"} </h3>
      </div>
    </div>
  );
}

export default ConnectWallet;
