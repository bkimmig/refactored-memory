package app

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler func(c *gin.Context)

type Method int

const (
	Get Method = iota
	Post
	// Put
	// Delete
)

type endpoint struct {
	path    string
	name    string
	method  Method
	handler gin.HandlerFunc
}

type controller []endpoint
type versionCollection map[int]controller

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world"})
}

func createEndpoints(router *gin.Engine, versions versionCollection) {
	n_versions := len(versions)
	var paths controller
	var prefix string
	var path string

	for i := 0; i < n_versions; i++ {
		for j := 0; j < i+1; j++ {
			prefix = "/v" + strconv.Itoa(i)
			// add all endpoints
			paths = versions[j]
			for _, p := range paths {
				path = prefix + p.path
				if p.method == Get {
					router.GET(path, p.handler)
				} else if p.method == Post {
					router.POST(path, p.handler)
				} else {
					return
				}
			}
		}
	}
}

func Run() {
	router := gin.Default()
	// ---------------------------------------
	// version 1 endpoints
	v0 := make(controller, 1)
	v0[0] = endpoint{
		method:  Get,
		name:    "home",
		path:    "/",
		handler: home,
	}

	// ---------------------------------------

	// Endpoint versions
	// beginning of slug will be versioned `v1/user`
	// if we change the user endpoint in a major way we will add a version 2 of
	// that and it will move over all other endpoints so the developers using our
	// API can just update the slug in one place and keep the rest of their logic
	// consistent.

	versions := make(versionCollection)
	versions[0] = v0
	// <add more versions here>

	createEndpoints(router, versions)
	router.Run()
}
