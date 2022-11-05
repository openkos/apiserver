// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kiaedev/kiae/api/graph"
	"github.com/kiaedev/kiae/internal/app/server/dao"
	"github.com/kiaedev/kiae/internal/app/server/service"
	"github.com/kiaedev/kiae/internal/app/server/watch"
	"github.com/kiaedev/kiae/internal/pkg/config"
	"github.com/kiaedev/kiae/internal/pkg/klient"
	"github.com/kiaedev/kiae/pkg/loki"
	"github.com/kiaedev/kiae/pkg/mongoutil"
	"github.com/kiaedev/kiae/pkg/oauth2"
	"github.com/oam-dev/kubevela-core-api/pkg/generated/client/clientset/versioned"
	versioned2 "github.com/pivotal/kpack/pkg/client/clientset/versioned"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Injectors from wire.go:

func buildInjectors(kubeconfig *rest.Config) (*Server, error) {
	router := mux.NewRouter()
	clientset, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	versionedClientset, err := versioned.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	clientset2, err := versioned2.NewForConfig(kubeconfig)
	if err != nil {
		return nil, err
	}
	client, err := klient.CtrRuntimeClient(kubeconfig)
	if err != nil {
		return nil, err
	}
	localClients := &klient.LocalClients{
		K8sCs:         clientset,
		VelaCs:        versionedClientset,
		KpackCs:       clientset2,
		RuntimeClient: client,
	}
	multiClusterInformers := watch.NewMultiClusterInformers(kubeconfig)
	watcher, err := watch.NewWatcher(localClients, multiClusterInformers)
	if err != nil {
		return nil, err
	}
	appPodsService := service.NewAppPodsService(multiClusterInformers)
	lokiClient, err := lokiConstructor()
	if err != nil {
		return nil, err
	}
	appEventService := service.NewAppEventService(lokiClient)
	resolver := &graph.Resolver{
		AppPodsSvc:   appPodsService,
		AppEventsSvc: appEventService,
	}
	database, err := dbConstructor()
	if err != nil {
		return nil, err
	}
	projectDao := dao.NewProject(database)
	appDao := dao.NewApp(database)
	entryDao := dao.NewEntryDao(database)
	routeDao := dao.NewRouteDao(database)
	middlewareInstance := dao.NewMiddlewareInstanceDao(database)
	middlewareClaim := dao.NewMiddlewareClaimDao(database)
	egressDao := dao.NewEgressDao(database)
	appService := service.NewAppService(projectDao, appDao, entryDao, routeDao, middlewareInstance, middlewareClaim, egressDao, clientset, versionedClientset)
	appStatusService := service.NewAppStatusService(client, versionedClientset, appService)
	clusterDao := dao.NewClusterDao(database)
	clusterService := service.NewClusterService(clusterDao, multiClusterInformers)
	egressService := service.NewEgressService(appService, egressDao)
	entryService := service.NewEntryService(appService, entryDao)
	projectImageDao := dao.NewProjectImageDao(database)
	builderDao := dao.NewBuilderDao(database)
	imageRegistryDao := dao.NewRegistryDao(database)
	imageRegistrySvc := service.NewImageRegistrySvc(imageRegistryDao, localClients)
	providerDao := dao.NewProviderDao(database)
	providerService := service.NewProviderService(providerDao, localClients)
	projectImageSvc := service.NewProjectImageSvc(projectDao, projectImageDao, builderDao, imageRegistrySvc, providerService, localClients)
	imageWatcher := service.NewImageWatcher(projectImageSvc, localClients)
	middlewareService := service.NewMiddlewareService(client, clientset, middlewareInstance, middlewareClaim, appService)
	serviceOauth2 := service.NewProviderOauth2Svc(providerService)
	builderSvc := service.NewBuilderSvc(builderDao, imageRegistrySvc, localClients)
	projectService := service.NewProjectService(projectDao, builderSvc)
	routeService := service.NewRouteService(appService, routeDao)
	deploymentDao := dao.NewDeploymentDao(database)
	deploymentService := service.NewDeploymentService(deploymentDao, projectImageSvc, appService)
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	oidcConfig := configConfig.OIDC
	oidc := oauth2.NewOIDC(oidcConfig)
	userDao := dao.NewUserDao(database)
	userSvc := service.NewUserSvc(userDao)
	session := service.NewSession(oidc, userSvc)
	serviceSets := &service.ServiceSets{
		AppService:        appService,
		AppPodsService:    appPodsService,
		AppStatusService:  appStatusService,
		AppEventService:   appEventService,
		ClusterService:    clusterService,
		EgressService:     egressService,
		EntryService:      entryService,
		ImageWatcher:      imageWatcher,
		MiddlewareService: middlewareService,
		Oauth2:            serviceOauth2,
		ProjectService:    projectService,
		ProjectImageSvc:   projectImageSvc,
		ProviderService:   providerService,
		RouteService:      routeService,
		DeploymentService: deploymentService,
		BuilderSvc:        builderSvc,
		ImageRegistrySvc:  imageRegistrySvc,
		Session:           session,
	}
	proxy := klient.NewProxy(kubeconfig)
	server := &Server{
		Router:        router,
		watcher:       watcher,
		graphResolver: resolver,
		svcSets:       serviceSets,
		proxy:         proxy,
	}
	return server, nil
}

// wire.go:

func dbConstructor() (*mongo.Database, error) {
	dbClient, err := mongoutil.New(viper.GetString("mongodb.dsn"))
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to mysql: %v", err)
	}

	return dbClient.DB.Database("kiae"), nil
}

func lokiConstructor() (*loki.Client, error) {
	return loki.NewLoki(viper.GetString("loki.endpoint")), nil
}
