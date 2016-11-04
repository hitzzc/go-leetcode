class Solution {
public:
    bool isMatch(string s, string p) {
        return isMatch(s, p, s.size(), p.size());
    }
    bool isMatch(string& s, string& p, int i, int j) {
    	if (i == 0 && j == 0) return true;
    	if (i != 0 && j == 0) return false;
    	if (i == 0 && j != 0) {
    		if (p[j-1] == '*') {
    			return isMatch(s, p, i, j-2);
    		}
    		return false;
    	}
    	if (s[i-1] == p[j-1] || p[j-1] == '.') {
    		return isMatch(s, p, i-1, j-1);
    	} else if (p[j-1] == '*') {
    		if (isMatch(s, p, i, j-2)) return true;
    		if (s[i-1] == p[j-2] || p[j-2] == '.') {
    			return isMatch(s, p, i-1, j);
    		}
    		return false;
    	}
    	return false;
    }
};