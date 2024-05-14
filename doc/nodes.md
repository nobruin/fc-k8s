# Nodes no Kubernetes

Em termos técnicos, os **nodes** (nós) no Kubernetes são instâncias individuais de máquinas físicas ou virtuais em um cluster. Cada nó possui os serviços necessários para executar pods e é capaz de se comunicar com os componentes mestres do cluster para coordenar e gerenciar as cargas de trabalho.

Cada nó geralmente possui os seguintes componentes:

- **Kubelet**: O Kubelet é um agente que roda em cada nó do cluster e garante que os contêineres estejam em execução nos pods. Ele gerencia o ciclo de vida dos pods e executa as instruções fornecidas pelo API Server do Kubernetes.

- **Kube-proxy**: O Kube-proxy é um serviço de rede que roda em cada nó e é responsável por implementar o encaminhamento de rede para os serviços do Kubernetes. Ele gerencia as regras de encaminhamento de rede (como o encaminhamento de porta) e permite a comunicação de rede entre os pods dentro do cluster e com o mundo exterior.

- **Container Runtime**: O Container Runtime é o software responsável por executar os contêineres dentro dos pods. Docker é um exemplo popular de um container runtime, mas existem outros como containerd, CRI-O, entre outros.

Esses componentes trabalham juntos para fornecer a infraestrutura necessária para a execução e gerenciamento de pods em um cluster Kubernetes. Cada nó contribui com recursos de computação, armazenamento e rede para o cluster e é gerenciado de forma centralizada pelo sistema Kubernetes.
