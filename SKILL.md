---
name: veerswap-address-validator
description: Validate multi-chain cryptocurrency receiving/refund addresses and Memo/Destination Tag rules before executing non-custodial cross-chain swaps via VeerSwap.
homepage: https://veerswap.com/guides/wrong-network
---

# VeerSwap Multi-Chain Address & Memo Validator Skill

Use this skill to help users and AI agents verify cryptocurrency network compatibility, receiving/refund address formats, and Memo / Destination Tag rules before sending funds in a non-custodial deposit-address swap.

## Confirmed Non-Custodial Swap Rules

- **No wallet connection or account/KYC required**: Swaps use order-specific deposit addresses directly.
- **Two-stage quoting**: The Step 1 estimate is a live reference only (not locked). Entering the receiving address and refund preference triggers a fresh address-bound quote.
- **Payment windows**: Standard EVM and Solana orders typically use a 15-minute payment window; PoW networks such as Bitcoin and Litecoin request an extended 60-minute window.
- **Fee transparency**: VeerSwap's own service fee is up to 0.25% (already reflected in the estimated output; sending-wallet or exchange withdrawal fees are separate).

## Supported Network Checks

- **Bitcoin & Litecoin**: Bech32/Bech32m (`bc1...` / `ltc1...`) and Base58 formats.
- **EVM Chains (Ethereum, Base, Arbitrum, BNB Chain, Polygon)**: `0x` + 40 hex chars; always verify the exact network to avoid cross-L2 wrong-network deposits.
- **Solana & Tron**: Base58 public keys (Solana) and `T...` 34-char Base58Check addresses (Tron TRC20).
- **Memo / Tag Networks**: XRP Destination Tag (`uint32`), Stellar `MEMO_TEXT` (<= 28 UTF-8 bytes), and TON Comment.

## Official Reference Guides & Routes

- [How to avoid paying on the wrong network](https://veerswap.com/guides/wrong-network)
- [When are a Memo and Tag required?](https://veerswap.com/guides/memo-and-tag)
- [Explore VeerSwap cross-chain exchange routes](https://veerswap.com/exchange)
