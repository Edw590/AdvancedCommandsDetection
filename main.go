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

package main

import (
	"Assist_Platforms_Unifier/APU_CmdDetection"
	"Assist_Platforms_Unifier/APU_GlobalUtils"
	"log"
	"os"
)

const ERR_NOT_ENOUGH_ARGS string = "3234_ERR_NOT_ENOUGH_ARGS"

func main() {
	log.Println(APU_GlobalUtils.ASSISTANT_NAME + "'s Platform Unifier module.")
	log.Println("---")
	log.Println("")
	log.Println("Version of the module: " + APU_GlobalUtils.VERSION)
	log.Println("")
	log.Println("//-->3234_BEGINNING<--\\\\")
	log.Println("")

	var all_allowed_cmds string = APU_CmdDetection.GenerateListAllCmds()
	log.Println(all_allowed_cmds)

	arguments := os.Args
	if len(arguments) > 1 {
		log.Println("To do: " + APU_CmdDetection.CmdsDetector(os.Args[1], all_allowed_cmds))
	} else {
		log.Println(ERR_NOT_ENOUGH_ARGS)
		//os.Exit(0) // Comment to dev mode (run stuff written below)
	}
	log.Println("")
	// Tests of good functioning of the commands detector.
	//var sentence_str string = "turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on"
	//var sentence_str string = "turn on turn wifi on please"
	var sentence_str string = "turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and the wifi please"
	//var sentence_str string = "turn on wifi and and the airplane mode and the flashlight"

	//var sentence_str string = "shutdown turn it on that's mom's turn on the wifi and and the airplane mode get it it on no don't turn it on turn airplane mode off and the wifi please"
	//var sentence_str string = "what's the time and the date please"
	//var sentence_str string = "the airplane mode turn it freaking on"

	log.Println("To do: " + APU_CmdDetection.CmdsDetectorInternal(sentence_str, all_allowed_cmds))

	// todo None of these below work decently... Fix them.

	//log.Println("To do: " + APU_CmdDetection.CmdsDetector("what's the time and the date please", all_allowed_cmds))
	//log.Println("To do: " + APU_CmdDetection.CmdsDetector("the wifi turn it on", all_allowed_cmds))

	log.Println("")
	log.Println("\\\\-->3234_END<--//")
}
