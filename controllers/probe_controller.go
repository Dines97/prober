/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"prober.io/prober/api/v1alpha1"
	probev1alpha1 "prober.io/prober/api/v1alpha1"
	"prober.io/prober/probes"
)

// ProbeReconciler reconciles a Probe object
type ProbeReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=probe.prober.io,resources=probes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=probe.prober.io,resources=probes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=probe.prober.io,resources=probes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Probe object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *ProbeReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here
	var probe v1alpha1.Probe
	if err := r.Get(ctx, req.NamespacedName, &probe); err != nil {
		log.Log.Error(err, "Unable to fetch Probe")

		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Log.Info("Reconciling", "Name", probe.Name)

	// pkg, err := importer.Default().Import("probes")
	// if err != nil {
	// 	log.Log.Error(err, "could not get probes")
	// }
	//
	// for _, declName := range pkg.Scope().Names() {
	//
	// }

	// Iterate over fields of probe
	v := reflect.ValueOf(probe.Spec)
	for i := 0; i < v.NumField(); i++ {

		// Find fields that dont equal to nil
		if s := v.Field(i).Interface(); s != nil {

			// fmt.Println(s)

			// Try to convert this field to ProbeRunner
			probeRunner, ok := s.(probes.ProbeRunner)
			if !ok {
				//fmt.Println("Not a type of Probe")
				continue
			}

			probe.Status.RunResult, _ = probeRunner.Run()

			if err := r.Status().Update(ctx, &probe); err != nil {
				log.Log.Error(err, "Unable to update Probe status")
				return ctrl.Result{}, err
			}
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ProbeReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&probev1alpha1.Probe{}).
		Complete(r)
}
