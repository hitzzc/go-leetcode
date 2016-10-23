#include <unordered_map>
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
    bool wordPattern(string pattern, string str) {
        unordered_map<char, string> pattern_to_str;
        unordered_map<string, char> str_to_pattern;
        int i = 0;
        int start = 0;
        for (int j = 0; j <= str.size(); ++j){
        	if (i >= pattern.size()) return false;
        	if (j == str.size() || str[j] == ' ') {
        		string sub = str.substr(start, j-start);
        		if (pattern_to_str.count(pattern[i]) == 0 && str_to_pattern.count(sub) == 0) {
        			pattern_to_str[pattern[i]] = sub;
        			str_to_pattern[sub] = pattern[i];
        		}else if (pattern_to_str.count(pattern[i]) != 0 && str_to_pattern.count(sub) != 0) {
        			if (pattern_to_str[pattern[i]] != sub || str_to_pattern[sub] != pattern[i]) {
        				return false;
        			}
        		}else {
        			return false;
        		}
        		++i;
        		start = j+1;
        	}
        }
        if (i != pattern.size()) return false;
        return true;
    }
};