class Solution {
public:
    int countRangeSum(vector<int>& nums, int lower, int upper) {
        vector<long> sums(nums.size()+1, 0);
        for (int i = 0; i < nums.size(); ++i) {
        	sums[i+1] = sums[i] + nums[i];
        }
        return count_when_merging_sort(sums, 0, sums.size(), lower, upper);
    }
    int count_when_merging_sort(vector<long>& array, int start, int end, int lower, int upper) {
    	if (end-start <= 1) return 0;
    	int mid = start + (end-start)/2;
    	int j = mid, k = mid, t = mid;
    	int count = count_when_merging_sort(array, start, mid, lower, upper) + count_when_merging_sort(array, mid, end, lower, upper);
    	vector<long> cache(end-start);
    	for (int i = start, r = 0; i < mid; ++i, ++r) {
    		while (k < end && array[k]-array[i] <= upper) ++k;
    		while (j < end && array[j]-array[i] < lower) ++j;
    		while (t < end && array[t] < array[i]) cache[r++] = array[t++];
    		cache[r] = array[i];
    		count += k - j; 
    	}
    	for (int i = start; i < t; ++i) {
    		array[i] = cache[i-start];
    	}
    	return count;
    }
};
