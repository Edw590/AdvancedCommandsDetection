/*
 * Copyright 2021 DADi590
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"log"
	"os"
	"strings"

	"AdvancedCommandsDetection/AdvancedCommandsDetection"
)

const ERR_NOT_ENOUGH_ARGS string = "3234_ERR_NOT_ENOUGH_ARGS"

func main() {
	log.Println("V.I.S.O.R.'s Advanced Commands Detection module.")
	log.Println("---")
	log.Println("")
	log.Println("Version of the module: " + AdvancedCommandsDetection.VERSION)
	log.Println("")
	log.Println("//-->3234_BEGINNING<--\\\\")
	log.Println("")

	const CMD_TOGGLE_FLASHLIGHT string = "1"
	const CMD_ASK_TIME string = "2"
	const CMD_ASK_DATE string = "3"
	const CMD_TOGGLE_WIFI string = "4"
	const CMD_TOGGLE_MOBILE_DATA string = "5"
	const CMD_TOGGLE_BLUETOOTH string = "6"
	const CMD_ANSWER_CALL string = "7"
	const CMD_END_CALL string = "9"
	const CMD_TOGGLE_SPEAKERS string = "10"
	const CMD_TOGGLE_AIRPLANE_MODE string = "11"
	const CMD_ASK_BATTERY_PERCENT string = "12"
	const CMD_SHUT_DOWN_DEVICE string = "13"
	const CMD_REBOOT_DEVICE string = "14"
	const CMD_TAKE_PHOTO string = "15"
	const CMD_RECORD_MEDIA string = "16"
	const CMD_SAY_AGAIN string = "17"
	const CMD_MAKE_CALL string = "18"
	const CMD_TOGGLE_POWER_SAVER_MODE string = "19"
	const CMD_STOP_RECORD_MEDIA string = "20"
	const CMD_STOP_MEDIA string = "21"
	const CMD_PAUSE_MEDIA string = "22"
	const CMD_PLAY_MEDIA string = "23"
	const CMD_NEXT_MEDIA string = "24"
	const CMD_PREVIOUS_MEDIA string = "25"

	var commands = [...][]string{
		{CMD_TOGGLE_WIFI, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "wifi"},
		{CMD_TOGGLE_MOBILE_DATA, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "mobile data"},
		{CMD_TOGGLE_BLUETOOTH, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "bluetooth"},
		{CMD_TOGGLE_FLASHLIGHT, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "flashlight/lantern"},
		{CMD_TOGGLE_SPEAKERS, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "speaker/speakers"},
		{CMD_TOGGLE_AIRPLANE_MODE, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "airplane mode"},
		{CMD_TOGGLE_POWER_SAVER_MODE, AdvancedCommandsDetection.CMDi_TYPE_TURN_ONFF, "", "", "power/battery saver"},
		{CMD_ASK_TIME, AdvancedCommandsDetection.CMDi_TYPE_ASK, "", "", "time"},
		{CMD_ASK_DATE, AdvancedCommandsDetection.CMDi_TYPE_ASK, "", "", "date"},
		{CMD_ASK_BATTERY_PERCENT, AdvancedCommandsDetection.CMDi_TYPE_ASK, "", "", "battery percentage", "battery status", "battery level"},
		{CMD_SAY_AGAIN, AdvancedCommandsDetection.CMDi_TYPE_REPEAT_SPEECH, "", "", "again", "say", "said"},
		{CMD_END_CALL, AdvancedCommandsDetection.CMDi_TYPE_STOP, "", "", "call"},
		{CMD_STOP_RECORD_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_STOP, "", "", "recording audio/sound|recording video/camera"},
		{CMD_ANSWER_CALL, AdvancedCommandsDetection.CMDi_TYPE_ANSWER, "", "", "call"},
		{CMD_SHUT_DOWN_DEVICE, AdvancedCommandsDetection.CMDi_TYPE_SHUT_DOWN, "", "", "device/phone"},
		{CMD_REBOOT_DEVICE, AdvancedCommandsDetection.CMDi_TYPE_REBOOT, "", "", "device/phone safe mode|device/phone recovery|device/phone"},
		{CMD_STOP_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "stop", "", "media/song/songs/music/musics/video/videos"},
		{CMD_PAUSE_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "pause", "", "media/song/songs/music/musics/video/videos"},
		{CMD_PLAY_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "play continue resume", "", "media/song/songs/music/musics/video/videos"},
		{CMD_NEXT_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "next", "", "media/song/songs/music/musics/video/videos"},
		{CMD_PREVIOUS_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "previous", "", "media/song/songs/music/musics/video/videos"},
		{CMD_TAKE_PHOTO, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "take", "", "frontal picture/photo|picture/photo"},
		{CMD_RECORD_MEDIA, AdvancedCommandsDetection.CMDi_TYPE_RECORD, "", "", "audio/sound|frontal video/camera|video/camera"},
		{CMD_MAKE_CALL, AdvancedCommandsDetection.CMDi_TYPE_MANUAL, "make place", "", "call"},
	}

	var commands_almost_str []string = nil
	for _, array := range commands {
		commands_almost_str = append(commands_almost_str, strings.Join(array, "||"))
	}
	var commands_str string = strings.Join(commands_almost_str, "\\")

	log.Println(commands_str)

	AdvancedCommandsDetection.PrepareCmdsArray(commands_str)

	arguments := os.Args
	if len(arguments) > 1 {
		log.Println("To do: " + AdvancedCommandsDetection.Main(os.Args[1]))
	} else {
		log.Println(ERR_NOT_ENOUGH_ARGS)
		//os.Exit(0) // Comment to dev mode (run stuff written below without needing CMD parameters)
	}
	log.Println("")

	var sentence_str string = "shut down the phone and reboot it"
	// todo None of these below work decently... Fix them.
	//var sentence_str string = "stop and play the song"
	// This above needs the change on the TO DO file. It needs to know it's to STOP_MEDIA. "song" is more than 3 words
	// away from "stop" - no detection.
	//var sentence_str string = "the video stop it and then play it again"
	// This above needs a change in the NLPAnalyzer...
	// The 1st "it" is "video", so it's replaced when the sentence_counter gets to "it". At that time, the "and"
	// function has "stop" stored as a verb, but the counter is on the "it" place. When "video" is deleted after being
	// replaced and the "and" checker goes to check what's on the current word, what's on its place is "video", but
	// I think it thinks it's "it" - so it's stored too as a "non-name" --> wrong!
	// Reset the counters every time or something.

	log.Println(sentence_str) // Just to also see it on the terminal (better than getting back here just to read it)
	log.Println("To do: " + AdvancedCommandsDetection.MainInternal(sentence_str))
	log.Println("")
	log.Println("\\\\-->3234_END<--//")

	// Uncomment to test if the commands detection is still functioning well after modifications to the engine.
	//testCommandsDetection()
}

func testCommandsDetection() {
	log.Println("Running commands detection tests...")

	// Tests of good functioning of the commands detector (the sentence and the expected output).
	// Only put commands here that have once worked, and so they must continue to work even after updates to the
	// detection engine.
	var sentences [][]string = [][]string{
		{"turn off airplane mode on", "11.02"},
		{"turn on turn off the wifi", "4.02"},
		{"turn on wifi and the bluetooth no don't turn it on", "4.01"},
		{"turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on", "11.01, 11.02, 4.01"},
		{"turn on turn wifi on please", "4.01"},
		{"turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please", "-10, 4.01, 11.02, 4.02"},
		{"turn on wifi and and the airplane mode and the flashlight", "4.01, 11.01, 1.01"},
	}

	var successes int = 0
	var problems []string = nil
	for _, j := range sentences {
		var detected_commands string = AdvancedCommandsDetection.Main(j[0])
		if detected_commands != j[1] {
			problems = append(problems, "PROBLEM DETECTED: "+j[0]+" / "+j[1]+" --> "+detected_commands)
		} else {
			successes++
		}
	}
	log.Println("Results (successes/total):", successes, "/", len(sentences))
	for _, j := range problems {
		log.Println(j)
	}
}
