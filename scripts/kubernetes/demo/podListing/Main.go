package main

import (
	"fmt"
	"strconv"

	"flag"
	apiv1 "k8s.io/api/core/v1"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {
	var kubeconfig *string
	var numOfOptional, numOfMandatory int
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

	fmt.Printf("Please enter number of optional and mandotory pods!\n")
	_, err = fmt.Scanf("%d %d", &numOfOptional, &numOfMandatory)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(numOfOptional+ numOfMandatory)

	for i := 0;i<numOfOptional;i++{
		newPod := getPodObject("optional",strconv.Itoa(i))
		newPod, err = clientset.CoreV1().Pods(newPod.Namespace).Create(newPod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Optional Pod created successfully...")
	}

	for i := 0;i< numOfMandatory;i++{
		newPod := getPodObject("mandatory",strconv.Itoa(i))
		newPod, err = clientset.CoreV1().Pods(newPod.Namespace).Create(newPod)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Mandatory Pod created successfully...")
	}

	//newPod := getPodObject()
	//newPod, err = clientset.CoreV1().Pods(newPod.Namespace).Create(newPod)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Pod created successfully...")


	podClient := clientset.CoreV1().Pods(apiv1.NamespaceDefault)

	fmt.Printf("Listing pods in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := podClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		fmt.Printf("* name: %s status: %s\n", d.Name, d.Status.Phase)
	}

}

func getPodObject(label, i string) *core.Pod {
	return &core.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:       label+"pod"+i,
			Namespace: "default",
			Labels: map[string]string{
				"app": label,
			},
		},
		Spec: core.PodSpec{
			Containers: []core.Container{
				{
					Name:            "nginx",
					Image:           "nginx:1.12",
					ImagePullPolicy: core.PullIfNotPresent,
					Ports: []apiv1.ContainerPort{
						{
							Name:          "http",
							Protocol:      apiv1.ProtocolTCP,
							ContainerPort: 80,
						},
					},
				},
			},
		},
	}
}

