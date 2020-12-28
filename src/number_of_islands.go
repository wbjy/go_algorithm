/**
 *岛屿数量
 *
 * 解题思路：dfs
 * 
 */
func numIslands(grid [][]byte) int {
    if grid == nil {
        return 0
    }
    res := 0
    for i := 0;i<len(grid);i++{
        for j := 0;j<len(grid[0]);j++{
            if grid[i][j] == '1'{
                res ++
                dfs(grid, i, j)
                
            }
        }
    }
    return res
}

func dfs(grid [][]byte, x, y int) {
    if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == '0' {
        return 
    }

    grid[x][y] = '0'
    dfs(grid, x - 1, y)
    dfs(grid, x, y - 1)
    dfs(grid, x +1, y)
    dfs(grid, x, y+1)

}