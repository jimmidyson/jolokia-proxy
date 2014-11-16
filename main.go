package main

import (
	"fmt"
	flags "github.com/jessevdk/go-flags"
	"log"
	"os"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
)

func main() {
	var opts struct {
		Verbose             bool     `short:"v" long:"verbose" description:"Verbose output" env:"VERBOSE"`
		ListenPort          uint16   `short:"p" long:"port" description:"What port to listen on" default:"8080" env:"JOLOKIA_PROXY_PORT"`
		KubernetesMaster    string   `short:"k" long:"kubernetes" description:"Kubernetes master URL" default:"http://localhost:8080" env:"KUBERNETES_MASTER"`
		KubernetesUsername  string   `short:"u" long:"kubernetes-user" description:"Username to authenticate to Kubernetes master" env:"KUBERNETES_USER"`
		KubernetesPassword  string   `short:"P" long:"kubernetes-password" description:"Password to authenticate to Kubernetes master" env:"KUBERNETES_PASSWORD"`
		KubernetesNamespace string   `short:"N" long:"kubernetes-namespace" description:"The namespace to search by default" env:"KUBERNETES_NAMESPACE" env-delim:","`
		JolokiaPorts        []uint16 `short:"j" long:"jolokia-port" description:"The Jolokia port number" default:"8778" env:"JOLOKIA_PORT" env-delim:","`
		JolokiaPortNames    []string `short:"n" long:"jolokia-port-name" description:"The Jolokia port name" default:"jolokia" env:"JOLOKIA_PORT_NAME" env-delim:","`
	}

	// parse said flags
	_, err := flags.Parse(&opts)
	if err != nil {
		if e, ok := err.(*flags.Error); ok {
			if e.Type == flags.ErrHelp {
				os.Exit(0)
			}
		}

		fmt.Println(err)

		os.Exit(1)
	}

	log.Printf("Listening on port %d", opts.ListenPort)
	log.Printf("Using Kubernetes master at %v", opts.KubernetesMaster)
	if len(opts.KubernetesUsername) > 0 && len(opts.KubernetesPassword) > 0 {
		log.Printf("Authenticating to Kubernetes with %v:********", opts.KubernetesUsername)
	}
	log.Printf("Posible Jolokia ports: %v", opts.JolokiaPorts)
	log.Printf("Possible Jolokia port names: %v", opts.JolokiaPortNames)

	config := client.Config{
		Host:     "http://localhost:8080",
		Username: "test",
		Password: "password",
	}
	client, err := client.New(&config)
	if err != nil {
		// handle error
	}
	version, err := client.ServerVersion()
	if err != nil {
		log.Panicf("Could not retrieve server version: %v", err)
	} else {
		log.Printf("Kubernetes server version: %v", version)
	}

	apiVersions, err := client.ServerAPIVersions()
	if err != nil {
		log.Printf("Could not retrieve server API versions: %v", err)
	} else {
		log.Printf("Kubernetes server supported API versions: %v", apiVersions)
	}

	podList, err := client.Pods(opts.KubernetesNamespace).List(labels.Everything())
	if err != nil {
		log.Printf("Could not retrieve pods: %v", err)
	} else {
		log.Printf("%v", podList.Items)
	}
}
