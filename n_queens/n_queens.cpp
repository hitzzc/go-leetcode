#include <vector>
#include <string>
#include <iostream>
using namespace std;

bool isValid(vector<int> &solution, int depth,  int val){
	for (int j = 0; j < solution.size(); ++j){
		if (solution[j] == val || j-depth == solution[j]-val || j-depth == val-solution[j]){
			return false;
		}
	}
	return true;
}

vector<vector<string>> solveNQueens(int n) {
	vector<vector<int>> results;
	vector<int> solution;
	int depth = -1;
	int max_depth = n - 1;
	vector<int> next_index(n, 0);
	while (true){
		while (depth < max_depth){
			int child_depth = depth + 1;
			bool found = false;
			while (true){
				if (next_index[child_depth] == n){
					next_index[child_depth] = 0;
					break;
				}
				int index = next_index[child_depth];
				++next_index[child_depth];
				if (isValid(solution, child_depth, index)){
					solution.push_back(index);
					found = true;
					break;
				}
			}
			if (!found){
				break;
			}
			++depth;
		}
		if (depth == max_depth){
			results.push_back(solution);
		}
		if (depth == -1){
			break;
		}
		--depth;	
		solution.pop_back();
	}
	vector<vector<string>> rets;
	vector<string> i_to_string;
	for (int i = 0; i < n; ++i){
		string s;
		for (int j = 0; j < n; ++j){
			if (j == i){
				s += "Q";
			}else{
				s += ".";
			}
		}
		i_to_string.push_back(s);
	}
	for (int i = 0; i < results.size(); ++i){
		vector<string> ret;
		for (int j = 0; j < results[i].size(); ++j){
			ret.push_back(i_to_string[results[i][j]]);
		}
		rets.push_back(ret);
	}
	return rets;
}	

int main(){
	for (auto i: solveNQueens(4)){
		for (auto j: i){
			cout << j << endl;
		}
	}
	return 0;
}