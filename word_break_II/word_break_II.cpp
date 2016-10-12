#include <string>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
public:
	vector<string> wordBreak(string s, unordered_set<string>& wordDict) {
		vector<bool> history(s.size()+1, true);
		vector<string> results;
		vector<string> path;
		DFS(s, 0, results, path, history, wordDict);
		return results;
	}

	void DFS(string s, int start, vector<string>& results, vector<string>& path, vector<bool>& history, unordered_set<string>& wordDict){
		if (start==s.size()){
			if (path.empty()) return;
			string ret = path[0];
			for (int i = 1; i < path.size(); ++i)
				ret = ret + " " + path[i];
			results.push_back(ret);
			return;
		}
		int pre_num = results.size();
		for (int end = start+1; end <= s.size(); ++end){
			if (!history[end] || wordDict.find(s.substr(start, end-start))==wordDict.end()) continue;
			path.push_back(s.substr(start, end-start));
			DFS(s, end, results, path, history, wordDict);
			path.pop_back();
		}
		if (results.size()==pre_num) history[start] = false;
		return;
	}
};