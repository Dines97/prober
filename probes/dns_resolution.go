package probes

import (
	"net"
	// "sigs.k8s.io/controller-runtime/pkg/log"
)

// +kubebuilder:object:generate:=true

type Resolution struct {
	// +kubebuilder:object:root:=true

	Host          string   `json:"host"`
	Registrations []string `json:"registrations"`
}

func (r *Resolution) Run() (RunResult, error) {

	_, err := net.LookupIP(r.Host)

	if err != nil {
		return Fail, err
	}

	// for _, ip := range ips {
	// 	log.Log.Info("%s. IN A %s\n", r.Host, ip.String())
	// }

	return Success, nil
}
