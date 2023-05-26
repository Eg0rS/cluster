FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["curator-service.csproj", "."]
RUN dotnet restore "./curator-service.csproj"
COPY . .
WORKDIR "/src/."
RUN dotnet build "curator-service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "curator-service.csproj" -c Release -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "curator-service.dll"]