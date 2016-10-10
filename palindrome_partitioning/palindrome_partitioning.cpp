#include <vector>
using namespace std;

class Solution {
public:
	enum
	{
		INIT = 0,
		YES = 1,
		NO = 2,
	};
	vector<vector<int> > matrix;
public:
	vector<vector<string>> partition(string s) {
		matrix = vector<vector<int> >(s.size(), vector<int>(s.size(), INIT));
		vector<vector<string>> res;
		vector<string> path;
		DFS(s, 0, res, path);
		return res;
	}

	void DFS(string s, int start, vector<vector<string>>& res, vector<string>& path){
		if (start==s.size()){
			res.push_back(path);
			return;
		}
		for (int i = start; i < s.size(); ++i){
			if (matrix[start][i]==INIT){
				if (is_partition(s, start, i))
					matrix[start][i] = YES;
				else
					matrix[start][i] = NO;
			}
			if (matrix[start][i]==YES){
				path.push_back(s.substr(start, i-start+1));
				DFS(s, i+1, res, path);
				path.pop_back();
			}
		}
	}

	bool is_partition(string s, int i, int j){
		while (i<j){
			if (s[i]!=s[j]) return false;
			++i;
			--j;
		}
		return true;
	}
};