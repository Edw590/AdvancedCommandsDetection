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

package CommandsDetection_APU

import (
	"log"
	"strconv"
	"strings"

	"Assist_Platforms_Unifier/GlobalUtilsInt_APU"
)

const DEFAULT_INDEX string = "3234_DEFAULT_INDEX"
const ALL_SUB_VERIFS_STR string = "3234_ALL_SUB_VERIFS_STR"
const ALL_SUB_VERIFS_INT int = -1
const INDEX_EVEN string = "3234_INDEX_EVEN"
const INDEX_ODD string = "3234_INDEX_ODD"
const IS_DIGIT string = "3234_IS_DIGIT"
const INDEX_WORD_FOUND string = "3234_INDEX_WORD_FOUND"
const INDEX_DEFAULT string = "3234_INDEX_DEFAULT"
const NONE string = "3234_NONE"
const WVD_ERR_1 string = "-1"
const WVD_ERR_2 string = "-2"
const WVD_ERR_3 string = "-3"

/*
wordsVerificationFunction iterates a sentence and searches for keywords provided on a list and returns the words it found.

The function searches through intervals of words, so other words can be between the keywords, and the function will
still find the correct words. Aside from that, various parameters can be set to adjust the behavior of the function. A
good idea might be to automate the decision of some parameters, if possible.

-----CONSTANTS-----

- ALL_SUB_VERIFS_STR – for 'left_intervs', 'right_intervs': used to indicate it applies to all sub-verifications

- ALL_SUB_VERIFS_INT – for 'exclude_found_word': used to indicate it applies to all sub-verifications

- INDEX_EVEN – for 'left_intervs', 'right_intervs': used to indicate the next value applies only for even indexes

- INDEX_ODD – for 'left_intervs', 'right_intervs': same as for INDEX_EVEN but for odd indexes

- DEFAULT_INDEX – for 'left_intervs', 'right_intervs': used to indicate the function to use the default index

- IS_DIGIT – for 'words_list': used to indicate the "word" is a digit and not really a word

- INDEX_WORD_FOUND - for 'init_indexes_sub_verifs': used to indicate the function to use the index of the last found
word

- INDEX_DEFAULT for 'init_indexes_sub_verifs': used to indicate it's to use the default index calculation

- NONE – for the returning value: used on the 1st index of each slice to indicate that the word that should
be on that index was not found

- WVD_ERR_1 – for the returning value: used to indicate the function found repeated words, and the verification stopped
(which can only happen if 'ignore_repets_main_words' and/or 'ignore_repets_cmds' are set to true). In this case, NONE
will be returned as the found word on all sub-slices and this constant as index.

- WVD_ERR_2 – for the returning value: used to indicate the function stopped the verification because a word was not
found (which can only happen with 'stop_first_not_found' set to true). In this case, NONE will be returned as the found
word on all sub-slices and this constant as index.

- WVD_ERR_3 – for the returning value: used to indicate the function stopped the verification because it detected a
command repetition (which can only happen with 'ignore_repets_cmds' set to false). In this case, NONE will be returned
as the found word on all sub-slices and this constant as index.

-----CONSTANTS-----

-----------------------------------------------------------

> Params:

- sentence – same as in sentenceCmdsDetector()

- sentence_index – index on the sentence where to start the search

- main_words – 1D slice with the words that activated the command detection. Example: for commands pair
"set the alarm"/"new alarm", 'main_words' is {"set", "new"}

- words_list – a 2D slice on which each slice contains words to be checked on the 'sentence' on the sub-verification
corresponding to the index of such slice. Example:
	{{"word_1_1st_sub_verif", "word_2_1st_sub_verif"}, {"word_1_2nd_sub_verif"}}

- left_intervs – a map in which each key is either a sub-verification number or one of the constants. This slice
provides the word search intervals for the left of the found word for each sub-verification. Leave empty to apply the
default value to all sub-verifications (default is 0). Example:
	{"2":"1", ALL_SUB_VERIFS_STR:"3", "4": DEFAULT_INDEX}
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

- exclude_found_word – 1D slice in which each element is a number of a sub-verification on which to exclude the word
found from the subsequent sub-verifications. One of the constants can also be used but only in the first index, and in
this case, any other elements in the slice will be ignored. Use a nil or empty slice to disregard this feature.

- return_last_match – true to instead of returning the first match in the 'words_list', keep searching until the end
and return the last match; false to return the first match on the 'words_list'

- ignore_repets_main_words – ignore repetitions of the words in the 'main_words' in the 'sentence' in an interval given
by 'right_intervs' for the current sub-verification (for example "turn on on wifi" - if this parameter is false, the
verification will stop immediately after finding the repeated word).

- ignore_repets_cmds – ignore repetitions of the original 'main_words' that activate the command detection. In
the case of "set alarm"/"new alarm", ignore repetitions of "set" and "new". Example: "set alarm set 2 reminders" -->
ignore the repetition and don't treat as being an example of "set ah... set 2 alarms" - treat as being 2 different
commands (or don't ignore and stop the verification immediately after finding the repeated word). The function checks if
the repeated word is before any of the found words. If it is, the verification will stop. For example, "turn on turn
wifi on please" - if it's to ignore 2 commands to start the Wi-Fi will be returned. Else, only one.

- order_words_list – true to iterate the 'words_list' to find an occurrence in the 'sentence', false to do the
opposite

- stop_first_not_found – true to stop the verification at the first time one word from the 'words_list' is not found,
false to continue even with words not found until the end

- exclude_original_words – true to exclude all the 'main_words' from the 'words_list' to be sure the function doesn't
detect them by accident (which already happened), false to not exclude them.

- continue_with_words_slice_number – after having no more slices of the 'words_list' to continue the search, continue
searching using the slice of index specified here. Use -1 to not use this parameter. Good use with
'stop_first_not_found' to search more words on the same list until one is not found. For example to simulate key
presses: "Control Control F1 F4 F3 ok done", and it will stop on "ok".


> Returns:

- a 2D slice in which each inner slice has the word found (or one of the non-error constants) on the 1st index, and on
the 2nd index, the index on which the word was found (or one of the error constants). In case a word was not found,
NONE will be used and on the index will be a non-negative index - discard that index, it's wrong (no word found, how
could there be an index)
*/
func wordsVerificationFunction(sentence []string, sentence_index int, main_words []string, words_list [][]string,
	left_intervs map[string]string, right_intervs map[string]string,
	init_indexes_sub_verifs map[string]string, exclude_found_word []int, return_last_match bool,
	ignore_repets_main_words bool, ignore_repets_cmds bool, order_words_list bool,
	stop_first_not_found bool, exclude_original_words bool,
	continue_with_words_slice_number int) [][]string {
	// Note: this function was created recursive (not exactly sure why I found it easier to do as recursive, but now I
	// don't find it that easy...) and is now a loop for easier debugging. Useful to know because I didn't rewrite all
	// the comments that explained the function when it was recursive.

	var ret_var [][]string = nil

	// Make a copy of the 'words_list', so it doesn't get modified by this function as the copy will be. Must be a real
	// copy (elements from sub-slices will be removed/added), so CopySlice().
	var words_list_int [][]string = nil
	GlobalUtilsInt_APU.CopySlice(&words_list_int, words_list)
	// And make a copy of the original words to use in the repeated words check. CopyOuterSlice() suffices, as it's just
	// to copy each value of the slice (which are pointers - no problem with that as the contents won't be modified).
	var original_words []string = GlobalUtilsInt_APU.CopyOuterSlice(main_words).([]string)

	// If it's to exclude all the original words from the 'words_list_int', do it here, before the sub-verifications
	// starts.
	if exclude_original_words {
		for _, main_word := range main_words {
			for counter, words_slice := range words_list_int {
				for counter1, word := range words_slice {
					if word == main_word {
						GlobalUtilsInt_APU.DelElemInSlice(&words_list_int[counter], counter1)

						break
					}
				}
			}
		}
	}

	var init_index_int int = sentence_index
	var index_previous_word_found int = -1 // Don't change this unless you check the implementation

	var words_list_length int = len(words_list_int) // In a variable because it can be changed, as it's done inside
	// the loop when an element is added to it.
	for sub_verification := 0; sub_verification < words_list_length; sub_verification++ {
		//log.Println("!!!!!!!!!!!!!!!!!")
		// (sub_verification is also the iterating index of 'words_list_int')

		// Internal function to choose the correct interval value based on the given intervals slices, the default value,
		// and the sub-verification number.
		chooseCustomIntervals := func(intervs_map map[string]string, sub_verif_number int, default_interv int) int {
			var interv_str string = ""
			// Check for an index corresponding to this specific sub-verification.
			if interv, ok := intervs_map[strconv.Itoa(sub_verif_number)]; ok {
				interv_str = interv
			} else {
				// Check for an index corresponding to either even or odd sub-verifications.
				var string_to_find string = INDEX_ODD
				if 0 == sub_verif_number%2 {
					string_to_find = INDEX_EVEN
				}
				if interv, ok = intervs_map[string_to_find]; ok {
					interv_str = interv
				} else {
					// Check for an ALL_SUB_VERIFS index.
					if interv, ok = intervs_map[ALL_SUB_VERIFS_STR]; ok {
						interv_str = interv
					} else {
						// If nothing found of the above, return the default value.

						return default_interv
					}
				}
			}

			if DEFAULT_INDEX == interv_str {
				return default_interv
			} else {
				interv, _ := strconv.Atoi(interv_str)

				return interv
			}
		}

		var left_interv int = chooseCustomIntervals(left_intervs, sub_verification, 0)
		var right_interv int = chooseCustomIntervals(right_intervs, sub_verification, 3)

		//log.Println("sub_verif_number -", sub_verif_number)

		//log.Println("left_interv -", left_interv)
		//log.Println("right_interv -", right_interv)

		//log.Println("words_list_int[sub_verification+1:] -->", words_list_int[sub_verification+1:])

		var sentence_len int = len(sentence)

		var word_found string = NONE
		var index_word_found int = init_index_int // Can't be a negative number as those are reserved for error codes...

		/*
			Given a 'word', the 'sentence' slice, and the current 'index' of iterating the 'sentence', this checks if the
			'word' is equal to the word on the current 'index' of the 'sentence', or in case it's a special command as those
			on the function documentation, it checks if the word on the index is equivalent to said special command.
			Example: in case the 'word' is a special command requesting a number, this checks if the word on the 'index' of
			'sentence' is a number.
			Returns true if it found a match, false otherwise.
		*/
		check_word := func(word string, index int) bool {
			// Checking special commands here.
			switch word {
			case IS_DIGIT:
				{
					if _, err := strconv.Atoi(sentence[index]); nil == err {
						word_found = sentence[index]
						index_word_found = index
						if !return_last_match {
							return true
						}
					}
				}
			}

			// If it's not a special command, just find the word normally.
			if sentence[index] == word {
				word_found = word
				index_word_found = index
				if !return_last_match {
					return true
				}
			}

			// No match, return false.
			return false
		}

		//log.Println("1---")

		// Depending on 'order_words_list', iterate in 2 different ways as explained in the function documentation.
		if order_words_list {
			for _, word := range words_list_int[sub_verification] {
				for sentence_counter := init_index_int - left_interv; sentence_counter <= init_index_int+right_interv; sentence_counter++ {
					// The last part ensures it's not checking the current word (not very useful)
					if sentence_counter >= 0 && sentence_counter < sentence_len && sentence_counter != init_index_int {
						if check_word(word, sentence_counter) {
							break
						}
					}
				}
				if word_found != NONE && !return_last_match {
					// In case a match was found and in case it's not to return the last match, stop now. If it's to return
					// the last match, carry on until the end of both for loops.
					break
				}
			}
		} else {
			// For each word in the current 'words_list_int' slice and within the words interval specified, this looks
			// for the word in the 'sentence'. When it finds one of the words in the slice, it notes down the word and
			// the index.
			for sentence_counter := init_index_int - left_interv; sentence_counter <= init_index_int+right_interv; sentence_counter++ {
				if sentence_counter >= 0 && sentence_counter < sentence_len && sentence_counter != init_index_int {
					for _, word := range words_list_int[sub_verification] {
						if check_word(word, sentence_counter) {
							break
						}
					}
				}
				if word_found != NONE && !return_last_match {
					break
				}
			}
		}

		//log.Println("2---")

		// If it's to stop the verification at the first word not found (no match), stop the verification and put on the
		// return slice, every word as NONE and with an error index.
		if stop_first_not_found && NONE == word_found {
			ret_var = nil
			for counter := 0; counter < len(words_list_int); counter++ {
				ret_var = append(ret_var, []string{NONE, WVD_ERR_2})
			}

			return ret_var
		}

		//log.Println("2.5---")

		// Check if the command seems to be repeated in the 'sentence' (for example, could be "set ah... set 2 alarms")
		// by verifying if any original word is repeated between the previous and current found words - if we're not on
		// the 1st sub-verification, because if we are, we can't check that (no previously found word).
		// - If it's not between both, all is alright with the repetition because it's not a case like said example
		//   (could be "turn on wifi turn on airplane mode" - both "turn"s are very near each other, but still, the
		//   verification will continue, because the "turn" is not between any consecutive found words, so that should
		//   mean it's an actual command).
		// - If the repeated original word is between *any* consecutive 'word_found's, then it might mean it's a case
		//   like the above and the verification will stop right away.
		if !ignore_repets_cmds && sub_verification != 0 {
			// No idea which index is the highest one (depends on the intervals and if the left one is > 0), so put the
			// function checking that by itself.
			var a, b int = index_previous_word_found, index_word_found
			var highest_index, lowest_index int = a, b
			if b > a {
				highest_index = b
				lowest_index = a
			}
			for sentence_counter := lowest_index + 1; sentence_counter <= highest_index-1; sentence_counter++ {
				// The last part ensures the 'sentence' word being checked is not the original one.
				if sentence_counter >= 0 && sentence_counter < sentence_len && sentence_counter != sentence_index {
					for _, original_word := range original_words {
						if sentence[sentence_counter] == original_word {
							// An original word is between 2 consecutive found words, so command repetition detected.
							ret_var = nil
							for counter := 0; counter < len(words_list_int); counter++ {
								ret_var = append(ret_var, []string{NONE, WVD_ERR_3})
							}

							return ret_var
						}
					}
				}
			}
		}

		//log.Println("3---")

		// Here will be checked if a main word is repeated inside the interval "index_word_found + right_interv" and
		// "index_word_found - left_interv". If it is, exit the function.
		if !ignore_repets_main_words {
			for sentence_counter := index_word_found - left_interv + 1; sentence_counter <= index_word_found+right_interv-1; sentence_counter++ {
				for _, word := range main_words {
					// The last part ensures the word being checked is not the found one.
					if sentence_counter >= 0 && sentence_counter < sentence_len && sentence_counter != index_word_found {
						if sentence[sentence_counter] == word {
							ret_var = nil
							for counter := 0; counter < len(words_list_int); counter++ {
								ret_var = append(ret_var, []string{NONE, WVD_ERR_1})
							}

							return ret_var
						}
					}
				}
			}
		}

		//log.Println("4---")

		// Put in the return slice an slice with the 'word_found' and its index.
		ret_var = append(ret_var, []string{word_found, strconv.Itoa(index_word_found)})

		//log.Println(ret_var)
		//log.Println("5---")

		// In case it's to continue searching with one of the words slices in the 'words_list_int', add it to the
		// 'words_list_int' to simulate it having already one more to continue the search - this in case there is no
		// more slices in the 'words_list_int' for next sub-verifications.
		if 1 == len(words_list_int[sub_verification:]) && continue_with_words_slice_number != -1 {
			words_list_int = append(words_list_int, words_list_int[continue_with_words_slice_number])
			words_list_length++ // Also increment the length of the list
		}

		//log.Println("6---")

		if len(exclude_found_word) > 0 {
			// If it's to exclude from the 'words_list_int' the word found in this sub-verification, check to see if
			// it's the ALL_SUB_VERIFS_INT command, or if it's for a specific sub-verification.
			var exclude_word_found_now bool = false
			if ALL_SUB_VERIFS_INT == exclude_found_word[0] {
				exclude_word_found_now = true
			} else {
				for _, number := range exclude_found_word {
					if number == sub_verification {
						exclude_word_found_now = true

						break
					}
				}
			}
			//log.Println("7---")
			//log.Println(words_list_int[sub_verification+1:])
			// Here, actually exclude the word from the next sub-verifications.
			if exclude_word_found_now && len(words_list_int[sub_verification:]) > 1 {
				//log.Println("7.1---")
				//log.Println(words_list_int[sub_verification:])
				//log.Println(words_list_int[sub_verification+1:])
				for counter, words_slice := range words_list_int[sub_verification+1:] {
					for counter1, word := range words_slice {
						if word == word_found {
							GlobalUtilsInt_APU.DelElemInSlice(&words_list_int[sub_verification+1+counter], counter1)

							break
						}
					}
				}
			}
		}

		//log.Println(words_list_int[sub_verification:])

		//log.Println("8---")

		// The function works by calling itself (recursively) while there are slices of words to check in the
		// 'words_list'. It calls itself with a list with one less slice than the previous call ("[1:]") - it removes
		// the first slice of the list in each recursive call.

		// The if statement below checks if there is more than 1 slice left in the list. If there is, that means there
		// is still at least one more sub-verification to make, and hence, continue with the recursion.
		if len(words_list_int[sub_verification:]) > 1 {
			//log.Println("9---")
			var init_index_next_sub_verif string = ""

			// Here it will check if the next index was specified in the function's parameters.
			if next_index, ok := init_indexes_sub_verifs[strconv.Itoa(sub_verification+1)]; ok {
				init_index_next_sub_verif = next_index
			} else {
				// If no index was specified specifically for the current sub-verification, check with ALL_SUB_VERIFS_STR.
				if next_index, ok = init_indexes_sub_verifs[ALL_SUB_VERIFS_STR]; ok {
					init_index_next_sub_verif = next_index
				}
			}
			// Now check if the index is a number or a special command.
			if strings.Contains(init_index_next_sub_verif, INDEX_WORD_FOUND) {
				if INDEX_WORD_FOUND == init_index_next_sub_verif {
					init_index_next_sub_verif = strconv.Itoa(index_word_found)
				} else if strings.Contains(init_index_next_sub_verif, "+") {
					number, _ := strconv.Atoi(strings.Split(init_index_next_sub_verif, "+")[1])
					init_index_next_sub_verif = strconv.Itoa(index_word_found + number)
				} else if strings.Contains(init_index_next_sub_verif, "-") {
					number, _ := strconv.Atoi(strings.Split(init_index_next_sub_verif, "-")[1])
					init_index_next_sub_verif = strconv.Itoa(index_word_found - number)
				}
			} else if strings.Contains(init_index_next_sub_verif, INDEX_DEFAULT) || "" == init_index_next_sub_verif {
				// If no index was specified or the default one was specified, calculate the next initial index by
				// calculating an average and summing 0.5.
				// If it was index 3, it's now 3.5, which is 3 when converted to int. If it was 3.5, it's now 4.
				// It's just a way of increasing the index sometimes. Sometimes not - random. The idea of increasing a
				// bit is to increase the search interval that bit. But not as much as to put the new index, the index
				// of the word found. Like a middle term, but with more weight on the right.
				//init_index_next_sub_verif = strconv.Itoa(int(
				//	(float32(init_index_int) + float32(index_word_found))/float32(2) + float32(0.5)))

				// UPDATE: I've just got it back to the older and first way --> the index of the word found.
				// I believe the way above was because there existed no way of knowing what a "it" means - which is not
				// true anymore.
				init_index_next_sub_verif = strconv.Itoa(index_word_found)
			}

			//log.Println("init_index_next_sub_verif_str -", init_index_next_sub_verif_str)
			//log.Println("words_list_int[sub_verification:] -->", words_list_int[sub_verification:])

			//log.Println("^^^^^^^^^^^^^^^^^^^^^")

			init_index_int, _ = strconv.Atoi(init_index_next_sub_verif)
			main_words = []string{sentence[init_index_int]}
		}

		index_previous_word_found = index_word_found
	}

	//log.Println("________________________")

	return ret_var
}

const A_WORD_IN_LIST string = "3234_ANY_WORD_IN_LIST"

/*
checkResultsWordsVerifFunc checks if the results coming from the wordsVerificationFunction() function are acceptable or not,
depending on the given parameters.

If there are conditions (described below), those will be checked. If not (leave 'conditions_continue' nil or empty),
the functions checks if ALL the words on the results are on the 'words_list' (NONE as a "detected" word on the results
would exclude that - example for "reboot phone", but "phone" was not found: {{"reboot", 2},{ NONE, 3}} --> this will
make this function return false, as not all the results' words are on the 'words_list').

If the results of the verification return all words as NONE, this function will not check anything else and will return
false immediately at the beginning before doing anything else.

Naming convention:
	- list: conditions_continue;
	- sub-list: condition
	- sub-sub-list: sub-condition
	- string(s) of the sub-sub-list: string(s) or word(s) of the sub-condition
	- -----
	- list: conditions_not_continue
	- sub-list: main condition
	- sub-sub-list: condition
	- sub-sub-sub-list: sub-condition
	- string(s) of the sub-sub-sub-list: string(s) or word(s) of the sub-condition

----------------

Format of the 'conditions_continue':
	var conditions_continue [][][]string = [][][]string{
		{{"turn"}, {"on","off"}, {"flashlight","lantern"}},
		{{"flashlight","lantern"}, {"on","off"}, {}},
	}
Here, {"flashlight","lantern"} are the main words. The rest is what comes in the results of the verification.

On the example above, for the 1st sub-condition, on the results of the verification must be, for example,
	{{"on",index},{"[doesn't matter here]", index}},
and "flashlight" or "lantern" must have been the main word. On the 2nd index of the example above, the main word must
have been "turn", then can be "on" or "off", and then can be "flashlight" or "lantern" (all parts of the results matter
in this case).

To indicate that any combination of words is allowed (and stop the function to check if all the found words are in the
list of words provided like it does where there are no conditions of continuation), an empty condition may be used: {}.
This can be paired with a no continuation condition if wanted.

Another option is to put { A_WORD_IN_LIST} various times (as many as the length of 'words_list' - or more, but those in
excess will be ignored) to do exactly what's described below of putting an empty slice --> except this way you can also
use a no continuation condition, unlike using an empty continuation condition slice. Btw, if put on the 1st index of the
condition, it will be ignored safely (the main word is on the list of course, or the command wouldn't have been detected
in the first place).

To allow ALL possible combinations of ANY words (the 'conditions_not_continue' will be IGNORED), put the
'conditions_continue' empty on the function call and that will make this function check if all the words on the results
are in the 'words_list' ( NONE would exclude that on the results, for example, as said in the beginning).
It's not possible to put only conditions of no continuation currently. Only if the corresponding continuation ones are
put too.

----------------

Format of the 'conditions_not_continue':
	var conditions_not_continue [][][][]string = [][][][]string{
		{},
		{  {{}, {"off"}, {"on"}},  {{}, {"on"}, {"off"}}  },
		{},
	}
	conditions_continue := [][][]string{
		{{}, {"whatever"}, {"doesn't", "matter"}},
		{{}, {"on", "off", "speaker", "speakers"}, {"on", "off", "speaker", "speakers"}},
		{{}, {"whatever"}, {"doesn't", "matter"}},
	}

In the case above, IF it's detected a match in the 'conditions_continue' with the SECOND condition, the SECOND main
condition of the 'conditions_not_continue' will be checked to see if there is a match in one of its conditions.

The example above says "on" and "off" can't be in both outputs of the verification. It can't return, for example,
{{"on",index}, {"off",index}} - which it might sometimes.

It must have the same number of main conditions as the number of conditions on 'conditions_continue', and each condition
must have the same number of sub-conditions as the number of sub-conditions on the conditions of 'conditions_continue'.
It can have any number of conditions.

-------- Sum up --------

The verification only stops if it:

- iterates ALL the continuation conditions and *none* is verified, independently of the non-continuation ones, OR if it
has non-continuation conditions invalidating *all* the continuation ones (returns false);

- iterates the continuation conditions until it finds one that is verified, and in that one, iterates *all* the no
continuation conditions but *none* is verified (returns true).

-----------------------------------------------------------

> Params:

- words_list – same as in wordsVerificationFunction() (the same that was sent to that function)

- main_word – the word that triggered the command detection (in "turn the flashlight on", would be "turn")

- results_wordsVerificationDADi – the output of wordsVerificationFunction()

- conditions_continue – a 3D slice with a set of conditions on which the results are acceptable. Empty to allow
any combination of words on the results to say they are acceptable (as long as the words on the results are inside the
'words_list' set). See the format above.

- conditions_not_continue – a 4D slice with a set of conditions on which each condition of continuation is not
acceptable. See the format above.


> Returns:

- true if the results are acceptable for the given parameters, false otherwise
*/
func checkResultsWordsVerifFunc(words_list [][]string, main_word string, results_wordsVerificationDADi [][]string,
	conditions_continue [][][]string, conditions_not_continue [][][][]string) bool {
	//log.Println("-------------------------------")
	//log.Println(words_list)
	//log.Println(main_word)
	//log.Println(results_wordsVerificationDADi)
	//log.Println(conditions_continue)
	//log.Println(conditions_not_continue)
	//log.Println("-------------")

	var all_none bool = true
	for _, result := range results_wordsVerificationDADi {
		if result[0] != NONE {
			all_none = false
		}
	}
	if all_none {
		return false
	}

	if len(conditions_continue) > 0 {
		// If there are conditions of continuation, check them and their corresponding conditions of no continuation.

		// The variable below is a copy of the results of the verification function with the main word added in the
		// index 0, so the conditions can be checked by their index (0 corresponding to the main word and 1 to the first
		// word of the results).
		var modified_results_verifDADi [][]string = GlobalUtilsInt_APU.CopyOuterSlice(results_wordsVerificationDADi).([][]string) // Will only add a new slice to the outer slice, so no problem in using CopyOuterSlice().
		GlobalUtilsInt_APU.AddElemSlice(&modified_results_verifDADi, []string{main_word, "-1"}, 0)

		var modified_results_verifDADi_len int = len(modified_results_verifDADi) // Optimization

		/*
			checkConditionNotContinueMatch checks if any condition on the main condition of the conditions of no
			continuation corresponding to a given condition of continuation has a match.

			Which means, if there is a match, it's NOT to continue with that continuation condition because it can't be
			applied --> go check the next condition of continuation, and the corresponding conditions of no continuation.

			-----------------------------------------------------------

			> Params:

			- cond_index – the index of the condition of continuation currently in analysis


			> Returns:

			- true if there is a match in a condition of no continuation, false if there is no match
		*/
		checkConditionNotContinueMatch := func(cont_cond_index int) bool {
			if len(conditions_not_continue) > 0 {
				// Main note: this function is almost copy-paste of what's below it. So to understand it, read below it first.
				for _, condition := range conditions_not_continue[cont_cond_index] {

					var number_sub_conds_must_match int = 0
					var number_sub_conds_matched int = 0
					for sub_cond_index, sub_cond := range condition {

						var any_word_match = false
						for _, sub_cond_word := range sub_cond {
							if modified_results_verifDADi[sub_cond_index][0] == sub_cond_word {
								any_word_match = true
								break
							}
						}

						if len(sub_cond) != 0 {
							// If there's nothing on the sub-condition, nothing is done - that result is ignored. On the
							// other hand, if it's not empty, then it's to consider that sub-condition. So it's added to the
							// total number of sub-conditions that must match.
							number_sub_conds_must_match++
							if any_word_match {
								// If any word on the sub-condition matched, then the sub-condition also matched. So
								// increment the number of sub-conditions matched.
								number_sub_conds_matched++
							}
						}
					}
					if number_sub_conds_matched == number_sub_conds_must_match {
						// In the end, if all sub-conditions inside the condition corresponding to the continuation
						// condition in analysis match (the number of must match is equal to the number of matches), then it
						// means there was a match. So return true.
						return true
					}
				}
			}

			// If no condition had a match or there were no conditions at all (no match also then...), return false.
			return false
		}

		for cond_index, condition := range conditions_continue {
			//log.Println("+++++++++++++++")

			// Start true and AND it with internal match check variable. This true also helps with putting only {} at
			// the end of the list of conditions to allow any other combination - if the length of the condition is 0,
			// no sub-conditions will be checked (there are none) and the value of true will not be changed.
			var all_sub_conds_match = true
			for sub_cond_index, sub_cond := range condition {
				if sub_cond_index >= modified_results_verifDADi_len {
					// One can put as many sub-conditions as wanted - the ones in excess will not be verified and will
					// be ignored.
					// Useful to put {A_WORD_IN_LIST} many times without worrying if it's enough or not.
					break
				}
				//log.Println("++++++")
				//log.Println(sub_cond_index)
				var any_word_match = false // To check if any word in the sub-condition matches the result's word.
				if 0 == len(sub_cond) {
					// If there is nothing in the sub-condition, it doesn't matter and allow anything, including NONE.
					any_word_match = true
				} else {
					if A_WORD_IN_LIST == sub_cond[0] {
						if 0 == sub_cond_index {
							// No need to check if it's a word on the list with the main word - of course it's on the list,
							// or the command wouldn't have been detected in the first place.
							any_word_match = true
						} else {
							for _, sub_cond_word := range words_list[sub_cond_index-1] {
								// (-1 above because 1 on sub_cond is 0 in words_list)
								if modified_results_verifDADi[sub_cond_index][0] == sub_cond_word {
									any_word_match = true
								}
							}
						}
					} else {
						// If there are things in the sub-condition, check if at any word matches the one on the results.
						for _, sub_cond_word := range sub_cond {
							//log.Println(sub_cond_word)
							//log.Println(sub_cond_number)
							//log.Println(modified_results_verifDADi[sub_cond_number][0])
							//log.Println(modified_results_verifDADi[sub_cond_number][0] == sub_cond_word)
							if modified_results_verifDADi[sub_cond_index][0] == sub_cond_word {
								any_word_match = true
								break
							}
						}
					}
				}
				all_sub_conds_match = all_sub_conds_match && any_word_match
				if !all_sub_conds_match {
					break
				}
			}
			if all_sub_conds_match {
				//log.Println("RRRRRRRRRRRRR")
				// If the condition of continuation matched, the corresponding conditions of no continuation are checked.
				// If any matches, the condition of no continuation cannot be applied and the next one is checked. If
				// none matches, there was a complete match in the condition of continuation - so the results are
				// acceptable.
				if !checkConditionNotContinueMatch(cond_index) {
					return true
				}
			}
		}

		log.Println("***************")
		// No condition of continuation matched completely (meaning the condition of continuation and no match in the
		// corresponding conditions of no continuation), so return false - the results are not acceptable.
		return false

	} else {
		// If there are no conditions to continue, just check if all the words in the results are in the 'words_list'.
		// For example, if one of the results has NONE on it, this will return false - will mean the function didn't
		// find all possible words and, therefore, the results are not acceptable.

		var all_match bool = true
		for result_index, result := range results_wordsVerificationDADi {
			var match_found bool = false
			for _, word := range words_list[result_index] {
				if word == result[0] {
					match_found = true

					break
				}
			}
			all_match = all_match && match_found
			if !all_match {
				break
			}
		}

		return all_match
	}
}
