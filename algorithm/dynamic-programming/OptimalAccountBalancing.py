"""
Description

Given a directed graph where each edge is represented by a tuple, such as [u, v, w]represents an edge with a weight w from u to v.
You need to calculate at least the need to add the number of edges to ensure that each point of the weight are balancing. That is, the sum of weight of the edge pointing to this point is equal to the sum of weight of the edge of the point that points to the other point.

Example

For example:
Given a graph [[0,1,10],[2,0,5]]
Return 2
Two edges are need to added. There are [1,0,5] and [1,2,5]

Given a graph [[0,1,10],[1,0,1],[1,2,5],[2,0,5]]
Return 1
Only one edge need to added. There is [1,0,4]

Idea

First, we get the account balance on each node and exclude those were already balanced, that is 0 account balance. Then, we consider each set of the rest accounts and check if sum of account balances in this set is 0. If true, we can say that this set of account can be balanced. And if this set has n nodes, we only need make a circle of these ndoes so needs n -1 edges to balance all of them to 0. Beside, if this set can be divided into two subset and sum of these subsets' additional edges is smaller than the previous result, then the current set should be the sum.
"""
from sys import maxsize

class Solution:
    """
    @param edges: a directed graph where each edge is represented by a tuple
    @return: the number of edges
    """

    def balanceGraph(self, edges):
        dept = {}
        account = []
        for e in edges:
            dept[e[0]] = dept.get(e[0], 0) - e[2]
            dept[e[1]] = dept.get(e[1], 0) + e[2]
        for v in dept.values():
            if v != 0: #exclude balanced accounts
                account.append(v)
        accountNum = len(account)
        if accountNum == 0:
            return 0
        setNum = 1 << accountNum
        dp = [maxsize for _ in range(setNum)]
        for _set in range(1, setNum):
            _sum = 0
            count = 0
            for node in range(0, accountNum):
                if ((1 << node) & _set) != 0: # node in the set
                    _sum += account[node]
                    count += 1
            if _sum == 0:
                dp[_set] = count - 1
                for subset in range(1, _set):
                    
                    if (subset & _set) == subset \
                        and dp[subset] + dp[_set - subset] < dp[_set]:
                        # subset is in set and these two subset can reduce number of edges
                        dp[_set] = dp[subset] + dp[_set - subset]
        return dp[setNum - 1]
        