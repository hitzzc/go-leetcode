class Solution {
public:
    int longestSubstring(string s, int k) {
    	int ret = 0;
    	int i = 0;
        while (i+k < s.size()) {
        	int mask = 0, i_max = i;
        	vector<int> cnt(26, 0);
        	for (int j = i; j < s.size(); ++j) {
        		int idx = s[j] - 'a';
        		++cnt[idx];
        		if (cnt[idx] < k) mask |= (1 << idx);
        		else mask &= ~(1 << idx); 
        		if (mask == 0) {
        			ret = max(ret, j-i+1);
        			i_max = j;
        		}
        	}
        	i = i_max + 1;
        }
        return ret;
    }
};