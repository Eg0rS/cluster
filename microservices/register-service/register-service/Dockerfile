FROM mcr.microsoft.com/dotnet/aspnet:7.0 AS base
WORKDIR /app
EXPOSE 80
EXPOSE 443

FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /src
COPY ["register-service/register-service.csproj", "register-service/"]
RUN dotnet restore "register-service/register-service.csproj"
COPY . .
WORKDIR "/src/register-service"
RUN dotnet build "register-service.csproj" -c Release -o /app/build

FROM build AS publish
RUN dotnet publish "register-service.csproj" -c Release -o /app/publish

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
ENTRYPOINT ["dotnet", "register-service.dll"]
