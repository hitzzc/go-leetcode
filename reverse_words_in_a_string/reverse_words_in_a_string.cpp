#include <string>
#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
	void reverseWords(string &s) {
		vector<string> words;
		bool found = false;
		int start = -1;
		const char space = ' ';
		for (int i = 0; i < s.size(); ++i){
			if (found && s[i]==space){
				words.push_back(s.substr(start, i-start));
				found = false;
			}else if (!found && s[i]!=space){
				start = i;
				found = true;
			}
		}
		if (found) words.push_back(s.substr(start));
		s.clear();
		if (words.empty()) return;
		for (int i = words.size()-1; i > 0; --i){
			s += words[i];
			s.push_back(space);
		}
		s += words[0];
		return;
	}
};