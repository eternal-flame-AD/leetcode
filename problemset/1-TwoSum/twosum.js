var twoSum = function(nums, target) {
    for (var index in nums) {
        index = parseInt(index)
        var num = nums[index];
        var res = nums.indexOf(target-num, index+1);
        if (res!==-1) {
            return [parseInt(index), res];
        }
    }
};