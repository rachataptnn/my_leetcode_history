let nums = [1,6,1]
let k = 3
console.log(smallestDistancePair(nums, k))

function smallestDistancePair(nums: number[], k: number): number {
    nums.sort((a, b) => a - b);

    // Binary Search:
    let minNum = nums[0];
    let maxNum = nums[nums.length - 1];

    let minDistance = 0;
    let maxDistance = maxNum - minNum;

    while (minDistance < maxDistance) {
        const midDistance = minDistance + Math.floor((maxDistance - minDistance) / 2);

        const pairsCount = countPairsWithDistanceLessThanOrEqual(nums, midDistance);
        if (pairsCount >= k) {
            // If we have at least k pairs
            // the k-th smallest distance is in the lower half, so narrow down the search
            maxDistance = midDistance;
        } else {
            // If we have fewer than k pairs
            // the k-th smallest distance is in the upper half, so narrow up the search
            minDistance = midDistance + 1;
        }
    }

    return minDistance;
}

function countPairsWithDistanceLessThanOrEqual(nums: number[], targetDistance: number): number {
    let count = 0;
    let right = 0;

    // Two Pointer:
    for (let left = 0; left < nums.length; left++) {
        // Move the right pointer while maintaining the distance condition
        while (right < nums.length && nums[right] - nums[left] <= targetDistance) {
            right++;
        }
        count += right - left - 1; // Count the pairs between left and right (excluding the pair with itself)
    }

    return count;
}
