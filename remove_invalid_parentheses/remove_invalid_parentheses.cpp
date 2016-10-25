#include <string>
#include <unordered_set>
using namespace std;

class Solution {
public:
    vector<string> removeInvalidParentheses(string s) {
        int left = 0;
        int right = 0;
        for (auto ch: s) {
        	if (ch == '(') ++left;
        	if (ch == ')') {
        		if (left > 0) --left;
        		else ++right;
        	}
        }
        vector<string> results;
        unordered_set<string> unique;
        DFS(s, 0, 0, left, right, "", results, unique);
        return results;
    }

    void DFS(string&s, int idx, int pair, int left, int right, string solution, vector<string>& results, unordered_set<string>& unique) {
    	if (idx == s.size()) {
    		if (left == 0 && right == 0 && pair == 0 && unique.find(solution) == unique.end()) {
    			results.push_back(solution);
    			unique.insert(solution);
    		}
    		return;
    	}
    	if (s[idx] == '(') {
    		if (left > 0) DFS(s, idx+1, pair, left-1, right, solution, results, unique);
    		DFS(s, idx+1, pair+1, left, right, solution+"(", results, unique);
    	}else if (s[idx] == ')') {
    		if (right > 0) DFS(s, idx+1, pair, left, right-1, solution, results, unique);
    		if (pair > 0) DFS(s, idx+1, pair-1, left, right, solution+")", results, unique);
    	}else {
    		DFS(s, idx+1, pair, left, right, solution+s[idx], results, unique);
    	}
    	return;
    }
};