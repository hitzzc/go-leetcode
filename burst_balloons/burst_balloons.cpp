class Solution {
public:
    int maxCoins(vector<int>& nums) {
        nums.insert(nums.begin(), 1);
        nums.push_back(1);
        vector<vector<int>> matrix(nums.size(), vector<int>(nums.size(), 0));
        for (int k = 2; k < nums.size(); ++k) {
        	for (int low = 0; low < nums.size()-k; ++low) {
        		int high = low+k;
        		for (int i = low+1; i < high; ++i) {
        			matrix[low][high] = max(matrix[low][high], nums[low]*nums[i]*nums[high] + matrix[low][i] + matrix[i][high]);
        		}
        	}
        }
        return matrix[0][matrix.size()-1];
    }
};