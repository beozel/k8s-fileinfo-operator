# Go by Example

Content and build toolchain for [Go by Example](https://gobyexample.com),
a site that teaches Go via annotated example programs.

### Overview

The Go by Example site is built by extracting code and
comments from source files in `examples` and rendering
them using `templates` into a static `public`
directory. The programs implementing this build process
are in `tools`, along with dependencies specified in
the `go.mod`file.

The built `public` directory can be served by any
static content system. The production site uses S3 and
CloudFront, for example.

### Building

[![test](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml/badge.svg)](https://github.com/mmcgrana/gobyexample/actions/workflows/test.yml)

To build the site you'll need Go installed. Run:

```console
$ tools/build
```

To build continuously in a loop:

```console
$ tools/build-loop
```

To see the site locally:

```console
$ tools/serve
```

and open `http://127.0.0.1:8000/` in your browser.

### Publishing

To upload the site:

```console
$ export AWS_ACCESS_KEY_ID=...
$ export AWS_SECRET_ACCESS_KEY=...
$ tools/upload
```

### License

This work is copyright Mark McGranaghan and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

The Go Gopher is copyright [Renée French](https://reneefrench.blogspot.com/) and licensed under a
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).


### Translations

Contributor translations of the Go by Example site are available in:

* [Chinese](https://gobyexample-cn.github.io/) by [gobyexample-cn](https://github.com/gobyexample-cn)
* [French](http://le-go-par-l-exemple.keiruaprod.fr) by [keirua](https://github.com/keirua/gobyexample)
* [Japanese](http://spinute.org/go-by-example) by [spinute](https://github.com/spinute)
* [Korean](https://mingrammer.com/gobyexample/) by [mingrammer](https://github.com/mingrammer)
* [Russian](https://gobyexample.com.ru/) by [badkaktus](https://github.com/badkaktus)
* [Ukrainian](https://butuzov.github.io/gobyexample/) by [butuzov](https://github.com/butuzov/gobyexample)
* [Brazilian Portuguese](https://lcslitx.github.io/GoEmExemplos/) by [lcslitx](https://github.com/LCSLITX)
* [Burmese](https://setkyar.github.io/gobyexample) by [Set Kyar Wa Lar](https://github.com/setkyar/gobyexample)

### Thanks

Thanks to [Jeremy Ashkenas](https://github.com/jashkenas)
for [Docco](http://jashkenas.github.io/docco/), which
inspired this project.

### FAQ

#### I found a problem with the examples; what do I do?

We're very happy to fix problem reports and accept contributions! Please submit
[an issue](https://github.com/mmcgrana/gobyexample/issues) or send a Pull Request.
See `CONTRIBUTING.md` for more details.

#### What version of Go is required to run these examples?

Given Go's strong [backwards compatibility guarantees](https://go.dev/doc/go1compat),
we expect the vast majority of examples to work on the latest released version of Go
as well as many older releases going back years.

That said, some examples show off new features added in recent releases; therefore,
it's recommended to try running examples with the latest officially released Go version
(see Go's [release history](https://go.dev/doc/devel/release) for details).

#### I'm getting output in a different order from the example. Is the example wrong?

Some of the examples demonstrate concurrent code which has a non-deterministic
execution order. It depends on how the Go runtime schedules its goroutines and
may vary by operating system, CPU architecture, or even Go version.

Similarly, examples that iterate over maps may produce items in a different order
from what you're getting on your machine. This is because the order of iteration
over maps in Go is [not specified and is not guaranteed to be the same from one
iteration to the next](https://go.dev/ref/spec#RangeClause).

It doesn't mean anything is wrong with the example. Typically the code in these
examples will be insensitive to the actual order of the output; if the code is
sensitive to the order - that's probably a bug - so feel free to report it.

# k8s-fileinfo-operator
// TODO(user): Add simple overview of use/purpose

## Description
// TODO(user): An in-depth paragraph about your project and overview of use

## Getting Started

### Prerequisites
- go version v1.23.0+
- docker version 17.03+.
- kubectl version v1.11.3+.
- Access to a Kubernetes v1.11.3+ cluster.

### To Deploy on the cluster
**Build and push your image to the location specified by `IMG`:**

```sh
make docker-build docker-push IMG=<some-registry>/k8s-fileinfo-operator:tag
```

**NOTE:** This image ought to be published in the personal registry you specified.
And it is required to have access to pull the image from the working environment.
Make sure you have the proper permission to the registry if the above commands don’t work.

**Install the CRDs into the cluster:**

```sh
make install
```

**Deploy the Manager to the cluster with the image specified by `IMG`:**

```sh
make deploy IMG=<some-registry>/k8s-fileinfo-operator:tag
```

> **NOTE**: If you encounter RBAC errors, you may need to grant yourself cluster-admin
privileges or be logged in as admin.

**Create instances of your solution**
You can apply the samples (examples) from the config/sample:

```sh
kubectl apply -k config/samples/
```

>**NOTE**: Ensure that the samples has default values to test it out.

### To Uninstall
**Delete the instances (CRs) from the cluster:**

```sh
kubectl delete -k config/samples/
```

**Delete the APIs(CRDs) from the cluster:**

```sh
make uninstall
```

**UnDeploy the controller from the cluster:**

```sh
make undeploy
```

## Project Distribution

Following the options to release and provide this solution to the users.

### By providing a bundle with all YAML files

1. Build the installer for the image built and published in the registry:

```sh
make build-installer IMG=<some-registry>/k8s-fileinfo-operator:tag
```

**NOTE:** The makefile target mentioned above generates an 'install.yaml'
file in the dist directory. This file contains all the resources built
with Kustomize, which are necessary to install this project without its
dependencies.

2. Using the installer

Users can just run 'kubectl apply -f <URL for YAML BUNDLE>' to install
the project, i.e.:

```sh
kubectl apply -f https://raw.githubusercontent.com/<org>/k8s-fileinfo-operator/<tag or branch>/dist/install.yaml
```

### By providing a Helm Chart

1. Build the chart using the optional helm plugin

```sh
kubebuilder edit --plugins=helm/v1-alpha
```

2. See that a chart was generated under 'dist/chart', and users
can obtain this solution from there.

**NOTE:** If you change the project, you need to update the Helm Chart
using the same command above to sync the latest changes. Furthermore,
if you create webhooks, you need to use the above command with
the '--force' flag and manually ensure that any custom configuration
previously added to 'dist/chart/values.yaml' or 'dist/chart/manager/manager.yaml'
is manually re-applied afterwards.

## Contributing
// TODO(user): Add detailed information on how you would like others to contribute to this project

**NOTE:** Run `make help` for more information on all potential `make` targets

More information can be found via the [Kubebuilder Documentation](https://book.kubebuilder.io/introduction.html)

## License

Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
