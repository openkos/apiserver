package service

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/openkos/openkos/api/settings"
)

type Settings struct {
	settings.UnimplementedSettingsServer

	cs *kubernetes.Clientset
}

func NewSettings(cs *kubernetes.Clientset) *Settings {
	return &Settings{cs: cs}
}

func (s *Settings) List(ctx context.Context, request *settings.ListRequest) (*settings.ListReply, error) {
	_, err := s.cs.CoreV1().ConfigMaps("openkos").List(ctx, metav1.ListOptions{})
	return &settings.ListReply{}, err
}

func (s *Settings) Update(ctx context.Context, request *settings.UpdateRequest) (*settings.UpdateReply, error) {
	cm := &v1.ConfigMap{}
	_, err := s.cs.CoreV1().ConfigMaps("openkos").Update(ctx, cm, metav1.UpdateOptions{})
	return &settings.UpdateReply{}, err
}
