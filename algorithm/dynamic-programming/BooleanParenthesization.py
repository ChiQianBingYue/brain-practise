"""
Description
Given a boolean expression with following symbols.

Symbols
    'T' ---> true 
    'F' ---> false 
And following operators filled between symbols

Operators
    &   ---> boolean AND
    |   ---> boolean OR
    ^   ---> boolean XOR 
Count the number of ways we can parenthesize the expression so that the value of expression evaluates to true.

Let the input be in form of two arrays one contains the symbols (T and F) in order and other contains operators (&, | and ^}


Example
Given symbol = ['T', 'F', 'T'], operator = ['^', '&']
return 2
The given expression is "T ^ F & T", it evaluates true, in two ways "((T ^ F) & T)" and "(T ^ (F & T))"

given symbol = ['T', 'F', 'F'], operator = ['^', '|']
return 2
The given expression is "T ^ F | F", it evaluates true, in two ways "( (T ^ F) | F )" and "( T ^ (F | F) )".

Ideal
Maintain two 2D matrix named T[][] and F[][],
Where T[i][j] indicates how many ways we can put parenthesize 
so the expression from i-th symbol to j-th symbol can 
evaluates to True. Where F[i][j] is the same but eval to False.

The base case is for for all T[i][i] = 1 and F[i][i] = 1 where 
i is range(0, len(symbol)). Then we can start from left to 
right with each base case (j), and sum the number of way to 
its beginning (i). And start with T[i][j] = 0 and F[i][j] = 0.
In each iteration, for iteration k in range(i, j). And for 
each k,
  if operator[k] is '&':
    T[i][j] += T[i][k] * T[k + 1][j]
    F[i][j] += (T[i][k] + F[i][k]) * (T[k+1][j] + F[k+1][j]) 
               - T[i][k] * T[k + 1][j]
  if operator[k] is '|':
    T[i][j] += (T[i][k] + F[i][k]) * (T[k+1][j] + F[k+1][j]) 
               - F[i][k] * F[k + 1][j]
    F[i][j] += F[i][k] * F[k + 1][j]
  if operator[k] is '^':
    T[i][j] += T[i][k] * F[k + 1][j] + F[i][k] * T[k + 1][j]
    F[i][j] += T[i][k] * T[k + 1][j] + F[i][k] * F[k + 1][j]
Finally we can get T[0][len(symbols)]
"""
class Solution:
    """
    @param symb: the array of symbols
    @param oper: the array of operators
    @return: the number of ways
    """
    def countParenth(self, symb, oper):
        if len(symb) == 0 or len(oper) + 1 != len(symb):
            raise ValueError('Input expression invalide')
        symbLen = len(symb)
        T = [[0 for _ in range(symbLen)] for _ in range(symbLen)]
        F = [[0 for _ in range(symbLen)] for _ in range(symbLen)]
        for i in range(symbLen):
            if symb[i] == 'T':
                T[i][i] = 1
            else:
                F[i][i] = 1
        for j in range(symbLen):
            for i in range(j - 1, -1, -1):
                T[i][j] = 0
                F[i][j] = 0
                for k in range(i, j):
                    if oper[k] == '&':
                        T[i][j] += T[i][k] * T[k+1][j]
                        F[i][j] += (T[i][k] + F[i][k]) * (T[k+1][j] + F[k+1][j]) - T[i][k] * T[k + 1][j]
                    elif oper[k] == '|':
                        F[i][j] += F[i][k] * F[k+1][j]
                        T[i][j] += (T[i][k] + F[i][k]) * (T[k+1][j] + F[k+1][j]) - F[i][k] * F[k+1][j]
                    else:
                        T[i][j] += T[i][k] * F[k+1][j] + F[i][k] * T[k+1][j]
                        F[i][j] += T[i][k] * T[k+1][j] + F[i][k] * F[k+1][j]
        return T[0][symbLen - 1]
      
s = Solution()
r = s.countParenth("TFT", "^&")
assert(r == 2)
r = s.countParenth("TFF", "^|")
assert(r == 2)