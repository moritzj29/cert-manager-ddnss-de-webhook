package config

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	missingSecretKeyMsg = "key %q not found in secret \"%s/%s\""
	missingSecretErrMsg = "Unable to load secret %v %q"
)

func GetDDNSSToken(client kubernetes.Interface, cfg *DDNSSProviderConfig, namespace string) (*string, error) {
	secretName := cfg.APITokenSecretRef.LocalObjectReference.Name

	secret, err := client.CoreV1().Secrets(namespace).Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		return nil, errors.Wrapf(err, missingSecretErrMsg, secretName, namespace+"/"+secretName)
	}

	data, ok := secret.Data[cfg.APITokenSecretRef.Key]
	if !ok {
		return nil, fmt.Errorf(missingSecretKeyMsg, cfg.APITokenSecretRef.Key,
			cfg.APITokenSecretRef.LocalObjectReference.Name, namespace)
	}

	apiKey := string(data)
	return &apiKey, nil
}
