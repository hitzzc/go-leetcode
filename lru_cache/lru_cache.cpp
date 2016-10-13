#include <unordered_map>
#include <iostream>
using namespace std;

struct node {
	int key_;
	int value_;
	node* pre_;
	node* next_;
	node(int key, int val): key_(key), value_(val), pre_(NULL), next_(NULL) {}
};

class node_list {
public:
	node* head_;
	node* tail_;
	int length_;
	node_list(): head_(NULL), tail_(NULL), length_(0) {}

	void push(node* n){
		if (head_==NULL && tail_==NULL)
			head_ = tail_ = n;
		else {
			n->next_ = head_;
			n->pre_ = NULL;
			head_->pre_ = n;
			head_ = n;
		}
		++length_;
	}

	void pop(){
		if (head_==NULL && tail_==NULL) return;
		if (head_ == tail_) {
			head_ = tail_ = NULL;
			return;
		}
		tail_ = tail_->pre_;
		tail_->next_ = NULL;
		--length_;
	}

	void update(node* n){
		if (n==NULL || n==head_) return;
		if (n->pre_!=NULL) n->pre_->next_ = n->next_;
		if (n->next_!=NULL) n->next_->pre_ = n->pre_;
		if (n==tail_) tail_ = tail_->pre_;
		head_->pre_ = n;
		n->next_ = head_;
		n->pre_ = NULL;
		head_ = n;
		return;
	}

	void print() {
		node* p = head_;
		while (p!=NULL){
			cout << "print\tkey\t" << p->key_ << "\tvalue\t" << p->value_ << endl;
			p = p->next_;
		}
	}

	int len(){
		return length_;
	}
};

class LRUCache{
public:
	int capacity_;
	unordered_map<int, node*> mapping;
	node_list lru_list;
public:
	LRUCache(int capacity): capacity_(capacity) {}
	
	int get(int key) {
		if (mapping.find(key)!=mapping.end()) {
			lru_list.update(mapping[key]);
			return mapping[key]->value_;
		}
		else
			return -1;
	}
	
	void set(int key, int value) {
		if (mapping.find(key)!=mapping.end()){
			mapping[key]->value_ = value;
			lru_list.update(mapping[key]);
			return;
		}
		if (lru_list.len()==capacity_){
			node* tail = lru_list.tail_;
			lru_list.pop();
			mapping.erase(mapping.find(tail->key_));
			delete tail;
		}
		node* new_node = new node(key, value);
		lru_list.push(new_node);
		mapping[key] = new_node;
	}
};

int main(){
	LRUCache cache(2);
	cache.set(2, 1);
	cache.lru_list.print();
	cache.set(1, 1);
	cache.lru_list.print();
	cache.get(2);
	cache.lru_list.print();
	cache.set(4,1);
	cache.lru_list.print();
	cache.get(1);
	cache.lru_list.print();
	cache.get(2);
	cache.lru_list.print();
	return 0;
}
