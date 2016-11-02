struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
	unordered_map<TreeNode*, int> records;
public:
    int rob(TreeNode* root) {
        if (root == NULL) return 0;
        if (records.count(root) == 0) {
        	int left = rob(root->left);
        	int right = rob(root->right);
        	int without_left = 0, without_right = 0;
        	if (root->left) without_left = rob(root->left->left) + rob(root->left->right);
        	if (root->right) without_right = rob(root->right->left) + rob(root->right->right);
        	records[root] = max(root->val+without_left+without_right, left+right);
        }
        return records[root];
    }
};