class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        seen = {}
        for num in nums:
            if seen.get(num) == 1:
                return True
            seen[num] = 1
            
        return False
