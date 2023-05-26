# I got an error when doing a dotnet restore on the container, 
# so I compile the application first, then I copy it to the docker builder
FROM nginx:1.25-alpine-slim AS final
WORKDIR /usr/share/nginx/html
COPY  wasm-frontend/app/publish/wwwroot .
COPY additional_files/nginx-wasm-frontend.conf /etc/nginx/nginx.conf
