class Solution {
public:
    int numSquares(int n) {
        if ( n<=0 ) return 0;
        static vector<int> dp(1, 0);
        if (dp.size() >n) return dp[n];
        for (int i = dp.size(); i <= n; ++i) {
        	int m = n;
        	for (int j = 1; i-j*j>=0; ++j) {
        		m = min(m, dp[i-j*j]+1);
        	}
        	dp.push_back(m);
        }
        return dp[n];
    }
    int min(int i, int j) {
    	if (i < j) return i;
    	return j;
    }
};