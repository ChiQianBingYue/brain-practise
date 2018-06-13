"""
Medium

Palindrome Partitioning II

Description
Given a string s, cut s into some substrings such that every substring is a palindrome.

Return the minimum cuts needed for a palindrome partitioning of s.


Example
Given s = "aab",

Return 1 since the palindrome partitioning ["aa", "b"] could be produced using 1 cut.

Better Example
Given s = "abcddcba"

Ideal:
Use f[i] to store the min cut for the string from its beginning to its (i)-th char.
First element in f, f[0] should be -1.
for s = "aab",
f = [-1, 0, 0, 1]

Use p[i][j]  to cache if s[i:j] is a palindrome

To get the answer, we first consider there is not palindrome in the string, so for string with length n,
f should be [-1, 0, 1, 2, 3, ... , n - 1]

Then we start from left to right, when new element with index i comes in, 
we check if it is the same with each previous char from right to left(with index i),
if so, and i,j is adjecent, or p[j + 1][i - 1] is True,
f[i] shoud be min(f[i], f[j - 1] + 1), and we set p[i][j] to be True.

After this n^2 loop finish, we return last element in f.
"""
class Solution:
    """
    @param s: A string
    @return: An integer
    """
    def minCut(self, s):
        n = len(s)
        f = []
        p = [[False for x in range(n)] for x in range(n)]
        #the worst case is cutting by each char
        for i in range(n+1):
            f.append(i - 1) # the first one, f[n]=-1
        for i in range(n):
            for j in range(i, -1, -1):
                if (s[i] == s[j] and (i - j < 2 or p[j + 1][i - 1])):
                    p[j][i] = True
                    f[i + 1] = min(f[i + 1], f[j] + 1)
        return f[-1]

s = Solution()
s.minCut('a')