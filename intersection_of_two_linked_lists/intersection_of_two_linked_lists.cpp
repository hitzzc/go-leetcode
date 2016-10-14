struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(NULL) {}
 };

class Solution {
public:
	ListNode *getIntersectionNode(ListNode *headA, ListNode *headB) {
		int count1 = 0;
		int count2 = 0;
		ListNode* p1 = headA;
		ListNode* p2 = headB;
		while (p1 != NULL) {
			++count1;
			p1 = p1->next;
		}
		while (p2 != NULL) {
			++count2;
			p2 = p2->next;
		}
		int diff;
		if (count1 > count2) {
			diff = count1 - count2;
			p1 = headA;
			p2 = headB;
		}else {
			diff = count2 - count1;
			p1 = headB;
			p2 = headA;
		}
		for (; diff > 0; --diff) p1 = p1->next;
		while (p1 != p2){
			p1 = p1->next;
			p2 = p2->next;
		}
		return p1;
	}
};