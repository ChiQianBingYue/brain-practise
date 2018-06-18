"""
Description
Given a set of positive integers, write a function to divide it into two sets S1 and S2 such that the absolute difference between their sums is minimum.

If there is a set S with n elements, then if we assume Subset1 has m elements, Subset2 must have n-m elements and the value of abs(sum(Subset1) – sum(Subset2)) should be minimum.

Example
Given nums = [1, 6, 11, 5], return 1

// Subset1 = [1, 5, 6], sum of Subset1 = 12 
// Subset2 = [11], sum of Subset2 = 11   
// abs(11 - 12) = 1     

Ideal

Use boolean dp[i] to mark if one set in nums can added to i,
then we only need to find dp[i] is true and sum of nums - i is minnimal.

To get dp[i], we get number from nums one by one, from left to right. For 
each number n, we test if dp[j - n], where sum >= j >= n, is True, if yes,
then dp[j] is True. We don't want duplicate count of n, so we start from right
to left.
"""
class Solution:
    """
    @param nums: the given array
    @return: the minimum difference between their sums 
    """
    def findMin(self, nums):
        # write your code here
        numsSum = sum(nums)
        dp = [False] * (numsSum + 1)
        dp[0] = True
        res = numsSum
        for n in nums:
            for j in range(numsSum, n - 1, -1):
                if dp[j - n]:
                    dp[j] = True
                if dp[j]:
                    res = min(res, abs((numsSum - j) - j))
        return res

    """
    optimize:
    Since if dp[i] is Ture, then dp[numsSum - i] will be Ture, we 
    only need to track from dp[0] to dp[numsSum // 2]
    """
    def findMinOptimized(self, nums):
        # write your code here
        numsSum = sum(nums)
        dp = [False] * (numsSum // 2 + 1)
        dp[0] = True
        res = numsSum
        for n in nums:
            for j in range(numsSum // 2, n - 1, -1):
                if dp[j - n]:
                    dp[j] = True
                if dp[j]:
                    res = min(res, abs((numsSum - j) - j))
        return res


        