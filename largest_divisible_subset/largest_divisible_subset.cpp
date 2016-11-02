class Solution {
public:
    vector<int> largestDivisibleSubset(vector<int>& nums) {
        if (nums.size() == 0) return {};
    	sort(nums.begin(), nums.end());
        vector<int> dp(nums.size(), 1), pre(nums.size());
        for (int i = 0; i < pre.size(); ++i) {
        	pre[i] = i;
        }
        int max_v = 1, k = 0;
        for (int i = 1; i < dp.size(); ++i) {
        	for (int j = i-1; j >= 0; --j) {
        		if (nums[i] % nums[j] != 0) continue;
        		if (dp[i] < dp[j]+1) {
        			pre[i] = j;
        			dp[i] = dp[j]+1;
        		}
        		if (dp[i] > max_v) {
        			max_v = dp[i];
        			k = i;
        		}
        	}
        }
        vector<int> result;
        while (k != pre[k]) {
        	result.push_back(nums[k]);
        	k = pre[k];
        }
        result.push_back(nums[k]);
        return result;
    }
};