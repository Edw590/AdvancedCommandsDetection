/*
 * Copyright 2021 DADi590
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
Package CommandsDetection_APU contains constants to be used outside the package for commands recognition.

All the possible recognizable commands are in this package. All constants started by CMD_ are to be sent to the
recognition function. All the constants started by RET_ are returned from the function to indicate the recognized
commands in the given sentence of words.

Format of CMD_-started constants: CMD_ + an indicative name of the command. The documentation of the constant will tell
more about the command.

Format of RET_-started constants: RET_ + value of the corresponding CMD_-started constants + _ + anything that indicates
the recognized command more specifically, if needed. For example, RET_0_ON means it was detected a command to turn the
flashlight on.
*/
package CommandsDetection_APU

////////////////////////////////////////////////////////////////////////////////////
// Constants of the various commands in the arrays below and in the other parts
//
// The CMD_-started constants have a number which is the index of the command place in all the arrays below this.
// --- WARNING ---
// The value 0 is reserved for function-related processing!!! (Want to know for what? Read the comment about
// MARK_TERMINATION_FLOAT32 on TaskChecker) Also, all negative values are also reserved for special commands.
//
// Note: all RET_-started constants must be a float32 in a string which starts by the number on the corresponding
// CMD_-started constant and must advance by increments of 0.1 and continues like if it were an integer > 1 (1.9,
// 1.10...). The first float can't end in .0. No reason in specific, it's just in case it's ever needed to use the main
// integer. So start for example with 0.1 and 1.1.

// CMD_TOGGLE_MEDIA is the command to stop, pause, play, or go to the next or previous media element
const CMD_TOGGLE_MEDIA string = "1"
const RET_1_STOP string = "1.1"
const RET_1_PAUSE string = "1.2"
const RET_1_PLAY string = "1.3"
const RET_1_NEXT string = "1.4"
const RET_1_PREVIOUS string = "1.5"

// CMD_ASK_TIME is the command to request the time
const CMD_ASK_TIME string = "2"
const RET_2 string = "2.1"

// CMD_ASK_DATE is the command to request the date
const CMD_ASK_DATE string = "3"
const RET_3 string = "3.1"

// CMD_TOGGLE_WIFI is the command to turn the Wi-Fi on or off
const CMD_TOGGLE_WIFI string = "4"
const RET_4_ON string = "4.1"
const RET_4_OFF string = "4.2"

// CMD_TOGGLE_MOBILE_DATA is the command to turn the Mobile Date on or off
const CMD_TOGGLE_MOBILE_DATA string = "5"
const RET_5_ON string = "5.1"
const RET_5_OFF string = "5.2"

// CMD_TOGGLE_BLUETOOTH is the command to turn the Bluetooth or off
const CMD_TOGGLE_BLUETOOTH string = "6"
const RET_6_ON string = "6.1"
const RET_6_OFF string = "6.2"

// CMD_ANSWER_CALL is the command to answer a generic call
const CMD_ANSWER_CALL string = "7"
const RET_7 string = "7.1"

// CMD_TOGGLE_FLASHLIGHT is the command to turn the flashlight on or off
const CMD_TOGGLE_FLASHLIGHT string = "8"
const RET_8_ON string = "8.1"
const RET_8_OFF string = "8.2"

// CMD_END_CALL is the command to end a generic call
const CMD_END_CALL string = "9"
const RET_9 string = "9.1"

// CMD_TOGGLE_SPEAKERS is the command to turn the speaker(s) on or off
const CMD_TOGGLE_SPEAKERS string = "10"
const RET_10_ON string = "10.1"
const RET_10_OFF string = "10.2"

// CMD_TOGGLE_AIRPLANE_MODE is the command to turn the airplane mode on or off
const CMD_TOGGLE_AIRPLANE_MODE string = "11"
const RET_11_ON string = "11.1"
const RET_11_OFF string = "11.2"

// CMD_ASK_BATTERY_PERCENT is the command to request the battery percentage
const CMD_ASK_BATTERY_PERCENT string = "12"
const RET_12 string = "12.1"

// CMD_SHUT_DOWN_DEVICE is the command to shut down a device
const CMD_SHUT_DOWN_DEVICE string = "13"
const RET_13 string = "13.1"

// CMD_REBOOT_DEVICE is the command to reboot a device, either normally, or into safe mode
const CMD_REBOOT_DEVICE string = "14"
const RET_14_NORMAL string = "14.1"
const RET_14_SAFE_MODE string = "14.2"

// CMD_TAKE_PHOTO is the command to take a rear or frontal photo. If neither rear nor frontal is specified, the default
// is RET_15_REAR.
const CMD_TAKE_PHOTO string = "15"
const RET_15_REAR string = "15.1"
const RET_15_FRONTAL string = "15.2"

// CMD_RECORD_MEDIA is the command to record media, like video or sound, and frontal or rear in case of video. Again,
// in case neither rear nor frontal is specified for video recording, the default is RET_16_VIDEO_REAR.
const CMD_RECORD_MEDIA string = "16"
const RET_16_AUDIO string = "16.1"
const RET_16_VIDEO_REAR string = "16.2"
const RET_16_VIDEO_FRONTAL string = "16.3"

// CMD_SAY_AGAIN is the command to request the assistant to repeat what it just said
const CMD_SAY_AGAIN string = "17"
const RET_17 string = "17.1"

// CMD_MAKE_CALL is the command to make a call (no names, just make a generic call)
const CMD_MAKE_CALL string = "18"
const RET_18 string = "18.1"

// HIGHEST_CMD_INT is a constant which has an always-updated value of the highest CMD_-started constant. This can be
// used to build a slice of integers from 1 to this value to use with Main(), and it will always have all the
// possible commands allowed for detection.
const HIGHEST_CMD_INT int = 18

// Special WARN_-started commands returned by the CmdsDetector() - must not collide with spec_-started constants on
// TaskChecker!!!

// WARN_WHATS_IT is the constant that signals that an "it" was said but there seems to be nothing that it refers to, so
// the assistant warns it didn't understand the meaning of the "it".
const WARN_WHATS_IT string = "-2"

////////////////////////////////////////////////////////////////////////////////////
// ATTENTION: keep the format of the below slices as it is. Each index must correspond to the value of the CMD_-started
// constant corresponding to the slice index contents in question.
//
// ----- NEVER EVER REMOVE AN ELEMENT FROM THE SLICES UNLESS IT'S THE LAST ONE!!! -----
// If it's to deactivate one of the elements, delete everything from it and put a // to do without the space in the
// front to be seen well that that index is not being used (if it's to be used again, it's a notification on the IDE
// warning about unused slots, so it's good to be there).

var main_words_GL = [...][]string{
	{}, // Ignored
	{"stop", "pause", "continue", "play", "resume", "next", "previous"}, // 1
	{"what's", "tell", "say"},    // 2
	{"what's", "tell", "say"},    // 3
	{"turn", "get"},              // 4
	{"turn", "get"},              // 5
	{"turn", "get"},              // 6
	{"answer", "pick"},           // 7
	{"turn", "get"},              // 8
	{"end", "stop", "terminate"}, // 9
	{"turn", "get"},              // 10
	{"turn", "get"},              // 11
	{"how's", "what's", "tell"},  // 12
	{"shutdown", "power"},        // 13
	{"reboot", "restart"},        // 14
	{"take"},                     // 15
	{"record"},                   // 16
	{"say", "what"},              // 17
	{"make"},                     // 18
}

var words_list_GL = [...][][]string{
	{}, // Ignored
	{ // 1
		{"media", "song", "songs", "music", "musics", "video", "videos"},
	},
	{ // 2
		{"time"},
	},
	{ // 3
		{"date"},
	},
	{ // 4
		{"on", "off", "wifi", "wi-fi"},
		{"on", "off", "wifi", "wi-fi"},
	},
	{ // 5
		{"on", "off", "data", "connection", "mobile"},
		{"data", "connection", "mobile"},
		{"on", "off", "data", "connection", "mobile"},
		{"on", "off", "connection"},
	},
	{ // 6
		{"on", "off", "bluetooth"},
		{"on", "off", "bluetooth"},
	},
	{ // 7
		{"call"},
	},
	{ // 8
		{"on", "off", "lantern", "flashlight", "flash"},
		{"on", "off", "lantern", "flashlight", "flash"},
	},
	{ // 9
		{"call"},
	},
	{ // 10
		{"on", "off"},
		{"speaker", "speakers"},
	},
	{ // 11
		{"on", "off", "mode", "airplane"},
		{"mode", "airplane"},
		{"on", "off", "mode"},
	},
	{ // 12
		{"battery", "percentage", "status", "state"},
	},
	{ // 13
		{"off", "down", "phone", "tablet", "device", "computer", "pc"},
		{"off", "down", "phone", "tablet", "device", "computer", "pc"},
	},
	{ // 14
		{"phone", "tablet", "device", "computer", "pc"},
		{"safe"},
		{"mode"},
	},
	{ // 15
		{"frontal", "front", "rear"},
		{"selfie", "picture", "photo", "photograph"},
	},
	{ // 16
		{"frontal", "front", "rear"},
		{"video", "camera", "audio", "sound"},
	},
	{ // 17
		{"again", "said", "say"},
	},
	{ // 18
		{"call"},
	},
}
