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
	"strings"

	"ACD/ACD"
)

type commandTestsInfo struct {
	sentence               string
	exp_cmd_list           string
	remove_repet_cmds      bool
	invalidate_detec_words bool
	prev_cmd_info          string
	exp_cmd_info           string
}

func testCommandsDetection() {
	log.Println("Running commands detection tests...")

	var successes int = 0
	var problems []string = nil
	for _, j := range commands_tests {
		var output string = ACD.MainInternal(j.sentence, j.remove_repet_cmds, j.invalidate_detec_words, j.prev_cmd_info)
		var output_list []string = strings.Split(output, ACD.INFO_CMDS_SEPARATOR)
		var cmd_info string = output_list[0]
		var detected_commands string = ""
		if len(output_list) > 1 {
			detected_commands = output_list[1]
		}
		if detected_commands != j.exp_cmd_list || cmd_info != j.exp_cmd_info {
			problems = append(problems, "PROBLEM DETECTED: "+j.sentence+" / "+j.exp_cmd_list+" / "+j.exp_cmd_info+" -----> "+output)
		} else {
			successes++
		}
	}
	log.Println("Results (successes/total):", successes, "/", len(commands_tests))
	for _, j := range problems {
		log.Println(j)
	}
}

// Tests of good functioning of the commands detector.
// Only put commands here that have once worked, and so they must continue to work even after updates to the detection
// engine.
var commands_tests = [...]commandTestsInfo{
	{ // 1
		"turn off airplane mode on",
		"11.00002",
		false, true,
		"|",
		"airplane mode|turn off on|",
	}, { // 2
		"turn on turn off the wifi",
		"4.00002",
		false, true,
		"|",
		"wifi|turn off the|",
	}, { // 3
		"turn on wifi and the bluetooth no don't turn it on",
		"4.00001",
		false, true,
		"|",
		"bluetooth|turn on|",
	}, { // 4
		"turn wifi and on get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on",
		"11.00001, 11.00002, 4.00001",
		false, true,
		"|",
		"wifi|turn the on|",
	}, { // 5
		"turn on turn wifi on please",
		"4.00001",
		false, true,
		"|",
		"wifi|turn on please|",
	}, { // 6
		"turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please",
		"-10, 4.00001, 11.00002, 4.00002",
		false, true,
		"|",
		"wifi|turn off|",
	}, { // 7
		"turn wifi on and and the airplane mode and the flashlight",
		"4.00001, 11.00001, 1.00001",
		false, true,
		"|",
		"flashlight|turn on|",
	}, { // 8
		"shut down the phone and then reboot it",
		"13.00001, 14.00002",
		false, true,
		"|",
		"phone|reboot|",
	}, { // 9
		"fast reboot the phone",
		"14.00001",
		false, true,
		"|",
		"phone|fast reboot the|",
	}, { // 10
		"fast phone recovery",
		"",
		false, true,
		"|",
		"phone recovery||",
	}, { // 11
		"the video stop it and then play it again",
		"21.00003, 21.00001",
		false, true,
		"|",
		"video|play again|",
	}, { // 12
		"stop the song and play the next one",
		"21.00003, 21.00004",
		false, true,
		"|",
		"song|play the next|",
	}, { // 13
		"and the airplane mode too",
		"11.00001",
		false, true,
		"wifi|turn the on|",
		"airplane mode|turn the on|",
	}, { // 14
		"and now turn it off",
		"4.00002",
		false, true,
		"wifi|turn on the|",
		"wifi|turn off|",
	},
}
