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

package AdvancedCommandsDetection

import (
	"strconv"
	"strings"
)

// The value of each TYPE constant is its index on the cmds_types_keywords array

const CMDi_TYPE_MANUAL string = "0"
const CMDi_TYPE_TURN_ONFF string = "1"
const CMDi_TYPE_ASK string = "2"
const CMDi_TYPE_STOP string = "3"
const CMDi_TYPE_ANSWER string = "4"
const CMDi_TYPE_SHUT_DOWN string = "5"
const CMDi_TYPE_REBOOT string = "6"
const CMDi_TYPE_REPEAT_SPEECH string = "7"
const CMDi_TYPE_RECORD string = "8"

// Each list of type keywords can have at most 2 arrays inside it (if more are needed, change the implementation, maybe
// even generalize it - for now it's made for case of 1 array and case of 2 arrays).
// The 2 arrays are of words that must be mixed with the command keywords to create the command ("turn", "on" + "wifi",
// for example).
var cmds_types_keywords = [...][][]string{
	{}, // 0 - Ignored
	{ // 1
		{"turn", "get", "switch", "put"}, // Default main words for this type to put be on the main_words array
		{"on", "off"},                    // Other optional words to continue the first ones if necessary (turn... what? something. what?
		// on or off)
	},
	{ // 2
		{"what's", "tell", "say"},
	},
	{ // 3
		{"stop", "end", "finish", "cease", "conclude", "terminate"},
	},
	{ // 4
		{"answer", "reply", "respond", "acknowledge"},
	},
	{ // 5
		{"shut", "power"},
		{"down", "off"},
	},
	{ // 6
		{"reboot", "restart"},
	},
	{ // 7
		{"what", "say", "come", "go", "repeat"},
	},
	{ // 8
		{"record"},
	},
}

func PrepareCmdsArray(commands_str string) {
	// Reset the commands array
	cmds_GL = nil

	var commands_info []string = strings.Split(commands_str, "\\")

	for i := 0; i < len(commands_info); i++ {
		var cmd_info []string = strings.Split(commands_info[i], "||")

		cmd_id, _ := strconv.Atoi(cmd_info[0])

		cmds_GL = append(cmds_GL, commandInfo{
			cmd_id:                           cmd_id,
			main_words:                       nil,
			main_words_ret_conds:             nil,
			words_list:                       nil,
			left_intervs:                     nil,
			right_intervs:                    nil,
			init_indexes_sub_verifs:          nil,
			exclude_word_found_group:         nil,
			ignore_repets_cmds:               false,
			exclude_mutually_exclusive_words: true,
		})

		prepareCmdArray(&cmds_GL[i], strings.Split(cmd_info[1], " "), strings.Split(cmd_info[2], " "), cmd_info[3],
			strings.Split(cmd_info[4], "|"))
	}

	//log.Println(len(cmds_GL))
	//log.Println("===========")
}

func prepareCmdArray(cmd_info_GL *commandInfo, types_str []string, main_words_alt []string,
	main_words_ret_conds_str string, words_list_param []string) {
	var types_int []int = nil
	for _, j := range types_str {
		type_int, _ := strconv.Atoi(j)
		types_int = append(types_int, type_int)
	}

	/////////////////////////////////
	// Arrays processing

	// main_words
	for i, j := range types_str {
		if CMDi_TYPE_MANUAL == j {
			cmd_info_GL.main_words = append(cmd_info_GL.main_words, main_words_alt...)
		} else {
			cmd_info_GL.main_words = append(cmd_info_GL.main_words, cmds_types_keywords[types_int[i]][0]...)
		}
	}
	//log.Println(cmd_info_GL.main_words)

	// words_list
	// "device/phone safe mode|device/phone recovery|device/phone"
	var words_list [][][][]interface{} = nil
	for condition_str_num, condition_str := range words_list_param {
		words_list = append(words_list, nil)
		for ii, words_group := range strings.Split(condition_str, " ") {
			words_list[condition_str_num] = append(words_list[condition_str_num], nil)
			words_list[condition_str_num][ii] = append(words_list[condition_str_num][ii], []interface{}{-1})
			words_list[condition_str_num][ii] = append(words_list[condition_str_num][ii], nil)
			for _, word := range strings.Split(words_group, "/") {
				words_list[condition_str_num][ii][1] = append(words_list[condition_str_num][ii][1], word)
			}
		}
	}
	for i, j := range types_str {
		switch j {
		case CMDi_TYPE_TURN_ONFF:
			{
				words_list[0] = append(words_list[0], nil)
				var words_list_0_len int = len(words_list[0])
				words_list[0][words_list_0_len-1] = append(words_list[0][words_list_0_len-1], []interface{}{-1})
				words_list[0][words_list_0_len-1] = append(words_list[0][words_list_0_len-1], []interface{}{"on"})

				words_list = append(words_list, nil)
				copySlice(&words_list[1], words_list[0])
				words_list[1][words_list_0_len-1][1][0] = "off"
			}
		default:
			{
				if 2 == len(cmds_types_keywords[types_int[i]]) {
					for ii := 0; ii < len(words_list); ii++ {
						words_list[ii] = append(words_list[ii], nil)
						var words_list_i_len int = len(words_list[ii])
						words_list[ii][words_list_i_len-1] = append(words_list[ii][words_list_i_len-1],
							[]interface{}{-1})
						words_list[ii][words_list_i_len-1] = append(words_list[ii][words_list_i_len-1], nil)
						for _, j := range cmds_types_keywords[types_int[i]][1] {
							words_list[ii][words_list_i_len-1][1] = append(words_list[ii][words_list_i_len-1][1], j)
						}
					}
				}
			}
		}
	}
	cmd_info_GL.words_list = words_list
	//log.Println(cmd_info_GL.words_list)

	// main_words_ret_conds
	if "" == main_words_ret_conds_str {
		var words_list_len int = len(cmd_info_GL.words_list)
		for i := 0; i < words_list_len; i++ {
			cmd_info_GL.main_words_ret_conds = append(cmd_info_GL.main_words_ret_conds, []string{ANY_MAIN_WORD})
		}
	} else {
		for _, j := range strings.Split(main_words_ret_conds_str, "|") {
			cmd_info_GL.main_words_ret_conds = append(cmd_info_GL.main_words_ret_conds, strings.Split(j, "/"))
		}
	}
	//log.Println(cmd_info_GL.main_words_ret_conds)

	// exclude_word_found
	cmd_info_GL.exclude_word_found_group = append(cmd_info_GL.exclude_word_found_group, ALL_SUB_VERIFS_INT)

	//log.Println("---------")
}
