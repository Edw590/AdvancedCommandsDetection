# Advanced Commands Detection
The Advanced Commands Detection module of my virtual assistant, V.I.S.O.R.

## Notice
This project is a part of a bigger project, consisting of the following:
- [V.I.S.O.R. - A better Android assistant](https://github.com/DADi590/VISOR---A-better-Android-assistant)
- [Advanced Commands Detection](https://github.com/DADi590/Advanced-Commands-Detection)

## Table of Contents
- [Background](#background)
- [How it works](#how-it-works)
- [For developers](#for-developers)
- - [To compile the module](#--to-compile-the-module)
- [About](#about)
- - [Roadmap](#--roadmap)
- - [Project status](#--project-status)
- - [License](#--license)
- [Support](#support)
- [Final notes](#final-notes)

## Background
This is a command detection module that I began in 2017 or 2018 still when I didn't know what a function was and I made many copy/pastes of the small code that I made for 2 specific cases of use. From what I saw in older files I have around, I made it a function in 2019, recursive (who knows why. I had just learned about that, maybe I thought it was better that way(?)), and I've been improving ever since. And in this project, I've put together that function alongside various other methods of detecting the commands in a string.

I call it Advanced(?) because it's not just a simple "turn on wifi" thing. No - it detects each keyword in an interval of words. It also knows what "don't" and "it" mean ("turn on wifi. no, don't turn it on" --> 0 commands detected here). Or knows what an "and" means ("turn on the wifi and the airplane mode" - no need to say "turn on wifi turn on airplane mode"). And there are still improvements to be made to this. I don't use AI for this, at most I use NLP (so far), so that might also explain why the module is so complex (aside from the reasons I said above).

Currently this may be found to be sort of a mess (I have a fivefold array in this thing.... a few fourfold, various threefold, many twofold, infinite single ones), but surprisingly enough, it works XD. All those infinite arrays have served a purpose, even if they don't serve anymore - but might serve again the future, so I haven't removed any of them. With time, I've improved the main word verifications function and the other arrays started to be less and less needed - though there are special types of commands that may required them again.

It can also probably be very much improved to be more efficient (as I said, this began years ago - when I had just learned programming, didn't know what a struct was, didn't know a string was harder to check for equality than an int or float). But that will take time and I rather improve the detection than improve the code (also, linking this to Android with Gomobile means I can only have basic types available as said below, so there's that to complicate a few things...).

I'm also not really wanting to use C/C++ for this, not unless Go stops being fast enough - else I have to pay attention to infinity that can can wrong on a C/C++ program... Waste of time if Go is fast enough. It's also in Go and not in Java as VISOR is because then I can use this for any other platform without worrying about the supported languages nor reimplementing all this infinity.

## How it works
It outputs a list of detected commands in a given sentence of words. For example, give it (without the punctuation, as Speech Recognition engines don't put it, so it's not used here and must not be present): "turn it on. turn on the wifi, and and the airplane mode, get it it on. no, don't turn it on. turn off airplane mode and also the wifi, please." - this string will make the module output orders to (in order of given commands), request an explanation of the first "it" (which has no meaning), turn on the Wi-Fi, then turn off the airplane mode, and also the Wi-Fi. And it does: "-2, 4.1, 11.2, 4.2", which means the same, according to the way the submodule works.

## For developers
### - To compile the module
- To run on PC, either use an IDE which does it automatically (I use GoLand, for example), or run the following command in the project folder as working directory: "go run AdvancedCommandsDetection".
- To compile for Android and create an AAR package, have a look on the Build_AAR_Android.bat file and execute the command inside it. If you use the file, make sure to change the ANDROID_HOME variable. For some reason, I can't use relative paths here, so I used an absolute one (must be doing something wrong).

And a small explanation of the module structure:
- Each main package has its name ending with "\_ACD". The reason is on Android, without that, appears, for example "CmdDetection" only. And that comes from where? No idea. No indication. So "CmdDetection_ACD" seems better to differentiate where the class came from (ACD comes from AdvancedCommandsDetection - the name of the project).
- Each submodule is inside a package. Like the Commands Detector module, which is inside CmdDetection_ACD.
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
Any new ideas and/or improvements are welcomed! (Just giving the idea and/or making a pull request)
