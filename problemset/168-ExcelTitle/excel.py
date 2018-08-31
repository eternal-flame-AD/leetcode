import string
def form(n):
    while n>0:
        yield string.ascii_uppercase[n%26-1]
        n = (n-1)//26

class Solution:
    def convertToTitle(self, n):
        """
        :type n: int
        :rtype: str
        """
        return "".join([x for x  in form(n)][::-1])
        