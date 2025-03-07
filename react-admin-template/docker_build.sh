#!/bin/bash

# $ docker network create template_go_react_network
# $ docker network ls
# $ docker network inspect template_go_react_network

# Запускается уже на запущенном контейнере
# docker exec -it template_go_react_react_cui sh


# ping works:
# $ docker exec -it template_go_react_react_cui sh
# /app # ping template_go_react_golang

# But when I open application in web browser and try to login in the app, it still cannot reach backend:
# http://template_go_react_golang:8080/api/v1/pub/login


    # When you access the React application from your browser, 
    #     the frontend code runs in the browser's context, not inside the Docker container. 
    #     Therefore, it cannot resolve Docker container names directly. 
    #     The browser needs to communicate with the backend via a URL that it can resolve, 
    #     typically localhost or the host machine's IP address.



    # Solution: Use a Proxy in the React Development Server
    #     One common solution is to use a proxy in the React development server 
    #         to forward API requests to the backend container. 
    #         This way, you can keep the frontend API calls consistent, and the proxy will handle the resolution.

    # Here's how you can set up the proxy:


            # Add Proxy to Vite Configuration:

            # Edit your Vite configuration file (vite.config.js or vite.config.ts) to include a proxy setup:

            # javascript
            # Copy code
            # // vite.config.js
            # export default {
            #   server: {
            #     proxy: {
            #       '/api': {
            #         target: 'http://template_go_react_golang:8080',
            #         changeOrigin: true,
            #         secure: false,
            #         rewrite: (path) => path.replace(/^\/api/, ''),
            #       },
            #     },
            #   },
            # };

            # This configuration proxies any request starting with /api to http://template_go_react_golang:8080.            


# Update API Calls in React App:

# Ensure your API calls in the React app use the /api prefix:

# javascript
# Copy code
# fetch('/api/v1/pub/login', {
#   method: 'POST',
#   headers: {
#     'Content-Type': 'application/json',
#   },
#   body: JSON.stringify({ username, password }),
# })
# .then(response => response.json())
# .then(data => {
#   // Handle the response data
# })
# .catch(error => {
#   console.error('Error:', error);
# });



docker build -t template_go_react_react_cui .

docker run --name template_go_react_react_cui --network template_go_react_network -p 3000:3000 template_go_react_react_cui
