class Solution {
public:
    string removeDuplicateLetters(string s) {
    	if (s.size() == 0) return "";
        unordered_map<char, int> count;
        for (int i = 0; i < s.size(); ++i) {
        	++count[s[i]];
        }
        unordered_set<char> existed = {s[0]};
        vector<char> vec = {s[0]};
        --count[s[0]];
        for (int i = 1; i < s.size(); ++i) {
       		--count[s[i]];
        	if (existed.count(s[i]) == 0) {
        		while (!vec.empty()) {
        			char back = vec.back();
        			if (back <= s[i] || count[back] == 0) break;
        			vec.pop_back();
        			existed.erase(back);
        		}
	        	vec.push_back(s[i]);
				existed.insert(s[i]);
        	}
        }
        string ret;
        for (auto ch: vec) {
        	ret += ch;
        }
        return ret;
    }
};