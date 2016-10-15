class Solution {
public:
	uint32_t reverseBits(uint32_t n) {
		uint32_t ret = 0;
		for (int i = 0; i < 32; ++i){
			ret = ret*2 + (n&0x1);
			n /= 2;
		}
		return ret;
	}
};