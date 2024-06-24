// https://leetcode.com/problems/minimum-number-of-k-consecutive-bit-flips/description/?envType=daily-question&envId=2024-06-24

package main

func main() {

}


k เป็น int ที่จำกำหนดความยาวของ sub array
ทีนี้โจทย์คือ 
ต้องทำให้ array nums เป็น 1 ทั้งก้อน [1, 1, 1, ..., 1]
แต่เปลี่ยนได้ทีละ [] sub array ที่ความยาวจำกัดที่ k นะ

เช่นๆ
Example 1
=========
Input: nums = [0,1,0], k = 1
Output: 2
Explanation: Flip nums[0], then flip nums[2].
ไม่มีไรอันนี้ให้ sub array len = 1 สบายเลย
เปลี่ยนทีละตัวให้เป็น 1
therefore -> minimum flip = 2 

Example 2
=========
Input: nums = [1,1,0], k = 2
Output: -1
Explanation: No matter how we flip subarrays of size 2, we cannot make the array become [1,1,1].
อันนี้ก็ไม่มีไร คือแบบ มึงมี array num ยาว 3 อะ
เราอยากทำ 1 หมด
สมมติเขียนมือนะ 
มึงฟลิปทีละ 2 element นะ
ฟลิปทีนึงได้ 000 งี้
ฟลิกอีกที 011, 110 งี้
ยังไงก็ 111 ไม่ได้
therefore -> minimum flip = -1 กุทำ all 1 ไม่ได้

Example 3
=========
               0 1 2 3 4 5 6 7
Input: nums = [0,0,0,1,0,1,1,0], k = 3
Output: 3
Explanation: 
Flip nums[0],nums[1],nums[2]: nums becomes [1,1,1,1,0,1,1,0]

Flip nums[4],nums[5],nums[6]:
อันนี้ flip nums[4] (0) -> 1
flip nums[5] (1) -> 0
flip nums[6] (1) -> 0
nums ก็จะเป็น [1,1,1,1,1,0,0,0]

Flip nums[5],nums[6],nums[7]: nums becomes [1,1,1,1,1,1,1,1]
 

