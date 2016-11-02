class Solution {
public:
    bool isPowerOfFour(int num) {
        static int mask = 0b01010101010101010101010101010101;
        
        //edge case
        if (num<=0) return false;
        
        // there are multiple bits are 1
        if ((num & num-1) != 0) return false;
        
        // check which one bit is zero, if the place is 1 or 3 or 5 or 7 or 9...,
        // then it is the power of 4
        if ((num & mask) != 0) return true;
        return false;
    }
};