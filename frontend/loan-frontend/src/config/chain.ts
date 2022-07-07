
import { ChainInfo } from "@keplr-wallet/types";

export const chainInfo: ChainInfo = {
  rpc: "http://localhost:26657",
  rest: "http://localhost:1317",
  chainId: "loan-1",
  chainName: "loan",
  stakeCurrency: {
    coinDenom: "loan",
    coinMinimalDenom: "uloan",
    coinDecimals: 6,
  },
  bip44: {
    coinType: 118,
  },
  bech32Config: {
    bech32PrefixAccAddr: "loan",
    bech32PrefixAccPub: "loan" + "pub",
    bech32PrefixValAddr: "loan" + "valoper",
    bech32PrefixValPub: "loan" + "valoperpub",
    bech32PrefixConsAddr: "loan" + "valcons",
    bech32PrefixConsPub: "loan" + "valconspub",
  },
  currencies: [
    {
      coinDenom: "LOAN",
      coinMinimalDenom: "uloan",
      coinDecimals: 6,
    },
    {
      coinDenom: "GREEN",
      coinMinimalDenom: "ugreen",
      coinDecimals: 6,
    },
  ],
  feeCurrencies: [
    {
      coinDenom: "LOAN",
      coinMinimalDenom: "uloan",
      coinDecimals: 6,
    },
  ],
  gasPriceStep: {
    low: 0.01,
    average: 0.025,
    high: 0.03,
  },
  features: [],
};
