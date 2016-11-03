class Solution {
public:
    char findTheDifference(string s, string t) {
        vector<int> vec(26, 0);
        for (auto& c: s) {
        	++vec[c-'a'];
        }
        char ret;
        for (auto& c: t) {
        	if (vec[c-'a'] == 0) {
        		ret = c;
        		break;
        	}
        	--vec[c-'a'];
        }
        return ret;
    }
};