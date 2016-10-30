class Solution {
public:
    int longestIncreasingPath(vector<vector<int>>& matrix) {
        int row = matrix.size();
        int col = row ? matrix[0].size() : 0;
        vector<vector<int>> path(row, vector<int>(col, 0));
        vector<vector<bool>> visited(row, vector<bool>(col, 0));
        int ret = 0;
        for (int i = 0; i < row; ++i) {
            for (int j = 0; j < col; ++j) {
                ret = max(ret, helper(matrix, path, row, col, i , j, visited));
            }
        }
        return ret;
    }
    
    int helper(vector<vector<int>>& matrix, vector<vector<int>>& path, const int row, const int col, int r, int c, vector<vector<bool>>& visited) {
        if (path[r][c] > 0) return path[r][c];
        visited[r][c] = true;
        int ret = 0;
        if (r > 0 && !visited[r-1][c] && matrix[r][c] < matrix[r-1][c]) {
            ret = max(ret, helper(matrix, path, row, col, r-1, c, visited));
        }
        if (r < row-1 && !visited[r+1][c] && matrix[r][c] < matrix[r+1][c]) {
            ret = max(ret, helper(matrix, path, row, col, r+1, c, visited));
        }
        if (c > 0 && !visited[r][c-1] && matrix[r][c] < matrix[r][c-1]) {
            ret = max(ret, helper(matrix, path, row, col, r, c-1, visited));
        }
        if (c < col-1 && !visited[r][c+1] && matrix[r][c] < matrix[r][c+1]) {
            ret = max(ret, helper(matrix, path, row, col, r, c+1, visited));
        }
        visited[r][c] = false;
        path[r][c] = ret+1;
        return path[r][c];
    }
};