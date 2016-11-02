class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        unordered_set<int> s, visited;
        for (auto& num: nums1) {
        	s.insert(num);
        }
        vector<int> res;
        for (auto& num: nums2) {
        	if (s.count(num) && visited.count(num) == 0) res.push_back(num);
        	visited.insert(num);
        }
        return res;
    }
};