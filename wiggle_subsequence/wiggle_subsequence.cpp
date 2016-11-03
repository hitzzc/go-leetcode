class Solution {
public:
    int wiggleMaxLength(vector<int>& nums) {
    	if (nums.size() < 2) return nums.size();
        int inc_length = 1, dec_length = 1;
        int pre_inc_v = nums[0], pre_dec_v = nums[0];
        for (int i = 1; i < nums.size(); ++i) {
        	if (nums[i] < pre_inc_v && dec_length <= inc_length) {
        		dec_length = inc_length + 1;
        		pre_dec_v = nums[i];
        	}
        	if (nums[i] > pre_dec_v && inc_length <= dec_length) {
        		inc_length = dec_length + 1;
        		pre_inc_v = nums[i];
        	}
        	if (nums[i] > pre_inc_v) pre_inc_v = nums[i];
        	if (nums[i] < pre_dec_v) pre_dec_v = nums[i];
        }
        return max(inc_length, dec_length);
    }
};