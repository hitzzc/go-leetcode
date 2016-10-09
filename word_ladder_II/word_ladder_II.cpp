#include <vector>
#include <string>
#include <map>
#include <queue>
#include <unordered_set>
#include <algorithm>
using namespace std;

class Solution {
public:
	map<string, int> distance;
	vector<vector<string>> results;
public:
	vector<vector<string>> findLadders(string beginWord, string endWord, unordered_set<string> &wordList) {
		bfs(beginWord, endWord, wordList);
		vector<string> path;
		path.push_back(endWord);
		dfs(endWord, path, wordList);
		return results;
	}

	void bfs(string beginWord, string endWord, unordered_set<string> &wordList){
		distance[beginWord] = 1;
		queue<string> q;
		q.push(beginWord);
		string word;
		bool found = false;
		while (!q.empty() && !found){
			word = q.front();
			q.pop();
			for (int i = 0; i < word.size() && !found; ++i){
				string temp = word;
				for (char j = 'a'; j <= 'z' && !found; ++j){
					temp[i] = j;
					if (temp==endWord) {
						found = true;
						distance[temp] = distance[word]+1;
					}else if (wordList.find(temp)!=wordList.end() && distance.find(temp)==distance.end()){
						distance[temp] = distance[word]+1;
						q.push(temp);
					}
				}
			}
		}
		return;
	}

	void dfs(string currentWord, vector<string>& path, unordered_set<string> &wordList){
		if (distance[currentWord]==1){
			vector<string> res = path;
			reverse(res.begin(), res.end());
			results.push_back(res);
			return;
		}
		for (int i = 0; i < currentWord.size(); ++i){
			string temp = currentWord;
			for (char j = 'a'; j <= 'z'; ++j){
				temp[i] = j;
				if (distance.find(temp)!=distance.end() && distance[temp]+1==distance[currentWord]){
					path.push_back(temp);
					dfs(temp, path, wordList);
					path.pop_back();
				}
			}
		}
		return;
	}
};