/**
 * @param {number[]} nums
 * @return {number}
 */
var findLHS = function(nums) {
    var sorted = nums.slice()
    sorted.sort((a,b)=>a-b)
    var numminplace = {}
    var nummaxplace = {}
    for (var index=0;index<nums.length;index++) {
        if (numminplace[nums[index]]===undefined) {
            numminplace[nums[index]] = [index, 1]
            nummaxplace[nums[index]] = [index, 1]
        } else {
            nummaxplace[nums[index]][1]++
        }
    }
    var maxlength = 0
    for (var i=0;i<sorted.length-1;i++) {
        if (sorted[i+1]-1==sorted[i]) {
            maxlength = Math.max(maxlength, nummaxplace[sorted[i]][1]+nummaxplace[sorted[i+1]][1])
        }
    }
    return maxlength
};