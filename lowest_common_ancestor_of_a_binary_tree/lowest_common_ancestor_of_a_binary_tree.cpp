#include <vector>
using namespace std;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        vector<TreeNode*> path_p = {root};
        this->searchNode(root, p, path_p);
        vector<TreeNode*> path_q = {root};
        this->searchNode(root, q, path_q);
        int idx = path_p.size()<path_q.size() ? path_p.size()-1 : path_q.size()-1;
        for (;idx >= 0; --idx) {
        	if (path_p[idx] == path_q[idx]) return path_p[idx];
        }
        return NULL;
    }

    bool searchNode(TreeNode* root, TreeNode* target, vector<TreeNode*>& path) {
    	if (root == target) {
    		return true;
    	}
    	if (root->left == NULL && root->right == NULL) {
    		return false;
    	}
    	if (root->left != NULL){
    		path.push_back(root->left);
    		if (searchNode(root->left, target, path)){
    			return true;
    		}
    		path.pop_back();
    	}
    	if (root->right != NULL){
    		path.push_back(root->right);
    		if (searchNode(root->right, target, path)){
    			return true;
    		}
    		path.pop_back();
    	} 
    	return false;
    }
};