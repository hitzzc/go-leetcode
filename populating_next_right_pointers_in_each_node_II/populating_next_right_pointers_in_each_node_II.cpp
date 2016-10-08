struct TreeLinkNode {
	int val;
	TreeLinkNode *left, *right, *next;
	TreeLinkNode(int x) : val(x), left(NULL), right(NULL), next(NULL) {}
};

void connect(TreeLinkNode *root) {
	TreeLinkNode *head, *tail, *last;
#undef PUSH
#define PUSH(node)\
	if (head == NULL) { head = tail = node;}\
	else { tail->next = node; tail = node;}

	head = tail = NULL;
	last = root;
	PUSH(root);
	while (head != NULL){
		TreeLinkNode *p = head;
		head = head->next;
		if (p->left) PUSH(p->left);
		if (p->right) PUSH(p->right);
		if (p == last){
			last = tail;
			p->next = NULL;
		}
	}
	return;
}