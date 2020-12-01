/**
 *课程表---dfs
 *
 * 解题思路：首先理解题意中的numCourses，这个和课程科目有关，如果要上numCourses门课，那么每一门课程都要上，不能上不了的课，
 * 也就是不能出现有环的情况，例如：上课程1需要先上课程0，上课程0需要上课程1，[[1,0],[0,1]]
 *
 * 了解了问题的本质之后就是准备解法，也就是遍历整个numCourses，然后判断每个i的前置中是否有环
 *
 * 方法：深度搜索，广度搜索
 *             
 */
func canFinish(numCourses int, prerequisites [][]int) bool {
    var dfs func(u int)
    var edg = make([][]int, numCourses)
    var visit = make([]int, numCourses)
    var flag = true
    for _, nums := range prerequisites {
        edg[nums[1]] = append(edg[nums[1]], nums[0])
    }
    dfs = func(u int) {
        visit[u] = 1
        for _, v := range edg[u] {
            if visit[v] == 0 {
                dfs(v)
                if !flag {  //表示已经检查出有环，直接返回
                    return
                }
            }else if visit[v] == 1{ // 表示有环，直接返回
                flag = false
                return
            }
        }

        visit[u] = 2  // 表示u已经检查完毕，u的前置都不需要检查
    }
    for i := 0; i<numCourses && flag;i++{
        if visit[i] == 0 {
            dfs(i)
        } 
    }
    return flag
}