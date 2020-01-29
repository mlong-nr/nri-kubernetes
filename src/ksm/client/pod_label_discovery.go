package client

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/newrelic/nri-kubernetes/src/client"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
)

type podLabelDiscoverer struct {
	ksmPodLabel string
	logger      *logrus.Logger
	k8sClient   client.Kubernetes
}

func (p *podLabelDiscoverer) findSingleKSMPodByLabel() (*v1.Pod, error) {

	pods, err := p.k8sClient.FindPodsByLabel(p.ksmPodLabel, "true")

	if err != nil {
		return nil, errors.Wrap(err, "could not query api server for pods")
	}
	if len(pods.Items) == 0 {
		return nil, errors.Wrapf(errNoKSMPodsFound, "no KSM pod found with label: '%s'", p.ksmPodLabel)
	}

	// In case there are multiple pods, we must be be sure to deterministically select the same Pod on each node
	// So we chose, for example, the HostIp with highest precedence in alphabetical order
	var chosenPod v1.Pod
	for _, pod := range pods.Items {

		if pod.Status.HostIP == "" {
			continue
		}

		if chosenPod.Status.HostIP == "" || pod.Status.HostIP > chosenPod.Status.HostIP {
			chosenPod = pod
		}
	}

	return &chosenPod, nil
}

// Discover will find a single KSM pod using the provided label.
func (p *podLabelDiscoverer) Discover(timeout time.Duration) (client.HTTPClient, error) {

	pod, err := p.findSingleKSMPodByLabel()
	if err != nil {
		return nil, err
	}

	endpoint := url.URL{
		Scheme: "http",
		Host:   fmt.Sprintf("%s:8080", pod.Status.PodIP),
	}

	return &ksm{
		nodeIP:   pod.Status.HostIP,
		endpoint: endpoint,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		logger: p.logger,
	}, nil
}

// NewPodLabelDiscoverer creates a new KSM discoverer that will find KSM pods using k8s labels
func NewPodLabelDiscoverer(ksmPodLabel string, logger *logrus.Logger, k8sClient client.Kubernetes) client.Discoverer {
	return &podLabelDiscoverer{
		logger:      logger,
		k8sClient:   k8sClient,
		ksmPodLabel: ksmPodLabel,
	}
}
