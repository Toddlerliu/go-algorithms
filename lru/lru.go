package main

// 缓存淘汰算法：
// 思想：如果数据最近被访问过，那么将来被访问的几率也很高
// 当限定的空间已存满数据时，应当把最久没有被访问到的数据淘汰，在尾部的数据就是最近最久未访问的数据

//  实现方式：
//1.用一个数组来存储数据，给每一个数据项标记一个访问时间戳，每次插入新数据项的时候，先把数组中存在的数据项的时间戳自增，
//  并将新数据项的时间戳置为0并插入到数组中。每次访问数组中的数据项的时候，将被访问的数据项的时间戳置为0。当数组空间已满时，将时间戳最大的数据项淘汰。
//2.利用一个链表来实现，每次新插入数据的时候将新数据插到链表的头部；每次缓存命中（即数据被访问），则将数据移到链表头部；
//  那么当链表满的时候，就将链表尾部的数据丢弃。
//3.利用链表和hashmap。插入数据项：如果新数据项在链表中存在（命中），则把该节点移到链表头部，如果不存在，
//  则新建一个节点，放到链表头部，若缓存满了，则把链表最后一个节点删除即可；访问数据：如果数据项在链表中存在，
//  则把该节点移到链表头部，否则返回-1。
