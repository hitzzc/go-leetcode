class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        int j = 0;
        int tmp;
        for (int i = 0; i < nums.size(); ++i){
        	if (nums[i] != 0){
        		tmp = nums[i];
        		nums[i] = nums[j];
        		nums[j] = tmp;
        		j++;
        	}
        }
        return;
    }
};