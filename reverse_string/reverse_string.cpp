class Solution {
public:
    string reverseString(string s) {
     	string ret(s.size(), '\0');
     	for (int i = 0, j = s.size()-1; i <= j; ++i, --j) {
     		ret[i] = s[j];
     		ret[j] = s[i];
     	}
     	return ret;
    }
};