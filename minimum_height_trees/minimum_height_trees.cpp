class Solution {
public:
    vector<int> findMinHeightTrees(int n, vector<pair<int, int>>& edges) {
        if (n <= 1) return {0};
        vector<unordered_set<int>> graph(n);
        for (auto& edge: edges) {
        	graph[edge.first].insert(edge.second);
        	graph[edge.second].insert(edge.first);
        }
        vector<int> current;
        for (int i = 0; i < graph.size(); ++i) {
        	if (graph[i].size() == 1) current.push_back(i);
        }
        while (true) {
        	vector<int> next;
        	for (auto& node: current) {
        		for (auto& neighbor: graph[node]) {
        			graph[neighbor].erase(node);
        			if (graph[neighbor].size() == 1) next.push_back(neighbor);
        		}
        	}
        	if (next.size() == 0) break;
        	current.swap(next);
        }
        return current;
    }
};