struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL) {}
};
class Solution {
public:
	ListNode* detectCycle(ListNode *head) {
		ListNode* fast = head;
		ListNode* slow = head;
		bool first_step = true;
		while (first_step || fast!=slow){
			first_step = false;
			for (int i = 0; i < 2 && fast!=NULL; ++i) fast = fast->next;
			if (fast==NULL) return NULL;
			slow = slow->next;
		}
		int steps = 1;
		while (fast->next != slow){
			fast = fast->next;
			++steps;
		}
		fast = slow = head;
		for (int i = 0; i < steps; ++i) fast = fast->next;
		while (fast!=slow){
			fast = fast->next;
			slow = slow->next;
		}
		return slow;
	}
};