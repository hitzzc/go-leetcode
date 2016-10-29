class Solution {
public:
    int nthSuperUglyNumber(int n, vector<int>& primes) {
        vector<int> dp(n, 1);
        vector<int> idxs(primes.size(), 0);
        for (int i = 1; i < n; ++i) {
        	dp[i] = INT_MAX;
        	for (int idx = 0; idx < idxs.size(); ++idx) {
        			dp[i] = min(dp[i], primes[idx]*dp[idxs[idx]]);
        	}
        	for (int idx = 0; idx < idxs.size(); ++idx) {
        		if (dp[i] == primes[idx] * dp[idxs[idx]]) ++idxs[idx];
        	}
        }
        return dp.back();
    }
};