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
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	api    pb.VirtualMachineServiceClient
}

// +kubebuilder:rbac:groups=infrastructure.nokamoto.github.com,resources=virtualmachines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infrastructure.nokamoto.github.com,resources=virtualmachines/status,verbs=get;update;patch

func (r *VirtualMachineReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("virtualmachine", req.NamespacedName)

	polling := func() ctrl.Result {
		now := time.Now()
		next := 10 * time.Second
		res := ctrl.Result{RequeueAfter: next}
		log.Info("polling", "now", now, "next run", next)
		return res
	}

	var virtualMachine infrav1.VirtualMachine
	if err := r.Get(ctx, req.NamespacedName, &virtualMachine); err != nil {
		log.Error(err, "unable to fetch VirtualMachine state")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	log.Info("found", "virtualMachine", virtualMachine)

	// 1. find by name
	res, err := r.api.FindByName(ctx, &pb.VirtualMachine{
		Name: virtualMachine.Spec.Name,
	})
	if status.Code(err) == codes.NotFound {
		log.Info("create if not exists")

		typ, _ := pb.MachineType_value[string(virtualMachine.Spec.MachineType)]

		res, err := r.api.Create(ctx, &pb.VirtualMachine{
			Name:        virtualMachine.Spec.Name,
			MachineType: pb.MachineType(typ),
		})
		if err != nil {
			log.Error(err, "unable to create VirtualMachine")
			return ctrl.Result{}, err
		}

		log.Info("create", "res", *res)

		return polling(), nil
	} else if err != nil {
		log.Error(err, "unable to fetch VirtualMachine status")
		return ctrl.Result{}, err
	}

	log.Info("found", "res", *res)

	return polling(), nil
}

func (r *VirtualMachineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	cc, err := grpc.Dial("todo", grpc.WithInsecure())
	if err != nil {
		return err
	}

	r.api = pb.NewVirtualMachineServiceClient(cc)

	return ctrl.NewControllerManagedBy(mgr).
		For(&infrav1.VirtualMachine{}).
		Complete(r)
}
