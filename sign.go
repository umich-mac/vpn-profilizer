// From: https://github.com/micromdm/micromdm/blob/main/pkg/crypto/profileutil/sign.go
// MIT License

// Copyright (c) 2016 Victor Vrantchan

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"crypto"
	"crypto/x509"
	"fmt"

	"go.mozilla.org/pkcs7"
)

// Sign takes an unsigned payload and signs it with the provided private key and certificate.
func Sign(key crypto.PrivateKey, cert *x509.Certificate, mobileconfig []byte) ([]byte, error) {
	sd, err := pkcs7.NewSignedData(mobileconfig)
	if err != nil {
		return nil, fmt.Errorf("create signed data for mobileconfig: [%w]", err)
	}

	if err := sd.AddSigner(cert, key, pkcs7.SignerInfoConfig{}); err != nil {
		return nil, fmt.Errorf("add crypto signer to mobileconfig signed data: [%w]", err)
	}

	signedMobileconfig, err := sd.Finish()
	return signedMobileconfig, err
}
