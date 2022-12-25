module Assist_Platforms_Unifier

go 1.17

require github.com/jdkato/prose/v2 v2.0.0

// Do NOT exclude golang.org/x/mobile even if it seems unused (Gomobile - Android AAR)
require (
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/mingrammer/commonregex v1.0.1 // indirect
	gonum.org/v1/gonum v0.12.0 // indirect
	gopkg.in/neurosnap/sentences.v1 v1.0.7 // indirect
)
