class Solution {
public:
    vector<string> findItinerary(vector<pair<string, string>> tickets) {
		unordered_map<string, multiset<string>> graph;
		for (auto& kv: tickets) {
			graph[kv.first].insert(kv.second);
		}
		vector<string> solution;
		DFS(graph, solution, "JFK");
		return vector<string>(solution.rbegin(), solution.rend());
    }
    void DFS(unordered_map<string, multiset<string>>& graph, vector<string>& solution, string start) {
    	while (!graph[start].empty()) {
    		string city = *graph[start].begin();
    		graph[start].erase(city);
    		(DFS(graph, solution, city));
    	}
    	solution.push_back(start);
    	return;
    }
};