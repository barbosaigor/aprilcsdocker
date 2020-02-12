package dockercs

type DockerDestryr struct{}

func (d DockerDestryr) Destroy(nodes []string) error {
	return StopAll(nodes)
}