"""
Description:

Given a rows x cols screen and a sentence represented by a list of non-empty words, find how many times the given sentence can be fitted on the screen.

A word cannot be split into two lines.
The order of words in the sentence must remain unchanged.
Two consecutive words in a line must be separated by a single space.
Total words in the sentence won't exceed 100.
Length of each word is greater than 0 and won't exceed 10.
1 ≤ rows, cols ≤ 20,000.

Example:
Given rows = 2, cols = 8, sentence = ["hello", "world"], retrun 1.

Explanation:
hello---
world---

The character '-' signifies an empty space on the screen.
Given rows = 3, cols = 6, sentence = ["a", "bcd", "e"], return 2.

Explanation:
a-bcd- 
e-a---
bcd-e-

The character '-' signifies an empty space on the screen.
Given rows = 4, cols = 5, sentence = ["I", "had", "apple", "pie"], return 1.

Explanation:
I-had
apple
pie-I
had--

The character '-' signifies an empty space on the screen.

Idea:
Given a sentence, say ["I", "had", "apple", "pie"], transform it to be sentenceString-"I had apple pie " with len-16. Then, we count how many pixel in the screen we will use to fit the sentence, called cursor. And for each row, we add our cursor with the col length cursor += col. And we get the last char by 
sentenceString[cursor % len]. If the last char is " ", we can switch to next row and add cursor by 1: cursor += 1. If not, we need to subscract cursor untile sentenceString[cursor % len] become " ". Finally, we use cursor // len to get count.
"""
class Solution:
    """
    @param sentence: a list of string
    @param rows: an integer
    @param cols: an integer
    @return: return an integer, denote times the given sentence can be fitted on the screen
    """
    def wordsTyping(self, sentence, rows, cols):
        # Write your code here
        sentenceString = ""
        for word in sentence:
            sentenceString += word
            sentenceString += " "
        length = len(sentenceString)
        cursor = 0
        for _ in range(rows):
            cursor += cols
            if sentenceString[cursor % length] == " ":
                cursor += 1
            else:
                while cursor > 0 and sentenceString[(cursor - 1) % length] != " ":
                    cursor -= 1
        return cursor // length

s = Solution()
s.wordsTyping(["I", "had", "apple", "pie"], 4, 5)