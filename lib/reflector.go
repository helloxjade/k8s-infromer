package lib

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

func reflector(){
	client := InitClient()
	//因为没有gvk 所以需要用coreV1()
	podListWatch := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods",
		"default", fields.Everything())
	store:=cache.NewStore(cache.MetaNamespaceKeyFunc)
	df:=cache.NewDeltaFIFOWithOptions(
		cache.DeltaFIFOOptions{
			KeyFunction: cache.MetaNamespaceKeyFunc,
			KnownObjects: store,//存进indexer,是个threadsafemap
		})
	rf:=cache.NewReflector(podListWatch,&v1.Pod{},df,0)

	ch:=make(chan struct{})
	//启动df，不会阻塞
	go func() {
		rf.Run(ch)
	}()
	//取出deltafifo中的值，这里只有pod类型 ，好比informer不断消费队列
	for {
		df.Pop(func(i interface{}) error {
			for _,delta:=range i.(cache.Deltas){
				fmt.Println(delta.Type,":",delta.Object.(*v1.Pod).Name,
					":",delta.Object.(*v1.Pod).Status.Phase)
				switch delta.Type {
				case cache.Sync,cache.Added:
					store.Add(delta.Object)
				case cache.Updated:
					store.Update(delta.Object)
				case cache.Deleted:
					store.Delete(delta.Object)
				}
			}
			return nil
		})
	}
}
