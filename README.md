# Advanced Commands Detection
The Advanced Commands Detection module of my virtual assistant, V.I.S.O.R.

## Notice
This project is a part of a bigger project, consisting of the following:
- [V.I.S.O.R. - A better Android assistant](https://github.com/DADi590/VISOR---A-better-Android-assistant)
- [Advanced Commands Detection](https://github.com/DADi590/Advanced-Commands-Detection)

## Examples of successful detections
A list of sentences sent to the module and what it successfully understands at its output:
```
- "turn off airplane mode on": turn off airplane mode
- "turn on turn off the wifi": turn off Wi-Fi and airplane mode
- "turn on wifi and the bluetooth no don't turn it on": turn on Wi-Fi
- "turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on": turn on airplane mode, then turn if off again, and turn on Wi-Fi
- "turn on turn wifi on please": turn on Wi-Fi
- "turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please": warn about a meaningless "it", turn on Wi-Fi, turn off both airplane mode and Wi-Fi
- "turn on wifi and and the airplane mode and the flashlight": turn on Wi-Fi, airplane mode, and flashligh
- "shut down the phone and reboot it": shut down and reboot device
- "fast reboot the phone": fast reboot device (this test exists because "fast" and "reboot" are both command triggers, but here only "fast" is used to trigger and "reboot" is ignored)
```
These are automated test sentences that are tested each time modifications are made to the engine, to be sure it at least remains working as good as it was before the modifications.

## Table of Contents
- [Background](#background)
- [How it works](#how-it-works)
- - [To compile the module](#--to-compile-the-module)
- - [Small explanation of the project structure](#--small-explanation-of-the-project-structure)
- [About](#about)
- - [Roadmap](#--roadmap)
- - [Project status](#--project-status)
- - [License](#--license)
- [Support](#support)
- [Final notes](#final-notes)

## Background
This is a command detection module that I began in 2017 or 2018 still when I didn't know what a function was and I made many copy/pastes of the small code that I made for 2 specific cases of use. From what I saw in older files I have around, I made it a function in 2019, recursive (who knows why. I had just learned about that, maybe I thought it was better that way(?)), and I've been improving ever since. And in this project, I've put together that function alongside various other methods that in the end help detecting the commands in a string.

I call it Advanced(?) because it's not just a simple "turn on wifi" thing. Instead, it detects each keyword in an interval of words. It also knows what "don't" and "it" mean ("turn on wifi. no, don't turn it on" --> 0 commands detected here). Or knows what an "and" means ("turn on the wifi and the airplane mode" - no need to say "turn on wifi turn on airplane mode"). And there are still improvements to be made to this. I don't use AI for this, at most I use NLP (so far), so that might also explain why the module is complex.

I'm also not really wanting to use C/C++ for this, not unless Go stops being fast enough - else I have to pay attention to infinity that can can wrong on a C/C++ program... Waste of time if Go is fast enough. It's also in Go and not in Java as VISOR is because then I can use this for any other platform without worrying about the supported languages nor reimplementing all this infinity (VISOR is supposed to be multi-platform, not just Android - but I lack the time to make that happen...).

## How it works
The `ACD.Main()` function outputs a list of detected commands in a given sentence of words. For example, give it (without the punctuation, as Speech Recognition engines don't put it, so it's not used here and must not be present): `"turn it on. turn on the wifi, and and the airplane mode, get it it on. no, don't turn it on. turn off airplane mode and also the wifi, please."` - this string will make the module output orders to (in order of given commands), request an explanation of the first "it" (which has no meaning), turn on the Wi-Fi, then turn off the airplane mode, and also the Wi-Fi. And it does: `"-10, 4.01, 11.02, 4.02"`, which means the same, according to the way the module works.

### - To compile the module
- To run on PC, either use an IDE which does it automatically (I use GoLand, for example), or run the following command in the project folder as working directory: "go run AdvancedCommandsDetection".
- To compile for Android and create an AAR package, have a look on the Build_AAR_Android.bat file and execute the command inside it. If you use the file, make sure to change the ANDROID_HOME variable. For some reason, I can't use relative paths here, so I used an absolute one (must be doing something wrong). You might also want to run VersionUpdater.py before the batch script to update the ACD's VERSION constant to the current date/time (just to keep track of which version is being used on the AAR).

### - Small explanation of the project structure
- All that belongs to the module remains inside the ACD folder. ACD because on Java is much easier to write ACD.function() than AdvancedCommandsDetection.function() (also much less space taken).
- Outside that, only things to make the library work as a main package for testing (like main.go or TryCatchFinally, which doesn't belong to the project and is just a "utility").
- Each command provided to the engine must be given a unique ID greater than 0. Those IDs are the ones returned by the `ACD.Main()` function, along with a decimal part, which is the number of the variation of the command. Example: `"turn on/off wifi` with ID 4. This is a command, with 2 variations. The `"on"` variation outputs 4.01 and the `"off"` variation outputs 4.02 (increments of 0.01 to the ID integer, beginning in 0.01).
- It also seems that after the library is loaded, all global variables remain with their last value, until it's unloaded. This is useful because of the types limitation below and still needing to pass a big array into the engine for it to load all possible commands --> a big string must be sent to the library for processing, but only in the beginning, which will make the performance not matter (it's only hurt on the entire program's beginning).

As this module is compiled for Android with Gomobile, it's limited to the supported types by go/build: https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions, so all the exported elements must follow those rules (some, as for example if a slice is exported, no error is thrown, so doesn't seem to be bad to export those to be accessible across packages - won't be accessible on Android though). So for example, to pass an array to the functions of the library, it must be encoded into a string and decoded on the function again.

## About
### - Roadmap
Have a look on the "TODO.md" file.

### - Project status
Ongoing, but possibly slowly since I'm a student, so I may not have that much time to work on this.

### - License
This project is licensed under Apache 2.0 License - http://www.apache.org/licenses/LICENSE-2.0.

## Support
If you have any questions, try the options below:
- Create an Issue here: https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Platforms-Unifier/issues
- Create a Discussion here: https://github.com/DADi590/V.I.S.O.R.---A-real-assistant--Platforms-Unifier/discussions

## Final notes
Any new ideas and/or improvements are welcomed! (Just giving the idea and/or making a pull request)
