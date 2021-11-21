# wordsVerificationDADi()

- Put the assistant calculating some parameters for the function.

For example, it could see if with "set an alarm", "set" appears only in the first sub-slice, it would set the
parameter to exclude that word from the list, but not if words are repeated on the sub-slices!
Infinite parameters on the function but some automatic! Adaptable assistant by itself...!

Based on the content of some function slices, it would adjust other parameters automatically without needing to go there
manually and hardcode it.


# NLP

- Also, ONLY USE NLP WHEN IT IS REALLY NEEDED! If not REALLY needed, use the function without its help (it may fail, like
with "turn on the wifi", on which it thinks "turn" is a name and not a verb - works in a bigger sentence though).

- Use NLP to exclude other adjectives from the words_list if the only one was said already, for example? (Or a verb or
whatever) Of course, works only (as said) when there is only and only 1 adjective to form the command.

- "turn airplane mode off and the wifi too" - he doesn't know "turn" is linked to "off". how will he know? as soon as he
understands that the command is to {TURN_OFF, something}... --> do this

Another example: "turn wifi on and the airplane mode and the flashlight" - only the Wi-Fi is detected there. Only works
if "on" is before "wifi" (which is expected, by the current implementation).


# NLP + wordsVerificationDADi()

- Re-implement the verification function to *also* use NLP: "turn on the wifi" --> "verb [no idea] name" or "verb name
[no idea]". Then it checks if the words on the 'words_list' are in the intervals of words that it found. For example,
a verb is at index 12. Then from index 13 until the next name, there must be an "on". Then from the next name to the
next non-name (since we're looking for a name), "wifi" must be there. And so on.


# General

"turn on the wifi, the airplane mode and the flashlight" --> ????? There are no commas on speech recognizers...
