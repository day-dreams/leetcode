# leetcode


- 链表
    - `medium`[链表求和](002.go)，补齐元素即可
    - `medium`[删除倒数第N个元素](019.go)，维护快慢指针
    - `easy`[合并有序链表](021.go)，可以递归
    - `hard`[合并K个排序链表](023.go)
        - 复用21题的代码，两两合并`NlogK`
        - 朴素的比较，一次扫K个值，每次拿出找到首元素最小的链表x，拿出x的首元素，继续下一次扫描 `NK`
        - 利用其他算法，比如堆排序，`NlogN`
