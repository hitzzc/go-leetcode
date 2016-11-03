class Solution {
public:
    Solution(vector<int> nums): nums(nums), current(nums) {
        srand(time(NULL));
    }
    
    /** Resets the array to its original configuration and return it. */
    vector<int> reset() {
        return current = nums;
    }
    
    /** Returns a random shuffling of the array. */
    vector<int> shuffle() {
    	int i = current.size();
    	while (--i > 0) {
    		int j = rand() % (i+1);
    		swap(current[i], current[j]);
    	}
    	return current;
    }
private:
    vector<int> nums, current;
};

/**
 * Your Solution object will be instantiated and called as such:
 * Solution obj = new Solution(nums);
 * vector<int> param_1 = obj.reset();
 * vector<int> param_2 = obj.shuffle();
 */