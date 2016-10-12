#include <string>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
public:
	bool wordBreak(string s, unordered_set<string>& wordDict) {
		vector<bool> history(s.size(), false);
		for (int i = 0; i < s.size(); ++i){
			bool yes = false;
			for (int j = i; j >=0 && !yes; --j){
				if (wordDict.find(s.substr(j, i-j+1))!=wordDict.end() && (j == 0 || history[j-1]))
					yes = true;
			}
			history[i] = yes;
		}
		return history[s.size()-1];
	}
};