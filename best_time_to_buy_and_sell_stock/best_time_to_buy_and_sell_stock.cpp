class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if (prices.size() == 0) return 0;
        int min = prices[0];
        int max = 0;
        for (int i = 1; i < prices.size(); ++i) {
        	int diff = prices[i] - min;
        	if (diff > max) max = diff;
        	if (prices[i] < min) min = prices[i];
        }
        return max;
    }
};