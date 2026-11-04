impl Solution {
 pub fn is_anagram(s: String, t: String) -> bool {
    if s.len() != t.len() {
        return false;
    }

    let mut counts = [0i32; 26];

    for (a, b) in s.bytes().zip(t.bytes()) {
        counts[usize::from(a - b'a')] += 1;
        counts[usize::from(b - b'a')] -= 1;
    }

    counts.into_iter().all(|count| count == 0)
    }
}
