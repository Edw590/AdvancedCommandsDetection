# V.I.S.O.R. - A real assistant [Platforms Unifier]

## Table of Contents
- [Notice](#notice)
- [Background](#background)
- [Current submodules](#current-submodules)
- [For developers](#for-developers)
- - [To compile the module](#--to-compile-the-module)
- [About](#about)
- - [Roadmap](#--roadmap)
- - [Project status](#--project-status)
- - [License](#--license)
- [Support](#support)
- [Final notes](#final-notes)

## Notice
This project is a part of a bigger project, consisting of the following:
- [V.I.S.O.R. - A real assistant [Android/Client]](https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Android-Client)
- [V.I.S.O.R. - A real assistant [Platforms Unifier]](https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Platforms-Unifier)

## Background
This is a module which is supposed to connect the assistant in different platforms. Something is coded in Go and does not need to be coded specifically in Java for Android, then for iOS, for example, Objective C, then C++ or C (Windows/Linux), or Python (Raspberry Pi), or any other. This should compile to all needed architectures and be present as a global utilities library.

## Current submodules
- **Commands Detection** --> Outputs a list of detected commands in a given sentence of words. For example, give it (without the punctuation, as Speech Recognition engines don't put it, so it's not used here and must not be present), "turn it on. turn on the wifi, and and the airplane mode, get it it on. no, don't turn it on. turn off airplane mode and also the wifi, please.". This should output orders to (in order of given commands), request an explanation of the first "it" (which has no meaning), turn on the Wi-Fi, then turn off the airplane mode, and also the Wi-Fi. And it does: "-2, 4.1, 11.2, 4.2", which means the same, according to the way the submodule works.

## For developers
### - To compile the module
- To run on PC, either use an IDE which does it automatically (I use GoLand, for example), or run the following command in the project folder as working directory: "go run Assist_Platforms_Unifier".
- To compile for Android and create an AAR package, have a look on the Build_AAR_Android.bat file and execute the command inside it. If you use the file, make sure to change the ANDROID_HOME variable. For some reason, I can't use relative paths here, so I used an absolute one (must be doing something wrong).

And a small explanation of the module structure:
- Each main package has its name ending with "\_APU". The reason is on Android, without that, appears, for example "CmdDetection" only. And that comes from where? No idea. No indication. So "CmdDetection_APU" seems better to differentiate where the class came from (APU comes from Assist_Plaforms_Unifier - the name of the project).
- Each submodule is inside a package. Like the Commands Detector module, which is inside APU_CmdDetection.
- There are packages which end in "Int". That means Internal - not to be exported like the packages which don't end in Int. With Gomobile it's possible to choose which packages are exported. The others will be compiled but will not be exported. So that's the idea of those Int-ended packages.
- There's also a Python program which updates the VERSION constant each the module is compiled for Android. That way forgetting to update the version doesn't happen.
- As this module is compiled for Android with Gomobile, it's limited to the supported types by go/build: https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions, so all the exported elements must follow those rules (some, as for example if a slice is exported, no error is thrown, so doesn't seem to be bad to export those to be accessible across packages - won't be accessible on Android though).

## About
### - Roadmap
Have a look on the "TODO.md" file.

### - Project status
Ongoing, but possibly slowly since I'm a student, so I may not have that much time to work on this (even though I'd love to have more time) - except on Holidays xD.

### - License
This project is licensed under Apache 2.0 License - http://www.apache.org/licenses/LICENSE-2.0.

## Support
If you have any questions, try the options below:
- Create an Issue here: https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Platforms-Unifier/issues
- Create a Discussion here: https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Platforms-Unifier/discussions

## Final notes
Any new ideas are welcomed! (I just may or may not implement them that fast - student)
