struct TreeLinkNode {
	int val;
	TreeLinkNode *left, *right, *next;
	TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
};
 
void connect(TreeLinkNode *root){
	if (!root){
		return;
	}
	TreeLinkNode* currentNode;
	while (root->left != NULL){
		currentNode = root;
		while (currentNode != NULL){
			currentNode->left->next = currentNode->right;
			if (currentNode->next != NULL){
				currentNode->right->next = currentNode->next->left;
			}
			currentNode = currentNode->next;
		}
		root = root->left;
	}
	return;
}