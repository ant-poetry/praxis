
// Below is the interface for Iterator, which is already defined for you.
// **DO NOT** modify the interface for Iterator.

#include<vector>

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
private:
	bool peek_;
	int  next_;
public:
	PeekingIterator(const vector<int>& nums) : Iterator(nums) {
		this->peek_ = false;
	}

	int peek() {
		if(this->peek_){
			return this->next_;
		}
		this->peek_ = true;
		return this->next_ = Iterator::next();
	}
	
	int next() {
		if(this->peek_){
			this->peek_ = false
			return this->next_;
		}
		return Iterator::next();
	}

	bool hasNext() const {
		return this->peek_ || Iterator::hasNext();
	}
};