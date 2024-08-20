package main

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

var crdName = "demos.example.com"
var namespace = "default"
var gvr = schema.GroupVersionResource{
	Group:    "example.com",
	Version:  "v1",
	Resource: "demos",
}

func main() {
	// Load kubeconfig and create a client
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	config.QPS = 1000
	config.Burst = 1000
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Fatalf("Failed to create dynamic client: %v", err)
	}

	//// Create CRD
	//createCRD(dynamicClient)

	// Start goroutines
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//for i := 0; i < 10; i++ {
	//	go getCRs(ctx, dynamicClient)
	//}
	//for i := 0; i < 10; i++ {
	//	go updateCRs(ctx, dynamicClient)
	//}
	for i := 0; i < 10; i++ {
		go getAndUpdateCRs(ctx, dynamicClient)
	}

	//go deleteCRs(ctx, dynamicClient)

	// Handle termination signals
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	fmt.Println("Received termination signal, shutting down...")
}

func createCRD(client dynamic.Interface) {
	crd := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apiextensions.k8s.io/v1",
			"kind":       "CustomResourceDefinition",
			"metadata": map[string]interface{}{
				"name": crdName,
			},
			"spec": map[string]interface{}{
				"group": "example.com",
				"versions": []interface{}{
					map[string]interface{}{
						"name":    "v1",
						"served":  true,
						"storage": true,
						"schema": map[string]interface{}{
							"openAPIV3Schema": map[string]interface{}{
								"type": "object",
								"properties": map[string]interface{}{
									"spec": map[string]interface{}{
										"type": "object",
										"properties": map[string]interface{}{
											"foo": map[string]interface{}{
												"type": "string",
											},
										},
									},
								},
							},
						},
					},
				},
				"scope": "Namespaced",
				"names": map[string]interface{}{
					"plural":     "demos",
					"singular":   "demo",
					"kind":       "Demo",
					"shortNames": []interface{}{"dm"},
				},
			},
		},
	}

	_, err := client.Resource(schema.GroupVersionResource{
		Group:    "apiextensions.k8s.io",
		Version:  "v1",
		Resource: "customresourcedefinitions",
	}).Create(context.TODO(), crd, metav1.CreateOptions{})
	if err != nil {
		log.Fatalf("Failed to create CRD: %v", err)
	}
	fmt.Println("CRD created successfully")
}

func createCRs(ctx context.Context, client dynamic.Interface) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			cr := &unstructured.Unstructured{
				Object: map[string]interface{}{
					"apiVersion": "example.com/v1",
					"kind":       "Demo",
					"metadata": map[string]interface{}{
						"name": fmt.Sprintf("demo-%d", time.Now().UnixNano()),
					},
					"spec": map[string]interface{}{
						"foo": "bar",
					},
				},
			}

			_, err := client.Resource(gvr).Namespace(namespace).Create(context.TODO(), cr, metav1.CreateOptions{})
			if err != nil {
				log.Printf("Failed to create CR: %v", err)
			} else {
				fmt.Println("CR created successfully")
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func deleteCRs(ctx context.Context, client dynamic.Interface) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			crList, err := client.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Printf("Failed to list CRs: %v", err)
			}

			for _, cr := range crList.Items {
				err = client.Resource(gvr).Namespace(namespace).Delete(context.TODO(), cr.GetName(), metav1.DeleteOptions{})
				if err != nil {
					log.Printf("Failed to delete CR: %v", err)
				} else {
					fmt.Println("CR deleted successfully")
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func getCRs(ctx context.Context, client dynamic.Interface) {
	crList, err := client.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Failed to list CRs: %v", err)
	}
	for {
		select {
		case <-ctx.Done():
			return
		default:
			for _, cr := range crList.Items {
				get, err := client.Resource(gvr).Namespace(namespace).Get(context.TODO(), cr.GetName(), metav1.GetOptions{})
				if err != nil {
					log.Printf("Failed to delete CR: %v", err)
				} else {
					fmt.Sprintf("CR get successfully, Name: %s", get.GetName())
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func updateCRs(ctx context.Context, client dynamic.Interface) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			crList, err := client.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Printf("Failed to list CRs: %v", err)
			}

			for _, cr := range crList.Items {
				// Update the "foo" field in the spec
				cr.Object["spec"].(map[string]interface{})["foo"] = fmt.Sprintf("updated-%d", time.Now().UnixNano())

				_, err = client.Resource(gvr).Namespace(namespace).Update(context.TODO(), &cr, metav1.UpdateOptions{})
				if err != nil {
					log.Printf("Failed to update CR: %v", err)
				} else {
					fmt.Println("CR updated successfully")
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}

func getAndUpdateCRs(ctx context.Context, client dynamic.Interface) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			crList, err := client.Resource(gvr).Namespace(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				log.Printf("Failed to list CRs: %v", err)
			}

			for _, cr := range crList.Items {
				// get
				get, err := client.Resource(gvr).Namespace(namespace).Get(context.TODO(), cr.GetName(), metav1.GetOptions{})
				if err != nil {
					log.Printf("Failed to delete CR: %v", err)
				} else {
					fmt.Sprintf("CR get successfully, Name: %s", get.GetName())
				}

				// Update the "foo" field in the spec
				get.Object["spec"].(map[string]interface{})["foo"] = fmt.Sprintf("updated-%d", time.Now().UnixNano())

				_, err = client.Resource(gvr).Namespace(namespace).Update(context.TODO(), get, metav1.UpdateOptions{})
				if err != nil {
					log.Printf("Failed to update CR: %v", err)
				} else {
					fmt.Println("CR updated successfully")
				}
			}
			time.Sleep(2 * time.Second)
		}
	}
}
