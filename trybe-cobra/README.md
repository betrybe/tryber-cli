## Trybe Cobra
https://github.com/spf13/cobra

### Instalação
1. Baixe o binário do seu sistema operacional que está na pasta bin
2. Jogue o binário para a pasta `/usr/local/bin`
    - ex: `mv trybe-cli /usr/local/bin`
3. Entre um pasta de projeto JS
4. Execute o comando `trybe-cli test`

### Build

Para realizar o build do CLI basta executar o comando `go build`

```sh
~ GOOS=linux GOARCH=amd64 go build -o bin/mac/trybe-cli
# ou
~ GOOS=darwin GOARCH=amd64 go build -o bin/linux/trybe-cli
```