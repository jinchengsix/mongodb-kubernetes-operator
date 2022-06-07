package rbac

import (
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type RBAC interface {
	SAGetCreator
	RoleGetCreator
	RoleBindingGetCreator
}

type SAGetCreator interface {
	GetServiceAccount(objectKey client.ObjectKey) (corev1.ServiceAccount, error)
	CreateServiceAccount(sa corev1.ServiceAccount) error
}

type RoleGetCreator interface {
	GetRole(objectKey client.ObjectKey) (rbacv1.Role, error)
	CreateRole(role rbacv1.Role) error
}

type RoleBindingGetCreator interface {
	GetRoleBinding(objectKey client.ObjectKey) (rbacv1.RoleBinding, error)
	CreateRoleBinding(rb rbacv1.RoleBinding) error
}
