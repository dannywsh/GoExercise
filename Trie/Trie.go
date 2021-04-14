package main

type Trie struct {
	isEnd bool      //结束标志
	next  [26]*Trie //26个英文小写字母的指针
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{} //构建返回一个空Trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] != nil { //如果指向不为空，则说明已插入该字符
			this = this.next[v-'a'] //直接跳到下一个字符
		} else { //如果为空
			this.next[v-'a'] = new(Trie) //则新增Trie节点
			this = this.next[v-'a']      //并指向该节点
		}
	}
	this.isEnd = true //将最后的节点标志位置为结束
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this != nil { //如果节点不为空
			this = this.next[v-'a'] //跳到下一个节点
		} else { //如果为空，则搜索失败
			return false
		}
	}
	if this != nil {
		return this.isEnd //最后节点为结束位则搜索成功
	} else {
		return false
	}
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this != nil {
			this = this.next[v-'a']
		} else {
			return false
		}
	}
	if this != nil {
		return true //最后节点不为空则存在前缀
	} else {
		return false
	}
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
