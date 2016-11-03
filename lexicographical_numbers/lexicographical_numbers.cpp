class Solution {
public:
    vector<int> lexicalOrder(int n) {
        vector<int> result;
        for (int i = 1; i <= n && i <= 9; ++i) {
        	result.push_back(i);
        	helper(i, n, result);
        }
        return result;
    }

    void helper(int num, int n, vector<int>& result) {
    	for (int i = 0; i <= 9; ++i) {
    		int tmp = num*10 + i;
    		if (tmp > n) break;
    		result.push_back(tmp);
    		helper(tmp, n, result);
    	}
    	return;
    }
};