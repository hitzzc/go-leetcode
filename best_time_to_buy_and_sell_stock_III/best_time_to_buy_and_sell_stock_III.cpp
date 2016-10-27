class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if (prices.size() == 0) return 0;
        vector<int> fn1(prices.size(), 0);
        vector<int> fn2(prices.size(), 0);
        int minV = prices[0];
        for (int i = 1; i < prices.size(); ++i) {
        	fn1[i] = max(fn1[i-1], prices[i]-minV);
        	minV = min(minV, prices[i]);
        }
        int maxV = prices[prices.size()-1];
        for (int i = prices.size()-2; i >= 0; --i) {
        	fn2[i] = max(fn2[i+1], maxV-prices[i]);
        	maxV = max(maxV, prices[i]);
        }
        int ret = 0;
        for (int i = 0; i < fn1.size(); ++i) {
        	ret = max(ret, fn1[i] + fn2[i]);
        }
        return ret;
    }
};