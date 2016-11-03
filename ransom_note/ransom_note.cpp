class Solution {
public:
    bool canConstruct(string ransomNote, string magazine) {
        vector<int> count(26, 0);
        for (auto& ch: magazine) {
        	++count[ch-'a'];
        }
        for (auto& ch: ransomNote) {
        	if (count[ch-'a'] <= 0) return false;
        	--count[ch-'a'];
        }
        return true;
    }
};