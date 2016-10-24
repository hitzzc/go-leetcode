#include <vector>
using namespace std;

class Solution {
public:
    int lengthOfLIS(vector<int>& nums) {
        if (nums.size() == 0) return 0;
        int max = 1;
        vector<int> fn(nums.size(), 0);
        fn[0] = 1;
        for (int i = 0; i < nums.size(); ++i) {
        	for (int j = 0; j < i; ++j) {
        		if (nums[i] > nums[j] && fn[j]+1 > fn[i])
        			fn[i] = fn[j]+1;
        	}
        	if (fn[i] > max) max = fn[i];
        }
        return fn.back();
    }

    int lengthOfLIS2(vector<int>& nums) {
    	if (nums.size() < 2) return nums.size();
       	vector<int> vec(1, nums[0]);
       	for (int i = 1; i < nums.size(); ++i) {
       		if (nums[i] > vec.back()) {
       			vec.push_back(nums[i]);
       			continue;
       		}else if (nums[i] == vec.back()) {
       			continue;
       		}else{
	       		int pos = binary_search(vec, nums[i]);
	       		vec[pos] = nums[i];
       		}
       	}
       	return vec.size();
    }

    int binary_search(vector<int>& nums, int target) {
		int start = 0;
		int end = nums.size();
		int idx = -1;
		int mid;
	   	while (start < end) {
	   		mid = start + (end-start)/2;
	   		if (nums[mid] == target) {
	   			return mid;
	   		}else if (nums[mid] < target) {
	   			idx = mid+1;
	   			start = mid+1;
	   		}else{
	   			idx = mid;
	   			end = mid;
	   		}
		}
		return idx;
	}
};