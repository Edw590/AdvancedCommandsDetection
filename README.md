# VISOR-AdvancedCommandsDetection

I made it for my virtual assistant, VISOR, but it can be used anywhere, not just with my assistant, and in any OS/architecture that Go can compile to (I only have releases for Android in AAR package form because for the other project I just use it as normal code and don't need a library file - only on Android).

This repository is a submodule on the [V.I.S.O.R. - Server Version Assistant](https://github.com/Edw590/VISOR---Server-Version-Assistant) project (the main project) and is compiled and used on its [Android version](https://github.com/DADi590/VISOR---A-better-Android-assistant) as an AAR package.

## Notice
This project is a part of a bigger project, consisting of the following:
- [V.I.S.O.R. - Android Version Assistant](https://github.com/DADi590/VISOR---A-better-Android-assistant)
- [V.I.S.O.R. - Server Version Assistant](https://github.com/Edw590/VISOR---Server-Version-Assistant)

## Examples of successful detections
A list of sentences sent to the module and what it successfully understands at its output:
```
- "turn off airplane mode on"  -->  turn off airplane mode
- "turn on turn off the wifi"  -->  turn off Wi-Fi
- "turn on wifi and the bluetooth no don't turn it on"  -->  turn on Wi-Fi
- "turn wifi on and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on"  -->  turn on airplane mode, then turn if off again, and turn on Wi-Fi
- "turn on turn wifi on please"  -->  turn on Wi-Fi
- "turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please"  -->  warn about a meaningless "it", turn on Wi-Fi, turn off both airplane mode and Wi-Fi
- "turn wifi on and and the airplane mode and the flashlight"  -->  turn on Wi-Fi, airplane mode, and flashlight
- "shut down the phone and reboot it"  -->  shut down and reboot device
- "fast reboot the phone"  -->  fast reboot device (this test exists because "fast" and "reboot" are both command triggers, but here only "fast" is used to trigger and "reboot" is ignored)
- "fast phone recovery"  -->  nothing (because of the way the reboot command is configured, this is a useful test)
- "the video stop it and then play it again"  -->  stop and play the video
- "stop the song and play the next one"  -->  stop the current song and play the next one
- "and the airplane mode too", with last cmd info being "turn on the wifi"  -->  turn on the airplane mode
- "and now turn it off", with last cmd info being "turn on the wifi"  -->  turn off the Wi-Fi
```
These are automated test sentences that are tested each time modifications are made to the engine, to be sure it at least remains working as good as it was before the modifications (can only improve or maintain, but never go back).

## Table of Contents
- [Background](#background)
- [How it works](#how-it-works)
- - [Small explanation of the project structure](#--small-explanation-of-the-project-structure)
- [To compile the module](#to-compile-the-module)
- [About](#about)
- - [Roadmap](#--roadmap)
- - [Project status](#--project-status)
- - [License](#--license)
- [Support](#support)
- [Final notes](#final-notes)

## Background
This is a command detection module that I began in 2017 or 2018 still when I didn't know what a function was and I made many copy/pastes of the small code that I made for 2 specific cases of use. From what I saw in older files I have around, I made it a function in 2019, recursive (who knows why. I had just learned about that, maybe I thought it was better that way(?)), and I've been improving ever since. And in this project, I've put together that function alongside various other methods that in the end help detecting the commands in a string.

I call it Advanced(?) because it's not just a simple "turn on wifi" thing. Instead, it detects each keyword in an interval of words. It also knows what "don't" and "it" mean ("turn on wifi. no, don't turn it on" --> 0 commands detected here). Or knows what an "and" means ("turn on the wifi and the airplane mode" - no need to say "turn on wifi turn on airplane mode"). And there are still improvements to be made to this. I don't use ML for this, at most I use NLP (so far), so that might also explain why this is a bit complex.

I'm also not really wanting to use C/C++ for this, not unless Go stops being fast enough - else I have to pay attention to infinity that can can wrong on a C/C++ program... Waste of time if Go is fast enough. It's also in Go and not in Java as VISOR is because then I can use this for any other platform without worrying about the supported languages nor reimplementing all this infinity (VISOR is supposed to be multi-platform, not just Android - but I lack the time to make that happen...).

## How it works
The `ACD.Main()` function outputs a list of detected commands in a given sentence of words. For example, give it (without the punctuation, as Speech Recognition engines don't put it, so it's not used here and must not be present): `"turn it on. turn on the wifi, and and the airplane mode, get it it on. no, don't turn it on. turn off airplane mode and also the wifi, please."` - this string will make the module output orders to (in order of given commands), request an explanation of the first "it" (which has no meaning), turn on the Wi-Fi, then turn off the airplane mode, and also the Wi-Fi. And it does: `"-10, 4.00001, 11.00002, 4.00002"`, which means the same, according to the way the module works.

Take a look at main.go to know how to actually use this. You need to call a function to prepare the library - you give it commands, it stores them, and then you call `ACD.Main()` how many times you want with different command strings and the commands you told it to store will be used to detect commands in the given string.

Also, previous command information can be given to `ACD.Main()` to make it know what to do if "and now turn it off" is sent to it, knowing the last executed command had as name "wifi" and action "turn on the" (though here the action is ignored - it's not in "and the bluetooth too" though - will use "turn on the" here), and it will replace "it" with "wifi" and continue the execution. This command information is also returned on the function, to be used for further calls if it's wanted.

### - How the engine works
Each word of the provided sentence is compared to all commands' `main_words` list. Those are the words that trigger the command detection. There are also the rest of the command words (`words_list`). For example, for the reboot command:
```go
main_words = []string{"fast", "reboot", "restart"}  
words_list = [][][][]interface{}{
	{{{-1}, {"reboot", "restart"}}, {{-1}, {"device", "phone"}}},
	{{{-1}, {"device", "phone"}}},
	{{{-1}, {"device", "phone"}}, {{-1}, {"safe"}}, {{-1}, {"mode"}}},
	{{{-1}, {"device", "phone"}}, {{-1}, {"recovery"}}},
}
main_words_ret_conds = [][]string{
	{"fast"},
	{ANY_MAIN_WORD, "-fast"},
}

// What is sent to the function for command configuration (much easier than the arrays above)
commands_info = [][]string{CMD_REBOOT_DEVICE, CMDi_TYPE_REBOOT+"+"+CMDi_TYPE_MANUAL, "fast", "fast|"+ANY_MAIN_WORD+" -fast", "reboot/restart device/phone|device/phone|device/phone safe mode|device/phone recovery"},

// To know exactly what each parameter here means, that's all on the PrepareCmdsArray function on CmdsArrayPreparation. When I have more time I'll do a full documentation about it. If I could send a struct/object through Gomobile would be much easier, but it's not possible.
```
The -1 on the arrays means the words in it can be in any detected word position. If for example on the "safe" were the index 0, then the "safe" word would have to be the 1st (not counting the main word) to be detected (something like "reboot into safe mode the phone" had to be said). Also any words on the same group are mutually exclusive (it can detect either "device" or "phone" and it will accept the word detection for that position).

The return is based on the number of the detected `words_list` array/condition. If the first command condition is detected ("fast reboot phone"), the command ID is returned (for reboot, CMD_REBOOT_DEVICE, it's 14) plus a decimal number which is the number of the detected condition. For the first one would be 14.00001. If the last one ("reboot into recovery") is detected, 14.00004 is returned.

The main words return conditions (`main_words_ret_conds`) serve the purpose of extending the word verification to the main command words. On the example above, the first condition on the `words_list` will only be detected successfully if the corresponding 1st condition on the `main_words_ret_conds` agrees - and in this case it says the main word must have been "fast" ("fast reboot/restart the device/phone"). Else, it uses any main word for any other condition on the `words_list`. The 2nd condition says that any main word can be detected ("fast", "reboot", or "restart"), but right after there's a "-" which indicates the word is to be excluded, so that leaves "reboot" and "restart" as possibilities for all the other command conditions to be accepted.

The last thing is the (current) sort of simple way the command is configured (2022-01-18 - check the main.go file which has the most updated way always). No need to manually create or generate the other arrays. There are also other command parameters automatically set that are not present on the `commands_info` array. The drawback of trying to simplify the commands configuration is that customizations are lost (those parameters and the arrays of position numbers (the -1s above)), but so far they haven't been needed.

If there are multiple detected conditions ("reboot device into recovery" makes the 2nd and the 4th conditions return true because all their words have been found), then the biggest of them is returned (the ones with more words have higher priority). If there are multiple biggest ones (various detected ones with the same highest length), the first of them on the `words_list` will be picked.

### - Small explanation of the project structure
- All that belongs to the module remains inside the ACD folder. ACD because on Java is much easier to write ACD.function() than AdvancedCommandsDetection.function() (also much less space taken).
- Outside that, only things to make the library work as a main package for testing (like main.go or TryCatchFinally, which doesn't belong to the project and is just a "utility").
- Each command provided to the engine must be given a unique ID greater than 0. Those IDs are the ones returned by the `ACD.Main()` function, along with a decimal part, which is the number of the variation of the command. Example: `"turn on/off wifi` with ID 4. This is a command, with 2 variations. The `"on"` variation outputs 4.00001 and the `"off"` variation outputs 4.00002 (increments of 0.00001 to the ID integer, beginning in 0.00001 - 1 000 000 variations possible).
- It also seems that after the library is loaded, all global variables remain with their last value, until it's unloaded. This is useful because of the types limitation below and still needing to pass a big array into the engine for it to load all possible commands --> a big string must be sent to the library for processing, but only in the beginning, which will make the performance not matter (it's only hurt on the entire program's beginning).

As this module is compiled for Android with Gomobile, it's limited to the supported types by go/build: https://pkg.go.dev/golang.org/x/mobile/cmd/gobind#hdr-Type_restrictions, so all the exported elements must follow those rules (some, as for example if a slice is exported, no error is thrown, so doesn't seem to be bad to export those to be accessible across packages - won't be accessible on Android though). So for example, to pass an array to the functions of the library, it must be encoded into a string and decoded on the function again.

## To compile the module
- To run on PC, either use an IDE which does it automatically (I use GoLand, for example), or run the following command in the project folder as working directory: "go run ACD".
- To compile for Android and create an AAR package, have a look on the Build_AAR_Android.bat file and execute the command inside it. If you use the file, make sure to change the ANDROID_HOME variable. For some reason, I can't use relative paths here, so I used an absolute one (must be doing something wrong). You might also want to run VersionUpdater.py before the batch script to update the ACD's VERSION constant to the current date/time (just to keep track of which version is being used on the AAR).

## About
### - Roadmap
Have a look on the "TODO.md" file.

### - Project status
Ongoing, but possibly slowly since I'm a student, so I may not have that much time to work on this. I also come here more rarely, mostly when I need VISOR to do something that requires an update here. Else I normally prefer to keep adding features to VISOR. Though, there's the TODO list... Not much time either. Heh.

### - License
This project is licensed under Apache 2.0 License - http://www.apache.org/licenses/LICENSE-2.0.

## Support
If you have any questions, try the options below:
- Create an Issue here: https://github.com/DADi590/Advanced-Commands-Detection/issues
- Create a Discussion here: https://github.com/DADi590/Advanced-Commands-Detection/discussions

## Final notes
Any new ideas and/or improvements are welcomed! (Just giving the idea and/or making a pull request)
