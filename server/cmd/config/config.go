package config

type Config struct {
	Port int
	UploadsPath string
	ProtoPoseDetectFileName string
	WeightsPoseDetectFileName string
}

func NewConfig() Config {
	return Config{
		Port: 8000,
		UploadsPath: "uploads",
		ProtoPoseDetectFileName: "pose_deploy_linevec.prototxt",
		WeightsPoseDetectFileName: "pose_iter_440000.caffemodel",
	}
} 