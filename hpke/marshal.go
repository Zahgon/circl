package hpke

// marshal serializes an HPKE context.
func (c *encdecContext) marshal() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// unmarshalContext parses an HPKE context.
func unmarshalContext(raw []byte) (*encdecContext, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarshalBinary serializes an HPKE sealer according to the format specified
// below. (Expressed in TLS syntax.) Note that this format is not defined by
// the HPKE standard.
//
// enum { sealer(0), opener(1) } HpkeRole;
//
//	struct {
//	    HpkeKemId kem_id;   // draft-irtf-cfrg-hpke-07
//	    HpkeKdfId kdf_id;   // draft-irtf-cfrg-hpke-07
//	    HpkeAeadId aead_id; // draft-irtf-cfrg-hpke-07
//	    opaque exporter_secret<0..255>;
//	    opaque key<0..255>;
//	    opaque base_nonce<0..255>;
//	    opaque seq<0..255>;
//	} HpkeContext;
//
//	struct {
//	  HpkeRole role = 0; // sealer
//	  HpkeContext context;
//	} HpkeSealer;
func (c *sealContext) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalSealer parses an HPKE sealer.
func UnmarshalSealer(raw []byte) (Sealer, error) {
	_ = "STUB: not implemented"
	return *new(Sealer), nil
}

// MarshalBinary serializes an HPKE opener according to the format specified
// below. (Expressed in TLS syntax.) Note that this format is not defined by the
// HPKE standard.
//
//	struct {
//	  HpkeRole role = 1; // opener
//	  HpkeContext context;
//	} HpkeOpener;
func (c *openContext) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalOpener parses a serialized HPKE opener and returns the corresponding
// Opener.
func UnmarshalOpener(raw []byte) (Opener, error) {
	_ = "STUB: not implemented"
	return *new(Opener), nil
}
