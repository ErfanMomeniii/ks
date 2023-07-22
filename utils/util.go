package utils

import (
	"github.com/rodaine/table"
	v1core "k8s.io/api/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(configPath string) (*rest.Config, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	return kubeConfig, nil
}

func ShowInTable(pods []v1core.Pod) {
	t := table.New("Name", "NameSpace", "Status")
	for _, pod := range pods {
		t.AddRow(pod.Name, pod.Namespace, pod.Status)
	}

	t.Print()
}
