class binary_index_tree {
	vector<int> array;
public:
	binary_index_tree(int n): array(n+1, 0) {}
	int lowbit(int x) {
		return x & (-x);
	}
	int sum(int idx) {
		int ret = 0;
		while (idx > 0) {
			ret += array[idx];
			idx -= lowbit(idx);
		}
		return ret;
	}
	void add(int idx, int val) {
		while (idx < array.size()) {
			array[idx] += val;
			idx += lowbit(idx);
		}
	}
}; 

class Solution {
public:
	vector<int> countSmaller(vector<int>& nums) {
		vector<int> sorted = nums;
		sort(sorted.begin(), sorted.end());
		unordered_map<int, int> mapping;
		for (int i = 0; i < sorted.size(); ++i) {
			mapping[sorted[i]] = i+1;
		}
		binary_index_tree my_bit(nums.size());
		vector<int> ret(nums.size());
		for (int i = nums.size()-1; i >=0; --i) {
			ret[i] = my_bit.sum(mapping[nums[i]]-1);
			my_bit.add(mapping[nums[i]], 1);
		}
		return ret;
	}
};