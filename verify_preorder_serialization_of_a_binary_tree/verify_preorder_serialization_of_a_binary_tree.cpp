class Solution {
public:
    bool isValidSerialization(string preorder) {
    	int idx = 0;
        return helper(preorder, idx) ? idx == preorder.size() ? true : false : false;
    }

    bool helper(string& preorder, int& idx) {
    	if (idx == preorder.size()) return false;
    	int start = idx;
    	while (idx < preorder.size() && preorder[idx] != ',') ++idx;
    	if (idx == preorder.size()) {
    		return preorder.substr(start) == "#" ? true : false;
    	}
    	++idx;
    	if (preorder.substr(start, idx-start-1) == "#") 
    		return true;
    	if (!helper(preorder, idx)) return false;
    	return helper(preorder, idx);
    }
};