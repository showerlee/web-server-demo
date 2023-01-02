package client

import (
	"errors"
	"flag"
	"path/filepath"
	"sync"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

var onceClient sync.Once
var onceConfig sync.Once
var KubeConfig *rest.Config
var KubeClientSet *kubernetes.Clientset

func Getk8sClientSet() (*kubernetes.Clientset, error) {
	onceClient.Do(func() {
		config, err := GetRestConfig()
		if err != nil {
			return
		}
		KubeClientSet, err = kubernetes.NewForConfig(config)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	return KubeClientSet, nil
}

func GetRestConfig() (*rest.Config, error) {
	onceConfig.Do(func() {
		var kubeConfig *string
		var err error
		if home := homedir.HomeDir(); home != "" {
			kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "")
		} else {
			err = errors.New("read config error, config is empty")
			klog.Fatal(err)
			return
		}
		flag.Parse()
		KubeConfig, err = clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			klog.Fatal(err)
			return
		}
	})
	return KubeConfig, nil
}
