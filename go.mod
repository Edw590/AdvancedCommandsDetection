module Assist_Platforms_Unifier

go 1.17

require github.com/jdkato/prose/v2 v2.0.0

// Do NOT exclude golang.org/x/mobile even if it seems unused (Gomobile - Android AAR)
require (
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/mingrammer/commonregex v1.0.1 // indirect
	golang.org/x/mobile v0.0.0-20211109191125-d61a72f26a1a // indirect
	gonum.org/v1/gonum v0.7.0 // indirect
	gopkg.in/neurosnap/sentences.v1 v1.0.6 // indirect
)
