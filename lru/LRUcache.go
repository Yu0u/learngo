package lru

type LRUCache struct {
	// 缓存的size
	size int
	// 缓存的容量
	capacity int
	// 缓存
	cache map[int]*Node
	// 头，尾节点
	head, tail *Node
}

type Node struct {
	// key 和 value 都是int类型
	key, value int
	// 双向链表
	prev *Node
	next *Node
}

func initNode(key, value int) *Node {
	// 将节点初始化
	return &Node{
		key:   key,
		value: value,
	}
}

// 构造函数
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*Node{},
		head:     initNode(0, 0),
		tail:     initNode(0, 0),
		capacity: capacity,
	}
	// 头指针的下一个指向尾指针
	l.head.next = l.tail
	// 尾指针的前一个指向头指针
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	// 查找是否存在key，如果不存在返回-1
	if _, ok := l.cache[key]; !ok {
		return -1
	}
	// 将访问过的数据移动到head
	node := l.cache[key]
	l.moveToHead(node)
	// 返回key对应的value
	return node.value
}

func (l *LRUCache) Put(key int, value int) {
	// 根据key查找map中是否存在value
	// 如果不存在就将数据存到head
	// 如果存在key就将可以对应的值修改为新的值，然后放到head
	if _, ok := l.cache[key]; !ok {
		node := initNode(key, value)
		l.cache[key] = node
		// 将新增的数据放到head
		l.addToHead(node)
		l.size++
		// 如果size大于capacity就将末尾的数据移除
		if l.size > l.capacity {
			removed := l.removeTail()
			// 用delete函数删除map存放的数据
			delete(l.cache, removed.key)
			l.size--
		}
	} else {

		node := l.cache[key]
		node.value = value
		l.moveToHead(node)
	}
}

// 双向链表的操作
func (l *LRUCache) moveToHead(node *Node) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *Node) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) removeTail() *Node {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
