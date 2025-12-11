package consts

const (
	Batch       int64  = 5000
	StandardMax int64  = 5000
	Less        uint64 = 10000
	RateLimit   int64  = 1000
)

type Plans string

// Starter → Standard → Ultimate → Ultimate
const (
	// Reserved : less than 10000
	Reserved Plans = "reserved"

	// Starter : contains 4 , contains 13
	Starter Plans = "starter"

	// Standard scene
	Standard Plans = "standard"

	// Premium : abc , cba , aaa , abab
	Premium Plans = "premium"

	// Ultimate : aaaa , abcd , dcba , ab...
	Ultimate Plans = "ultimate"
)
