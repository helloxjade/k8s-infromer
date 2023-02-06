package main

import (
	"fmt"
	"k8s-informer/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
	"log"
)

func main() {
	client:=lib.InitClient()
    podListWatch:=cache.NewListWatchFromClient(client.CoreV1().RESTClient(),"pods",
   	"default", fields.Everything())
   	podsList,err:=podListWatch.List(metav1.ListOptions{})
   	if err != nil {
   		log.Fatal("the err is :",err)
		return
	}
	fmt.Printf("%T\n",podsList)
    podList:=podsList.(*v1.PodList)
    for _,pod := range podList.Items {
     fmt.Println(pod.Name)
	}
}


