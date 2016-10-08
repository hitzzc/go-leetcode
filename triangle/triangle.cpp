#include <vector>
using namespace std;

int minimumTotal(vector<vector<int>>& triangle) {
	int length = triangle.size();
	if (length == 0) return 0;
	if (length == 1) return triangle[0][0];
	vector<int> dp(length);
	dp[0] = triangle[0][0];
	int min;
	bool first = true;
	for (int i = 1; i < length; ++i){
		for (int j = i; j >= 0; --j){
			if (j != 0 && (dp[j-1] < dp[j] || j == i)){
				dp[j] = dp[j-1];
			}
			dp[j] += triangle[i][j];
			if (i == length-1 && (first || dp[j] < min)){
				first = false;
				min = dp[j];
			}
		}
	}
	return min;
}
