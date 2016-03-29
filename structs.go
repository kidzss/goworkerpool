package main

type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
}

//Payload指消息本身内容
//单词原意是载荷
//这里可以根据业务来任意定义
type Payload struct {
}

//Job是Payload的上层封装
type Job struct {
	Payload Payload
}

//JobQueue里存储的是所有待处理的任务
var JobQueue chan Job

//全局的workerpool
//每一个worker都会把他的jobchannel push到这个chan里
//并且会监听他们自己的这条jobchannel
//所以这个channel里存的东西可以认为是它们本身，虽然不是。。
var WorkerPool chan chan Job

//具体执行任务的Worker
type Worker struct {
	//workerpool是对全局的workerpool的一个引用
	//和c里的存储全局指针有点类似，初始化之后会把自己放到池子里
	WorkerPool chan chan Job
	//JobChannel是对全局的JobChannel的一个引用
	JobChannel chan Job
	WorkerID   int
	quit       chan bool
}
