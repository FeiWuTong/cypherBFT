// Copyright 2015 The go-ethereum Authors
// Copyright 2017 The cypherBFT Authors
// This file is part of the cypherBFT library.
//
// The cypherBFT library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cypherBFT library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cypherBFT library. If not, see <http://www.gnu.org/licenses/>.

package enr

import (
	"crypto/ecdsa"
	"math/big"
	"testing"
)

// Checks that failure to sign leaves the record unmodified.
func TestSignError(t *testing.T) {
	invalidKey := &ecdsa.PrivateKey{D: new(big.Int), PublicKey: *pubkey}

	var r Record
	if err := SignV4(&r, invalidKey); err == nil {
		t.Fatal("expected error from SignV4")
	}
	if len(r.pairs) > 0 {
		t.Fatal("expected empty record, have", r.pairs)
	}
}
