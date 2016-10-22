// Below is the interface for Iterator, which is already defined for you.
// **DO NOT** modify the interface for Iterator.
class Iterator {
    struct Data;
	Data* data;
public:
	Iterator(const vector<int>& nums);
	Iterator(const Iterator& iter);
	virtual ~Iterator();
	// Returns the next element in the iteration.
	int next();
	// Returns true if the iteration has more elements.
	bool hasNext() const;
};


class PeekingIterator : public Iterator {
	int cache;
	bool is_peeked;
public:
	PeekingIterator(const vector<int>& nums) : Iterator(nums), is_peeked(false) {
	    
	}

    // Returns the next element in the iteration without advancing the iterator.
	int peek() {
        if (!is_peeked) {
        	cache = Iterator::next();
        	is_peeked = true;
        }
        return cache;
	}

	// hasNext() and next() should behave the same as in the Iterator interface.
	// Override them if needed.
	int next() {
		if (is_peeked) {
			is_peeked = false;
			return cache;
		}
		return Iterator::next();
	}

	bool hasNext() const {
	    if (is_peeked) {
	    	return true;
	    }
	    return Iterator::hasNext();
	}
};