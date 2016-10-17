struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
	int countNodes(TreeNode* root) {
		int depth = 0;
		TreeNode *left, *right;
		left = right = root; 
		while (left!=NULL && right!=NULL){
			left = left->left;
			right = right->right;
			++depth;
		}
		if (left==NULL && right==NULL)
			return (1<<depth) - 1;
		int cnt1 = countNodes(root->left);
		int cnt2 = countNodes(root->right);
		return cnt1 + cnt2 + 1;
	}
};