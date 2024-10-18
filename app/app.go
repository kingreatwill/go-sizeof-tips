package app

const DefaultHttpPort = ":7777"

// Represents simple zero-cost message that can be used
// as signal between goroutines.
type sig struct{}
