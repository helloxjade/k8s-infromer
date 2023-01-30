package main

import (
	"fmt"
	"k8s.io/client-go/tools/cache"
)
type pod struct {
	Name string
	Value float64
}
func newPod(name string, value float64)pod{
	return pod{Name: name,Value:value }
}

func main() {
   df:=cache.NewDeltaFIFOWithOptions(cache.DeltaFIFOOptions{KeyFunction: podKeyFunc})
   pod1:=newPod("pod1",1)
   pod2:=newPod("pod2",1)
   pod3:=newPod("pod3",1)
   df.Add(pod1)
   df.Add(pod2)
   df.Add(pod3)
   fmt.Println(df.List())


   pod1.Value=1.1
   df.Update(pod1)
   df.Delete(pod1)

   //执行三次 先入先出
   df.Pop(func(i interface{}) error {
	   fmt.Printf("%T\n", i)
	   for _,deltas:=range i.(cache.Deltas){
		   fmt.Println(deltas.Type ,":" ,deltas.Object.(pod).Name,"value:",deltas.Object.(pod).Value)
		   switch deltas.Type{
		   case cache.Added:
		   	    fmt.Println("added")
		   case cache.Updated:
		   	    fmt.Println("updated")
		   case cache.Deleted:
		   	    fmt.Println("deleted")
		   }

	   }


	   return nil
   })
	df.Pop(func(i interface{}) error {
		fmt.Printf("%T\n", i)
		for _,deltas:=range i.(cache.Deltas){
			fmt.Println(deltas.Type ,":" ,deltas.Object.(pod).Name,"value:",deltas.Object.(pod).Value)

		}
		return nil
	})
	df.Pop(func(i interface{}) error {
		fmt.Printf("%T\n", i)
		for _,deltas:=range i.(cache.Deltas){
			fmt.Println(deltas.Type ,":" ,deltas.Object.(pod).Name)
		}
		return nil
	})

}

func podKeyFunc(obj interface{})(string,error) {
	return obj.(pod).Name,nil
}
