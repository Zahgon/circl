package common

// Keeps mapping: SIDH prime field ID to domain parameters
var sidhParams = make(map[uint8]SidhParams)

// Params returns domain parameters corresponding to finite field and identified by
// `id` provided by the caller. Function panics in case `id` wasn't registered earlier.
func Params(id uint8) *SidhParams { _ = "STUB: not implemented"; return nil }

// Registers SIDH parameters for particular field.
func Register(id uint8, p *SidhParams) { _ = "STUB: not implemented"; return }
