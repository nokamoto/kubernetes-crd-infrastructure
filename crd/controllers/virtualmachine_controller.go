/*

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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "github.com/nokamoto/kubernetes-crd-infrastructure/crd/api/v1"

	pb "github.com/nokamoto/kubernetes-crd-infrastructure/api/protobuf"
)

// VirtualMachineReconciler reconciles a VirtualMachine object
type VirtualMachineReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	pb.VirtualMachineServiceClient
}

// +kubebuilder:rbac:groups=infrastructure.nokamoto.github.com,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.nokamoto.github.com,resources=virtualmachines/status,verbs=get;update;patch

func (r *VirtualMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("virtualmachine", req.NamespacedName)

	var virtualMachine infrav1.VirtualMachine
	if err := r.Get(ctx, req.NamespacedName, &virtualMachine); err != nil {
		log.Error(err, "unable to fetch CronJob")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("found", "virtualMachine", virtualMachine)

	now := time.Now()
	next := 10 * time.Second
	polling := ctrl.Result{RequeueAfter: next}
	log.Info("polling", "now", now, "next run", next)

	return polling, nil
}

func (r *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.VirtualMachine{}).
		Complete(r)
}
