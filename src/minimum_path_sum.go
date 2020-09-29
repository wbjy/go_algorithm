/**
 *最小路径和
 *
 * 动态规划
 */
func minPathSum(grid [][]int) int {
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if i==0&&j==0{
                continue
            }
            
            tmp := 1<<31
            if i>0 {
                tmp = min(grid[i-1][j], tmp)
            }
            if j >0{
                tmp = min(grid[i][j-1], tmp)
            }
            
            grid[i][j]+= tmp
        }
    }
    
    return grid[len(grid)-1][len(grid[0])-1]
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}