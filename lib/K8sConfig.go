package lib

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

func K8sRestConfig()*rest.Config {
	config,err:=clientcmd.BuildConfigFromFlags("","config")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return config
}

//初始化k8s client-go客户端
func InitClient()*kubernetes.Clientset{
	client,err:=kubernetes.NewForConfig(K8sRestConfig())
	if err != nil{
		log.Fatal(err)
		return nil
	}
	return client
}