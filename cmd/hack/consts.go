package main

const (
	name    = "hack"
	version = "v0.3.1"
)

var (
	prebuilt bool

	// set via ldflags
	revision  string
	buildDate string
)

func init() {
	prebuilt = revision != ""
}
