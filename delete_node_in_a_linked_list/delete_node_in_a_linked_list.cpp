struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    void deleteNode(ListNode* node) {
        while (node->next != NULL) {
        	node->val = node->next->val;
        	if (node->next->next == NULL)
        		node->next = NULL;
        	else
        		node = node->next;
        }
        return;
    }
};