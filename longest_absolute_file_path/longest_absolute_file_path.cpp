class Solution {
public:
    int lengthLongestPath(string input) {
    	unordered_map<int, int> depth;
    	int level = 0, res = 0;
    	for (int i = 0; i < input.size(); ++i) {
    		int start = i;
    		while (i < input.size() && input[i] != '\n' && input[i] != '\t') ++i;
    		if (i == input.size() || input[i] == '\n') {
    			string name = input.substr(start, i-start);
    			if (name.find('.') != string::npos) {
    				res = max(res, depth[level] + (int)name.size());
    			}else {
    				++level;
    				depth[level] = depth[level-1] + (int)name.size() + 1;	 
    			}
    			level = 0;
    		}else {
    			++level;
    		}
    	}
    	return res;
    }
};