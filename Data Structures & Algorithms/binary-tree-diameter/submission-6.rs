// Definition for a binary tree node.
// #[derive(Debug, PartialEq, Eq)]
// pub struct TreeNode {
//     pub val: i32,
//     pub left: Option<Rc<RefCell<TreeNode>>>,
//     pub right: Option<Rc<RefCell<TreeNode>>>,
// }
//
// impl TreeNode {
//     #[inline]
//     pub fn new(val: i32) -> Self {
//         TreeNode {
//             val,
//             left: None,
//             right: None,
//         }
//     }
// }

use std::rc::Rc;
use std::cell::RefCell;

impl Solution {
    pub fn diameter_of_binary_tree(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
		let mut diameter = 0;
		dfs(root, &mut diameter);
		diameter

    }


}
	fn dfs(root: Option<Rc<RefCell<TreeNode>>>, diameter: &mut i32) -> i32 {
		if let Some(node) = root {
			let left = dfs(node.borrow().left.clone(), diameter);
			let right = dfs(node.borrow().right.clone(), diameter);
			*diameter = (*diameter).max(left + right);
			return left.max(right) + 1;
		}
		0
	}
