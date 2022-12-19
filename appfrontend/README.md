# **Frontend**

## App-vote-react

Provides a web based frontend written in React. The web application provides a platforms voting feature where end users can vote for 1 of 3 platforms (Kubernetes, Rancher, and OpenShift). The React based web application is designed to be compiled and containerised, and eventually deployed into an OpenShift based Kubernetes cluster, exposed using an Ingress Controller. The web application generates AJAX requests which are sent to a publicly exposed API hosted on the same cluster. The API is written in Go and reads/writes to a MongoDB database, also hosted on the same cluster using a StatefulSet setup.

