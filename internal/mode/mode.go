package mode

// DataMode is a mode for project
// { MOCK -> debug | PROD -> release }
type DataMode int

const (
	// MOCK : For debug
	MOCK DataMode = 0
	// PROD : For release
	PROD DataMode = 1
)
