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

	"AdvancedCommandsDetection/ACD"
)

const ERR_NOT_ENOUGH_ARGS string = "3234_ERR_NOT_ENOUGH_ARGS"

func main() {
	log.Println("V.I.S.O.R.'s Advanced Commands Detection module.")
	log.Println("---")
	log.Println("")
	log.Println("Version of the module: " + ACD.VERSION)
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
	const CMD_CONTROL_MEDIA string = "21"
	const CMD_CONFIRM string = "22"

	var commands = [...][]string{
		{CMD_TOGGLE_FLASHLIGHT, ACD.CMDi_TYPE_TURN_ONFF, "", "", "flashlight/lantern"},
		{CMD_ASK_TIME, ACD.CMDi_TYPE_ASK, "", "", "time"},
		{CMD_ASK_DATE, ACD.CMDi_TYPE_ASK, "", "", "date"},
		{CMD_TOGGLE_WIFI, ACD.CMDi_TYPE_TURN_ONFF, "", "", "wifi"},
		{CMD_TOGGLE_MOBILE_DATA, ACD.CMDi_TYPE_TURN_ONFF, "", "", "mobile data"},
		{CMD_TOGGLE_BLUETOOTH, ACD.CMDi_TYPE_TURN_ONFF, "", "", "bluetooth"},
		{CMD_ANSWER_CALL, ACD.CMDi_TYPE_ANSWER, "", "", "call"},
		{CMD_END_CALL, ACD.CMDi_TYPE_STOP, "", "", "call"},
		{CMD_TOGGLE_SPEAKERS, ACD.CMDi_TYPE_TURN_ONFF, "", "", "speaker/speakers"},
		{CMD_TOGGLE_AIRPLANE_MODE, ACD.CMDi_TYPE_TURN_ONFF, "", "", "airplane mode"},
		{CMD_ASK_BATTERY_PERCENT, ACD.CMDi_TYPE_ASK, "", "", "battery percentage", "battery status", "battery level"},
		{CMD_SHUT_DOWN_DEVICE, ACD.CMDi_TYPE_SHUT_DOWN, "", "", "device/phone"},
		{CMD_REBOOT_DEVICE, ACD.CMDi_TYPE_REBOOT, "fast", "fast|;4; -fast", "reboot/restart device/phone|device/phone|device/phone recovery|device/phone safe mode|device/phone bootloader"},
		{CMD_TAKE_PHOTO, ACD.CMDi_TYPE_NONE, "take", "", "picture/photo|frontal picture/photo"},
		{CMD_RECORD_MEDIA, ACD.CMDi_TYPE_START, "record", "record|record|;4; -record", "audio/sound|video/camera|recording audio/sound|recording video/camera"},
		{CMD_SAY_AGAIN, ACD.CMDi_TYPE_REPEAT_SPEECH, "", "", "again", "say", "said"},
		{CMD_MAKE_CALL, ACD.CMDi_TYPE_NONE, "make place", "", "call"},
		{CMD_TOGGLE_POWER_SAVER_MODE, ACD.CMDi_TYPE_TURN_ONFF, "", "", "power/battery saver"},
		{CMD_STOP_RECORD_MEDIA, ACD.CMDi_TYPE_STOP, "", "", "recording audio/sound|recording video/camera"},
		{CMD_CONTROL_MEDIA, ACD.CMDi_TYPE_NONE, "play continue resume pause stop next previous", "play continue resume|pause|stop|next|previous", "media/song/songs/music/audio/musics/video/videos"},
		{CMD_CONFIRM, ACD.CMDi_TYPE_NONE, "i", "", "do/confirm/approve/certify"},
	}

	var commands_almost_str []string = nil
	for _, array := range commands {
		commands_almost_str = append(commands_almost_str, strings.Join(array, "||"))
	}
	var commands_str string = strings.Join(commands_almost_str, "\\")

	log.Println(commands_str)

	ACD.ReloadCmdsArray(commands_str)

	arguments := os.Args
	if len(arguments) > 1 {
		log.Println("To do: " + ACD.Main(os.Args[1], false, true))
	} else {
		log.Println(ERR_NOT_ENOUGH_ARGS)
		//os.Exit(0) // Comment to dev mode (run stuff written below without needing CMD parameters)
	}
	log.Println("")

	var sentence_str string = "start recording audio"
	//var sentence_str string = "record audio"
	// todo None of these below work decently... Fix them.
	//var sentence_str string = "stop and play the song"
	// This above needs the change on the TO DO file. It needs to know it's to STOP_MEDIA. "song" is more than 3 words
	// away from "stop" - no detection.

	log.Println(sentence_str) // Just to also see it on the terminal (better than getting back here just to read it)
	log.Println("To do: " + ACD.MainInternal(sentence_str, false, true))
	log.Println("")
	log.Println("\\\\-->3234_END<--//")

	// Uncomment to test if the commands detection is still functioning well after modifications to the engine.
	testCommandsDetection()
}

func testCommandsDetection() {
	log.Println("Running commands detection tests...")

	// Tests of good functioning of the commands detector (the sentence and the expected output).
	// Only put commands here that have once worked, and so they must continue to work even after updates to the
	// detection engine.
	var sentences [][]interface{} = [][]interface{}{
		{"turn off airplane mode on", true, "11.00002"},
		{"turn on turn off the wifi", true, "4.00002"},
		{"turn on wifi and the bluetooth no don't turn it on", true, "4.00001"},
		{"turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on", false, "11.00001, 11.00002, 4.00001"}, // false because this is not realistic (why turn on and off on the same command...?). This is just to test a complex example.
		{"turn on turn wifi on please", true, "4.00001"},
		{"turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please", true, "-10, 4.00001, 11.00002, 4.00002"},
		{"turn on wifi and and the airplane mode and the flashlight", true, "4.00001, 11.00001, 1.00001"},
		{"shut down the phone and then reboot it", true, "13.00001, 14.00002"},
		{"fast reboot the phone", true, "14.00001"},
		{"fast phone recovery", true, ""},
		{"the video stop it and then play it again", true, "21.00003, 21.00001"},
		{"stop the song and play the next one", true, "21.00003, 21.00004"},
	}

	var successes int = 0
	var problems []string = nil
	for _, j := range sentences {
		var detected_commands string = ACD.MainInternal(j[0].(string), false, true) //j[1].(bool), true)
		if detected_commands != j[2] {
			problems = append(problems, "PROBLEM DETECTED: "+j[0].(string)+" / "+j[2].(string)+" --> "+detected_commands)
		} else {
			successes++
		}
	}
	log.Println("Results (successes/total):", successes, "/", len(sentences))
	for _, j := range problems {
		log.Println(j)
	}
}
