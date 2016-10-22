class Solution {
public:
    vector<string> addOperators(string num, int target) {
        vector<string> results;
        DFS(num, target, 0, "", 0, " ", 0, results);
        return results;
    }

    void DFS(string& num, int target, int idx, string solution, long long current_val, string pre_op, long long pre_val, vector<string>& results) {	
    	if (current_val == target && idx == num.size()) {
    		results.push_back(solution);
    		return;
    	}
    	if (idx == num.size()) return;
    	string n;
    	long long v = 0;
    	for (int i = idx; i < num.size(); ++i) {
    		if (n =="0") {
    			return;
    		}
    		n += num[i];
    		v = 10*v + num[i] - '0';
    		if (solution.size() == 0) {
    			DFS(num, target, i+1, n, v, " ", 0, results);
    		}else {
    			DFS(num, target, i+1, solution + "+" + n, current_val + v, "+", v, results);
    			DFS(num, target, i+1, solution + "-" + n, current_val - v, "-", v, results);
    			if (pre_op == "+") {
    				DFS(num, target, i+1, solution + "*" + n, current_val-pre_val+pre_val*v, pre_op, pre_val*v, results);
    			}else if (pre_op == "-") {
    				DFS(num, target, i+1, solution + "*" + n, current_val+pre_val-pre_val*v, pre_op, pre_val*v, results);
    			}else {
    				DFS(num, target, i+1, solution + "*" + n, current_val*v, "*", v, results);
    			}
    		}
    	}
    	return;
    }
};