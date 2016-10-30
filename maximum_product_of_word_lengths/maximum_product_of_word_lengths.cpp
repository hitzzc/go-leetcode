class Solution {
public:
    int maxProduct(vector<string>& words) {
        vector<int> helper(words.size(), 0);
        for (int i = 0; i < words.size(); ++i) {
        	for (auto& ch: words[i]) {
        		helper[i] |= 1 << (ch-'a');
        	}
        }
        int ret = 0;
        for (int i = 0; i < helper.size(); ++i) {
        	for (int j = i+1; j < helper.size(); ++j) {
        		if (helper[i] & helper[j]) continue;
        		if (words[i].size() * words[j].size() > ret) ret = words[i].size() * words[j].size();
        	}
        }
        return ret;
    }
};