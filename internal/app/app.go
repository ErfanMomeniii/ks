package app

import (
	"context"

	v1core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	"github.com/erfanmomeniii/ks/utils"
)

func kubeClient(kubeConfig *rest.Config) (*kubernetes.Clientset, error) {
	client, err := kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func namespaces(client *kubernetes.Clientset, ctx context.Context, options v1.ListOptions) ([]v1core.Namespace, error) {
	ns, err := client.CoreV1().Namespaces().List(ctx, options)
	if err != nil {
		return nil, err
	}

	return ns.Items, nil
}

func pods(client *kubernetes.Clientset, namespace v1core.Namespace, ctx context.Context, options v1.ListOptions) ([]v1core.Pod, error) {
	ps, err := client.CoreV1().Pods(namespace.Name).List(ctx, options)
	if err != nil {
		return nil, err
	}

	return ps.Items, nil
}

func Run(configPath string) error {
	config, err := utils.GetConfig(configPath)
	if err != nil {
		return err
	}

	client, err := kubeClient(config)
	if err != nil {
		return err
	}

	allNamespaces, err := namespaces(client, context.Background(), v1.ListOptions{})
	if err != nil {
		return err
	}

	allPods := make([]v1core.Pod, len(allNamespaces))

	// get all pods of namespace
	for _, ns := range allNamespaces {
		ps, _ := pods(client, ns, context.Background(), v1.ListOptions{})
		allPods = append(allPods, ps...)
	}

	utils.ShowInTable(allPods)

	return nil
}
