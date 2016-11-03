class Solution {
public:
    vector<pair<int, int>> kSmallestPairs(vector<int>& nums1, vector<int>& nums2, int k) {
        priority_queue<pair<int, int>, vector<pair<int, int>>, CMP> q;
        vector<pair<int, int>> result;
        for (int i = 0; i < min(k, (int)nums1.size()); ++i) {
        	for (int j = 0; j < min(k, (int)nums2.size()); ++j) {
        		auto node = make_pair(nums1[i], nums2[j]);
        		if (q.size() == k && node.first + node.second < q.top().first + q.top().second) {
        			q.pop();
        		}
        		if (q.size() < k) q.push(node);
        	}
        }
        while (!q.empty()) {
        	result.push_back(q.top());
        	q.pop();
        }
        return result;
    }

    struct CMP {
		bool operator () (const pair<int, int>& a, const pair<int, int>& b) {
	    	return a.first + a.second < b.first + b.second;
	    }
	};
};