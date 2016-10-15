class Solution {
public:
    int hammingWeight(uint32_t n) {
        int x = 1;
        int count = 0;
        for (int i = 0; i < 32; ++i){
        	if (n & x) count++;
        	x <<= 1;
        }
        return count;
    }
};