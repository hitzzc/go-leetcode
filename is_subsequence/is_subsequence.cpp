class Solution {
public:
    bool isSubsequence(string s, string t) {
    	if (s.size() == 0) return true;
        int ps = 0, pt = 0;
        for (; pt < t.size() && ps < t.size(); ++pt) {
        	if (s[ps] == t[pt]) ++ps;
        }
        return ps == s.size();
    }
};