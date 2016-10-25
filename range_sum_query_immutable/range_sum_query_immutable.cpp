#include <vector>
using namespace std;

class NumArray {
	vector<int> fn;
public:
    NumArray(vector<int> &nums): fn(vector<int>(nums.size()+1)) {
        fn[0] = 0;
        for (int i = 0; i < nums.size(); ++i) {
        	fn[i+1] = fn[i] + nums[i];
        }
    }

    int sumRange(int i, int j) {
        return fn[j] - fn[i-1];
    }
};