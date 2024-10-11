/*******************************************************************************
 * Copyright 2023-2024 Edw590
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
		sentence:               "turn off airplane mode on",
		exp_cmd_list:           "11.00002",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "airplane mode|turn off on|",
	}, { // 2
		sentence:               "turn on turn off the wifi",
		exp_cmd_list:           "4.00002",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "wifi|turn off the|",
	}, { // 3
		sentence:               "turn on wifi and the bluetooth no don't turn it on",
		exp_cmd_list:           "4.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "bluetooth|turn on|",
	}, { // 4
		sentence:               "turn wifi and on get the airplane mode on no don't turn the wifi on turn off airplane mode and turn the wifi on",
		exp_cmd_list:           "11.00001, 11.00002, 4.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "wifi|turn the on|",
	}, { // 5
		sentence:               "turn on turn wifi on please",
		exp_cmd_list:           "4.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "wifi|turn on please|",
	}, { // 6
		sentence:               "turn it on turn on the wifi and and the airplane mode get it it on no don't turn it on turn off airplane mode and also the wifi please",
		exp_cmd_list:           "-10, 4.00001, 11.00002, 4.00002",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "wifi|turn off|",
	}, { // 7
		sentence:               "turn wifi on and and the airplane mode and the flashlight",
		exp_cmd_list:           "4.00001, 11.00001, 1.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "flashlight|turn on|",
	}, { // 8
		sentence:               "shut down the phone and then reboot it",
		exp_cmd_list:           "13.00001, 14.00002",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "phone|reboot|",
	}, { // 9
		sentence:               "fast reboot the phone",
		exp_cmd_list:           "14.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "phone|fast reboot the|",
	}, { // 10
		sentence:               "fast phone recovery",
		exp_cmd_list:           "",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "phone recovery||",
	}, { // 11
		sentence:               "the video stop it and then play it again",
		exp_cmd_list:           "21.00003, 21.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "video|play again|",
	}, { // 12
		sentence:               "stop the song and play the next one",
		exp_cmd_list:           "21.00003, 21.00004",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "song|play the next|",
	}, { // 13
		sentence:               "and the airplane mode too",
		exp_cmd_list:           "11.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "wifi|turn the on|",
		exp_cmd_info:           "airplane mode|turn the on|",
	}, { // 14
		sentence:               "and now turn it off",
		exp_cmd_list:           "4.00002",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "wifi|turn on the|",
		exp_cmd_info:           "wifi|turn off|",
	}, { // 15
		sentence:               "tell me the weather and the news",
		exp_cmd_list:           "26.00001, 27.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "news|tell me the|",
	}, { // 16
		sentence:               "turn on the mobile data and the bluetooth never mind don't do it turn on the wifi",
		exp_cmd_list:           "4.00001",
		remove_repet_cmds:      false,
		invalidate_detec_words: true,
		prev_cmd_info:          "|",
		exp_cmd_info:           "wifi|turn on the|",
	},
}
