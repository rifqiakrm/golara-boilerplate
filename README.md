# GOLARA Boilerplate

Golang  project that can connect to laravel with passport library

It comes pre-configured with :

1. Gin Gonic Router (https://github.com/gin-gonic/gin)
2. JWT-GO (https://github.com/dgrijalva/jwt-go)
3. Viper (https://github.com/spf13/viper)
4. Cobra (https://github.com/spf13/cobra)
5. Testify (https://github.com/stretchr/testify)

## Setup

Use this command to install the blueprint

```bash
go get github.com/rifqiakrm/golara-boilerplate
```

or manually clone the repo and then run `go run main.go`.

## Quick Note

Before you start the main service, you may want to set your environtment variables. You can choose it on the config, fill the env key and then set the env path file on `cmd/root.go`

```golang
rootCMD.PersistentFlags().StringVar(&cfgFile, "configs", "configs/config.{the_choosen_env}.toml", "configs file (example is $HOME/configs.toml)")
```

## Good ol oauth-public.key and oauth-private.key
To make this golang project can connect the token from laravel passport you have to copy the `oauth-private.key` and `oauth-public.key` from your laravel project and then paste it to `config/rsa-key`

## Step by step deploying to server?
Thanks to the internet and Tabvn, you can follow this guide to deploy your application to the server 

> https://medium.com/@tabvn/deploy-golang-application-on-digital-ocean-server-ubuntu-16-04-b7bf5340ccd9