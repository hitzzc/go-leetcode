#include <vector>
#include <queue>
#include <unordered_set>
#include <unordered_map>
using namespace std;

struct UndirectedGraphNode {
	int label;
	vector<UndirectedGraphNode *> neighbors;
	UndirectedGraphNode(int x) : label(x) {};
};

class Solution {
public:
	UndirectedGraphNode *cloneGraph(UndirectedGraphNode *node) {
		if (node==NULL) return NULL;
		queue<UndirectedGraphNode*> q;
		unordered_set<UndirectedGraphNode*> visited;
		unordered_map<UndirectedGraphNode*, UndirectedGraphNode*> mapping;
		q.push(node);
		while (!q.empty()){
			UndirectedGraphNode* p = q.front();
			q.pop();
			if (p==NULL) continue;
			UndirectedGraphNode* copy;
			if (mapping.find(p)!=mapping.end()){
				copy = mapping[p];
			}else {
				copy = new UndirectedGraphNode(p->label);
				mapping[p] = copy;
			}
			visited.insert(p);
			for (auto child: p->neighbors){
				if (mapping.find(child)!=mapping.end()){
					copy->neighbors.push_back(mapping[child]);
				}else{
					UndirectedGraphNode* copy_child = new UndirectedGraphNode(child->label);
					copy->neighbors.push_back(copy_child);
					mapping[child] = copy_child;
					if (visited.find(child)==visited.end()) q.push(child);
				}
			}
		}
		return mapping[node];
	}
};