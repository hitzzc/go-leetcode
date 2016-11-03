class Solution {
public:
    string decodeString(string s) {
    	int i = 0;
    	return decode(s, i);
    }

    string decode(string& s, int& idx) {
    	string ret;
    	while (idx < s.size() && s[idx] != ']') {
    		if (s[idx] < '0' || s[idx] > '9') {
    			ret += s[idx++];
    		} else {
    			int num = 0;
    			while (s[idx] != '[') {
    				num = 10*num + s[idx++] - '0';
    			}
    			++idx;
    			string tmp = decode(s, idx);
    			while (num-- > 0) {
    				ret += tmp;
    			}
    			++idx;
    		}
    	}
    	return ret;
    }
};