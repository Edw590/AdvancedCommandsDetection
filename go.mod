module ACD

// Keep it on 1.20, so that it can be compiled for Windows 7 too if it's compiled with Go 1.20 (it's the last version
// supporting it).
go 1.20

require github.com/jdkato/prose/v2 v2.0.0

// Do NOT exclude golang.org/x/mobile,mod,sys,tools from require() even if they seem unused (Gomobile - Android AAR).
// If they go away, run `go get golang.org/x/mobile/bind` on the main ACD folder. Before that, come here to install
// Gomobile: https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile.
// If the lines are not in require(), the app may hang on startup saying.

// golang.org/x/mobile v0.0.0-20231108233038-35478a0c49da // indirect
// golang.org/x/mod v0.14.0 // indirect
// golang.org/x/sys v0.14.0 // indirect
// golang.org/x/tools v0.15.0 // indirect

require (
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/mingrammer/commonregex v1.0.1 // indirect
	golang.org/x/mobile v0.0.0-20231108233038-35478a0c49da // indirect
	golang.org/x/mod v0.14.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/tools v0.15.0 // indirect
	gonum.org/v1/gonum v0.15.0 // indirect
	gopkg.in/neurosnap/sentences.v1 v1.0.7 // indirect
)
