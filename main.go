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
		//os.Exit(0) // Comment to dev mode (run stuff put below)
	}
	log.Println("")
	// Test of good functioning of the commands detector.
	//log.Println("To do: " + APU_CmdDetection.CmdsDetector("turn on wifi and get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on", all_allowed_cmds))

	// todo None of these below work decently... Fix them.

	sentenceAnalyzer()

	//log.Println("To do: " + APU_CmdDetection.CmdsDetector("what's the time and the date please", all_allowed_cmds))
	// todo The keyword there is "and" --> do something about that? NLP... verb name AND name (so it's the same verb)
	// Also beware that the "and" is referring to some verb before or after it. So put the first names before and after
	// the "and" to the same action.
	// Except.... "airplane mode" --> both are names... Name immediately after name behaves as one name/command?
	// ONLY USE NLP WHEN IT IS REALLY NEEDED! Else, use the function without its help (it may fail, like with "turn on
	// the wifi", on which it thinks "turn" is a name and not a verb, but works in a bigger sentence).
	//log.Println("To do: " + APU_CmdDetection.CmdsDetector("the wifi turn it on", all_allowed_cmds))
	// todo Use NLP to exclude an adjective if the only one was said already, for example? (Or a verb or whatever)
	// todo Besides that, use it to know what a "it" means... Last name said?

	log.Println("")
	log.Println("\\\\-->3234_END<--//")
}
