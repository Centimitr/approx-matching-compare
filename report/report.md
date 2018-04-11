# The practicality of approxmate matching methods in spelling correction

## Introduction

Spellings are not always correct, so the users need suggestions. Only the users know what the correct words they want should be and what computer systems can do is to guess the possible candidates. An appropriate start for this problem is to make the assumption that a wrong spelling might be similar in some way to the correct one. This report analyses several methods which can be used to test the similarity between two words. Using these methods, a list of candidates can be produced from a dictionary. Evaluations of these methods of their effectiveness and efficiency are explained to measure their practical value to use to reccommand a suggestion.

## Dataset

The dataset includes a list of misspelled words and a list of its corresponding expected corrections with 719 items. A dictionary from UrbanDictionary[] which has 393954 items is provided to help the correction process. This is a real-world dataset, so the expected corrections are not guaranteed. A test directly match the items in misspelled list and corrected list with the dictionary has been done to check how many cases in the misspelled list are possibly solvable using the dictionary. 

| Dataset     | measure   | value            |
| ----------- | --------- | ---------------- |
| corrections | accuracy  | 594/716 (82.96%) |
| misspells   | recall    | 6/716 (0.84%)    |
| misspells   | precision | 6/175 (3.43%)    |

From the table, it can be seen that the accuracy of the corrections is 594/716, which means the 122 correct spellings do not exist in the dictionary. For these cases, providing correction suggestions from the dictionary can never success. 

There is also another assumption that if a word in the misspelled list exists in the dictionary, it is not misspelled. It is because the dictionary cannot produce a better match than a perfect match and providing corrections for a probably right word seems to be strange. However, from the test result above, 175 words are considered right spellings while only 3.43% of them are correct indeed. This assumption can solve 0.84% of the problem and other 23.6% of the problem are wrong speculations.

More data sources such as frequency statistics are meaningful. These data can provide a wider cover of the solutions and can support decision makings when the system do not know whether a word is misspelled or not. 

## Methods Overview

### Neiboughhood Search

The neighbourhood search method is to enumerate possible variants from a given spelling and then verify them. The variants are generated with one or more modification. In this report, the method adopts insertion, deletion and replacement. More times of modification will produce much larger candidate sets and its recall might be improved. K is the times of modification.N-grams Distance

### Global Edit Distance

### Soundex

## Evaluation

Recall & Precision

All bad methods

is the best

## Application

## Conclusions



# Refrences

Zobel, Justin and Philip Dart. (1996). Phonetic String Matching: Lessons from Informa-tion Retrieval. In Proceedings of the Eighteenth International ACM SIGIR Conference on Research and Development in Information Retrieval. Zu ̈rich, Switzerland. pp. 166–173.

Naomi Saphra and Adam Lopez (2016) Evaluating Informal-Domain Word Representations with UrbanDictionary. In Proceedings of the 1st Workshop on Evaluating Vector-Space Representations for NLP, Berlin, Germany. pp. 94–98.