"""
Description
Imagine you have a special keyboard with the following keys:

Key 1: (A): Print one 'A' on screen.

Key 2: (Ctrl-A): Select the whole screen.

Key 3: (Ctrl-C): Copy selection to buffer.

Key 4: (Ctrl-V): Print buffer on screen appending it after what has already been printed.

Now, you can only press the keyboard for N times (with the above four keys), find out the maximum numbers of 'A' you can print on screen.

Example
Given N = 3, return 3.

Explanation: 
We can at most get 3 A's on screen by pressing following key sequence:
A, A, A
Given N = 7, return 9.

Explanation: 
We can at most get 9 A's on screen by pressing following key sequence:
A, A, A, Ctrl A, Ctrl C, Ctrl V, Ctrl V

Idea:

Let dp[i] to store the result for i presses,
then 
dp[0] = 0
dp[1] = 1
dp[2] = 2
dp[3] = 3
...
dp[6] = 6
dp[7] = 9

We can find that for each dp[i+1] with given dp[i]
we know dp[i + 1] = max( 
    dp[i] + 1, 
    dp[i-3] * 2,
    dp[i-4] * 3,
    .....
    dp[1] * i
    )

"""
class Solution:
    """
    @param N: an integer
    @return: return an integer
    """
    def maxA(self, N):
        # write your code here
        dp = list(range(N+1))
        for i in range(4, N + 1):
            prev = i - 3
            count = 2
            while prev > 0:
                dp[i] = max(dp[i], dp[prev] * count)
                prev -= 1
                count += 1 
        return dp[-1]