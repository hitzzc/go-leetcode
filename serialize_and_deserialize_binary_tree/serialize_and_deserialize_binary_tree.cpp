#include <string>
using namespace std;

struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Codec {
public:
    string serialize(TreeNode* root) {
        if (root == NULL) {
        	return "null";
        }
        string ret = to_string(root->val);
        ret += "," + serialize(root->left);
        ret += "," + serialize(root->right);
        return ret; 
    }

    TreeNode* deserialize(string data) {
    	int idx = 0;
    	return helper(data, idx);
    }

    TreeNode* helper(string& data, int& idx) {
    	int start = idx;
    	for (; idx < data.size(); ++idx) {
    		if (data[idx] == ',') {
    			break;
    		}
    	}
    	string sub = data.substr(start, idx-start);
    	++idx;
    	if (sub == "null") return NULL;
    	TreeNode* node = new TreeNode(stoi(sub));
    	if (idx == data.size()) return node;
    	node->left = helper(data, idx);
    	node->right = helper(data, idx);
    	return node;
    }
};