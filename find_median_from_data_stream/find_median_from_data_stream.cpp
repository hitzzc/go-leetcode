#include <vector>
#include <iostream>
#include <priority_queue>
#include <functional>
using namespace std;

class Heap {
private:
	vector<int> array;
	int (*compare) (int, int);
	void adjust(int idx) {
		int largest = idx;
		if (2*idx+1 < array.size() && compare(array[2*idx+1], array[idx]))
			largest = 2*idx+1;
		if (2*idx+2 < array.size() && compare(array[2*idx+2],array[largest]))
			largest = 2*idx+2;
		if (largest != idx) {
			int tmp = array[idx];
			array[idx] = array[largest];
			array[largest] = tmp;
			adjust(largest);
		}
		return;
	}
public:
	Heap(int (*compare) (int, int)): compare(compare) {}
	void push(int val) {
		array.push_back(val);
		for (int i = array.size()/2 - 1; i >= 0; --i){
			adjust(i);
		}
	}
	void print() {
		for (auto val: array) 
			cout << "print\t" << val << endl;
	}
	int pop() {
		int ret = array[0];
		array[0] = array.back();
		array.pop_back();
		adjust(0);
		return ret;
	}
	int empty() {
		return array.empty();
	}
	int size() {
		return array.size();
	}
	int top() {
		return array[0];
	}
};

int max(int i, int j) {
	if (i > j) return true;
	return false;
}

int min(int i, int j) {
	if (i < j) return true;
	return false;
}

class MedianFinder {
	priority_queue<int> max_heap;
	priority_queue<int, vector<int>, greater<int>> min_heap;
public:
    // Adds a number into the data structure.
    void addNum(int num) {
        if (max_heap.empty() || num <= max_heap.top()) {
        	max_heap.push(num);
        }else {
        	min_heap.push(num);
        }
        if (max_heap.size() > min_heap.size()+1){
        	int top = max_heap.top();
        	max_heap.pop();
        	min_heap.push(top);
        }
        if (min_heap.size() > max_heap.size()) {
        	int top = min_heap.top();
        	min_heap.pop();
        	max_heap.push(top);
        }
        return;
    }

    // Returns the median of current data stream
    double findMedian() {
    	return max_heap.size() == min_heap.size() ? (max_heap.top()+min_heap.top())/2.0 : max_heap.top();
    }
};

int main() {
	MedianFinder finder;
	finder.addNum(-1);
	finder.addNum(-2);
	finder.addNum(-3);
	finder.addNum(-4);
	finder.addNum(-5);
	return 0;
}
