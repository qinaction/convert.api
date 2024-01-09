package common

const (
	ServiceRootDir    = "convert.api/"
	ServiceConfDir    = ServiceRootDir + "services-conf/"
	ServiceAddressDir = ServiceRootDir + "services/"
)

const (
	CertificateDir      = "certificate/"
	CertificateCertFile = CertificateDir + "cert.pem"
	CertificateKeyFile  = CertificateDir + "key.pem"
	K8sKeyDir           = "k8s_keys/"
	K8sKeyCaFile        = K8sKeyDir + "ca.pem"
	K8sKeyServerFile    = K8sKeyDir + "server.pem"
	K8sKeyserverKeyFile = K8sKeyDir + "server-key.pem"
)

var CommandParams = map[string]interface{}{}
