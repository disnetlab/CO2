package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	var kubeconfig *string
	var deleteRate, numOfOptional,i int
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	fmt.Println("Please Enter Deletion Percentage: ")
	_, err = fmt.Scanf("%d", &deleteRate)
	if err != nil {
		panic(err)
	}

	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	fmt.Printf("Listing pods in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := podClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		if d.Labels["app"] == "optional" {
			numOfOptional++
		}
	}

	for _, d := range list.Items {
		if d.Labels["app"] == "optional" && i < int((numOfOptional*deleteRate)/100) {
			err = clientset.CoreV1().Pods(apiv1.NamespaceDefault).Delete(d.Name,&metav1.DeleteOptions{})
			if err != nil{
				panic(err)
			}
			fmt.Printf("%s Pod deleted successfully!\n",d.Name)
			i++
		}
		if i >= int((numOfOptional*deleteRate)/100){
			break
		}
	}
}

