class Solution {
public:
    int findNthDigit(int n) {
        long len = 1;
        long count = 9;
        long start = 1;
        while (n > len*count) {
            n -= len*count;
        	start += count;
        	++len;
        	count *= 10;
        }
        return to_string(start + (n-1)/len)[(n-1)%len] - '0';
    }
};