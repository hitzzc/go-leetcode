struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
	bool hasCycle(ListNode *head) {
		ListNode* fast = head;
		ListNode* slow = head;
		bool first_step = true;
		while (first_step || fast!=slow){
			first_step = false;
			for (int i = 0; i < 2 && fast!=NULL; ++i) fast = fast->next;
			if (fast==NULL) return false;
			slow = slow->next;
		}
		return true;
	}
};