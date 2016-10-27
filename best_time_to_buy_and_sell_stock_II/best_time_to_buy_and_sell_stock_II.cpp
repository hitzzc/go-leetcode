class Solution {
public:
    int maxProfit(vector<int>& prices) {
    	if (prices.size() == 0) return 0;
   		int ret = 0;
   		for (int i = 0; i < prices.size()-1; ++i) {
   			if (prices[i+1] > prices[i]) ret += prices[i+1] - prices[i];
   		}
   		return ret;
   	}
};