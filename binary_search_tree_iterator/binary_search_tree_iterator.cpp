#include <stack>

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class BSTIterator {
private:
	stack<TreeNode*> stk;
public:
	BSTIterator(TreeNode *root) {
		while (root != NULL){
			stk.push(root);
			root = root->left;
		}
	}
	
	/** @return whether we have a next smallest number */
	bool hasNext() {
		return !stk.empty();
	}

	/** @return the next smallest number */
	int next() {
		auto top = stk.top();
		stk.pop();
		int ret = top->val;
		top = top->right;
		while (top != NULL){
			stk.push(top);
			top = top->left;
		}
		return ret;
	}
};