package hpke

func (st state) keySchedule(ss, info, psk, pskID []byte) (*encdecContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (st state) verifyPSKInputs(psk, pskID []byte) error { _ = "STUB: not implemented"; return nil }

// Params returns the codepoints for the algorithms comprising the suite.
func (suite Suite) Params() (KEM, KDF, AEAD) {
	_ = "STUB: not implemented"
	return *new(KEM), *new(KDF), *new(AEAD)
}

func (suite Suite) String() string { _ = "STUB: not implemented"; return "" }

func (suite Suite) getSuiteID() (id [10]byte) { _ = "STUB: not implemented"; return nil }

func (suite Suite) isValid() bool { _ = "STUB: not implemented"; return false }

func (suite Suite) labeledExtract(salt, label, ikm []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (suite Suite) labeledExpand(prk, label, info []byte, l uint16) []byte {
	_ = "STUB: not implemented"
	return nil
}
