class Solution {
public:
    string reverseVowels(string s) {
    	unordered_set<char> vowels = {'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'};
        string ret(s.size(), '\0');
        for (int i = 0, j = s.size()-1; i <= j; ++i, --j) {
        	while (i < s.size()-1 && vowels.count(s[i]) == 0) {
        		ret[i] = s[i];
        		++i;
        	}
        	while (j >= 0 && vowels.count(s[j]) == 0) {
        		ret[j] = s[j];
        		--j;
        	}
        	if (i > j) break;
        	ret[i] = s[j];
        	ret[j] = s[i]; 
        }
        return ret;
    }
};