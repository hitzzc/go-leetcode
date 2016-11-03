class Solution {
public:
    int kthSmallest(vector<vector<int>>& matrix, int k) {
        int left = matrix[0][0], right = matrix.back().back();
        while (left < right) {
        	int mid = left + (right - left)/2;
        	int cnt = 0;
        	for (auto& vec: matrix) {
        		cnt += upper_bound(vec.begin(), vec.end(), mid) - vec.begin();
        	}
        	if (cnt < k) left = mid + 1;
        	else right = mid;
        }
        return left;
    }
};