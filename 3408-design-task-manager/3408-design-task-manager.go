

type Task struct {
	userId   int
	taskId   int
	priority int
	version  int
}

type MaxHeap []Task

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool {
	if h[i].priority != h[j].priority {
		return h[i].priority > h[j].priority
	}
	return h[i].taskId > h[j].taskId
}

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Task))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type TaskManager struct {
	pq         *MaxHeap
	curVersion map[int]int
	active     map[int]bool
	user       map[int]int
	priority   map[int]int
}

func Constructor(tasks [][]int) TaskManager {
	h := &MaxHeap{}
	heap.Init(h)

	tm := TaskManager{
		pq:         h,
		curVersion: make(map[int]int),
		active:     make(map[int]bool),
		user:       make(map[int]int),
		priority:   make(map[int]int),
	}

	for _, t := range tasks {
		u, id, p := t[0], t[1], t[2]
		tm.curVersion[id]++
		v := tm.curVersion[id]
		heap.Push(h, Task{u, id, p, v})
		tm.active[id] = true
		tm.user[id] = u
		tm.priority[id] = p
	}
	return tm
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	this.curVersion[taskId]++
	v := this.curVersion[taskId]
	heap.Push(this.pq, Task{userId, taskId, priority, v})
	this.active[taskId] = true
	this.user[taskId] = userId
	this.priority[taskId] = priority
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	this.curVersion[taskId]++
	v := this.curVersion[taskId]
	heap.Push(this.pq, Task{this.user[taskId], taskId, newPriority, v})
	this.priority[taskId] = newPriority
	this.active[taskId] = true
}

func (this *TaskManager) Rmv(taskId int) {
	this.active[taskId] = false
}

func (this *TaskManager) ExecTop() int {
	for this.pq.Len() > 0 {
		cur := heap.Pop(this.pq).(Task)
		if !this.active[cur.taskId] {
			continue
		}
		if cur.version != this.curVersion[cur.taskId] {
			continue
		}
		this.active[cur.taskId] = false
		return cur.userId
	}
	return -1
}
