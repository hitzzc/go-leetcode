class Solution {
public:
    int maxProfit(int k, vector<int>& prices) {
        if (prices.size() < 2) return 0;
        if (k > prices.size()/2) {
            int ret = 0;
            for (int i = 0; i < prices.size()-1; ++i) {
                if (prices[i+1] - prices[i] > 0) ret += prices[i+1] - prices[i];
            }
            return ret;
        }
        vector<vector<int>> dp(k+1, vector<int>(prices.size()+1, 0));
        for (int i = 1; i <= k; ++i) {
            int maxV = dp[i-1][0] - prices[0];
            for (int j = 1; j <= prices.size(); ++j) {
                dp[i][j] = max(dp[i][j-1], maxV + prices[j-1]);
                if (j < prices.size() && dp[i-1][j] - prices[j] > maxV) {
                    maxV = dp[i-1][j] - prices[j];
                }
            }
        }
        return dp[k][prices.size()];
    }
};