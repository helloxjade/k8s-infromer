package lib

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"log"
)

func ListAndWatch(){
		client := InitClient()
		//因为没有gvk 所以需要用coreV1()
		podListWatch := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "pods",
			"default", fields.Everything())
		podsList, err := podListWatch.List(metav1.ListOptions{})
		if err != nil {
			log.Fatal("the err is :", err)
			return
		}
		fmt.Printf("%T\n", podsList)
		podList := podsList.(*v1.PodList)
		for _, pod := range podList.Items {
			fmt.Println(pod.Name)
		}

		watcher, err := podListWatch.Watch(metav1.ListOptions{})
		if err != nil {
			log.Fatal("the err is :", err)
			return
		}

		for {
			select {
			case v, ok := <-watcher.ResultChan():
				if ok {
					fmt.Println(v.Type, ":", v.Object.(*v1.Pod).Name)
				}
			}
		}
}
