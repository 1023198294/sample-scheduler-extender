package controller

import (
	schedulerapi "k8s.io/kubernetes/pkg/scheduler/api"
	"log"
)

// It'd better to only define one custom priority per extender
// as current extender interface only supports one single weight mapped to one extender
// and also it returns HostPriorityList, rather than []HostPriorityList
const (
	// lucky priority gives a random [0, schedulerapi.MaxPriority] score
	// currently schedulerapi.MaxPriority is 10
	luckyPrioMsg = "pod %v/%v is lucky to get score %v\n"
)
var term int
// it's webhooked to pkg/scheduler/core/generic_scheduler.go#PrioritizeNodes()
// you can't see existing scores calculated so far by default scheduler
// instead, scores output by this function will be added back to default scheduler
func prioritize(args schedulerapi.ExtenderArgs) *schedulerapi.HostPriorityList {
	pod := args.Pod
	nodes := args.Nodes.Items
	//每轮优先级+1
	hostPriorityList := make(schedulerapi.HostPriorityList, len(nodes))
	for i, node := range nodes {
		term = (term + 1)%(schedulerapi.MaxPriority + 1)
		score := term
		log.Printf(luckyPrioMsg, pod.Name, pod.Namespace, score)
		hostPriorityList[i] = schedulerapi.HostPriority{
			Host:  node.Name,
			Score: score,
		}
	}

	return &hostPriorityList
}
