class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        for index in range(len(nums)):
            t = nums[index]
            try:
                res = nums.index(target-t, index+1)
                return [index,res]
            except:
                pass
        