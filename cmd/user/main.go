package main

import (
	"edgenet/pkg/authorization"
	"edgenet/pkg/controller/v1alpha/user"
)

func main() {
	// Set kubeconfig to be used to create clientsets
	authorization.SetKubeConfig()
	// Start the controller to provide the functionalities of user resource
	user.Start()
}
