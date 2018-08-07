package logdna

import (
    "errors"
    "log"
    "os"

    "github.com/gliderlabs/logspout/router"
    "github.com/smusali/logspout/logdna/adapter"
)

func init() {
    router.AdapterFactories.Register(NewLogDNAAdapter, "logdna")

    r := &router.Route{
        Adapter:    "logdna",
    }

    err := router.Routes.Add(r)
    if err != nil {
        log.Fatal("could not add route: ", err.Error())
    }
}

func NewLogDNAAdapter(route *router.Route) (router.LogAdapter, error) {
    endpoint := os.Getenv("LOGDNA_URL")
    token := os.Getenv("LOGDNA_KEY")
    tags := os.Getenv("TAGS")
    hostname := os.Getenv("HOSTNAME")

    if endpoint == "" {
        endpoint = "logs.logdna.com/logs/ingest"
    }

    if token == "" {
        return nil, errors.New(
            "could not find environment variable LOGDNA_KEY",
        )
    }

    if hostname == "" {
        host, err := os.Hostname()
        if err != nil {
            log.Fatal(err.Error())
        }
        hostname = host
    }

    return adapter.New(
        endpoint,
        token,
        tags,
        hostname,
    ), nil
}