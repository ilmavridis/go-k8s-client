package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	gvr := map[string]schema.GroupVersionResource{
		"Pods": {
			Group:    "",
			Version:  "v1",
			Resource: "pods",
		},
		"Services": {
			Group:    "",
			Version:  "v1",
			Resource: "services",
		},
		"Deployments": {
			Group:    "apps",
			Version:  "v1",
			Resource: "deployments",
		},
		"DaemonSets": {
			Group:    "apps",
			Version:  "v1",
			Resource: "daemonsets",
		},
		"ReplicaSets": {
			Group:    "apps",
			Version:  "v1",
			Resource: "replicasets",
		},
		"StatefulSets": {
			Group:    "apps",
			Version:  "v1",
			Resource: "statefulsets",
		},
		"Jobs": {
			Group:    "batch",
			Version:  "v1",
			Resource: "jobs",
		},
		"CronJobs": {
			Group:    "batch",
			Version:  "v1",
			Resource: "cronjobs",
		},
	}

	// Create the client to talk with the K8s cluster
	dynamicClient := connectToK8s()

	// Set the namespace
	nsSelected := setNamespace(dynamicClient)
	fmt.Println("\n Selected namespace:", nsSelected)
	fmt.Println(" Current time: ", time.Now())

	for resources, groupVersion := range gvr {
		fmt.Printf("\n\n  *** %s *** \n", resources)

		// Provide additional info if the resource is a pod
		if resources == "Pods" {

			runningPods := getPods("status.phase=Running", dynamicClient, groupVersion, nsSelected)

			nonRunningPods := getPods("status.phase!=Running", dynamicClient, groupVersion, nsSelected)

			totalPods := len(runningPods.Items) + len(nonRunningPods.Items)

			if totalPods > 0 { // If there are pods, print pods info
				fmt.Println(" In total there are", totalPods, "Pods in namespace", nsSelected, ": \n")
				fmt.Println("\t\t\t ---- Running Pods :)  ----")
				podPrinter(runningPods)
				fmt.Println("\n\t\t\t ---- NON Running Pods  :(  ----")
				podPrinter(nonRunningPods)
			} else {
				fmt.Printf("   none \n")
			}
		} else { // If the resource is not a pod
			runningResources, err := dynamicClient.Resource(gvr[resources]).Namespace(nsSelected).List(context.Background(), v1.ListOptions{})

			if err != nil {
				log.Printf("Error getting %s: %v\n", resources, err)
			}

			for _, runningResource := range runningResources.Items {
				fmt.Println(runningResource.Object["metadata"].(map[string]interface{})["name"])
			}

			if len(runningResources.Items) == 0 {
				fmt.Printf("   none \n")
			}
		}
	}
}

// Authenticate and connect with the cluster.
// There are 2 different options:
//	1: Application runs outside the K8s cluster
//	2: Application runs inside the K8s cluster (in comment)
func connectToK8s() dynamic.Interface {
	// ----- Option 1 -----
	// Outside the K8s ->
	// Using kubeconfig file to connect to the API
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalln("Error getting user home dir:", err)
	}

	kubeConfigPath := filepath.Join(userHomeDir, ".kube", "config")
	fmt.Printf("Kubeconfig file: %s\n", kubeConfigPath)

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatalln("Error getting kubeconfig file:", err)
	}
	// <- Outside the K8s

	// ----- Option 2 -----
	// Inside the K8s ->
	// Using rest to get the configuration from pod our application running in
	/*
		kubeConfig, err := rest.InClusterConfig()
		if err != nil {
			log.Fatalln("Error getting kubeconfig:", err)
		}
	*/
	// <- Inside the K8s

	// Create the dynamic client
	client, err := dynamic.NewForConfig(kubeConfig)
	if err != nil {
		log.Fatalln("Error creating client:", err)
	}

	return client
}

func setNamespace(dynamicClient dynamic.Interface) string {
	gvrns := map[string]schema.GroupVersionResource{
		"Namespaces": {
			Group:    "",
			Version:  "v1",
			Resource: "namespaces",
		},
	}

	var nsSelected string
	var nsExists bool

	if len(os.Args) > 1 {
		nsSelected = os.Args[1]
	}

	namespaces, err := dynamicClient.Resource(gvrns["Namespaces"]).List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatalln("Failed to get namespaces:", err)
	}

	fmt.Println("\nAvailable Namespaces =>")
	for i, namespace := range namespaces.Items {
		fmt.Printf("[%d] %s\n", i, namespace.GetName())
		if nsSelected == namespace.GetName() {
			nsExists = true
		}
	}

	// If a namespace is not provided as argument, it shows the avaliable namespaces and
	// awaits for user input
	if nsExists == false {
		var nsSelectedID int
		fmt.Println("\n-> Please select an id from the namespaces above.")

		if _, err := fmt.Scanf("%d", &nsSelectedID); err != nil {
			log.Printf("Failed to read namespace id: %v", err)
		}
		nsSelected = namespaces.Items[nsSelectedID].GetName()
	}
	return nsSelected
}

func getPods(fieldSelector string, dynamicClient dynamic.Interface, groupVersion schema.GroupVersionResource, nsSelected string) *unstructured.UnstructuredList {
	pods, err := dynamicClient.Resource(groupVersion).Namespace(nsSelected).List(context.Background(), v1.ListOptions{
		FieldSelector: fieldSelector,
	})
	if err != nil {
		log.Printf("Error getting %s: %v\n", fieldSelector, err)
	}

	return pods
}

func podPrinter(pods *unstructured.UnstructuredList) {
	var i int = 1

	if len(pods.Items) != 0 {
		fmt.Println("\t Name \t|\t CreationTimeStamp \t|\t ContainerImage")
		for _, pod := range pods.Items {
			for _, podSpec := range pod.Object["spec"].(map[string]interface{})["containers"].([]interface{}) {
				fmt.Printf(" [%d] \t %s \t %s \t %v \n", i, pod.GetName(), pod.GetCreationTimestamp(), podSpec.(map[string]interface{})["image"])
				i++
			}
		}
	} else {
		fmt.Printf("\t\t\t\t   none \n")
	}

}
