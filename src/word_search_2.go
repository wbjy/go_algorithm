/**
 *单词搜索 II
 *
 * 解题思路：字典树+dfs+回朔法
 *
 * 先将words中的单词构建成字典树
 * 然后再用回溯+dfs去遍历整个二维数组
 * 
 */

type TrieNode struct {
    children [26]*TrieNode
    word string
}


func findWords(board [][]byte, words []string) []string {
    root := &TrieNode{}
    for _,w := range words{
        node := root
        for _,v := range w{
            if node.children[v-'a'] == nil{
                node.children[v-'a'] = new(TrieNode)
            }
            node = node.children[v-'a']
        }
        node.word = w
    }
    
    res := make([]string, 0)
    for i:=0;i<len(board);i++{
        for j:=0;j<len(board[0]);j++{
            dfs(i,j, board, root, &res)
        }
    }
    
    return res
    
}

func dfs(i,j int, board [][]byte, node *TrieNode, res *[]string){
    if i<0||j<0||i==len(board)||j==len(board[0]){
        return
    }
    
    c := board[i][j]
    if c =='#'||node.children[c-'a'] ==nil{
        return
    }
    
    node = node.children[c-'a']
    
    if node.word != ""{
        *res = append(*res, node.word)
        node.word =""
    }
    board[i][j] = '#'
    dfs(i, j+1, board, node, res)
    dfs(i+1, j, board, node, res)
    dfs(i-1, j, board, node, res)
    dfs(i, j-1, board, node, res)
    board[i][j] =c
}