package fukabunsan

// import (
// 	"context"
// 	"net"
// 	"sort"
// 	"strings"
// 	"errors"

// 	"github.com/bonavadeur/katyusha/pkg/bonalib"
// 	corev1 "k8s.io/api/core/v1"
// 	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
// 	"k8s.io/apimachinery/pkg/runtime"
// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/rest"

// 	"k8s.io/client-go/util/retry"
// )

// type SourceIPinfo struct {
// 	Node nodeIPInfo
// 	Pod  podIPInfo
// }

// type nodeIPInfo struct {
// 	Name      string
// 	IPaddress string
// }

// type podIPInfo struct {
// 	IPRange      string
// 	PrefixLength int32
// }

// type PodCIDR struct {
// 	Nodename   string
// 	NodeIP     string
// 	PodIPRange string
// 	PodPrefix  int32
// }

// type IPAMBlock struct {
// 	Spec struct {
// 		CIDR     string  `json:"cidr"`
// 		Affinity *string `json:"affinity"`
// 	} `json:"spec"`
// }

// type IPAMBlockList struct {
// 	metav1.TypeMeta `json:",inline"`
// 	metav1.ListMeta `json:"metadata,omitempty"`
// 	Items           []IPAMBlock `json:"items"`
// }

// func (obj *IPAMBlockList) DeepCopyObject() runtime.Object {
// 	dst := &IPAMBlockList{}
// 	dst.TypeMeta = obj.TypeMeta
// 	dst.ListMeta = obj.ListMeta
// 	dst.Items = make([]IPAMBlock, len(obj.Items))
// 	objcopy := make([]IPAMBlock, len(obj.Items))
// 	copy(objcopy, obj.Items)
// 	for i := range objcopy {
// 		dst.Items[i] = obj.Items[i]
// 	}
// 	return dst
// }

// func GetKubeconfig() *rest.Config {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// 	return config
// }

// func GetClientSet() *kubernetes.Clientset {
// 	clientset, err := kubernetes.NewForConfig(KUBECONFIG)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return clientset
// }

// func GetNodenames() []string {
// 	nodes, err := CLIENTSET.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		bonalib.Warn("Error listing nodes: %v\n", err)
// 		return []string{}
// 	}

// 	ret := []string{}
// 	for _, node := range nodes.Items {
// 		ret = append(ret, node.Name)
// 	}
// 	sort.Strings(ret)
// 	return ret
// }

// func GetPodsCIDRs() []PodCIDR {
// 	nodes, err := CLIENTSET.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	var ipamBlockList IPAMBlockList
// 	retry.OnError(retry.DefaultRetry, func(error) bool { return true }, func() error {
// 		return CLIENTSET.RESTClient().
// 			Get().
// 			AbsPath("/apis/crd.projectcalico.org/v1/ipamblocks").
// 			Do(context.TODO()).
// 			Into(&ipamBlockList)
// 	})

// 	ret := make([]PodCIDR, 0, len(NODENAMES))

// 	for i := 0; i < len(NODENAMES); i++ {
// 		nodeip := []corev1.NodeAddress{}
// 		node := NODENAMES[i]
// 		for j := 0; j < len(nodes.Items); j++ {
// 			if nodes.Items[j].Name == node {
// 				nodeip = nodes.Items[j].Status.Addresses
// 			}
// 		}
// 		for _, item := range ipamBlockList.Items {
// 			if item.Spec.Affinity != nil && strings.Contains(*item.Spec.Affinity, node) {
// 				ip, ipNet, err := net.ParseCIDR(item.Spec.CIDR)
// 				ones, _ := ipNet.Mask.Size()
// 				if err != nil {
// 					panic(err)
// 				}
// 				ret = append(ret, PodCIDR{
// 					Nodename:   node,
// 					NodeIP:     nodeip[0].Address,
// 					PodIPRange: ip.String(),
// 					PodPrefix:  int32(ones),
// 				})
// 			}
// 		}
// 	}

// 	return ret
// }

// func isIPInPodCIDR(podCIDR PodCIDR, ipToCheck string) (bool, error) {
// 	// Chuyển PodIPRange thành CIDR
// 	_, podCIDRNet, err := net.ParseCIDR(podCIDR.PodIPRange)
// 	if err != nil {
// 		return false, err
// 	}

// 	// Chuyển IP cần kiểm tra thành kiểu net.IP
// 	ip := net.ParseIP(ipToCheck)
// 	if ip == nil {
// 		return false, errors.New("Invalid IP address")
// 	}

// 	// Kiểm tra xem IP có thuộc dải CIDR không
// 	return podCIDRNet.Contains(ip), nil
// }