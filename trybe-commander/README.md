## Trybe Commander.js
https://github.com/tj/commander.js

### Instalação
1. Baixe o binário do seu sistema operacional que está na pasta bin
2. Jogue o binário para a pasta `/usr/local/bin`
    - ex: `mv trybe-cli /usr/local/bin`
3. Entre um pasta de projeto JS
4. Execute o comando `trybe-cli test`

### Build

Para realizar o build do CLI basta executar o comando `npm run build-[linux,mac]`. Esse comando irá executar o `ncc` e posteriormente o `pkg`, criando a basta `bin` com o binário do CLI.