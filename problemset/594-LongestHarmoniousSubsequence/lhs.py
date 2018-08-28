class Solution:
    def findLHS(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        self.numminplace ={}
        self.nummaxplace ={}
        for i in range(len(nums)):
            if nums[i] in self.numminplace:
                self.nummaxplace[nums[i]][1] += 1
            else:
                self.numminplace[nums[i]] = [i, 1]
                self.nummaxplace[nums[i]] = [i, 1]
        maxlength = 0
        sort = sorted(nums)
        for i in range(len(sort)-1):
            if sort[i+1]-1 == sort[i]:
                maxplace = self.nummaxplace[sort[i]][1] + self.nummaxplace[sort[i+1]][1]
                maxlength = maxplace if maxplace>maxlength else maxlength
        return maxlength