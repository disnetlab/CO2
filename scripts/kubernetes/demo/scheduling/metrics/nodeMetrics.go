package main

import (
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"path/filepath"
)

func main() {
	var kubeconfig *string
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

	mc, err := metrics.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nodeMetrics, err := mc.MetricsV1beta1().NodeMetricses().List(metav1.ListOptions{})

	//podMetrics, err := mc.MetricsV1beta1().PodMetricses(metav1.NamespaceAll).List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	//for _, podMetric := range podMetrics.Items {
	//	podContainers := podMetric.Containers
	//	for _, container := range podContainers {
	//		cpuQuantity, ok := container.Usage.Cpu().AsInt64()
	//		memQuantity, ok := container.Usage.Memory().AsInt64()
	//		if !ok {
	//			return
	//		}
	//		msg := fmt.Sprintf("Container Name: %s \n CPU usage: %d \n Memory usage: %d", container.Name, cpuQuantity, memQuantity)
	//		fmt.Println(msg)
	//	}
	//}
	for _, nodeMetric := range nodeMetrics.Items {
		node := nodeMetric.Name
		cpuQuantity := nodeMetric.Usage.Cpu().MilliValue()
		memQuantity := nodeMetric.Usage.Memory().MilliValue()
		msg := fmt.Sprintf("Container Name: %s \n CPU usage: %v \n Memory usage: %v", node, cpuQuantity, memQuantity)
		fmt.Println(msg)
	}
}

