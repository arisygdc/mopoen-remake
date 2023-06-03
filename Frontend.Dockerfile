FROM mcr.microsoft.com/dotnet/sdk:6.0 AS build
WORKDIR /wasm-frontend
COPY wasm-frontend/wasm-frontend.csproj .
RUN dotnet restore wasm-frontend.csproj
COPY wasm-frontend/* .
RUN dotnet build wasm-frontend.csproj -c Release -o /app/build

FROM build AS publish
RUN dotnet publish wasm-frontend.csproj -c Release -o /app/publish

FROM nginx:1.25-alpine-slim AS final
WORKDIR /usr/share/nginx/html
COPY  --from=publish /app/publish/wwwroot .
COPY additional_files/nginx-wasm-frontend.conf /etc/nginx/nginx.conf
