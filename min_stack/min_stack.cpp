#include <stack>
using namespace std;

class MinStack {
public:

    MinStack() {
    
    }
    
    void push(int x) {
        normal_stack.push(x);
        if (min_stack.empty() || x <= min_stack.top()) min_stack.push(x);
        return;
    }
    
    void pop() {
        int t = normal_stack.top();
        normal_stack.pop();
        if (t == min_stack.top()) min_stack.pop();
        return;
    }
    
    int top() {
        return normal_stack.top();
    }
    
    int getMin() {
        return min_stack.top();
    }
private:
    stack<int> normal_stack;
    stack<int> min_stack;
};