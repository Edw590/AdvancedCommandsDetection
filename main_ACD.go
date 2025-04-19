/*******************************************************************************
 * Copyright 2023-2025 Edw590
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
 ******************************************************************************/

package main

import (
	"log"
	"os"
	"strings"

	"ACD/ACD"
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
	const CMD_REJECT string = "23"
	const CMD_STOP_LISTENING string = "24"
	const CMD_START_LISTENING string = "25"
	const CMD_TELL_WEATHER string = "26"
	const CMD_TELL_NEWS string = "27"
	const CMD_GONNA_SLEEP string = "28"
	const CMD_HELP_VISION string = "32"

	var commands = [...][]string{
		// {command ID, types separated by "+", manual main words, return conditions for the main words, list of words
		//  separated by "|" with optional words separated by "/"} - look at the examples below
		{CMD_TOGGLE_FLASHLIGHT, ACD.CMDi_TYPE_TURN_ONFF, "", "", "flashlight/lantern"},
		{CMD_ASK_TIME, ACD.CMDi_TYPE_ASK, "", "", "time"},
		{CMD_ASK_DATE, ACD.CMDi_TYPE_ASK, "", "", "date/day/month/year"},
		{CMD_TOGGLE_WIFI, ACD.CMDi_TYPE_TURN_ONFF, "", "", "wifi"},
		{CMD_TOGGLE_MOBILE_DATA, ACD.CMDi_TYPE_TURN_ONFF, "", "", "mobile data"},
		{CMD_TOGGLE_BLUETOOTH, ACD.CMDi_TYPE_TURN_ONFF, "", "", "bluetooth"},
		{CMD_ANSWER_CALL, ACD.CMDi_TYPE_ANSWER, "", "", "call"},
		{CMD_END_CALL, ACD.CMDi_TYPE_STOP, "", "", "call"},
		{CMD_TOGGLE_SPEAKERS, ACD.CMDi_TYPE_TURN_ONFF, "", "", "speaker/speakers"},
		{CMD_TOGGLE_AIRPLANE_MODE, ACD.CMDi_TYPE_TURN_ONFF, "", "", "airplane mode"},
		{CMD_ASK_BATTERY_PERCENT, ACD.CMDi_TYPE_ASK, "", "", "battery percentage/status/level/levels"},
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
		{CMD_REJECT, ACD.CMDi_TYPE_NONE, "i", "", "don't/reject/disapprove"},
		{CMD_STOP_LISTENING, ACD.CMDi_TYPE_STOP, "", "", "listening"},
		{CMD_START_LISTENING, ACD.CMDi_TYPE_START, "", "", "listening"},
		{CMD_TELL_WEATHER, ACD.CMDi_TYPE_ASK, "", "", "weather"},
		{CMD_TELL_NEWS, ACD.CMDi_TYPE_ASK, "", "", "news"},
		{CMD_GONNA_SLEEP, ACD.CMDi_TYPE_WILL_GO, "", "", "sleep"},
		{CMD_HELP_VISION, ACD.CMDi_TYPE_NONE, "help", "", "this image/picture|image/picture clipboard/copied"},
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
		log.Println("To do: " + ACD.Main(os.Args[1], false, true, "|"))
	} else {
		log.Println(ERR_NOT_ENOUGH_ARGS)
		//os.Exit(0) // Comment to dev mode (run stuff written below without needing CMD parameters)
	}
	log.Println("")

	var sentence_str string = "take a frontal picture and a rear picture"
	// TODO: None of these below work decently... Fix them.
	//var sentence_str string = "take a frontal picture and a rear picture"

	log.Println(sentence_str) // Just to also see it on the terminal (better than getting back here just to read it)
	log.Println("To do: " + ACD.MainInternal(sentence_str, false, true, "|"))
	log.Println("")
	log.Println("\\\\-->3234_END<--//")

	// Uncomment to test if the commands detection is still functioning well after modifications to the engine.
	testCommandsDetection()
}
