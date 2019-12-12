<p align="center">
    <img alt="terraform-provider-example Logo" src="https://www.terraform.io/assets/images/og-image-8b3e4f7d.png" width="150" />
    <h3 align="center">Terraform Provider Example.</h3>
    <p align="center">
        <a href="https://godoc.org/github.com/Clivern/terraform-provider-example"><img src="https://godoc.org/github.com/Clivern/terraform-provider-example?status.svg"></a>
        <a href="https://travis-ci.org/Clivern/terraform-provider-example"><img src="https://travis-ci.org/Clivern/terraform-provider-example.svg?branch=master"></a>
        <a href="https://github.com/Clivern/terraform-provider-example/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://goreportcard.com/report/github.com/Clivern/terraform-provider-example"><img src="https://goreportcard.com/badge/github.com/Clivern/terraform-provider-example?v=0.0.1"></a>
        <a href="https://github.com/Clivern/terraform-provider-example/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-MIT-orange.svg"></a>
    </p>
</p>


## Documentation

First we need to create a simple web service to do a CRUD operations.

```golang
package main

import (
	"fmt"
	"net/http"
	"os"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Id     int    `json:"id"`
	Image  string `json:"image"`
	Name   string `json:"name"`
	Size   string `json:"size"`
	Region string `json:"region"`
}

// LoadFromJSON update object from json
func (s *Server) LoadFromJSON(data []byte) (bool, error) {
	err := json.Unmarshal(data, &s)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ConvertToJSON convert object to json
func (s *Server) ConvertToJSON() (string, error) {
	data, err := json.Marshal(&s)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Store(file, data string)(bool, error){
    f, err := os.Create(file)
    if err != nil {
        return false, err
    }
    _, err = f.WriteString(data)
    if err != nil {
        f.Close()
        return false, err
    }
    err = f.Close()
    if err != nil {
        return false, err
    }
    return true, nil
}

func Retrieve(file string) string{
    b, err := ioutil.ReadFile(file) // just pass the file name

    if err != nil {
        return ""
    }

    return string(b)
}

func main() {

	gin.DisableConsoleColor()
	gin.DefaultWriter = os.Stdout

	r := gin.Default()

	r.GET("/favicon.ico", func(c *gin.Context) {
		c.String(http.StatusNoContent, "")
	})

	r.GET("/server/:id", func(c *gin.Context) {
		data := Retrieve("db.txt")

		if data == ""{
			c.JSON(http.StatusNotFound, gin.H{
				"status": "error",
				"error":  "Server not found",
			})
			return
		}

		server := &Server{}
		server.LoadFromJSON([]byte(data))

		c.JSON(http.StatusOK, gin.H{
			"id": server.Id,
			"image": server.Image,
			"name": server.Name,
			"size": server.Size,
			"region": server.Region,
		})
	})

	r.POST("/server", func(c *gin.Context) {
		rawBody, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request",
			})
			return
		}

		server := &Server{}
		server.LoadFromJSON([]byte(rawBody))

		server.Id = 1

		data, _ := server.ConvertToJSON()

		ok, err := Store("db.txt", data)

		if !ok || err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": server.Id,
			"image": server.Image,
			"name": server.Name,
			"size": server.Size,
			"region": server.Region,
		})
	})

	r.PUT("/server/:id", func(c *gin.Context) {
		rawBody, err := c.GetRawData()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": "error",
				"error":  "Invalid request",
			})
			return
		}

		server := &Server{}
		server.LoadFromJSON([]byte(rawBody))

		server.Id = 1

		data, _ := server.ConvertToJSON()

		ok, err := Store("db.txt", data)

		if !ok || err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  "Internal Server Error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"id": server.Id,
			"image": server.Image,
			"name": server.Name,
			"size": server.Size,
			"region": server.Region,
		})
	})

	r.DELETE("/server/:id", func(c *gin.Context) {
		Store("db.txt", "")
		c.Status(http.StatusNoContent)
	})

	r.Run(fmt.Sprintf(":%d", 8080))
}
```

Run this service on the background.

```bash
$ go run main.go
```

Then we can use our terraform provider to make changes to the web service resources.

```bash
$ git clone https://github.com/Clivern/terraform-provider-example.git

# Build the provider
$ go build -o terraform-provider-example

# Initialize a working directory containing Terraform configuration files
$ terraform init

# Create an execution plan.
$ terraform plan

# Apply the changes required to reach the desired state of the configuration
$ terraform apply
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, terraform-provider-example is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/Clivern/terraform-provider-example/releases) for changelogs for each release version of terraform-provider-example. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/Clivern/terraform-provider-example/issues


## Security Issues

If you discover a security vulnerability within terraform-provider-example, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2019, Clivern. Released under [MIT License](https://opensource.org/licenses/mit-license.php).

**terraform-provider-example** is authored and maintained by [@Clivern](http://github.com/Clivern).
