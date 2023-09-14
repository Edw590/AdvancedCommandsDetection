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
)

/*
sentenceCorrection abbreviates a sentence to reduce words on it, and corrects some mispronunciations that speech
recognizers may leave.

The idea is for the verification function to be able to have more probabilities of succeeding in a command detection,
also to be easier to list the commands (less word variations), and finally to check words in a normalized way (this way,
one knows "do not" doesn't happen and doesn't need to worry about checking that, only "don't", which is easier - only
one word, and in a "foreach word" loop is much easier).

This function should be the first thing to be called on mainInternal() to normalize the sentence for the entire
library.

For example, it finds all occurrences of "what is" and replaces them by "what's". It also corrects "whats" to "what's",
in case the speech recognizers put them like that.

To see all it does, take a look at the function (easy reading).

-----------------------------------------------------------

– Params:
  - sentence_str – same as in Main()

– Returns:
  - a string with everything replaced/corrected on it
*/
func sentenceCorrection(sentence_str string, sentence *[]string, before_nlp_analyzer bool) string {
	// Hopefully the words to replace here cannot be joint with others, because this doesn't check what's before the
	// first letter of each word... (could be part of the word before or after). Maybe it doesn't happen, to simplify.

	// NOTE ABOUT THIS FUNCTION VS THE FUNCTION BELOW
	// This one normalizes everything for the entire library. This way one doesn't need to worry, for example, about
	// checking "don't" or "do not" --> ("do" == sentence[counter] && "not" == sentence[counter+1]) - complication.
	// The one below makes some adjustments for the NLP analyzer to better understand and correct the sentence (have
	// "what" and "'s" on different tags is unhelpful when the sentence doesn't divide those, nor should it).

	//
	// SYNCHRONIZE THIS WITH THE FUNCTION BELOW!!!
	//

	if before_nlp_analyzer {
		sentence_str = strings.Replace(sentence_str, "next one", "next it", -1)
		sentence_str = strings.Replace(sentence_str, "previous one", "previous it", -1)
		sentence_str = strings.Replace(sentence_str, "that one", "that it", -1)
		sentence_str = strings.Replace(sentence_str, "this one", "this it", -1)

		sentence_str = strings.Replace(sentence_str, "what is", "what's", -1)
		sentence_str = strings.Replace(sentence_str, "whats", "what's", -1)

		sentence_str = strings.Replace(sentence_str, "who is", "who's", -1)
		sentence_str = strings.Replace(sentence_str, "whos", "who's", -1)

		sentence_str = strings.Replace(sentence_str, "how is", "how's", -1)
		sentence_str = strings.Replace(sentence_str, "hows", "how's", -1)

		sentence_str = strings.Replace(sentence_str, "that is", "that's", -1)
		sentence_str = strings.Replace(sentence_str, "thats", "that's", -1)

		sentence_str = strings.Replace(sentence_str, "there is", "there's", -1)
		sentence_str = strings.Replace(sentence_str, "theres", "there's", -1)

		sentence_str = strings.Replace(sentence_str, "do not", "don't", -1)
		sentence_str = strings.Replace(sentence_str, "dont", "don't", -1)

		// This may be incorrect sometimes, but the recognizers may not be able to distinguish, so one must treat them
		// as equal anyways.
		sentence_str = strings.Replace(sentence_str, "shutdown", "shut down", -1)
	} else {
		// Do these only after the NLP analyzer. Removing dashes may be bad for it, who knows. Better to let them stay
		// and remove only after it for the rest of the function analysis.

		for i, word := range *sentence {
			// No dashes ("wi-fi" == "wifi")
			if strings.Contains(word, "-") {
				(*sentence)[i] = strings.Replace(word, "-", "", -1)
			}
		}
	}

	return sentence_str
}

/*
sentenceNLPPreparation prepares the 'sentence' to be sent to the NLP analyzer, and also prepares it to be returned and
analyzed by the command detector.

-----------------------------------------------------------

– Params:
  - sentence – a pointer to the header of the created 'sentence' slice on the beginning of mainInternal()
  - before_sending – true if this function is being called before the NLP analyzer, false if it's being called after it

– Returns:
  - a string with the 'sentence' elements joined with a space between each (equivalent to 'sentence_str' on Main()).
*/
func sentenceNLPPreparation(sentence_str string, sentence *[]string, before_nlp_analyzer bool) string {
	//
	// SYNCHRONIZE THIS WITH THE FUNCTION ABOVE!!!
	//

	if before_nlp_analyzer {
		sentence_str = strings.Replace(sentence_str, "what's", "what is", -1)
		sentence_str = strings.Replace(sentence_str, "who's", "who is", -1)
		sentence_str = strings.Replace(sentence_str, "how's", "how is", -1)
		sentence_str = strings.Replace(sentence_str, "that's", "that is", -1)
		sentence_str = strings.Replace(sentence_str, "there's", "there is", -1)
		sentence_str = strings.Replace(sentence_str, "don't", "do not", -1)
	} else {
		sentence_str = strings.Replace(sentence_str, "what is", "what's", -1)
		sentence_str = strings.Replace(sentence_str, "who is", "who's", -1)
		sentence_str = strings.Replace(sentence_str, "how is", "how's", -1)
		sentence_str = strings.Replace(sentence_str, "that is", "that's", -1)
		sentence_str = strings.Replace(sentence_str, "there is", "there's", -1)
		sentence_str = strings.Replace(sentence_str, "do not", "don't", -1)
	}

	*sentence = strings.Split(sentence_str, " ")

	return sentence_str
}
