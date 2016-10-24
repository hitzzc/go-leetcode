#include <string>
using namespace std;

class Solution {
public:
    string getHint(string secret, string guess) {
        int secret_array[10] = {0};
        int bull = 0;
        for (int i = 0; i < secret.size(); ++i) {
        	if (secret[i] == guess[i]) {
        		++bull;
        		continue;
        	}
        	++secret_array[secret[i]-'0'];
        }
        int cow = 0;
        for (int i = 0; i < guess.size(); ++i) {
        	if (secret[i] == guess[i]) {
        		continue;
        	}
        	if (secret_array[guess[i]-'0'] !=0 ){
        		++cow;
        		--secret_array[guess[i]-'0'];
        	}
        }
        return to_string(bull) + "A" + to_string(cow) + "B";
    }
};