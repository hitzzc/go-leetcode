class Solution {
public:
    int findDuplicate(vector<int>& nums) {
  		int i = 0;
  		int j = 0;
  		while (true) {
  			i = nums[i];
  			j = nums[nums[j]];
  			if (i == j) break;
  		}
  		j = 0;
  		while (i != j) {
  			i = nums[i];
  			j = nums[j];
  		}
  		return i;
    }
};