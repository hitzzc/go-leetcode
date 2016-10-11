#include <unordered_map>
using namespace std;

struct RandomListNode {
	int label;
	RandomListNode *next, *random;
	RandomListNode(int x) : label(x), next(NULL), random(NULL) {}
};

class Solution {
public:
	RandomListNode *copyRandomList(RandomListNode *head) {
		if (head==NULL) return NULL;
		auto fake = new RandomListNode(0);
		unordered_map<RandomListNode*, RandomListNode*> mapping;
		auto ptr = head;
		auto pre_copy = fake;
		while (ptr!=NULL){
			RandomListNode* copy_node;
			if (mapping.find(ptr)!=mapping.end())
				copy_node = mapping[ptr];
			else{
				copy_node = new RandomListNode(ptr->label);
				mapping[ptr] = copy_node;
			}
			pre_copy->next = copy_node;
			pre_copy = pre_copy->next;

			if (ptr->random!=NULL){
				RandomListNode* copy_node_random;
				if (mapping.find(ptr->random)!=mapping.end()){
					copy_node_random = mapping[ptr->random];
				}
				else{
					copy_node_random = new RandomListNode(ptr->random->label);
					mapping[ptr->random] = copy_node_random;
				}
				copy_node->random = copy_node_random;
			}
			ptr = ptr->next;
		}
		return fake->next;
	}
};