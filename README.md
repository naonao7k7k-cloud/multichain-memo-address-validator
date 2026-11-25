# multichain-memo-address-validator

Zero-dependency Go library for validating cryptocurrency address formats and Memo / Destination Tag rules across major blockchain networks (Bitcoin, Ethereum, Base, Arbitrum, BNB Chain, Polygon, Solana, Tron, XRP Ledger, Stellar, TON, and Litecoin).

## Why Validate Network & Memo Rules Before Sending?

Cross-chain and same-chain crypto swaps frequently fail or require manual recovery when a sender:
1. Sends funds on the wrong network (for example, confusing `0x` addresses across Ethereum, Base, Arbitrum, or BNB Chain, or sending ERC20 instead of TRC20).
2. Omits a required **XRP Destination Tag**, **Stellar Memo**, or **TON Comment** when sending to an exchange deposit address.

For complete operational guides on non-custodial deposit-address swaps:
- [How to avoid paying on the wrong network](https://veerswap.com/guides/wrong-network)
- [When are a Memo and Tag required?](https://veerswap.com/guides/memo-and-tag)
- [Explore supported cross-chain swap routes on VeerSwap](https://veerswap.com/exchange)

## Installation

```bash
go get github.com/naonao7k7k-cloud/multichain-memo-address-validator@v1.0.0
```

## Quick Start

```go
package main

import (
	"fmt"
	validator "github.com/naonao7k7k-cloud/multichain-memo-address-validator"
)

func main() {
	res := validator.ValidateAddress(validator.NetworkXRP, "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe")
	fmt.Println("Valid:", res.Valid, "MemoOptional:", res.MemoOptional)

	if err := validator.ValidateMemo(validator.NetworkXRP, "123456"); err != nil {
		fmt.Println("Invalid Destination Tag:", err)
	}
}
```

## License

MIT License.
