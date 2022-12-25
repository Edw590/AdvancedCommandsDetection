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

	"Assist_Platforms_Unifier/CommandsDetection_APU"
	"Assist_Platforms_Unifier/Utils_APU"
)

const ERR_NOT_ENOUGH_ARGS string = "3234_ERR_NOT_ENOUGH_ARGS"

func main() {
	log.Println(Utils_APU.ASSISTANT_NAME + "'s Platform Unifier module.")
	log.Println("---")
	log.Println("")
	log.Println("Version of the module: " + Utils_APU.VERSION)
	log.Println("")
	log.Println("//-->3234_BEGINNING<--\\\\")
	log.Println("")

	var all_allowed_cmds string = CommandsDetection_APU.GenerateListAllCmds()
	log.Println(all_allowed_cmds)

	arguments := os.Args
	if len(arguments) > 1 {
		log.Println("To do: " + CommandsDetection_APU.Main(os.Args[1], all_allowed_cmds))
	} else {
		log.Println(ERR_NOT_ENOUGH_ARGS)
		//os.Exit(0) // Comment to dev mode (run stuff written below)
	}
	log.Println("")
	// Tests of good functioning of the commands detector.
	//var sentence_str string = "turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on" // 11.1, 11.2, 4.1
	//var sentence_str string = "turn on turn wifi on please" // 4.1
	//var sentence_str string = "turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please" // -2, 4.1, 11.2, 4.2
	//var sentence_str string = "turn on wifi and and the airplane mode and the flashlight" // 4.1, 11.1, 8.1

	// todo None of these below work decently... Fix them.

	var sentence_str string = "reboot to recovery"
	//var sentence_str string = "the video stop it and then play it again"
	// This above needs a change in the NLPAnalyzer...
	// The 1st "it" is "video", so it's replaced when the sentence_counter gets to "it". At that time, the "and"
	// function has "stop" stored as a verb, but the counter is on the "it" place. When "video" is deleted after being
	// replaced and the "and" checker goes to check what's on the current word, what's on its place is "video", but
	// I think it thinks it's "it" - so it's stored too as a "non-name" --> wrong!
	// Reset the counters every time or something.

	log.Println(sentence_str) // Just to also see it on the terminal (better than get back here to read it)
	log.Println("To do: " + CommandsDetection_APU.MainInternal(sentence_str, all_allowed_cmds))
	log.Println("")
	log.Println("\\\\-->3234_END<--//")
}
