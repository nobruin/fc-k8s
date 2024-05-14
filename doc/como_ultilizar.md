# Como subir Cluster Kubernetes, Deployment e Service neste projeto

## Passo 1: Instalar o Docker

Se você ainda não tiver o Docker instalado, você pode seguir as [instruções de instalação na documentação oficial do Docker](https://docs.docker.com/get-docker/) para o seu sistema operacional.

## Passo 2: Instalar o Kubernetes

Se você ainda não tiver o Kubernetes instalado, siga as [instruções de instalação na documentação oficial do Kubernetes](https://kubernetes.io/docs/setup/) para o seu ambiente. Dependendo das suas necessidades, você pode escolher entre diferentes ferramentas de provisionamento, como Minikube, kubeadm ou kind.

## Passo 3: Clonar o Repositório

```bash
git clone git@github.com:nobruin/fc-k8s.git
```

## Passo 4: Acessar o Diretório do Projeto

```bash
cd fc-k8s
```

## Passo 5: Criar os Recursos do Kubernetes

No diretório do projeto, você encontrará os arquivos de configuração YAML que definem os recursos do Kubernetes, como deployments e services. Você pode usar o comando `kubectl apply` para criar esses recursos:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## Passo 6: Criar os Recursos do Kubernetes

No diretório do projeto, você encontrará os arquivos de configuração YAML que definem os recursos do Kubernetes, como deployments e services. Você pode usar o comando `kubectl apply` para criar esses recursos:

```bash
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

## Passo 7: Verificar o Status dos Recursos

Após aplicar os arquivos de configuração, você pode usar o comando `kubectl get` para verificar o status dos recursos criados:

```bash
kubectl get pods
kubectl get services
```

## Passo 8: Testar a Aplicação

Se tudo estiver configurado corretamente, você pode testar a aplicação acessando os endpoints fornecidos pelo serviço.

Certifique-se de ajustar os arquivos de configuração YAML conforme necessário para corresponder às suas necessidades específicas.
