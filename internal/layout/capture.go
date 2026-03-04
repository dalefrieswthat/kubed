package layout

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// RunCapture builds a v2 index (infra paths, project structure, context sections, optional k8s layout) and writes .kubed/layout.json.
func RunCapture(allNamespaces bool) error {
	idx, err := RunCaptureV2(allNamespaces)
	if err != nil {
		return err
	}
	layoutPath, err := WriteIndexV2(idx)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %s\n", layoutPath)
	return nil
}

func buildLayout(ctx context.Context, c *kubernetes.Clientset, contextName string, allNamespaces bool) (*Layout, error) {
	layout := &Layout{Context: contextName}

	namespaces, err := getNamespaces(ctx, c, allNamespaces)
	if err != nil {
		return nil, err
	}
	layout.Namespaces = namespaces

	var resources []Resource
	var relationships []Relationship

	for _, ns := range namespaces {
		deployments, err := c.AppsV1().Deployments(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("list deployments in %s: %w", ns, err)
		}
		services, err := c.CoreV1().Services(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("list services in %s: %w", ns, err)
		}
		configMaps, err := c.CoreV1().ConfigMaps(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("list configmaps in %s: %w", ns, err)
		}
		secrets, err := c.CoreV1().Secrets(ns).List(ctx, metav1.ListOptions{})
		if err != nil {
			return nil, fmt.Errorf("list secrets in %s: %w", ns, err)
		}

		for _, d := range deployments.Items {
			replicas := d.Spec.Replicas
			var rep *int32
			if replicas != nil {
				rep = replicas
			}
			images := getContainerImages(d.Spec.Template.Spec.Containers)
			resources = append(resources, Resource{
				Kind:      "Deployment",
				Name:      d.Name,
				Namespace: d.Namespace,
				Replicas:  rep,
				Images:    images,
			})
			refs := refsFromPodSpec(&d.Spec.Template.Spec, d.Namespace)
			for _, toRef := range refs {
				relationships = append(relationships, Relationship{
					From: Ref{Kind: "Deployment", Name: d.Name, Namespace: d.Namespace},
					To:   toRef,
					Kind: refKind(toRef.Kind),
				})
			}
		}
		for _, s := range services.Items {
			resources = append(resources, Resource{
				Kind:      "Service",
				Name:      s.Name,
				Namespace: s.Namespace,
			})
			depRef := findDeploymentMatchingSelector(s.Spec.Selector, deployments.Items)
			if depRef != nil {
				relationships = append(relationships, Relationship{
					From: Ref{Kind: "Service", Name: s.Name, Namespace: s.Namespace},
					To:   *depRef,
					Kind: "service-selects-deployment",
				})
			}
		}
		for _, cm := range configMaps.Items {
			resources = append(resources, Resource{
				Kind:      "ConfigMap",
				Name:      cm.Name,
				Namespace: cm.Namespace,
			})
		}
		for _, sec := range secrets.Items {
			resources = append(resources, Resource{
				Kind:      "Secret",
				Name:      sec.Name,
				Namespace: sec.Namespace,
			})
		}
	}

	layout.Resources = resources
	layout.Relationships = relationships
	return layout, nil
}

func getNamespaces(ctx context.Context, c *kubernetes.Clientset, all bool) ([]string, error) {
	if !all {
		return []string{corev1.NamespaceDefault}, nil
	}
	list, err := c.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(list.Items))
	for _, n := range list.Items {
		names = append(names, n.Name)
	}
	return names, nil
}

func getContainerImages(containers []corev1.Container) []string {
	images := make([]string, 0, len(containers))
	for _, c := range containers {
		images = append(images, c.Image)
	}
	return images
}

func refKind(kind string) string {
	switch kind {
	case "ConfigMap":
		return "deployment-uses-configmap"
	case "Secret":
		return "deployment-uses-secret"
	default:
		return "deployment-uses-" + kind
	}
}

// refsFromPodSpec extracts ConfigMap and Secret refs from volumes and env.
func refsFromPodSpec(spec *corev1.PodSpec, namespace string) []Ref {
	var refs []Ref
	seen := make(map[string]bool)
	add := func(kind, name string) {
		key := kind + "/" + name
		if seen[key] {
			return
		}
		seen[key] = true
		refs = append(refs, Ref{Kind: kind, Name: name, Namespace: namespace})
	}
	for _, v := range spec.Volumes {
		if v.ConfigMap != nil {
			add("ConfigMap", v.ConfigMap.Name)
		}
		if v.Secret != nil {
			add("Secret", v.Secret.SecretName)
		}
	}
	for _, c := range spec.Containers {
		for _, e := range c.EnvFrom {
			if e.ConfigMapRef != nil {
				add("ConfigMap", e.ConfigMapRef.Name)
			}
			if e.SecretRef != nil {
				add("Secret", e.SecretRef.Name)
			}
		}
		for _, e := range c.Env {
			if e.ValueFrom != nil {
				if e.ValueFrom.ConfigMapKeyRef != nil {
					add("ConfigMap", e.ValueFrom.ConfigMapKeyRef.Name)
				}
				if e.ValueFrom.SecretKeyRef != nil {
					add("Secret", e.ValueFrom.SecretKeyRef.Name)
				}
			}
		}
	}
	return refs
}

func findDeploymentMatchingSelector(selector map[string]string, deployments []appsv1.Deployment) *Ref {
	if len(selector) == 0 {
		return nil
	}
	for _, d := range deployments {
		if matchSelector(d.Spec.Template.Labels, selector) {
			return &Ref{Kind: "Deployment", Name: d.Name, Namespace: d.Namespace}
		}
	}
	return nil
}

func matchSelector(labels, selector map[string]string) bool {
	for k, v := range selector {
		if labels[k] != v {
			return false
		}
	}
	return true
}
