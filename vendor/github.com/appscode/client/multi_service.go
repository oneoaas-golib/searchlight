package client

import (
	attic "github.com/appscode/api/attic/v1beta1"
	auth "github.com/appscode/api/auth/v1beta1"
	ca "github.com/appscode/api/certificate/v1beta1"
	ci "github.com/appscode/api/ci/v1beta1"
	db "github.com/appscode/api/db/v1beta1"
	glusterfs "github.com/appscode/api/glusterfs/v1beta1"
	kubernetesV1beta1 "github.com/appscode/api/kubernetes/v1beta1"
	kubernetesV1beta2 "github.com/appscode/api/kubernetes/v1beta2"
	namespace "github.com/appscode/api/namespace/v1beta1"
	"google.golang.org/grpc"
)

// multi client services are grouped by there main client. the api service
// clients are wrapped around with sub-service.
type multiClientInterface interface {
	Attic() *atticService
	Authentication() *authenticationService
	CA() *caService
	CI() *ciService
	DB() *dbService
	Namespace() *nsService
	GlusterFS() *glusterFSService
	Kubernetes() *versionedKubernetesService
}

type multiClientServices struct {
	atticClient               *atticService
	authenticationClient      *authenticationService
	caClient                  *caService
	ciClient                  *ciService
	glusterfsClient           *glusterFSService
	nsClient                  *nsService
	versionedKubernetesClient *versionedKubernetesService
	dbClient                  *dbService
}

func newMultiClientService(conn *grpc.ClientConn) multiClientInterface {
	return &multiClientServices{
		atticClient: &atticService{
			artifactClient: attic.NewArtifactsClient(conn),
			versionClient:  attic.NewVersionsClient(conn),
		},
		authenticationClient: &authenticationService{
			authenticationClient: auth.NewAuthenticationClient(conn),
			conduitClient:        auth.NewConduitClient(conn),
			projectClient:        auth.NewProjectsClient(conn),
		},
		caClient: &caService{
			certificateClient: ca.NewCertificatesClient(conn),
		},
		ciClient: &ciService{
			agentsClient:   ci.NewAgentsClient(conn),
			metadataClient: ci.NewMetadataClient(conn),
		},
		glusterfsClient: &glusterFSService{
			clusterClient: glusterfs.NewClustersClient(conn),
			volumeClient:  glusterfs.NewVolumesClient(conn),
		},
		versionedKubernetesClient: &versionedKubernetesService{
			v1beta1Service: &kubernetesV1beta1Service{
				clientsClient:      kubernetesV1beta1.NewClientsClient(conn),
				clusterClient:      kubernetesV1beta1.NewClustersClient(conn),
				eventsClient:       kubernetesV1beta1.NewEventsClient(conn),
				incidentClient:     kubernetesV1beta1.NewIncidentsClient(conn),
				loadBalancerClient: kubernetesV1beta1.NewLoadBalancersClient(conn),
				metdataClient:      kubernetesV1beta1.NewMetadataClient(conn),
			},
			v1beta2Service: &kubernetesV1beta2Service{
				clientsClient: kubernetesV1beta2.NewClientsClient(conn),
				diskClient:    kubernetesV1beta2.NewDisksClient(conn),
			},
		},
		nsClient: &nsService{
			teamClient:    namespace.NewTeamsClient(conn),
			billingClient: namespace.NewBillingClient(conn),
		},
		dbClient: &dbService{
			database: db.NewDatabasesClient(conn),
			snapshot: db.NewSnapshotsClient(conn),
		},
	}
}

func (s *multiClientServices) Attic() *atticService {
	return s.atticClient
}

func (s *multiClientServices) Authentication() *authenticationService {
	return s.authenticationClient
}

func (s *multiClientServices) Namespace() *nsService {
	return s.nsClient
}

func (s *multiClientServices) CA() *caService {
	return s.caClient
}

func (s *multiClientServices) CI() *ciService {
	return s.ciClient
}

func (s *multiClientServices) GlusterFS() *glusterFSService {
	return s.glusterfsClient
}

func (s *multiClientServices) Kubernetes() *versionedKubernetesService {
	return s.versionedKubernetesClient
}

func (s *multiClientServices) DB() *dbService {
	return s.dbClient
}

// original service clients that needs to exposed under grouped wrapper
// services.
type atticService struct {
	artifactClient attic.ArtifactsClient
	versionClient  attic.VersionsClient
}

func (a *atticService) Artifacts() attic.ArtifactsClient {
	return a.artifactClient
}

func (a *atticService) Versions() attic.VersionsClient {
	return a.versionClient
}

type authenticationService struct {
	authenticationClient auth.AuthenticationClient
	conduitClient        auth.ConduitClient
	projectClient        auth.ProjectsClient
}

func (a *authenticationService) Authentication() auth.AuthenticationClient {
	return a.authenticationClient
}

func (a *authenticationService) Conduit() auth.ConduitClient {
	return a.conduitClient
}

func (a *authenticationService) Project() auth.ProjectsClient {
	return a.projectClient
}

type ciService struct {
	agentsClient   ci.AgentsClient
	metadataClient ci.MetadataClient
}

func (a *ciService) Agents() ci.AgentsClient {
	return a.agentsClient
}

func (a *ciService) Metadata() ci.MetadataClient {
	return a.metadataClient
}

type nsService struct {
	teamClient    namespace.TeamsClient
	billingClient namespace.BillingClient
}

func (b *nsService) Team() namespace.TeamsClient {
	return b.teamClient
}

func (b *nsService) Billing() namespace.BillingClient {
	return b.billingClient
}

type caService struct {
	certificateClient ca.CertificatesClient
}

func (c *caService) CertificatesClient() ca.CertificatesClient {
	return c.certificateClient
}

type glusterFSService struct {
	clusterClient glusterfs.ClustersClient
	volumeClient  glusterfs.VolumesClient
}

func (g *glusterFSService) Cluster() glusterfs.ClustersClient {
	return g.clusterClient
}

func (g *glusterFSService) Volume() glusterfs.VolumesClient {
	return g.volumeClient
}

type versionedKubernetesService struct {
	v1beta1Service *kubernetesV1beta1Service
	v1beta2Service *kubernetesV1beta2Service
}

func (v *versionedKubernetesService) V1beta1() *kubernetesV1beta1Service {
	return v.v1beta1Service
}

func (v *versionedKubernetesService) V1beta2() *kubernetesV1beta2Service {
	return v.v1beta2Service
}

type kubernetesV1beta1Service struct {
	clientsClient      kubernetesV1beta1.ClientsClient
	clusterClient      kubernetesV1beta1.ClustersClient
	eventsClient       kubernetesV1beta1.EventsClient
	incidentClient     kubernetesV1beta1.IncidentsClient
	loadBalancerClient kubernetesV1beta1.LoadBalancersClient
	metdataClient      kubernetesV1beta1.MetadataClient
}

func (k *kubernetesV1beta1Service) Client() kubernetesV1beta1.ClientsClient {
	return k.clientsClient
}

func (k *kubernetesV1beta1Service) Cluster() kubernetesV1beta1.ClustersClient {
	return k.clusterClient
}

func (k *kubernetesV1beta1Service) Event() kubernetesV1beta1.EventsClient {
	return k.eventsClient
}

func (a *kubernetesV1beta1Service) Incident() kubernetesV1beta1.IncidentsClient {
	return a.incidentClient
}

func (a *kubernetesV1beta1Service) LoadBalancer() kubernetesV1beta1.LoadBalancersClient {
	return a.loadBalancerClient
}

func (k *kubernetesV1beta1Service) Metadata() kubernetesV1beta1.MetadataClient {
	return k.metdataClient
}

type kubernetesV1beta2Service struct {
	clientsClient kubernetesV1beta2.ClientsClient
	diskClient    kubernetesV1beta2.DisksClient
}

func (k *kubernetesV1beta2Service) Client() kubernetesV1beta2.ClientsClient {
	return k.clientsClient
}

func (k *kubernetesV1beta2Service) Disk() kubernetesV1beta2.DisksClient {
	return k.diskClient
}

type dbService struct {
	database db.DatabasesClient
	snapshot db.SnapshotsClient
}

func (p *dbService) Database() db.DatabasesClient {
	return p.database
}

func (p *dbService) Snapshot() db.SnapshotsClient {
	return p.snapshot
}
