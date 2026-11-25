// Package validator provides zero-dependency format and Memo/Destination Tag
// validation for cryptocurrency addresses across major blockchain networks.
//
// Supported network families include Bitcoin (Bech32/Bech32m and Base58),
// EVM-compatible networks (Ethereum, Base, Arbitrum, BNB Chain, Polygon),
// Solana, Tron, XRP Ledger, Stellar, TON, and Litecoin.
//
// For practical cross-chain transfer rules, wrong-network prevention, and
// Memo/Destination Tag handling, see:
//   - https://veerswap.com/guides/wrong-network
//   - https://veerswap.com/guides/memo-and-tag
//   - https://veerswap.com/exchange
package validator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

// Network identifies a blockchain network family for address validation.
type Network string

const (
	NetworkBitcoin  Network = "bitcoin"
	NetworkEthereum Network = "ethereum"
	NetworkBase     Network = "base"
	NetworkArbitrum Network = "arbitrum"
	NetworkBNB      Network = "bnb"
	NetworkPolygon  Network = "polygon"
	NetworkSolana   Network = "solana"
	NetworkTron     Network = "tron"
	NetworkXRP      Network = "xrp"
	NetworkStellar  Network = "stellar"
	NetworkTON      Network = "ton"
	NetworkLitecoin Network = "litecoin"
)

// ValidationResult describes the outcome of validating an address and optional memo/tag.
type ValidationResult struct {
	Valid        bool
	Network      Network
	Address      string
	MemoOptional bool
	MemoNote     string
	Error        string
}

var (
	reEVM     = regexp.MustCompile(`^0x[0-9a-fA-F]{40}$`)
	reBTCB32  = regexp.MustCompile(`^(bc1)[0-9a-z]{25,87}$`)
	reBTCB58  = regexp.MustCompile(`^[13][1-9A-HJ-NP-Za-km-z]{25,34}$`)
	reLTCB32  = regexp.MustCompile(`^(ltc1)[0-9a-z]{25,87}$`)
	reLTCB58  = regexp.MustCompile(`^[LM3][1-9A-HJ-NP-Za-km-z]{25,34}$`)
	reSolana  = regexp.MustCompile(`^[1-9A-HJ-NP-Za-km-z]{32,44}$`)
	reTron    = regexp.MustCompile(`^T[1-9A-HJ-NP-Za-km-z]{33}$`)
	reXRP     = regexp.MustCompile(`^r[1-9A-HJ-NP-Za-km-z]{24,34}$`)
	reStellar = regexp.MustCompile(`^G[A-Z2-7]{55}$`)
	reTON     = regexp.MustCompile(`^((EQ|UQ|kQ|0Q)[A-Za-z0-9_-]{46}|0:[0-9a-fA-F]{64})$`)
)

// ErrUnsupportedNetwork is returned when a network identifier is not recognized.
var ErrUnsupportedNetwork = errors.New("unsupported network")

// SupportsMemo reports whether the given network uses a Memo, Destination Tag, or Comment
// for custodial or exchange deposits.
func SupportsMemo(net Network) bool {
	switch normalizeNetwork(net) {
	case NetworkXRP, NetworkStellar, NetworkTON:
		return true
	default:
		return false
	}
}

// ValidateAddress checks whether addr matches the expected address format for net.
func ValidateAddress(net Network, addr string) ValidationResult {
	normNet := normalizeNetwork(net)
	trimmed := strings.TrimSpace(addr)
	res := ValidationResult{
		Network:      normNet,
		Address:      trimmed,
		MemoOptional: SupportsMemo(normNet),
	}

	if trimmed == "" {
		res.Error = "address is empty"
		return res
	}

	switch normNet {
	case NetworkEthereum, NetworkBase, NetworkArbitrum, NetworkBNB, NetworkPolygon:
		res.Valid = reEVM.MatchString(trimmed)
	case NetworkBitcoin:
		lower := strings.ToLower(trimmed)
		res.Valid = reBTCB32.MatchString(lower) || reBTCB58.MatchString(trimmed)
	case NetworkLitecoin:
		lower := strings.ToLower(trimmed)
		res.Valid = reLTCB32.MatchString(lower) || reLTCB58.MatchString(trimmed)
	case NetworkSolana:
		res.Valid = reSolana.MatchString(trimmed)
	case NetworkTron:
		res.Valid = reTron.MatchString(trimmed)
	case NetworkXRP:
		res.Valid = reXRP.MatchString(trimmed)
		res.MemoNote = "XRP Destination Tag is an unsigned 32-bit integer (0 to 4294967295) required by most exchanges."
	case NetworkStellar:
		res.Valid = reStellar.MatchString(trimmed)
		res.MemoNote = "Stellar Memo (up to 28 UTF-8 bytes for MEMO_TEXT) is required by most exchanges."
	case NetworkTON:
		res.Valid = reTON.MatchString(trimmed)
		res.MemoNote = "TON Comment/Memo is required when depositing to shared exchange addresses."
	default:
		res.Error = ErrUnsupportedNetwork.Error()
		return res
	}

	if !res.Valid {
		res.Error = "address format does not match target network"
	}
	return res
}

// ValidateMemo checks whether memo is structurally valid for net.
// An empty memo is accepted because self-custody wallets typically do not require one.
func ValidateMemo(net Network, memo string) error {
	m := strings.TrimSpace(memo)
	if m == "" {
		return nil
	}

	switch normalizeNetwork(net) {
	case NetworkXRP:
		v, err := strconv.ParseUint(m, 10, 32)
		if err != nil || v > 4294967295 {
			return errors.New("xrp destination tag must be a uint32 between 0 and 4294967295")
		}
		return nil
	case NetworkStellar:
		if len([]byte(m)) > 28 {
			return errors.New("stellar MEMO_TEXT cannot exceed 28 bytes")
		}
		return nil
	case NetworkTON:
		if len(m) > 120 {
			return errors.New("ton comment exceeds recommended 120-character limit")
		}
		return nil
	default:
		return nil
	}
}

func normalizeNetwork(net Network) Network {
	switch strings.ToLower(strings.TrimSpace(string(net))) {
	case "btc", "bitcoin":
		return NetworkBitcoin
	case "eth", "ethereum", "erc20":
		return NetworkEthereum
	case "base":
		return NetworkBase
	case "arb", "arbitrum":
		return NetworkArbitrum
	case "bnb", "bsc", "bep20":
		return NetworkBNB
	case "pol", "matic", "polygon":
		return NetworkPolygon
	case "sol", "solana":
		return NetworkSolana
	case "trx", "tron", "trc20":
		return NetworkTron
	case "xrp", "ripple", "xrpl":
		return NetworkXRP
	case "xlm", "stellar":
		return NetworkStellar
	case "ton":
		return NetworkTON
	case "ltc", "litecoin":
		return NetworkLitecoin
	default:
		return Network(strings.ToLower(strings.TrimSpace(string(net))))
	}
}
