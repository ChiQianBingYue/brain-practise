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

Then we start from left to right, when new element with index j comes in, 
we check if it is the same with each previous char from right to left(with index i),
if so, and i,j is adjecent(j - i < 2), or p[i + 1][j - 1] is True,
f[i + 1] shoud be min(f[i + 1], f[j] + 1), and we set p[i][j] to be True.

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
        p = [[False for _ in range(n)] for _ in range(n)]
        for i in range(n + 1):
            f.append(i - 1)
        for j in range(n):
            for i in range(j, -1, -1):
                if s[i] == s[j] and (j - i < 2 or p[i+1][j-1]):
                    p[i][j] = True
                    f[j + 1] = min(f[j + 1], f[i] + 1)
        return f[n]

s = Solution()
s.minCut('a')