class Solution {
public:
    int maxProfit(vector<int>& prices) {
        if (prices.size() < 2) return 0;
       	int buy = -prices[0];
       	int sell = -1;
       	int cold = 0;
       	int tmp;
       	for (int i = 1; i < prices.size(); ++i) {
       		tmp = buy;
       		buy = max(tmp, cold-prices[i]);
       		cold = max(cold, sell);
       		sell = tmp+prices[i];
       	}
       	return max(sell, cold);
    }
};