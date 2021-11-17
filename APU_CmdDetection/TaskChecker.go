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

package APU_CmdDetection

import (
	"Assist_Platforms_Unifier/APU_GlobalUtilsInt"
	"fmt"
	"log"
	"strconv"
	"strings"
)

/*
GenerateListAllCmds generates a string with a list of all the available commands specified on each of the CMD_-started
constants. This string can be used directly with CmdsDetector().

-----------------------------------------------------------

> Params:

- none


> Returns:

- a string with all the available commands separated by CMDS_SEPARATOR
*/
func GenerateListAllCmds() string {
	var ret_var string = ""
	for counter := 1; counter < HIGHEST_CMD_INT; counter++ {
		ret_var += strconv.Itoa(counter) + CMDS_SEPARATOR
	}

	return ret_var[:len(ret_var)-len(CMDS_SEPARATOR)]
}


const UNKNOWN_ERR string = "3234_UNKNOWN_ERR - "
/*
CmdsDetector is the function to be called to process tasks in a given sentence of words.

-----------------------------------------------------------

> Params:

- sentence_str – a sentence of words, for example coming directly from speech recognition

- allowed_cmds  a string containing a list of the CMD_-started constants of all the commands that are allowed to be
returned if found on the 'sentence', separated by CMDS_SEPARATOR


> Returns:

- a list of the detected commands in the form
	[CMD1][separator][CMD2][separator][CMD3]
being CMDS_SEPARATOR the separator; in case there were no detected commands, an empty string; in case there was an
unknown error, a string beginning with UNKNOWN_ERR.
*/
func CmdsDetector(sentence_str string, allowed_cmds string) string {
	var ret_var string = ""

	APU_GlobalUtilsInt.Tcf{
		Try: func() {
			ret_var = CmdsDetectorInternal(sentence_str, allowed_cmds)
		},
		Catch: func(e APU_GlobalUtilsInt.Exception) {
			ret_var = UNKNOWN_ERR + fmt.Sprint(e)
		},
	}.Do()

	return ret_var
}


const CMDS_SEPARATOR string = ", " // Leave the space (mess because of a char being a uint8 and not a string)
/*
CmdsDetectorInternal is the actual function that will do what's written on CmdsDetector() - continue reading there.

There is just one exception, which is this one doesn't return any error code in case anything goes wrong - it will
panic instead (no protection here), so always call the other one.

Note: if you find this function exported, know it's just for testing from the main package. Do not use it in production.
*/
func CmdsDetectorInternal(sentence_str string, allowed_cmds_str string) string {
	var sentence []string = strings.Split(sentence_str, " ")
	var allowed_cmds []int = nil
	for _, cmd_index := range strings.Split(allowed_cmds_str, CMDS_SEPARATOR) {
		number, _ := strconv.Atoi(cmd_index)
		allowed_cmds = append(allowed_cmds, number)
	}

	var sentence_cmds []float32 = sentenceCmdsChecker(sentence, allowed_cmds)

	taskFilter(&sentence_cmds)

	var ret_var string = ""
	for _, command := range sentence_cmds {
		ret_var += fmt.Sprintf("%g", command) + CMDS_SEPARATOR
	}

	log.Println("::::::::::::::::::::::::::::::::::")
	log.Println(ret_var)

	if strings.HasSuffix(ret_var, CMDS_SEPARATOR) {
		ret_var = ret_var[:len(ret_var)-len(CMDS_SEPARATOR)]
	}

	//ret_var = removeRepeatedCmds(ret_var) - let's see if the verification function can handle it without this...

	log.Println(ret_var)
	log.Println("::::::::::::::::::::::::::::::::::")

	return ret_var
}

/*
removeRepeatedCmds removes immediately repeated commands from the commands verification ([1, 3, 3, 4, 3, 4] will become
[1, 3, 4, 3, 4], for example).

The idea of this function is to kind of fix the problem of wrongly detected repeated commands.

For example, with the command (punctuation added for better understanding - remove it to test):
"turn on wifi and get the airplane mode on. no, don't turn the wifi on. turn off airplane mode and turn the wifi on.",
the result is the following:
	"3234_wifi(),on \\// 3234_wifi(),on \\// 3234_wifi(),on \\// 3234_wifi(),on \\// 3234_wifi(),on \\//
	3234_airplane_mode(),on \\// 3234_airplane_mode(),on \\// 3234_wifi(),on \\// 3234_wifi(),on \\// 3234_wifi(),on
	\\// ".
Awfully wrong. This function improves that to "3234_wifi(),on \\// 3234_airplane_mode(),on \\// 3234_wifi(),on".
The first command is wrong, the but idea here is to delete all the repeated elements (which improved MUCH in this
case).

Though, this also poses the problem of deleting purposefully repeated commands... Will be used until the
wordsVerificationDADi() can do the job better. In that case might be better (for now) to say the repeated commands in
another function call.

(As a curiosity, the overall CmdsDetector() function is now capable of knowing what to do in the example above, without
this function being executed at all! KEEP THE RIGHT SEARCH INTERVAL IN 3 AND THE LEFT ONE IN 0 AS DEFAULT!!!!!)
*/
func removeRepeatedCmds(ret_var string) string {
	var ret_var_list []string = strings.Split(ret_var, CMDS_SEPARATOR)

	const MARK_TERMINATION_STR string = "3234_MARK_TERMINATION_STR"

	var ret_var_list_len = len(ret_var_list) // Optimization
	for counter := 0; counter < ret_var_list_len; counter++ {
		if counter != ret_var_list_len - 1 {
			if ret_var_list[counter+1] == ret_var_list[counter] {
				ret_var_list[counter] = MARK_TERMINATION_STR
			}
		}
	}

	ret_var = ""
	for _, command := range ret_var_list {
		if command != MARK_TERMINATION_STR {
			ret_var += command + CMDS_SEPARATOR
		}
	}

	if strings.HasSuffix(ret_var, CMDS_SEPARATOR) {
		ret_var = ret_var[:len(ret_var)-len(CMDS_SEPARATOR)]
	}

	return ret_var
}


//const spec_cmd_dont_instead_CONST float32 = -1.1
//const spec_cmd_stop_CONST float32 = -2
//const spec_cmd_forget_CONST float32 = -3
const spec_cmd_dont_CONST float32 = -1

/*
sentenceCmdsChecker checks a sentence for commands whose indexes are listed in an array of numbers.

-----------------------------------------------------------

> Params:

- sentence – a 1D array of words on which the verification will be executed

- allowed_cmds – same as in CmdsDetector() but here it's in an array of integers and not as a string


> Returns:

- a slice on which each index is a command found in the 'sentence', represented by one of its RET_-started constant
*/
func sentenceCmdsChecker(sentence []string, allowed_cmds []int) []float32 {
	var ret_var []float32 = nil

	var sentence_len int = len(sentence)
	for main_counter, main_word := range sentence {

		if main_word == "don't" || main_word == "dont" || main_word == "do" {
			var carry_on bool = false
			// This below checks (in the 2nd part) if the next element exists or not in the 'sentence'
			if main_word == "do" && main_counter + 1 != sentence_len - 1 {
				if sentence[main_counter+1] == "not" {
					carry_on = true
				}
			} else {
				carry_on = true
			}
			if carry_on {
				ret_var = append(ret_var, spec_cmd_dont_CONST)
			}
		} else {
			for _, cmd_index := range allowed_cmds {
				for _, word := range main_words_GL[cmd_index] {
					if word == main_word {
						/*if cmd_index != 6 {
							// For testing
							continue
						}*/

						log.Println("==============")
						log.Println(word)
						log.Println(cmd_index)

						var results_WordsVerificationDADi [][]string = wordsVerificationDADi(sentence, main_counter,
							main_words_GL[cmd_index], words_list_GL[cmd_index], left_intervs_GL[cmd_index],
							right_intervs_GL[cmd_index], init_indexes_sub_verifs_GL[cmd_index],
							exclude_word_found_GL[cmd_index], return_last_match_GL[cmd_index],
							ignore_repets_main_words_GL[cmd_index], ignore_repets_original_word_GL[cmd_index],
							order_words_list_GL[cmd_index], stop_first_not_found_GL[cmd_index],
							exclude_original_words_GL[cmd_index], continue_with_words_array_number_GL[cmd_index])

						log.Println("-----------")
						log.Println(results_WordsVerificationDADi)

						if checkResultsWordsVerifDADi(words_list_GL[cmd_index], main_word,
								results_WordsVerificationDADi, conditions_continue_GL[cmd_index],
								conditions_not_continue_GL[cmd_index]) {
							log.Println("LLLLLLL")
							var sub_cond_match_found bool = false
							for _, sub_condition := range conditions_return_GL[cmd_index] {
								//log.Println("++++++++++++")
								if len(sub_condition) == 1 {
									float, _ := strconv.ParseFloat(sub_condition[0][0], 32)
									ret_var = append(ret_var, float32(float))

									break
								} else {
									var parameters_matched bool = true
									var sub_condition_len = len(sub_condition)
									for _, parameter := range sub_condition[:sub_condition_len-1] {
										results_index, _ := strconv.Atoi(parameter[0])
										var word_match bool = false
										//log.Println("-------")
										for _, word_1 := range parameter[1:] {
											//log.Println(word_1)
											if results_index == -1 {
												if word == word_1 {
													word_match = true

													break
												}
											} else {
												//log.Println(results_WordsVerificationDADi[results_index][0])
												if results_WordsVerificationDADi[results_index][0] == word_1 {
													//log.Println("KKKKKKKKKKKKK")
													word_match = true

													break
												}
											}
										}
										parameters_matched = parameters_matched && word_match
										if !parameters_matched {
											break
										}
									}
									if parameters_matched {
										float, _ := strconv.ParseFloat(sub_condition[sub_condition_len-1][0], 32)
										ret_var = append(ret_var, float32(float))

										sub_cond_match_found = true
									}
								}
								if sub_cond_match_found {
									log.Println("QQQQQQQ")
									log.Println(ret_var)
									break
								}
							}
						}
					}
				}
			}
		}
	}

	return ret_var
}

/*
taskFilter filters a sentence of commands depending on special commands present on it.

For example, "turn on the lights and play some music. no, don't turn on the lights" --> the special command here is
"don't", this function will only leave on the array the music command.

-----------------------------------------------------------

> Params:

- sentence_cmds – same as in sentenceCmdsChecker()


> Returns:

- nothing
*/
func taskFilter(sentence_cmds *[]float32) {
	// For testing
	//*sentence_filtered = [][]string{{"test"}, {"test"}, {"test 234 lkj"}, {"test"}, {"test"}, {"test"}, {"test"},
	//	{"test"}, {"test"}, {"test"}, {"test"}, {"test"}, {"test"}, {"test"}, {"test"}, }
	//*sentence_cmds = []float32{24, -1, 26, 25, -1, -1, -1, 25, 24}

	log.Println("==============================================")
	log.Println("*sentence_cmds -->", *sentence_cmds)

	// RESTRICTED VALUE ON THE sentence_cmds ARRAY - Used to mark elements for deletion on the 2 arrays. This way
	// they're deleted only in the end and on the main loop it doesn't get confusing about what indexes have been
	// deleted already and stuff.
	var MARK_TERMINATION_FLOAT32 float32 = 0

	for counter, number := range *sentence_cmds {
		if number == spec_cmd_dont_CONST {

			var delete_number_before_dont bool = false

			// Delete the "don't"
			(*sentence_cmds)[counter] = MARK_TERMINATION_FLOAT32

			//log.Println("1 -", *sentence_cmds)
			if counter != len(*sentence_cmds) - 1 {
				// If the next index is within the maximum index (which means, if the next number exists)...

				var next_number float32 = (*sentence_cmds)[counter+1]
				if next_number > 0 { // Means if it's a normal command. If it is, assume the below case.
					// Case: "do [1] and do [2]. no don't do [1]" - delete this, don't, and this. Also, if by any reason
					// there are more copies of [1], delete them also - if they're before the next element only.

					var number_mentioned bool = false
					var pos_next_number []int = nil
					for counter1, number1 := range *sentence_cmds {
						if number1 == next_number {
							pos_next_number = append(pos_next_number, counter1)
							number_mentioned = true
						}
						if counter1 == counter {
							// Stop when it gets to before the next element
							break
						}
					}
					if number_mentioned {
						// If the number was mentioned before (like [24, 25, 24, -1, 24]), delete all copies and the -1.
						(*sentence_cmds)[counter+1] = MARK_TERMINATION_FLOAT32

						//log.Println("2 -", *sentence_cmds)

						for _, index_element := range pos_next_number {
							(*sentence_cmds)[index_element] = MARK_TERMINATION_FLOAT32
						}
						//log.Println("3 -", *sentence_cmds)
					} else {
						// Else, delete only the element before the current "don't" (if there exists one).
						// Example: [24, -1, 26, 25, -1, 25, 24] will become [26, 24], because, "do 24, no don't do
						// it. do 26 and do 25. no don't do 25. do 24."
						delete_number_before_dont = true
					}
				}// else if next_number < 0 { // If it's not, assume the below case.
				// Case: "do this, no don't do it don't do it. do that". Delete only the current "don't" until
				// there's only one (done above), which will be the one used to decide what to delete.
			} else {
				// If there's no more elements, there can be previous ones. So delete the previous number to the "don't".
				// Which would be a "do [1]. no, never mind, don't do it".
				delete_number_before_dont = true
			}

			if delete_number_before_dont {
				if counter - 1 >= 0 {
					(*sentence_cmds)[counter-1] = MARK_TERMINATION_FLOAT32
					//log.Println("4 -", *sentence_cmds)
				}
			}
		}
	}

	//log.Println("5 -", *sentence_cmds)

	// Delete all elements marked for deletion
	for counter := 0; counter < len(*sentence_cmds); counter++ {
		// Don't forget (again) the length is checked every time on the loop
		if (*sentence_cmds)[counter] == MARK_TERMINATION_FLOAT32 {
			APU_GlobalUtilsInt.DelEleInSlice(sentence_cmds, counter)
			counter--
		}
	}

	log.Println("*sentence_cmds -->", *sentence_cmds)
	log.Println("==============================================")
}
