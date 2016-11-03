class Solution {
public:
    int getMoneyAmount(int n) {
        vector<vector<int>> dp(n+1, vector<int>(n+1, 0));
        for (int gap = 1; gap < n; ++gap) {
        	for (int low = 1; low <= n-gap; ++low) {
        		int high = low + gap;
        		dp[low][high] = INT_MAX;
        		for (int i = low; i < high; ++i) {
        			dp[low][high] = min(dp[low][high], i + max(dp[low][i-1], dp[i+1][high]));
        		}
        	}
        }
        return dp[1][n];
    }
};