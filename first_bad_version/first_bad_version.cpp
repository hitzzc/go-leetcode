bool isBadVersion(int version);

class Solution {
public:
    int firstBadVersion(int n) {
		int i = 1;
		int j = n;
		int ret;
		while (i <= j) {
			int mid = i + (j-i)/2;
			if (isBadVersion(mid)) {
				j = mid;
				ret = mid;
			}else {
				i= mid+1;
				ret = mid+1;
			}
		}
		return ret;   
    }
};