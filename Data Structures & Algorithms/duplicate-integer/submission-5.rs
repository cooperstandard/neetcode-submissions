impl Solution {
    pub fn has_duplicate(nums: Vec<i32>) -> bool {
        let mut set = HashSet::new();
        nums.into_iter().any(|n| !set.insert(n))
    }
}
