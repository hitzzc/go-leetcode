#include <stack>
using namespace std;

class Queue {
public:
    stack<int> push_stack;
    stack<int> pop_stack;
    // Push element x to the back of queue.
    void push(int x) {
        push_stack.push(x);
    }

    // Removes the element from in front of queue.
    void pop(void) {
        if (pop_stack.empty()){
            while (!push_stack.empty()) {
                pop_stack.push(push_stack.top());
                push_stack.pop();
            }
        }
        pop_stack.pop();
    }

    // Get the front element.
    int peek(void) {
        if (pop_stack.empty()){
            while (!push_stack.empty()) {
                pop_stack.push(push_stack.top());
                push_stack.pop();
            }
        }
        return pop_stack.top();
    }

    // Return whether the queue is empty.
    bool empty(void) {
        return pop_stack.empty() && push_stack.empty();
    }
};