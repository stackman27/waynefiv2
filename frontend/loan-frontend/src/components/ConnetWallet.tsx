import React, {useState, useEffect} from 'react';
import { getKeplrFromWindow } from "@keplr-wallet/stores"
import { Keplr } from "@keplr-wallet/types"
import { chainInfo } from '../config/chain';

function ConnectWallet(props: { connectWallet: () => void; signOut: () => void; userAddr: string; }) {
  return (
    <div className="App">
      <button onClick={() => props.connectWallet()}> Connect Wallet </button>
      <button onClick={() => props.signOut()}> SignOut </button>
      <h1> Address: {props.userAddr || "Address Not found"} </h1>
    </div>
  );
}

export default ConnectWallet;
