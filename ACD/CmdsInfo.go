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

package ACD

////////////////////////////////////////////////////////////////////////////////////
// Constants of the various commands in the arrays below and in the other parts
//
// --- WARNING ---
// The command ID 0 is reserved for function-related processing!!! (Want to know for what? Read the comment about
// MARK_TERMINATION_FLOAT32 on TaskChecker.) All negative values are also reserved for special commands.
//
// Note: all RET_-started constants must be a float32 in a string which starts by the number on the corresponding
// CMD_-started constant and must advance by increments of 0.01 (means 100 sub-commands at most). The first float can't
// end in .0 (99 then, not 100). No reason in specific, it's just in case it's ever needed to use the main integer. So
// start for example with 1.01 and 2.01.

var cmds_GL []commandInfo = nil

// commandInfo represents a command that this module detects. For use with wordsVerificationFunction() - check the
// meaning of each attribute there.
type commandInfo struct {
	// Positive (1+) integer
	cmd_id     int
	main_words []string

	/*
		Example of how it could be for some commands:
		{ // 4
			{";ANY;"},
			{";ANY;"},
		},
		{ // 14
			{";ANY;"},
			{";ANY;"},
			{"reboot restart"},
			{";ANY;"},
		},
	*/
	main_words_ret_conds [][]string

	/*
		Example of how it could be for some commands (it now includes the old conditions_continue and conditions_return):
		{ // 4
			{{{-1}, {"on"}}, {{-1}, {"wifi", "wi-fi"}}},
			{{{-1}, {"off"}}, {{-1}, {"wifi", "wi-fi"}}},
		},
		{ // 14
			{{{-1}, {"device", "phone"}}, {{-1}, {"safe"}}, {{-1}, {"mode"}}},
			{{{-1}, {"device", "phone"}}, {{-1}, {"recovery"}}},
			{{{-1}, {"device", "phone"}}},
			{{{-1}, {"device", "phone"}}},
		},
		{ // 16
			{{{-1}, {NONE, "rear"}}, {{-1}, {"video"}}},
			{{{-1}, {"frontal"}}, {{-1}, {"video"}}},
			{{{-1}, {"audio"}}},
		},
		{ // 17
			{{-1}, {{"again", "said", "say"}}},
		},
		{ // 19
			{{{-1}, {"on"}}, {-1: {"battery", "power"}}, {{-1}, {"saver"}}},
			{{{-1}, {"off"}}, {-1: {"battery", "power"}}, {{-1}, {"saver"}}},
		},
	*/
	words_list [][][][]interface{}

	left_intervs                     map[int]int
	right_intervs                    map[int]int
	init_indexes_sub_verifs          map[int]string
	exclude_word_found_group         []int
	ignore_repets_cmds               bool
	exclude_main_words               bool
	exclude_mutually_exclusive_words bool
}

// Special WARN_-started commands returned by the sentenceCmdsDetector() - must not collide with spec_-started constants
// on TaskChecker!!!

// WARN_WHATS_IT is the constant that signals that an "it" was said but there seems to be nothing that it refers to, so
// the assistant warns it didn't understand the meaning of the "it".
const WARN_WHATS_IT string = "-10"
