# Como subir Cluster Kubernetes, Deployment e Service neste projeto

## Passo 1: Instalar o Docker

Se você ainda não tiver o Docker instalado, você pode seguir as [instruções de instalação na documentação oficial do Docker](https://docs.docker.com/get-docker/) para o seu sistema operacional.

## Passo 2: Instalar o Kind e Kubernetes

### Instalando Kind (Kubernetes in Docker)

1. Instale o Kind:
   - Linux/macOS:
     ```bash
     curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.20.0/kind-linux-amd64
     chmod +x ./kind
     sudo mv ./kind /usr/local/bin/kind
     ```
   - Windows (com chocolatey):
     ```bash
     choco install kind
     ```

2. Instale o kubectl:
   - Linux/macOS:
     ```bash
     curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
     chmod +x kubectl
     sudo mv kubectl /usr/local/bin/
     ```
   - Windows (com chocolatey):
     ```bash
     choco install kubernetes-cli
     ```

### Configuração e Criação do Cluster

1. O projeto já inclui um arquivo de configuração do Kind em `k8s/kind.yml`. Este arquivo define a configuração necessária para o cluster, incluindo mapeamentos de porta e outras configurações específicas do projeto.

2. Crie o cluster usando o arquivo de configuração do projeto:
   ```bash
   kind create cluster --config k8s/kind.yml
   ```

3. Verifique se o cluster está funcionando:
   ```bash
   kubectl cluster-info
   ```

4. Verifique os nodes do cluster:
   ```bash
   kubectl get nodes
   ```

## Passo 3: Clonar o Repositório

```bash
git clone git@github.com:nobruin/fc-k8s.git
```

## Passo 4: Acessar o Diretório do Projeto

```bash
cd fc-k8s
```

## Passo 5: Criar os Recursos do Kubernetes

Antes de aplicar os recursos, verifique se o cluster está funcionando:

```bash
kubectl cluster-info
```

Se tudo estiver ok, aplique os recursos na seguinte ordem:
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## Passo 6: Verificar o Status dos Recursos

```bash
kubectl get pods
kubectl get services
```

## Passo 7: Testar a Aplicação

Se tudo estiver configurado corretamente, você pode testar a aplicação acessando os endpoints fornecidos pelo serviço.

## Troubleshooting para Kind

Se você encontrar problemas, aqui estão algumas soluções comuns:

1. Erro ao criar o cluster:
   - Certifique-se de que o Docker está rodando
   - Verifique se o arquivo `k8s/kind.yml` existe e está correto
   - Tente recriar o cluster:
     ```bash
     kind delete cluster
     kind create cluster --config k8s/kind.yml
     ```

2. Verificar o status do cluster Kind:
   ```bash
   kind get clusters
   kubectl cluster-info
   ```

3. Problemas com imagens locais:
   ```bash
   kind load docker-image minha-imagem:tag
   ```

4. Ver logs do cluster:
   ```bash
   kubectl logs -n kube-system kube-apiserver-kind-control-plane
   ```

5. Verificar eventos do cluster:
   ```bash
   kubectl get events --all-namespaces
   ```

6. Problemas de rede:
   - Verifique se as portas necessárias estão disponíveis
   - Confirme se o Docker tem permissões de rede adequadas
   - Verifique os logs do Docker:
     ```bash
     docker logs kind-control-plane
     ```

7. Para verificar os logs de um pod específico:
   ```bash
   kubectl logs <nome-do-pod>
   ```

8. Se necessário, reinicie o Docker e recrie o cluster:
   ```bash
   systemctl restart docker
   kind delete cluster
   kind create cluster --config k8s/kind.yml
   ```

Para mais informações sobre troubleshooting com Kind, consulte a [documentação oficial do Kind](https://kind.sigs.k8s.io/docs/user/quick-start/).

## Estrutura do Projeto

```
.
├── k8s/
│   ├── kind.yml           # Configuração do cluster Kind
│   ├── deployment.yaml    # Configuração do Deployment
│   └── service.yaml       # Configuração do Service
└── ...
```

Certifique-se de que todos os arquivos YAML estejam na pasta `k8s/` e que os caminhos nos comandos correspondam à estrutura do seu projeto.
