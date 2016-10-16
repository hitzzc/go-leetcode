#include <unordered_map>
using namespace std;

class TrieNode {
public:
    TrieNode(): is_word_(false) {}
public:
    bool is_word_;
    unordered_map<char, TrieNode*> children;
};

class Trie {
public:
    Trie() {
        root = new TrieNode();
    }

    // Inserts a word into the trie.
    void insert(string word) {
        if (word.size() <= 0) return;
        TrieNode* node = root;
        for (int i = 0; i < word.size(); ++i){
            if (node->children.find(word[i]) == node->children.end()){
                node->children[word[i]] = new TrieNode();
            }
            node = node->children[word[i]];
        }
        node->is_word_ = true;
        return;
    }

    // Returns if the word is in the trie.
    bool search(string word) {
        return retrieve(word, false);
    }

    // Returns if there is any word in the trie
    // that starts with the given prefix.
    bool startsWith(string prefix) {
        return retrieve(prefix, true);
    }

private:
    TrieNode* root;
private:
    bool retrieve(string& key, bool prefix){
        if (key.size() <= 0) return false;
        TrieNode* node = root;
        for (int i = 0; i < key.size(); i++){
            if (node->children.find(key[i]) == node->children.end())
                return false;
            node = node->children[key[i]];
        }
        return prefix ? true : node->is_word_;
    }
};