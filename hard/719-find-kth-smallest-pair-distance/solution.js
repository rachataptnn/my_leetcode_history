var nums = [1, 6, 1];
var k = 3;
console.log(smallestDistancePair(nums, k));

function smallestDistancePair(nums, k) {
    nums.sort(function (a, b) { return a - b; });
    // Binary Search:
    var minNum = nums[0];
    var maxNum = nums[nums.length - 1];
    var minDistance = 0;
    var maxDistance = maxNum - minNum;
    while (minDistance < maxDistance) {
        var midDistance = minDistance + Math.floor((maxDistance - minDistance) / 2);
        var pairsCount = countPairsWithDistanceLessThanOrEqual(nums, midDistance);
        if (pairsCount >= k) {
            // If we have at least k pairs
            // the k-th smallest distance is in the lower half, so narrow down the search
            maxDistance = midDistance;
        }
        else {
            // If we have fewer than k pairs
            // the k-th smallest distance is in the upper half, so narrow up the search
            minDistance = midDistance + 1;
        }
    }
    return minDistance;
}

function countPairsWithDistanceLessThanOrEqual(nums, targetDistance) {
    var count = 0;
    var right = 0;
    // Two Pointer:
    for (var left = 0; left < nums.length; left++) {
        // Move the right pointer while maintaining the distance condition
        while (right < nums.length && nums[right] - nums[left] <= targetDistance) {
            right++;
        }
        count += right - left - 1; // Count the pairs between left and right (excluding the pair with itself)
    }
    return count;
}
