class Solution {
public:
    vector<vector<int>> palindromePairs(vector<string>& words) {
        vector<vector<int>> results;
        unordered_map<string, int> mapping;
        set<int> index_set;
        for (int i = 0; i < words.size(); ++i) {
        	mapping[words[i]] = i;
        	index_set.insert(words[i].size());
        }
        for (int i = 0; i < words.size(); ++i) {
        	string word = words[i];
        	int len = word.size();
        	reverse(word.begin(), word.end());
        	if (mapping.count(word) != 0 && mapping[word] != i) {
        		results.push_back(vector<int>{i, mapping[word]});
        	}
        	for (auto j = index_set.begin(); j != index_set.find(word.size()); ++j) {
        		int d = *j;
        		if (is_palindrome(word, 0, len - d - 1) && mapping.count(word.substr(len - d))) {
                    results.push_back({i, mapping[word.substr(len - d)]});
                }
                if (is_palindrome(word, d, len - 1) && mapping.count(word.substr(0, d))) {
                    results.push_back({mapping[word.substr(0, d)], i});
                }
        	}
        }
        return results;
    }

    bool is_palindrome(string& s, int i, int j) {
    	while (i < j) {
    		if (s[i++] != s[j--]) return false;
    	}
    	return true;
    }
};