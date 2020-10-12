/** 
 *字典树（前缀树）
 *
 * Insert：存入单词时，遍历单词的字母来插入树中，并在单词的结尾字母标记表示这是一个单词的结尾（这个标记很重要）
 * Search：查找时遍历单词，看是否所有字母都在字典树中，并且单词结尾是否有单词标记
 * StartsWith：前缀查找，看传入的前缀是否全在字典树中
 */

type Trie struct {
    name [26]*Trie
    isWord bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    for _,v := range word{
        if this.name[v-'a'] == nil{
            this.name[v-'a'] = new(Trie)
            fmt.Println(this.name)
        }
        this = this.name[v-'a']
        
    }
    
    this.isWord = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    for _,v :=range word{
        if this.name[v-'a'] == nil{
            return false
        }
        
        this = this.name[v-'a']
    }
    
    if this.isWord == false{
        return false
    }
    
    return true
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    for _,v :=range prefix{
        if this.name[v-'a'] == nil{
            return false
        }
        
        this = this.name[v-'a']
    }
    return true
}