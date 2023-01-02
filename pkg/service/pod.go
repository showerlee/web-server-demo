package service

import (
	"context"
	"web-server-demo/pkg/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func GetPods(namespaceName string) ([]v1.Pod, error) {
	ctx := context.Background()
	clientSet, err := client.Getk8sClientSet()
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	podList, err := clientSet.CoreV1().Pods(namespaceName).List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	return podList.Items, nil
}
