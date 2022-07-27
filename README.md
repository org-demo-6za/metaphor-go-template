# Metaphor GO API

`metaphor` is an example application which serves as a demonstration space for how your applications hook into your 
infrastructure and tooling.

`metaphor` is deployed to all of your environments just like your other applications. This means that when you make 
changes to your ci/cd, you can test it out using an application that works just like your applications do.

`metaphor` currently demonstrates the following capabilities:
- building a docker container
 (WIP) - publishing a docker container to ecr
 (WIP) - publishing a prerelease helm chart
 (WIP) - gitops delivery of metaphor to different namespaces
 (WIP) - a release process that publishes a release chart and patches the chart version for the next release
 (WIP) - secrets sourced from vault
 (WIP) - certificate management using cert-manager
 (WIP) - automatic dns management using external-dns

# CI/CD

(WIP) `metaphor` has 5 ci stages defined in its .gitlab-ci.yml file.
(WIP) 
(WIP) - branch: branch jobs run when your branch pushes to origin and report status to your merge requests
(WIP) - publish: publishes your docker container to ecr and publishes your prerelease chart to chartmuseum
(WIP) - development: sets the desired chart version for development
(WIP) - staging: sets the desired chart version for staging
(WIP) - release: publishes a release chart, sets the desired chart version for production, and patches chart in source for the next release
(WIP) 
(WIP) `argocd` is the gitops tool responsible for autosyncing the desired state in each environment. It follows a pull model so our CI/CD ecosystem doesn't need to know how to connect to our kubernetes clusters.
(WIP) 
(WIP) We have `metaphor` set up to run its automation by invoking argo-workflows. Those submitted workflows are in the .argo directory of this repository.
