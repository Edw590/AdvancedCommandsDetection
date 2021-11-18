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
	"github.com/jdkato/prose/v2"
	"log"
	"strconv"
	"strings"
)

const WHATS_IT string = "3234_WHATS_IT"

/*
ReplaceIts replaces all "it"s that it finds on the sentence by their meaning, based on the names that appear before
them.

For example, "The Wi-Fi, turn it on please." --> the "it" refers to "Wi-Fi" - this function replaces "it" by "Wi-Fi" and
deletes "Wi-Fi" from the sentence.

-----CONSTANTS-----

- WHATS_IT – this may be in the 'sentence' after the return of this function. It means an "it" was detected with no
meaning. Could be a good idea for the assistant to return a warning about an "it" on the sentence with no meaning.

-----CONSTANTS-----

-----------------------------------------------------------

> Params:

- sentence – a pointer to the header of the created 'sentence' slice on the function CmdsDetectorInternal()

- sentence_str – same as in CmdsDetector()


> Returns:

- nothing
*/
func ReplaceIts(sentence *[]string, sentence_str string) {
	// Leave the 2 parameters because both exist on CmdsDetector(), so why create one of them again and not use the one
	// that already exists? Optimization.

	// Way of checking if the function is working well below (seems a good check - has all cases on it, I think).
	//var sentence_str string = "turn it on turn on wifi and the airplane mode get it it on no don't turn it on turn " +
	//	"off airplane mode and the wifi turn it on"

	//log.Println("-----")

	// Create a new document with the default configuration
	nlp_doc, _ := prose.NewDocument(sentence_str)

	// Iterate over the doc's tokens:
	for _, tok := range nlp_doc.Tokens() {
		log.Println(tok)
	}
	//log.Println("")

	var last_name_found [][]string = nil
	var non_name_passed_since_last_name bool = false
	var last_was_an_it bool = false
	var tokens []prose.Token = nlp_doc.Tokens()
	var sentence_counter = 0
	for _, token := range tokens {
		if token.Text == "it" {
			if last_was_an_it {
				// If the last word was an "it", it means there are repeated ones - delete all the repeated ones and use
				// only the first one. If they were not deleted, too many words would in between the command words -->
				// no detection.
				APU_GlobalUtilsInt.DelElemInSlice(sentence, sentence_counter)
				sentence_counter-- // And since an element was deleted, decrement the sentence_counter.
				//log.Println("*****")
				//log.Println(*sentence)

				continue // And go to the next word on the sentence.
			}
			last_was_an_it = true
			if len(last_name_found) > 0 {
				(*sentence)[sentence_counter] = last_name_found[0][0]
				if len(last_name_found) > 1 {
					for name_index, name_array := range last_name_found[1:] {
						// +1 below because we're starting from [1:].
						APU_GlobalUtilsInt.AddElemSlice(sentence, name_array[0], sentence_counter+name_index+1)
					}
				}
				// Increment the sentence_counter. -1 because it needs to stay at the element before the next. The next
				// one will be taken care of by the sentence_counter++ line. This is done because we just added elements
				// to the sentence. So the counter must be incremented to go to the next old element, not go through the
				// newly added ones (that would also not be in accordance with the tokens iteration).
				sentence_counter += len(last_name_found) - 1

				//log.Println("+++++")
				//log.Println(*sentence)
				// The original words on the 'sentence' get deleted so the function does not try to detect anything
				// based on them, because the words belong to where they were put (on the "if"'s place) - but they're
				// deleted also to reduce the number of words between the command main words - or no detection will
				// happen.
				for i := len(last_name_found) - 1; i >= 0; i-- {
					word_index, _ := strconv.Atoi(last_name_found[i][1])
					if word_index > 0 { // If the word hasn't been deleted already...
						APU_GlobalUtilsInt.DelElemInSlice(sentence, word_index)
						last_name_found[i][1] = "-1" // This signals the word has been deleted - can't delete it twice.
						sentence_counter--
					}
				}
				//log.Println(*sentence)
			} else {
				(*sentence)[sentence_counter] = WHATS_IT
			}
		} else {
			last_was_an_it = false
			if strings.HasPrefix(token.Tag, "N") { // Which means it's a name
				if non_name_passed_since_last_name {
					// If a non-name passed since the last name, first empty the array before appending - because on the
					// array are only consecutive names (like "airplane mode" - 2 names, that are put on the array).
					last_name_found = nil
					// Don't reset the name until a new name passes by. That way, this, for example, works: "the wifi
					// turn it on now turn it off".
				}
				non_name_passed_since_last_name = false

				last_name_found = append(last_name_found, []string{token.Text, strconv.Itoa(sentence_counter)})
			} else {
				if last_name_found != nil {
					// If a word that is a not a name passed since the last consecutive name, signal it to know that the
					// next time a name is detected, it's not just to add it to the array - first empty the array.
					non_name_passed_since_last_name = true
				}
			}
		}

		sentence_counter++
	}

	//log.Println("---")
	log.Println(*sentence)
	//log.Println("-----")
}
