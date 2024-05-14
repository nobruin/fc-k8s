# Sobre o projeto:


Este projeto consiste em provisionar um servidor simples em Golang, contendo apenas alguns endpoints básicos, com o intuito de demonstrar a configuração e funcionamento dos principais objetos no Kubernetes. Para isso, será necessário primeiro instalar o Kubernetes em seu ambiente de desenvolvimento.

A instalação do Kubernetes pode variar dependendo do seu sistema operacional e das ferramentas disponíveis. No entanto, uma das maneiras mais comuns de instalar o Kubernetes é usando ferramentas como Minikube, kubeadm ou kind, que permitem criar um cluster Kubernetes local em sua máquina.

Minikube é uma ferramenta que permite executar um único cluster Kubernetes localmente em sua máquina. É útil para desenvolvimento, testes e aprendizado do Kubernetes em um ambiente controlado.

kubeadm é uma ferramenta de linha de comando que facilita a inicialização de um cluster Kubernetes. É recomendado para cenários de produção ou quando você precisa de mais controle sobre a configuração do cluster.

kind (Kubernetes IN Docker) é uma ferramenta que permite criar clusters Kubernetes usando contêineres Docker como "nós". É uma ótima opção para desenvolvimento local e testes de integração.

Escolha a ferramenta que melhor se adapta às suas necessidades e siga as instruções de instalação fornecidas pela documentação oficial do Kubernetes ou da ferramenta escolhida. Uma vez que o Kubernetes esteja instalado em seu ambiente, você estará pronto para configurar e implantar o servidor Golang e explorar os conceitos do Kubernetes neste projeto.

Para instruções detalhadas sobre como instalar o Kubernetes, consulte a [documentação oficial do Kubernetes](https://kubernetes.io/docs/setup/).



Durante o desenvolvimento e implantação do servidor Golang no Kubernetes, você encontrará alguns comandos essenciais que ajudarão na configuração, gerenciamento e depuração do ambiente. Abaixo estão os principais comandos que você provavelmente utilizará:

1. **kubectl create**: Este comando é usado para criar recursos no cluster Kubernetes, como pods, serviços, deployments, entre outros.

2. **kubectl apply**: Semelhante ao comando 'create', mas é usado para aplicar alterações em recursos existentes ou criar novos recursos se ainda não existirem.

3. **kubectl get**: Este comando é usado para obter informações sobre os recursos no cluster, como pods, serviços, deployments, entre outros. Por exemplo, 'kubectl get pods' exibe uma lista de todos os pods no cluster.

4. **kubectl describe**: Fornece detalhes sobre um recurso específico no cluster. Por exemplo, 'kubectl describe pod <nome-do-pod>' exibe informações detalhadas sobre um pod específico.

5. **kubectl logs**: Permite visualizar os logs de um contêiner em um pod específico. Por exemplo, 'kubectl logs <nome-do-pod>' exibe os logs do contêiner principal em um pod.

6. **kubectl exec**: Este comando é usado para executar comandos dentro de um contêiner em um pod específico. Por exemplo, 'kubectl exec -it <nome-do-pod> -- <comando>' executa um comando dentro do contêiner especificado.

7. **kubectl delete**: Este comando é usado para excluir recursos no cluster. Por exemplo, 'kubectl delete pod <nome-do-pod>' exclui um pod específico do cluster.

Esses são apenas alguns dos principais comandos que você usará ao trabalhar com Kubernetes. Familiarizar-se com esses comandos será fundamental para a configuração e gerenciamento eficaz do seu ambiente Kubernetes durante o desenvolvimento e implantação do servidor Golang.
