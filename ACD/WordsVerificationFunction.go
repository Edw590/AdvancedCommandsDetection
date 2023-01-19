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

import (
	"regexp"
	"strconv"
	"strings"
)

var mutually_exclusive_words = [...][]string{
	{"on", "off"},
	{"stop", "continue", "play", "resume", "next", "previous"},
}

// Internal function to choose the correct interval value based on the given intervals slices, the default value,
// and the sub-verification number.
func chooseCustomIntervals(intervs_map map[int]int, sub_verif_number int, default_interv int) int {
	var interv int = -1
	// Check for an index corresponding to this specific sub-verification.
	if interv_temp, ok := intervs_map[sub_verif_number]; ok {
		interv = interv_temp
	} else {
		// Check for an index corresponding to either even or odd sub-verifications.
		var int_to_find int = INDEX_ODD
		if 0 == (sub_verif_number % 2) {
			int_to_find = INDEX_EVEN
		}
		if interv_temp, ok = intervs_map[int_to_find]; ok {
			interv = interv_temp
		} else {
			// Check for an ALL_SUB_VERIFS index.
			if interv_temp, ok = intervs_map[ALL_SUB_VERIFS_INT]; ok {
				interv = interv_temp
			} else {
				// If nothing found of the above, use the default value.
				interv = DEFAULT_INDEX
			}
		}
	}

	if DEFAULT_INDEX == interv {
		return default_interv
	} else {
		return interv
	}
}

type WordFoundInfo struct {
	word_found           string
	index_word_found     int
	index_word_found_map int
}

const DEFAULT_INDEX int = -1
const ALL_SUB_VERIFS_INT int = -1
const INDEX_EVEN int = -2
const INDEX_ODD int = -3
const IS_DIGIT string = ";1;"
const INDEX_WORD_FOUND string = ";2;"
const INDEX_DEFAULT string = ";3;"
const NONE string = ";0;"
const NOTHING_DETECTED = -1

/*
wordsVerificationFunction iterates a sentence and searches for keywords provided on a list and returns the index of the
words condition it detected.

The function searches through intervals of words, so other words can be between the keywords, and the function will
still find the correct words. Aside from that, various parameters can be set to adjust the behavior of the function.

Note: sub-verification means "position index verification", in which the function tries to find the "position-nth" word
(the meaning of "position index" is explained on the 'words_list' parameter documentation).

-----CONSTANTS-----

  - ALL_SUB_VERIFS_INT – for 'exclude_word_found_group', 'init_indexes_sub_verifs': used to indicate it applies to all
    sub-verifications
  - INDEX_EVEN – for 'left_intervs', 'right_intervs': used to indicate the next value applies only for even indexes
  - INDEX_ODD – for 'left_intervs', 'right_intervs': same as for INDEX_EVEN but for odd indexes
  - DEFAULT_INDEX – for 'left_intervs', 'right_intervs': used to indicate the function to use the default index

– IS_DIGIT – for 'words_list': used to indicate the "word" is a digit and not really a word. Example:

	{{{-1}, {IS_DIGIT}}, {{-1}, {"alarms"}}},

Here the 1st words group must contain a digit.
  - INDEX_WORD_FOUND – for 'init_indexes_sub_verifs': used to indicate the function to use the index of the last found
    word
  - INDEX_DEFAULT – for 'init_indexes_sub_verifs': used to indicate it's to use the default index calculation

– NONE – for 'words_list': to use if one of the words group is optional. Example:

	{{{-1}, {"audio"}}, {}},
	{{{-1}, {"frontal"}}, {{-1}, {"video"}}},
	{{{-1}, {NONE, "rear"}}, {{-1}, {"video"}}},

Here, the 1st words group is optional on the last condition. The NONE keyword will only be checked after all checks are
completed on the current condition, and only in case no word was found.

-----CONSTANTS-----

-----------------------------------------------------------

> Params:
  - sentence – same as in sentenceCmdsDetector()
  - sentence_index – index on the sentence where to start the search
  - main_words – 1D slice with the words that activated the command detection. Example: for commands pair
    "set the alarm"/"new alarm", 'main_words' is {"set", "new"}

– words_list – a [][][][]interface{} that contains the words that the command accepts; their variations and with them,
variations of the return command. Examples:

	{{{-1}, {"on"}},  {{-1}, {"wifi", "wi-fi"}}},
	{{{-1}, {"off"}}, {{-1}, {"wifi", "wi-fi"}}},
	---
	{{{-1}, {"device", "phone"}}},
	{{{-1}, {"device", "phone"}}, {{-1}, {"safe"}},     {-1: {"mode"}}},
	{{{-1}, {"device", "phone"}}, {{-1}, {"recovery"}}},
	---
	{{{-1}, {NONE, "rear"}}, {{-1}, {"video"}}},
	{{{-1}, {"frontal"}},    {{-1}, {"video"}}},
	{{{-1}, {"audio"}}},
	---
	{{{-1}, {"again", "said", "say"}}},
	---
	{{{-1}, {"on"}},  {{-1}, {"battery"}}, {{-1}, {"saver"}}},
	{{{-1}, {"off"}}, {{-1}, {"battery"}}, {{-1}, {"saver"}}},

Structure and naming convention:

- Each line (like "{{{-1}, {"on"}}, {{-1}, {"wifi", "wi-fi"}}},"): condition

- Each slice inside the line: words map (it was supposed to be a map, if the "key" weren't an array - to be possible to
give multiple values)

- Each slice inside a word map (there can only be 0 or 2): the 1st is the allowed indexes list, the 2nd is the words
group

If a word map is empty, the function will disregard it and only check the non-empty ones.

Each condition with word groups containing NONE must be above all the ones that don't have NONE. As a secondary ordering
rule, each condition with more non-empty maps than the others must be above them, or the function will not detect things
correctly.

The list of allowed indexes can be an array of various numbers, being each number the position index on which the
detected word must be for a successful detection. For example, "turn on the damn wifi" ("turn", a main word, doesn't
count): "on" has index 0, "wifi" has index 1, because that would be the detection order. For "turn the damn wifi on" it
would be "wifi" 0, "on" 1. Use ALL_SUB_VERIFS_INT to allow the words group in any position.

Note that this means that the order of the maps doesn't matter AT ALL. Only the indexes matter here, as they're the ones
that specify positions. So why are the words separated? Because the words in each words group must be mutually
exclusive. Like "device" and "phone". Doesn't matter which one is detected for the word to exist. Read about the
'exclude_word_found_group' parameter to understand the actual reason.

– left_intervs – a map in which each key is either a sub-verification number or one of the constants. This slice
provides the word search intervals for the left of the found word for each sub-verification. Leave empty to apply the
default value to all sub-verifications (default is 0). Example:

	{2:"1", ALL_SUB_VERIFS_INT:"3", 4: DEFAULT_INDEX}

In this case, the normal will be to use 3, except on the sub-verification 2 (3rd) in which 1 will be used, on the
number 4 (5th) the default index will be used. Selecting 3 for all sub-verifications, for example means it will check
the 3 words before the found word for the next word on the 'words_list'. Also, if ALL_SUB_VERIFS_STR is not used and a
number is not used for a specific sub-verification, the default one will be used. Use a nil or empty slice to disregard
this feature.

  - right_intervs – same as for 'left_intervs', but for the right side. Default is 3.
  - init_indexes_sub_verifs – a map in which each key X is the number of the sub-verification, and the value is the index
    on which to begin the specified sub-verification. Note: the 1st sub-verification (number 0) always starts on index 0, so
    attempts to change that sub-verification initial index will be ignored. The value can also be one of the constants. In
    case it's INDEX_WORD_FOUND, one can put a + or a - and a number to add or subtract said number to the index of the word
    found. Example: { INDEX_WORD_FOUND + "+1"}. The default index is the index of the word found in the sub-verification.
    Use a nil or empty slice to disregard this feature.
  - exclude_word_found_group – 1D slice in which each element is a number of a sub-verification on which to exclude the
    word group found from the subsequent sub-verifications. One of the constants can also be used but only in the first
    index, and in this case, any other elements in the slice will be ignored.
    Use a nil or empty slice to disregard this feature.
  - return_last_match – true to instead of returning the first match in the 'words_list', keep searching until the end
    and return the last match; false to return the first match on the 'words_list'
  - ignore_repets_cmds – ignore repetitions of the original 'main_words' that activate the command detection. In
    the case of "set alarm"/"new alarm", ignore repetitions of "set" and "new". Example: "set alarm set 2 reminders" -->
    ignore the repetition and don't treat as being an example of "set ah... set 2 alarms" - treat as being 2 different
    commands (or don't ignore and stop the verification immediately after finding the repeated word). The function checks if
    the repeated word is before any of the found words. If it is, the verification will stop. For example, "turn on turn
    wifi on please" - if it's to ignore 2 commands to start the Wi-Fi will be returned. Else, only one.
  - exclude_main_words – true to exclude all the 'main_words' from the 'words_list' to be sure the function doesn't detect
    them by accident (which already happened), false to not exclude them.

- exclude_mutually_exclusive_words – exclude words that are mutually exclusive, like "on" and "off". With this enabled,
this function will never return results with both "on" and "off" on the detection, which already happened. Real example:

	{{{-1}, {"on"}}, {{-1}, {"wifi", "wi-fi"}}},
	{{{-1}, {"off"}}, {{-1}, {"wifi", "wi-fi"}}},

With "turn off the wifi on", the function detected both conditions and returned the first (because of the order in which
they are on the array). With this parameter enabled, it will detect "off" first, remove "on" from all other conditions
in any of the word maps, and proceed with the verification --> will never return the first condition in this case,
because "on" is no longer on it.
Note that this will not remove the word found (if it detects "off", it will only remove the other mutually exclusive
words - which excludes "off"). For that, use the 'exclude_word_found_group' parameter.

> Returns:
  - a 3D array with the first arrays being the number of command variations (len(words_list)), the second arrays being
    the number of word maps in the command variation, and the third arrays are arrays of 2 elements only, with the 1st
    a bool indicating if the word was detected or not, and the 2nd the index in which it was detected, or -1 if the 1st
    index is false (no word detected)
*/
func wordsVerificationFunction(sentence []string, sentence_index int, cmd commandInfo) [][][]interface{} {

	// Make a copy of the 'words_list', so it doesn't get modified by this function as the copy will be. Must be a real
	// copy (elements from sub-slices will be removed/added), so CopySlice().
	var words_list [][][][]interface{} = nil
	copySlice(&words_list, cmd.words_list)
	// And make a copy of the original words to use in the repeated words check. CopyOuterSlice() suffices, as it's just
	// to copy each value of the slice (which are pointers - no problem with that as the contents won't be modified).
	var original_main_words []string = copyOuterSlice(cmd.main_words).([]string)

	// If it's to exclude all the original words from the 'words_list', do it here, before the sub-verifications begin.
	if cmd.exclude_main_words {
		for i, j := range words_list {
			for ii, jj := range j {
				if 0 == len(jj) {
					// If this map is empty, go to the next one
					continue
				}

				for iii, word := range jj[1] {
					for _, main_word := range cmd.main_words {
						if word == main_word {
							delElemInSlice(&words_list[i][ii][1], iii)

							break
						}
					}
				}
			}
		}
	}

	var sentence_len int = len(sentence)
	var num_conditions int = len(words_list)

	var success_detects [][][]interface{} = nil
	var left_intervs []int = nil
	var right_intervs []int = nil
	var max_sub_verifications int = 0
	for _, j := range words_list {
		if len(j) > max_sub_verifications {
			max_sub_verifications = len(j)
		}
		success_detects = append(success_detects, nil)
	}
	for i := 0; i < max_sub_verifications; i++ {
		left_intervs = append(left_intervs, chooseCustomIntervals(cmd.left_intervs, i, 0))
		right_intervs = append(right_intervs, chooseCustomIntervals(cmd.left_intervs, i, 3))
	}

	var init_indexes []int = nil
	var indexes_previous_word_found []int = nil
	for i := 0; i < num_conditions; i++ {
		// Don't change these initial indexes unless you check their usages
		init_indexes = append(init_indexes, sentence_index)
		indexes_previous_word_found = append(indexes_previous_word_found, sentence_index)
	}

	for sub_verification := 0; sub_verification < max_sub_verifications; sub_verification++ {

		//log.Println("!!!!!!!!!!!!!!!!!")
		//log.Println("Condition index:", curr_words_cond_index)

		for curr_words_cond_index, curr_words_condition := range words_list {
			if sub_verification >= len(curr_words_condition) {
				continue
			}

			success_detects[curr_words_cond_index] = append(success_detects[curr_words_cond_index], []interface{}{true, -1})

			var init_index int = init_indexes[curr_words_cond_index]
			var index_previous_word_found int = indexes_previous_word_found[curr_words_cond_index]

			var left_interv int = left_intervs[sub_verification]
			var right_interv int = right_intervs[sub_verification]

			//log.Println(init_index)
			//log.Println(left_interv)
			//log.Println(right_interv)

			var word_found_info WordFoundInfo = WordFoundInfo{
				word_found:           NONE,
				index_word_found:     init_index, // Can't be a negative number as those are reserved for error codes...
				index_word_found_map: -1,
			}

			//log.Println("1---")

			// Word discovery
			// For each word in the current 'words_list' slice and within the words interval specified, this looks
			// for the word in the 'sentence'. When it finds one of the words in the slice, it notes down the word and
			// the index.
			var search_NONE_now bool = false
			var word_detected bool = false
			var anything_searched bool = false
		restart_words_check:
			for index := init_index - left_interv; index <= (init_index + right_interv); index++ {
				if index >= sentence_len {
					break
				}
				if index < 0 || index == init_index {
					continue
				}

				//log.Println("SDFJLH")
				//log.Println(index)

				for index_words_map, words_map := range curr_words_condition {
					if 0 == len(words_map) {
						// Nothing to check here
						continue
					}

					//log.Println("JJJJJJJJJ")
					//log.Println(sentence[index])
					//log.Println(words_map)

					var use_words_here bool = false
					if ALL_SUB_VERIFS_INT == words_map[0][0] {
						use_words_here = true
					} else {
						for _, allowed_index := range words_map[0] {
							if allowed_index == sub_verification {
								use_words_here = true

								break
							}
						}
					}
					if !use_words_here {
						// Words group not allowed for this position index, so go to the next one
						continue
					}

					anything_searched = true

					for _, word := range words_map[1] {
						if (NONE == word && !search_NONE_now) || (NONE != word && search_NONE_now) {
							// Go to the next word if the current one on the list is NONE (can't check for that in the
							// beginning, as that's what's found before starting to search: nothing --> NONE) and if
							// it's not supposed to be checked for special commands right now, or vice-versa.
							continue
						}

						// Checking special commands here
						switch word {
						case IS_DIGIT:
							{
								if _, err := strconv.Atoi(sentence[index]); nil == err {
									word_detected = true
								}
							}
						default:
							{
								// If it's not a special command, just check the word normally
								if sentence[index] == word {
									word_detected = true
								}
							}
						}

						if word_detected {
							//log.Println("+++++++++++")
							//log.Println(sentence[index])
							//log.Println(word)
							word_found_info.word_found = sentence[index]
							word_found_info.index_word_found = index
							word_found_info.index_word_found_map = index_words_map

							// Word found for the current position (sub-verification). Exit the discovery.
							goto leave_loops
						}
					}
				}
			}
		leave_loops:

			//log.Println("2---")

			if !word_detected {
				if !search_NONE_now && anything_searched {
					// If no word was found, search NONE was not enabled, and anything was searched (there are words in
					// lists for search and the sentence indexes are still inside the sentence), try again now also
					// checking for the NONE command.
					search_NONE_now = true

					goto restart_words_check
				}

				// Else, output a false to the success array and go to the next condition since this was is garbage now.
				success_detects[curr_words_cond_index][sub_verification] = []interface{}{false, -1}

				goto end_condition
			}

			//log.Println(word_found_info.word_found)
			//log.Println(success_detects)

			//log.Println("3---")

			// Check if the command seems to be repeated in the 'sentence' (for example, could be "set ah... set 2 alarms")
			// by verifying if any original word is repeated between the previous and current found words - if we're not on
			// the 1st sub-verification, because if we are, we can't check that (no previously found word).
			// - If it's not between both, all is alright with the repetition because it's not a case like said example
			//   (could be "turn on wifi turn on airplane mode" - both "turn"s are very near each other, but still, the
			//   verification will continue, because the "turn" is not between any consecutive found words, so that should
			//   mean it's an actual command).
			// - If the repeated original word is between *any* consecutive 'word_found's, then it might mean it's a case
			//   like the above and the verification will stop right away.
			if !cmd.ignore_repets_cmds {
				// No idea which index is the highest one (depends on the intervals and if the left one is > 0), so put
				// the function checking that by itself.
				var a, b int = index_previous_word_found, word_found_info.index_word_found
				var highest_index, lowest_index int = a, b
				if b > a {
					highest_index = b
					lowest_index = a
				}
				for sentence_counter := lowest_index + 1; sentence_counter <= (highest_index - 1); sentence_counter++ {
					// The last part ensures the 'sentence' word being checked is not the original one.
					if sentence_counter >= 0 && sentence_counter < sentence_len && sentence_counter != sentence_index {
						for _, original_word := range original_main_words {
							if sentence[sentence_counter] == original_word {
								// An original word is between 2 consecutive found words --> command repetition detected
								// and we go to the next condition ("by verifying if any original word is repeated
								// between the previous and current found words" - the words for detection are in the
								// various conditions. We need to check them all first in this case.)

								// Set one of the word detections to false to exclude this condition.
								success_detects[curr_words_cond_index][sub_verification] = []interface{}{false, -1}

								goto end_condition
							}
						}
					}
				}
			}

			//
			// Word detection completed. Now to prepare the next sub-verification...
			//

			//log.Println("4---")

			// Detection successful, so update the index of the word found.
			success_detects[curr_words_cond_index][sub_verification][1] = word_found_info.index_word_found

			// If there are more sub-verifications, prepare the next one
			if sub_verification != max_sub_verifications-1 {
				//log.Println("5---")

				if cmd.exclude_mutually_exclusive_words {
					//log.Println(words_list)
					var idx_array_mut_excl_words int = -1
					for i, word_slice := range mutually_exclusive_words {
						for _, word := range word_slice {
							if word == word_found_info.word_found {
								idx_array_mut_excl_words = i

								goto end_of_loops
							}
						}
					}
				end_of_loops:
					if -1 != idx_array_mut_excl_words {
						for i := range words_list {
							for ii := range words_list[i] {
								if 0 == len(words_list[i][ii]) {
									continue
								}
								for iii := 0; iii < len(words_list[i][ii][1]); iii++ {
									jjj := words_list[i][ii][1][iii]
									for _, word_to_exclude := range mutually_exclusive_words[idx_array_mut_excl_words] {
										if (word_to_exclude == jjj) && (word_to_exclude != word_found_info.word_found) {
											delElemInSlice(&words_list[i][ii][1], iii)
											iii--
										}
									}
								}
							}
						}
					}
					//log.Println(idx_array_mut_excl_words)
					//log.Println(words_list)
				}

				if len(cmd.exclude_word_found_group) > 0 {
					// If it's to exclude from the 'words_list' the word group found in this sub-verification, check
					// to see if it's the ALL_SUB_VERIFS_INT command, or if it's only for specific sub-verifications.
					var exclude_word_found_now bool = false
					if ALL_SUB_VERIFS_INT == cmd.exclude_word_found_group[0] {
						exclude_word_found_now = true
					} else {
						for _, number := range cmd.exclude_word_found_group {
							if number == sub_verification {
								exclude_word_found_now = true

								break
							}
						}
					}
					if exclude_word_found_now && (-1 != word_found_info.index_word_found_map) {
						var words_to_exclude []interface{} = nil
						// Full copy of the slice here (I'll be deleting from the original)
						copySlice(&words_to_exclude, curr_words_condition[word_found_info.index_word_found_map][1])

						for ii := range curr_words_condition {
							if 0 == len(curr_words_condition[ii]) {
								continue
							}
							var word_map []interface{} = curr_words_condition[ii][1]
							for iii := 0; iii < len(word_map); iii++ {
								jjj := word_map[iii]
								for _, word_to_exclude := range words_to_exclude {
									if word_to_exclude == jjj {
										delElemInSlice(&word_map, iii)
										iii--
									}
								}
							}
						}
					}
				}

				var init_index_next_sub_verif_str string = ""

				if 0 != len(cmd.init_indexes_sub_verifs) {
					// Here it will check if the next index was specified in the function's parameters.
					if next_index, ok := cmd.init_indexes_sub_verifs[sub_verification+1]; ok {
						init_index_next_sub_verif_str = next_index
					} else {
						// If no index was specified specifically for the current sub-verification, check with
						// ALL_SUB_VERIFS_INT.
						if next_index, ok = cmd.init_indexes_sub_verifs[ALL_SUB_VERIFS_INT]; ok {
							init_index_next_sub_verif_str = next_index
						}
					}
				}
				// Now check if the index is a number or a special command.
				if "" == init_index_next_sub_verif_str || strings.Contains(init_index_next_sub_verif_str, INDEX_DEFAULT) {
					// If no index was specified or the default one was specified, calculate the next initial index by
					// calculating an average and summing 0.5.
					// If it was index 3, it's now 3.5, which is 3 when converted to int. If it was 3.5, it's now 4.
					// It's just a way of increasing the index sometimes. Sometimes not - random. The idea of increasing a
					// bit is to increase the search interval that bit. But not as much as to put the new index, the index
					// of the word found. Like a middle term, but with more weight on the right.
					//init_index_next_sub_verif = strconv.Itoa(int(
					//	(float32(init_index_int) + float32(index_word_found))/float32(2) + float32(0.5)))

					// UPDATE: I've just got it back to the older and first way --> the index of the word found.
					// I believe the way above was because there existed no way of knowing what a "it" meant, which is
					// not true anymore.
					init_index_next_sub_verif_str = strconv.Itoa(word_found_info.index_word_found)
				} else if strings.Contains(init_index_next_sub_verif_str, INDEX_WORD_FOUND) {
					if INDEX_WORD_FOUND == init_index_next_sub_verif_str {
						init_index_next_sub_verif_str = strconv.Itoa(word_found_info.index_word_found)
					} else if strings.Contains(init_index_next_sub_verif_str, "+") {
						number, _ := strconv.Atoi(strings.Split(init_index_next_sub_verif_str, "+")[1])
						init_index_next_sub_verif_str = strconv.Itoa(word_found_info.index_word_found + number)
					} else if strings.Contains(init_index_next_sub_verif_str, "-") {
						number, _ := strconv.Atoi(strings.Split(init_index_next_sub_verif_str, "-")[1])
						init_index_next_sub_verif_str = strconv.Itoa(word_found_info.index_word_found - number)
					}
				}

				//log.Println("^^^^^^^^^^^^^^^^^^^^^")

				init_indexes[curr_words_cond_index], _ = strconv.Atoi(init_index_next_sub_verif_str)
			}

			indexes_previous_word_found[curr_words_cond_index] = word_found_info.index_word_found

		end_condition:
		}
	}

	//log.Println("________________________")

	return success_detects
}

/*
checkMainWordsRetConds checks if the results coming from wordsVerificationFunction() agree with the return conditions
for the command 'main_words'.

-----------------------------------------------------------

> Params:
  - results_wordsVerifFunc – the direct return from wordsVerificationFunction()
  - sentence_word – the current 'sentence_word' on the sentenceCmdsDetector()
  - cmds_GL_index – the current 'cmds_GL' looping index on the sentenceCmdsDetector()

> Returns:

– the index of the final accepted 'words_list' condition for the current 'sentence_word'
*/
func checkMainWordsRetConds(results_wordsVerifFunc [][][]interface{}, sentence_word string, cmds_GL_index int) int {
	var final_condition int = -1
	// Must be the biggest condition because, for example "reboot phone" and "reboot phone into
	// recovery", and the sentence is "reboot phone into recovery". Both are successful
	// detections (all words are found in both variations). But only the 2nd (the *biggest*) is
	// correct, because more words were found, and more words has higher priority than fewer
	// words.
	var biggest_len int = -1

	//log.Println(success_detects)
	for ii, jj := range results_wordsVerifFunc {
		var all_true bool = true
		for _, jjj := range jj {
			all_true = all_true && jjj[0].(bool)
		}
		if all_true {
			var main_words_ret_conds [][]string = cmds_GL[cmds_GL_index].main_words_ret_conds
			var arr_id int = 0
			if ii >= len(main_words_ret_conds) {
				// In case there are not enough return conditions, use the last one present.
				arr_id = len(main_words_ret_conds) - 1
			} else {
				arr_id = ii
			}

			var any_main_word bool = false
			if (1 == len(main_words_ret_conds[arr_id])) &&
				(ANY_MAIN_WORD == main_words_ret_conds[arr_id][0]) {
				any_main_word = true
			}
			var words_exclude_anyway []string = nil
			//log.Println("SSSSSSSSSSSSSSS")
			//log.Println(main_words_ret_conds[arr_id])
			for _, word := range main_words_ret_conds[arr_id] {
				if ANY_MAIN_WORD == word {
					any_main_word = true
				} else if strings.HasPrefix(word, "-") {
					words_exclude_anyway = append(words_exclude_anyway, regexp.MustCompile("[^a-zA-Z]+").ReplaceAllString(word, ""))
				}
			}

			//log.Println("DDDDDDDDDDDDDDDDDDD")
			//log.Println(any_main_word)
			//log.Println(words_exclude_anyway)

			for _, word := range main_words_ret_conds[arr_id] {
				if strings.HasPrefix(word, "-") {
					// Don't do anything if it's a word that beings with a "-", which means it's to exclude it from the
					// accepted main words. The actual verification will be on the ANY_MAIN_WORD command or any other
					// main words on the list. The ones beginning with "-" are only added to an exclusion list to be
					// iterated in the end of each iteration of this loop.
					continue
				}

				//log.Println("++++++++++++++")
				//log.Println(word)
				//log.Println(sentence_word)

				var actual_word string = word
				if !isSpecialCommand(actual_word) {
					// If it's a special command like ;4;, keep it like that. Else, leave only alphanumeric characters.
					// Example, "-fast", to remove the word "fast" from the accepted main words.
					regexp.MustCompile("[^a-zA-Z0-9 ]+").ReplaceAllString(word, "")
				}

				// If any main word counts, then if the current word matches the sentence word or not doesn't matter,
				// because the command was triggered by a main word, and any main word is accepted.
				if any_main_word || actual_word == sentence_word {
					// Though, there can still be words that must be excluded ("All except these: [...]").
					var exclude_word bool = false
					for _, word_exclude := range words_exclude_anyway {
						if sentence_word == word_exclude {
							exclude_word = true

							break
						}
					}
					//log.Println("FFFFFFFFFFFFFFFFF")
					//log.Println(actual_word)
					//log.Println(sentence_word)
					//log.Println(exclude_word)
					// If the 'sentence_word' is not on the excluded list, carry on.
					if !exclude_word && len(jj) > biggest_len {
						//log.Println("QQQQQQQQQQQQQQQQQQ")
						final_condition = ii
						biggest_len = len(jj)

						break
					}
				}
			}
		}
	}

	return final_condition
}
