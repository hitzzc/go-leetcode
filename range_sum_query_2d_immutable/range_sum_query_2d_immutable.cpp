#include <vector>
using namespace std;

class NumMatrix {
	vector<vector<int>> records;
public:
    NumMatrix(vector<vector<int>> &matrix) {
    	int row = matrix.size();
    	int col = row == 0 ? 0 : matrix[0].size();
        records = vector<vector<int>>(row+1, vector<int>(col+1, 0));
        for (int i = 0; i < row; ++i) {
        	for (int j = 0; j < col; ++j) {
        		records[i+1][j+1] = records[i][j+1] + records[i+1][j] - records[i][j] + matrix[i][j];
        	}
        }
    }

    int sumRegion(int row1, int col1, int row2, int col2) {
    	return records[row2+1][col2+1] - records[row2+1][col1] - records[row1][col2+1] + records[row1][col1];
    }
};