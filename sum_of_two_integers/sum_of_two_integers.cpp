class Solution {
public:
    int getSum(int a, int b) {
    	int carry = 0;
        while (b != 0) {
        	carry = a & b;
        	a ^= b;
        	b = carry << 1;
        }
        return a;
    }
};