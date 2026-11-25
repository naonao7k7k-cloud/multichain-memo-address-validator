package validator

import "testing"

func TestValidateAddress(t *testing.T) {
	cases := []struct {
		net   Network
		addr  string
		valid bool
	}{
		{NetworkBitcoin, "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4", true},
		{NetworkEthereum, "0x742d35Cc6634C0532925a3b844Bc454e4438f44e", true},
		{NetworkBase, "0x742d35Cc6634C0532925a3b844Bc454e4438f44e", true},
		{NetworkTron, "TNPeeaaFB7K9cmo4uQpcU32zGK8G1NYqeL", true},
		{NetworkXRP, "rPT1Sjq2YGrBMTttX4GZHjKu9dyfzbpAYe", true},
		{NetworkBitcoin, "0x742d35Cc6634C0532925a3b844Bc454e4438f44e", false},
		{NetworkEthereum, "not-an-address", false},
	}

	for _, tc := range cases {
		res := ValidateAddress(tc.net, tc.addr)
		if res.Valid != tc.valid {
			t.Fatalf("ValidateAddress(%q, %q) = %v, want %v", tc.net, tc.addr, res.Valid, tc.valid)
		}
	}
}

func TestValidateMemo(t *testing.T) {
	if err := ValidateMemo(NetworkXRP, "123456"); err != nil {
		t.Fatalf("expected valid XRP tag, got %v", err)
	}
	if err := ValidateMemo(NetworkXRP, "not-a-number"); err == nil {
		t.Fatalf("expected error for non-numeric XRP tag")
	}
	if err := ValidateMemo(NetworkStellar, "order-ref-2026"); err != nil {
		t.Fatalf("expected valid Stellar memo, got %v", err)
	}
}
