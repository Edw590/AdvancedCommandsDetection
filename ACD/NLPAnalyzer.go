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
	"strings"

	"github.com/jdkato/prose/v2"

	"Utils"
)

//////////////////////////
// Global variables (the "nlp_" prefix is because the variables are global for the *package*)

var nlp_sentence_counter int
var nlp_token_counter int

// For replaceIts()

var nlp_last_was_an_it bool
var nlp_non_name_passed_since_last_name bool
var nlp_last_name_found []string
var nlp_last_it string
var prev_sentence_it string

// For replaceAnds()

var nlp_last_was_an_and bool
var nlp_non_allowed_tag_passed_since_last_allowed bool
var nlp_verbs_passed int
var nlp_second_last_to_last_non_allowed_tag []string
var nlp_last_and string
var prev_sentence_and string

// This is a list of words and tags to always be applied to the NLP tagging output. Sometimes it sees "mode" as a verb,
// for example. That's wtf, I think. In any case, for the purposes of the assistant at the moment, "mode" is always a
// name for the required commands (even though I think "mode" is always a name - like "do" is always a verb and so on).
// "turn" may be a name too ("number of turns" or "it's my turn"), but again, for the purposes of the assistant, "turn"
// is always a verb, so assume all "turn" instances on the 'sentence' are verbs.
// This applies for Prose in its current version (writing this on 2021-11-20).
// Note: I took the tags below from an online P.O.S. tagger (https://parts-of-speech.info) in sentences that would make
// it obvious what each word is (name, verb, adjective...).
// Note 2: I've just added all the verbs present cmds_types_keywords array, so that they are always recognized as verbs.
var nlp_static_word_tags map[string]string = map[string]string{
	//////////////////////////
	// Generic words
	"do":          "VBP",
	"mode":        "NN",
	"bluetooth":   "NN",
	"it":          "NN", // To recognize "it" as a name, not a pronoun, because it will be replaced by the name of the
	// last name, and so that would be included in the middle of the "and" meaning (not supposed to be names there).
	"please":      "RB",
	"fast":        "JJ",

	//////////////////////////
	// Commands
	"turn":   "VB",
	"get":   "VB",
	"switch":   "VB",
	"put":   "VB",

	"record": "VB", // For the commands, it's always a verb (recognized as name in "record the audio").

	"reboot": "VB", // For the commands, it's always a verb (recognized as name in "shut down the phone and reboot it").
	"restart":   "VB",

	"shut":   "VB",

	"start":  "VB", // For the commands, it's always a verb (recognized as name in "start recording it").
	"begin":   "VB",
	"initialize":   "VB",
	"commence":   "VB",

	"stop":   "VB", // For the commands, it's always a verb (recognized as name in "stop recording it").
	"finish":   "VB",
	"cease":   "VB",
	"conclude":   "VB",
	"terminate":   "VB",

	"answer":   "VB",
	"reply":   "VB",
	"respond":   "VB",
	"acknowledge":   "VB",
}

//////////////////////////

/*
nlpAnalyzer analyzes a 'sentence' with Natural Language Processing help to better understand its meaning.

Currently, it replaces all instances of "it"s and "and"s on the 'sentence' by their meaning, allowing the verification
function to better do its job.

WARNING: the "and" verification assumes the parameter 'ignore_repets_cmds' is set to false when calling
wordsVerificationFunction() - if it's set to true, please note this may not work as expected. Why? Refer to
replaceAnds().

-----CONSTANTS-----

  - WHATS_IT – this may be in the 'sentence' after the return of this function. It means an "it" was detected with no
    meaning.
  - WHATS_AND – same as WHATS_IT but for an "and".

-----CONSTANTS-----

-----------------------------------------------------------

– Params:
  - sentence – a pointer to the header of the created 'sentence' slice on the function mainInternal()
  - sentence_str – the string sent to mainInternal() but with the modifications done on sentenceNLPPreparation()

– Returns:
  - the updated 'sentence_str' according to the also updated 'sentence'
*/
func nlpAnalyzer(sentence *[]string, sentence_str string, nlp_meanings []string) []string {
	//log.Println("-----")

	resetVariables()

	// todo See how this file works and finish this. You were going to pass to the function an it_and variable saying
	//  what was "it" and what "and" was referring to, like "turn wifi on", and seconds later, "turn it off" --> "wifi"
	//  is the "it"
	//nlp_last_name_found = append(nlp_last_name_found, it_and)

	//log.Println("-----------------------------")
	//log.Println(sentence_str)

	// Create a new document with the default configuration
	nlp_doc, _ := prose.NewDocument(sentence_str)

	var tokens []prose.Token = nlp_doc.Tokens()

	// len(tokens) on the condition so it's checked every time what is the length and there's no need to decrement it or
	// increment it manually depending on the modifications to the tokens' slice.
	for counter, token_text := 0, ""; counter < len(tokens); counter++ {
		token_text = tokens[counter].Text
		if "'s" == token_text {
			// Remove all possessive cases - not helpful yet, so can remove them with no problems, as they just get
			// things harder because "what's" goes to "what" and "'s" on the NLP tokens detector, but not on the
			// 'sentence', so the loop doesn't work. Removing the "'s" from the tokens list might help. It doesn't
			// matter for the time being (no command requires any use of the possessive case or similar).
			// Note that nothing will be touched on the 'sentence'. Only on the tokens. This way the 'sentence' is still
			// intact. It's just removed from the tokens to synchronize them with the 'sentence' and nothing else. If
			// possible nothing would be removed... (don't like the idea too much, but no better ideas at the moment).

			Utils.DelElemSLICES(&tokens, counter)
			// Decrement the counter, so we go to the previous, to then be incremented by the loop and go to the
			// current one, which is now the old next one.
			counter--
		} else if new_tag, ok := nlp_static_word_tags[token_text]; ok {
			tokens[counter].Tag = new_tag
		}
	}

	//log.Println(*sentence)

	// Print all the tokens
	//for _, tok := range tokens {
	//	log.Println(tok)
	//}

	prev_sentence_it = nlp_meanings[0]
	prev_sentence_and = nlp_meanings[1]

	if "" != prev_sentence_it {
		nlp_last_name_found = strings.Split(prev_sentence_it, " ")
	}
	if "" != prev_sentence_and {
		nlp_second_last_to_last_non_allowed_tag = strings.Split(prev_sentence_and, " ")
	}

	//log.Println("prev_sentence_it:", prev_sentence_it)
	//log.Println("prev_sentence_and:", prev_sentence_and)

	// The sentence_counter was already set before, so no setting it here on the loop (empty part).
	for ; nlp_sentence_counter < len(*sentence); nlp_sentence_counter, nlp_token_counter = nlp_sentence_counter+1, nlp_token_counter+1 {
		replaceIts(sentence, &tokens)
		replaceAnds(sentence, &tokens)
	}

	//log.Println("---")
	//log.Println(*sentence)
	//log.Println("-----")

	return []string{nlp_last_it, nlp_last_and}
}

const WHATS_IT string = ";6;"
const WHATS_AND string = ";7;"

/*
replaceIts replaces all "it"s that it finds on the sentence by their meaning, based on the names that appear before
them.

For example, "The Wi-Fi, turn it on please." --> the "it" refers to "Wi-Fi" - this function replaces "it" by "Wi-Fi" and
deletes "Wi-Fi" from the sentence.
EDIT: not anymore deleting the word. Hopefully that's not bad, because with "shut down the phone and reboot it", "phone"
must not be removed.

-----CONSTANTS-----

  - WHATS_IT – same as in nlpAnalyzer()
  - WHATS_AND – same as in nlpAnalyzer()

-----CONSTANTS-----

-----------------------------------------------------------

– Params:
  - sentence – same as in nlpAnalyzer()
  - tokens – a list with the tokens of all the 'sentence' words

– Returns:
  - nothing
*/
func replaceIts(sentence *[]string, tokens *[]prose.Token) {
	// Leave the 2 parameters because both exist on CmdsDetector(), so why create one of them again and not use the one
	// that already exists? Optimization.

	// Way of checking if the function is working well below (seems to be a good check - has all cases on it, I think):
	// "turn it on turn on wifi and the airplane mode get it it on no don't turn it on turn off airplane mode and the " +
	//	"wifi turn it on"

	// Leave len(*sentence) there and don't assign a variable to it. That way it keeps checking the length, and it's not
	// needed to increase or decrease based on changes on the 'sentence' (it will calculate the length every time).
	if "it" == (*sentence)[nlp_sentence_counter] {
		//log.Println("-------")
		//log.Println(nlp_sentence_counter)
		//log.Println(nlp_last_was_an_it)
		if nlp_last_was_an_it {
			// If the last word was an "it", it means there are repeated ones - delete all the repeated ones and use
			// only the first one. If they were not deleted, too many words would be in between the command words -->
			// no detection.
			Utils.DelElemSLICES(sentence, nlp_sentence_counter)
			nlp_sentence_counter-- // And since an element was deleted, decrement the sentence_counter.
			//log.Println("*****")
			//log.Println(*sentence)

			return // And go to the next word on the sentence.
		}
		nlp_last_was_an_it = true
		if len(nlp_last_name_found) > 0 {
			//log.Println((*sentence)[nlp_sentence_counter])
			//log.Println(nlp_last_name_found[0][0])
			(*sentence)[nlp_sentence_counter] = nlp_last_name_found[0]
			if len(nlp_last_name_found) > 1 {
				for name_index, name := range nlp_last_name_found[1:] {
					// +1 below because we're starting from [1:].
					Utils.AddElemSLICES(sentence, name, nlp_sentence_counter+name_index+1)
				}
			}
			// Increment the sentence_counter. -1 because it needs to stay at the element before the next. The next
			// one will be taken care of by the sentence_counter++ line. This is done because we just added elements
			// to the sentence. So the counter must be incremented to go to the next old element, not go through the
			// newly added ones (that would also not be in accordance with the tokens iteration).
			nlp_sentence_counter += len(nlp_last_name_found) - 1

			//log.Println(*sentence)
		} else {
			//log.Println("RRRRRRRRRRRRRRRRRRRRRRRRRRRRR1")
			var whats_it = WHATS_IT
			if "" != prev_sentence_it {
				whats_it = prev_sentence_it
				prev_sentence_it = ""
			}

			whats_it_list := strings.Split(whats_it, " ")
			(*sentence)[nlp_sentence_counter] = whats_it_list[0]
			if len(whats_it_list) > 1 {
				for name_index, name := range whats_it_list[1:] {
					Utils.AddElemSLICES(sentence, name, nlp_sentence_counter+name_index+1)
				}
			}
			nlp_sentence_counter += len(whats_it_list) - 1
		}
	} else {
		nlp_last_was_an_it = false
		if strings.HasPrefix((*tokens)[nlp_token_counter].Tag, "N") { // Which means it's a name
			if nlp_non_name_passed_since_last_name {
				// If a non-name passed since the last name, first empty the slice before appending - because on the
				// slice are only consecutive names (like "airplane mode" - 2 names, that are put on the slice).
				nlp_last_name_found = nil
				// Don't reset the name until a new name passes by. That way, this, for example, works: "the wifi
				// turn it on now turn it off".
			}
			nlp_non_name_passed_since_last_name = false

			nlp_last_name_found = append(nlp_last_name_found, (*sentence)[nlp_sentence_counter])
		} else {
			if nlp_last_name_found != nil {
				// If a word that is a not a name passed since the last consecutive name, signal it to know that the
				// next time a name is detected, it's not just to add it to the slice - first empty the slice.
				nlp_non_name_passed_since_last_name = true
			}
		}
	}

	nlp_last_it = strings.Join(nlp_last_name_found, " ")
}

/*
replaceAnds replaces all "and"s that it finds on the sentence by the action they refer to, or simply deletes them in
case they don't mean anything.

For example, "Turn on the Wi-Fi and the airplane mode." --> the "and" refers to "turn on" - this function replaces "and"
by "Turn on", being the final sentence "Turn on the Wi-Fi Turn on the airplane mode".

Another example, "Shut down the phone and reboot it." --> the "and" means nothing here, because there's a verb (new
action just in front of it). So the final sentence is just "Shut down the phone reboot it.".

WARNING: this function assumes the parameter 'ignore_repets_cmds' is set to false when calling wordsVerificationFunction() -
if it's set to true, please note this may not work as expected. This is because the function replaces *any* "and", even
those that don't refer to actions, like this example: "turn on the wifi and turn off the airplane mode", which will
become "turn on the wifi turn on turn off the airplane mode" --> ignoring command repetitions and processing them
anyways will make this not work as expected. If the repetition is ignored, the sentence will be like "turn on the wifi
turn off the airplane mode".

-----------------------------------------------------------

– Params:
  - sentence – same as in nlpAnalyzer()
  - tokens – a list with the tokens of all the 'sentence' words

– Returns:
  - nothing
*/
func replaceAnds(sentence *[]string, tokens *[]prose.Token) {
	// Leave the 2 parameters because both exist on CmdsDetector(), so why create one of them again and not use the one
	// that already exists? Optimization.

	// Way of checking if the function is working well below (seems a good check - has all cases on it, I think):
	// "turn on wifi and and the airplane mode and the flashlight"
	// When the implementation is changed, swap the places of "on" and "wifi" and check if it still works.

	if "and" == (*sentence)[nlp_sentence_counter] {
		if nlp_last_was_an_and ||
			((len(*tokens) > nlp_sentence_counter+1) && strings.HasPrefix((*tokens)[nlp_sentence_counter+1].Tag, "VB")) ||
			((len(*tokens) > nlp_sentence_counter+2) && strings.HasPrefix((*tokens)[nlp_sentence_counter+2].Tag, "VB")) {
			// The same as for the "it" case.
			// Except here also delete if the next word is a verb: "shut down the phone and reboot it". Here, "and" is
			// not supposed to be replaced by "shut down". Instead, its presence is irrelevant. So just remove it,
			// because the next word is a verb (means after it is said the actual action and not to repeat the previous
			// one).
			// Also with +2 because "and then reboot". The verb is the 2nd word here, not the 1st.
			Utils.DelElemSLICES(sentence, nlp_sentence_counter)
			nlp_sentence_counter--
			//log.Println("*****")
			//log.Println(*sentence)

			return
		}
		nlp_last_was_an_and = true

		//log.Println("------")
		//log.Println(nlp_second_last_to_last_non_allowed_tag)

		if len(nlp_second_last_to_last_non_allowed_tag) > 0 {
			//log.Println(nlp_sentence_counter)
			(*sentence)[nlp_sentence_counter] = nlp_second_last_to_last_non_allowed_tag[0]
			if len(nlp_second_last_to_last_non_allowed_tag) > 1 {
				for non_name_index, non_name := range nlp_second_last_to_last_non_allowed_tag[1:] {
					// +1 below because we're starting from [1:].
					Utils.AddElemSLICES(sentence, non_name, nlp_sentence_counter+non_name_index+1)
				}
			}
			nlp_sentence_counter += len(nlp_second_last_to_last_non_allowed_tag) - 1

			// This -1 makes it so that as it found an "and", it will stop adding words to the list but will not discard
			// or erase them.
			nlp_verbs_passed = -1

			// No deletions here as with "it". What was before the "and" remains there to still have impact (unlike with
			// "it" in which the names are just in the wrong place for the verification function to work properly).
		} else {
			//log.Println("RRRRRRRRRRRRRRRRRRRRRRRRRRRRR2")
			var whats_and = WHATS_AND
			if "" != prev_sentence_and {
				whats_and = prev_sentence_and
				prev_sentence_and = ""
			}

			whats_and_list := strings.Split(whats_and, " ")
			(*sentence)[nlp_sentence_counter] = whats_and_list[0]
			if len(whats_and_list) > 1 {
				for non_name_index, non_name := range whats_and_list[1:] {
					Utils.AddElemSLICES(sentence, non_name, nlp_sentence_counter+non_name_index+1)
				}
			}
			nlp_sentence_counter += len(whats_and_list) - 1
		}
	} else {
		nlp_last_was_an_and = false
		var current_tag string = (*tokens)[nlp_token_counter].Tag
		if !strings.HasPrefix(current_tag, "N") {
			if strings.HasPrefix(current_tag, "VB") {
				if nlp_verbs_passed < 0 {
					nlp_verbs_passed = 0
				}
				nlp_verbs_passed++
				if nlp_verbs_passed == 1 {
					// Reset the slice if a new verb is found. Useful for the first time in which a verb is detected
					// and a slice had been passed as previous command information.
					nlp_second_last_to_last_non_allowed_tag = nil
				}
			}
			if nlp_non_allowed_tag_passed_since_last_allowed || nlp_verbs_passed > 1 {
				// If a non-allowed tag passed since the last allowed one, empty the slice before appending - because on
				// the slice are only consecutive allowed tags' words (like "turn on" - 2 allowed tags' words, that are
				// put on the slice).
				nlp_second_last_to_last_non_allowed_tag = nil
				// Don't reset the slice until a new allowed tags' word passes by. That way, this, for example, works:
				// "turn on the wifi and the airplane mode and the flashlight".
				if nlp_verbs_passed > 1 {
					nlp_verbs_passed = 1 // Verb just passed, so set to 1
				}
			}
			nlp_non_allowed_tag_passed_since_last_allowed = false

			if nlp_verbs_passed == 1 {
				if strings.HasPrefix(current_tag, "VB") {
					var adjectives_list []string = nil
					for i := nlp_sentence_counter - 1; i >= 0; i-- {
						if strings.HasPrefix((*tokens)[i].Tag, "J") {
							// Add all adjectives right behind the current word in case it's a verb.
							adjectives_list = append(adjectives_list, (*sentence)[i])
						} else {
							// Stop when a non-adjective is found (must be consecutive adjectives).
							break
						}
					}
					for i := len(adjectives_list)-1; i >= 0; i-- {
						// Add all adjectives in the order they were inserted in the sentence.
						nlp_second_last_to_last_non_allowed_tag = append(nlp_second_last_to_last_non_allowed_tag,
							adjectives_list[i])
					}
				}
				nlp_second_last_to_last_non_allowed_tag = append(nlp_second_last_to_last_non_allowed_tag,
					(*sentence)[nlp_sentence_counter])
			}
		}
	}

	nlp_last_and = strings.Join(nlp_second_last_to_last_non_allowed_tag, " ")
}

/*
resetVariables resets the global variables used in this file every time it's called.
*/
func resetVariables() {
	nlp_sentence_counter = 0
	nlp_token_counter = 0

	// For replaceIts()
	nlp_last_was_an_it = false
	nlp_non_name_passed_since_last_name = false
	nlp_last_name_found = nil
	nlp_last_it = ""
	prev_sentence_it = ""

	// For replaceAnds()
	nlp_last_was_an_and = false
	nlp_non_allowed_tag_passed_since_last_allowed = false
	nlp_verbs_passed = 0
	nlp_second_last_to_last_non_allowed_tag = nil
	nlp_last_and = ""
	prev_sentence_and = ""
}
