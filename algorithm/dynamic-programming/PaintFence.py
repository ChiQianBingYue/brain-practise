"""
EAZY

There is a fence with n posts, each post can be painted with one of the k colors.
You have to paint all the posts such that no more than two adjacent fence posts have the same color.
Return the total number of ways you can paint the fence.

Example:
Given n=3, k=2 return 6

      post 1,   post 2, post 3
way1    0         0       1 
way2    0         1       0
way3    0         1       1
way4    1         0       0
way5    1         0       1
way6    1         1       0

Ideal:
dp[i] is the total number of ways with only i posts.
Then dp[i] could be k different types of post, but it have to be different from post 1 or post 2.
So, dp[i] = dp[i - 1] * (k - 1) + dp[i - 2] * (k - 1).
For example, in terms of dp[i] and dp[i - 1], we get way3 and way6, whatever post 2 is.
And in terms of dp[i] and dp[i - 2], we get way 1, 2, 5, 6, whatever post 1 is.
"""
class Solution:
    """
    @param n: non-negative integer, n posts
    @param k: non-negative integer, k colors
    @return: an integer, the total number of ways
    """
    def numWays(self, n, k):
        # write your code here
        dp = [0, k, k * k]
        if n <= 2:
            return dp[n]
        if k == 1 and n >= 3:
            return 0
        for _ in range(2, n):
            dp.append((dp[-1] + dp[-2]) * (k - 1))
        return dp[-1]
