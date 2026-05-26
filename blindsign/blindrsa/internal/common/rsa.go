// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package common

import (
	"hash"
	"io"
	"math/big"

	"github.com/cloudflare/circl/blindsign/blindrsa/internal/keys"
)

var (
	bigZero = big.NewInt(0)
	bigOne  = big.NewInt(1)
)

// incCounter increments a four byte, big-endian counter.
func incCounter(c *[4]byte) { _ = "STUB: not implemented"; return }

// mgf1XOR XORs the bytes in out with a mask generated using the MGF1 function
// specified in PKCS #1 v2.1.
func mgf1XOR(out []byte, hash hash.Hash, seed []byte) { _ = "STUB: not implemented"; return }

func encrypt(c *big.Int, N *big.Int, e *big.Int, m *big.Int) *big.Int {
	_ = "STUB: not implemented"
	return nil

	// decrypt performs an RSA decryption, resulting in a plaintext integer. If a
	// random source is given, RSA blinding is used.
}

func decrypt(random io.Reader, priv *keys.BigPrivateKey, c *big.Int) (m *big.Int, err error) {
	_ = "STUB: not implemented"
	// TODO(agl): can we get away with reusing blinds?
	return nil, nil
}

// Blinding enabled. Blinding involves multiplying c by r^e.
// Then the decryption operation performs (m^e * r^e)^d mod n
// which equals mr mod n. The factor of r can then be removed
// by multiplying by the multiplicative inverse of r.

// N != 0

// Unblind.
