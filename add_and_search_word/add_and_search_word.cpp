#include <unordered_map>
using namespace std;

struct TrieNode {
	bool is_word_;
	unordered_map<char, TrieNode*> children;
	TrieNode(): is_word_(false) {}
};

class WordDictionary {
public:
	TrieNode* root;
	WordDictionary() {
		root = new TrieNode();
	}
    // Adds a word into the data structure.
    void addWord(string word) {
        if (word.size() <= 0) return;
        TrieNode* node = root;
        for (int i = 0; i < word.size(); ++i){
        	if (node->children.find(word[i]) == node->children.end()){
        		node->children[word[i]] = new TrieNode();
        	}
        	node = node->children[word[i]];
        }
        node->is_word_ = true;
    }

    // Returns if the word is in the data structure. A word could
    // contain the dot character '.' to represent any one letter.
    bool search(string word) {
        return retrieve(word, root, 0, word.size());
    }

    bool retrieve(string& word, TrieNode* node, int idx, int length) {
    	if (idx == length) return node->is_word_;
    	if (word[idx] == '.') {
    		if (node->children.size() == 0) {
    			return false;
    		}
    		for (auto& kv: node->children) {
    			if (retrieve(word, kv.second, idx+1, length)) {
    				return true;
    			}
    		}
    		return false;
    	}else if (node->children.find(word[idx]) != node->children.end()) {
    		return retrieve(word, node->children[word[idx]], idx+1, length);
    	}
    	return false;
    }
};
