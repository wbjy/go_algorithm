/**
 *除法求值
 *
 * 解题思路：dfs
 */
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
    cal := make(map[string]map[string]float64)
    for k, v :=range equations {
        if _, ok := cal[v[0]]; !ok {
            cal[v[0]] = make(map[string]float64)
        }
        cal[v[0]][v[1]] = values[k]
        cal[v[0]][v[0]] = 1

        if _,ok := cal[v[1]];!ok{
            cal[v[1]] = make(map[string]float64)
        }
        cal[v[1]][v[1]] = 1
        cal[v[1]][v[0]] = 1/values[k]
    }
    res := make([]float64, len(queries))
    for k, v := range queries {
        router := make(map[string]int)
        one := dfs(cal, router, v[0], v[1])
        res[k] = one
        if _, ok := cal[v[0]]; !ok {
            cal[v[0]] = make(map[string]float64)
        }
        cal[v[0]][v[1]] = one
        if _,ok := cal[v[1]]; !ok{
            cal[v[1]] = make(map[string]float64)
        }
        cal[v[1]][v[0]] = 1/one
    }

    return res
}

func dfs(cal map[string] map[string]float64, router map[string]int, from, target string) float64{
    //防止重复访问形成死循环
    if _, ok := router[from]; ok {
        return -1.0
    }
    router[from] = 1
    val,ok := cal[from]
    if !ok {
            return -1.0
    } 
    if finalV, ok := val[target];ok{
        return finalV
    }
    for k,v := range val {
        if v == -1.0{
            return -1.0
        }
        res := dfs(cal, router, k, target)
        if res == -1.0{
            continue
        }else{
            return res*v
        }
    }
    return -1.0
}