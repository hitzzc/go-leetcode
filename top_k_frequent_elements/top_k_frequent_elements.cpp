class Solution {
public:
    vector<int> topKFrequent(vector<int>& nums, int k) {
        unordered_map<int, int> mapping;
        for (auto num: nums) {
        	++mapping[num];
        }
        vector<vector<int>> freq(nums.size()+1);
        for (auto m: mapping) {
        	freq[m.second].push_back(m.first);
        }
        vector<int> result;
        for (int i = freq.size()-1; i >=0 && k > 0; --i) {
        	for (auto num: freq[i]) {
        		result.push_back(num);
        		--k;
        		if (k == 0) break;
        	}
        }
        return result;
    }
};