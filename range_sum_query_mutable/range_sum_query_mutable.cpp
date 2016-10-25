class NumArray {
    vector<int> records;
    vector<int> ori;
public:
    NumArray(vector<int> &nums): records(vector<int>(nums.size()+1, 0)), ori(nums) {
        for (int i = 0; i < nums.size(); ++i) {
            add(i+1, nums[i]);
        }
    }

    void update(int i, int val) {
        int diff = val - ori[i];
        ori[i] = val;
        add(i+1, diff);
    }

    int sumRange(int i, int j) {
        return sum(j+1) - sum(i);
    }

    void add(int x, int diff) {
        while (x < records.size()) {
            records[x] += diff;
            x += lowbit(x);
        }
    }

    int sum(int x) {
        int ret = 0;
        while (x > 0) {
            ret += records[x];
            x -= lowbit(x);
        }
        return ret;
    }

    int lowbit(int n) {
        return n&-n;
    }
};