class Solution {
public:
    int maxRotateFunction(vector<int>& A) {
    	if (A.size() == 0) return 0;
        int sum = 0;
        int pre = 0;
        for (int i = 0; i < A.size(); ++i) {
        	sum += A[i];
        	pre += i * A[i];
        }
        int ret = pre;
        int len = A.size();
        for (int i = 0; i < len-1; ++i) {
        	pre = pre - sum + len*A[i];
        	ret = max(ret, pre);
        }
        return ret;
    }
};