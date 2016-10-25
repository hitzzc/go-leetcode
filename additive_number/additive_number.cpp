#include <string>
using namespace std;

class Solution {
public:
    bool isAdditiveNumber(string num) {
        return DFS(num, 0, 0, 0, 0);
    }

    bool DFS(string& num, int start, int length, int first, int second) {
    	if (start == num.size()) {
    		if (length > 2) return true;
    		else return false;
    	}
    	int current_num = 0;
    	for (int i = start; i < num.size(); ++i) {
    	    if (num[start] == '0' && i != start) break;
    		current_num = 10*current_num + num[i]-'0';
    		if (length == 0) first = current_num;
    		else if (length == 1) second = current_num;
    		else if (current_num != first+second) continue;
    		else {
    			first = second;
    			second = current_num;
    		}
    		if (DFS(num, i+1, length+1, first, second)) return true;
    	}
    	return false;
    }
};