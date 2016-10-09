#include <queue>
#include <unordered_set>
#include <map>
#include <string>
using namespace std;

int ladderLength(string beginWord, string endWord, unordered_set<string>& wordList) {
	if (beginWord == endWord) 
		return 1;
	if (distance(beginWord, endWord)==1)
		return 2;
	map<string, int> dist;
	dist[beginWord] = 1;
	queue<string> q;
	q.push(beginWord);
	while (!q.empty()){
		string word = q.front();
		q.pop();
		if (word == endWord) return dist[word];
		if (wordList.size() < 26*word.size()){
			for (auto& w: wordList){
				if (dist.find(w)==dist.end() && distance(word, w)==1){
					dist[w] = dist[word] + 1;
					q.push(w);
				}
			}
		}else{
			for (int i = 0; i < word.size(); ++i){
				string temp = word;
				for (int j = 'a'; j <= 'z'; ++j){
					temp[i] = j;
					if (wordList.find(temp)!=wordList.end() && dist.find(temp)==dist.end()){
						dist[temp] = dist[word] + 1;
						q.push(temp);
					}
				}
			}
		}
	}
	return 0;
}

int distance(string dist, string src){
	int count = 0;
	for (int i = 0; i < dist.size(); ++i){
		if (dist[i]!=src[i]) ++count;
	}
	return count;
}