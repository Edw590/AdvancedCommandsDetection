# LEGION - A real assistant [Platforms Unifier]

## Notice
This project is connected to other projects:
- [LEGION - A real assistant [Android/Client]](https://github.com/DADi590/LEGION---A-real-assistant--Android-Client)

## Table of Contents
- [Background](#background)
- [Current sub-modules](#current-sub-modules)
- [For developers](#for-developers)
- - [To compile the module](#--to-compile-the-module)
- [About](#about)
- - [Roadmap](#--roadmap)
- - [Project status](#--project-status)
- - [License](#--license)
- [Support](#support)
- [Final notes](#final-notes)

## Background
This is a module which is supposed to connect the assistant in different platforms. Something is coded in Go and does not need to be coded specifically in Java, then for iOS, for example, Objective C, then C++ or C (Windows/Linux), or Python (Raspberry Pi), or any other. This should compile to all needed architectures and be present as a global utilities library.

## Current sub-modules
- **Commands Detector** --> A string is given to the main function of the detector and the output is a list of detected commands on the given string. For example (give it without the punctuation, as Speech Recognition engines don't put it, so it's not used here and must not be present), "turn on wifi and get the airplane mode on. no, don't turn the wifi on. turn off airplane mode and turn the wifi on". This should output an order to (in order of given commands), turn on the airplane mode, then turn it off, and finally turn on the Wi-Fi. And it does: "11.1, 11.2, 4.1", which means the same, according to the way the function works.

## For developers
### - To compile the module
- To run on PC, either use an IDE which does it automatically (I use GoLand, for example), or run the following command in the project folder as working directory: "go run Assist_Platforms_Unifier".
- To compile for Android and create an AAR package, have a look on the Build_AAR_Android.bat file and execute the command inside it. If you use the file, make sure to change the ANDROID_HOME variable. For some reason, I can't use relative paths here, so I used an absolute one (must be doing something wrong).

And a small explanation of the module structure:
- Each main package has its name beginning with "APU_". The reason is on Android, without that, appears, for example "CmdDetection" only. And that comes from where? No idea. No indication. So "APU_CmdDetection" seems better to differentiate where the class came from (APU comes from Assist_Plaforms_Unifier - the name of the project).
- Each sub-module is inside a package. Like the Commands Detector module, which is inside APU_CmdDetection.
- There's also a Python program which updates the VERSION constant each the module is compiled for Android. That way forgetting to update the version doesn't happen.
- As this module is compiled for Android with Gomobile, it's limited to the supported types by go/build: https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions, so all the exported elements must follow those rules (some, as for example if an array is exported, no error is thrown, so doesn't seem to be bad to export those to be accessible across packages - won't be accessible on Android though).

## About
### - Roadmap
Have a look on the "TODO.md" file.

### - Project status
Ongoing, but possibly slowly since I'm a student, so I may not have that much time to work on this (even though I'd love to have more time) - except on Holidays xD.

### - License
This project is licensed under Apache 2.0 License - http://www.apache.org/licenses/LICENSE-2.0.

## Support
If you have any questions, try the options below:
- Create an Issue here: https://github.com/DADi590/LEGION---A-real-assistant--PlatformsUnifier/issues
- Create a Discussion here: https://github.com/DADi590/LEGION---A-real-assistant--PlatformsUnifier/discussions

## Final notes
Hope you like the app! Any new ideas are welcomed! (I just may or may not implement them that fast - student)
