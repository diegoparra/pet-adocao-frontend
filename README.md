# Pet Adocao Frontend

### Ambiente de desenvolvimento:

Atualmente utilizamos a ferramenta [air](https://github.com/cosmtrek/air) para fazer auto reload das nossas mudancas. Existe um arquivo `.air.conf` com as configuracoes necessarias de incluir ou excluir pastas ou arquivos do reload.

### Build da Imagem local e deploy para Oracle Cloud

Atualmente estamos rodando uma maquina na Oracle Cloud e rodando o watch tower escutando mudancas no repositorio docker, entao ao simples fato de fazer um build e push para o docker registry sera o suficiente para realizar o deployment.

```
docker build -t diegoparra/reikianos:pet-frontend -f Dockerfile-arm . ; docker push diegoparra/reikianos:pet-frontend
```
