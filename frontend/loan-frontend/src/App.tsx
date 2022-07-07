import React, {useState, useEffect} from 'react';
import { getKeplrFromWindow } from "@keplr-wallet/stores"
import { Keplr } from "@keplr-wallet/types"
import { chainInfo } from './config/chain';
import ConnectWallet from './components/ConnetWallet'
import GetLoans from './components/GetLoans'

const KeyAccountAutoConnect = "acount_auto_connect"


function App() {
  const [keplr, setKeplr] = useState<Keplr | null> (null)
  const [bech32Address, setBech32Address] = useState<string>("")

  const connectWallet = async() => {
    try {
      const newKeplr = await getKeplrFromWindow();
      if (!newKeplr) {
        throw new Error("Keplr extension not found")
      }

      await newKeplr.experimentalSuggestChain(chainInfo);
      await newKeplr.enable(chainInfo.chainId);

      localStorage?.setItem(KeyAccountAutoConnect, "true")
      setKeplr(newKeplr)
      const key = await newKeplr.getKey(chainInfo.chainId)
      setBech32Address(key.bech32Address) 
    } catch(e) {
      alert(e)
    }
  }

  const signOut = () => {
    localStorage?.removeItem(KeyAccountAutoConnect);
    setKeplr(null);
    setBech32Address("");
  };
  return (
    <div style = {{textAlign: 'center'}}>  
      <ConnectWallet userAddr = {bech32Address} connectWallet = {connectWallet} signOut = {signOut}/>
      <GetLoans userAddr = {bech32Address} keplr = {keplr || null}/>
    </div>
    )
}

export default App
