# Nodes x Pods

### Nodes são a mesma coisa que pods no kubernestes ?

Não, em Kubernetes, nós temos conceitos distintos entre **(nodes)** e **pods**.

    * Nodes (nós): São máquinas físicas ou virtuais, onde o Kubernetes é instalado. Cada nó pode ser uma máquina física ou uma máquina virtual e é onde os pods serão executados. Os nós fornecem a capacidade computacional e de armazenamento para executar os aplicativos. Eles são os blocos básicos da infraestrutura Kubernetes.

    * Pods: São a unidade básica de implantação em Kubernetes. Um pod pode conter um ou mais contêineres Docker, que são executados juntos em um ambiente compartilhado. Esses contêineres compartilham recursos de rede e armazenamento, e podem se comunicar entre si de forma eficiente. Os pods são implantados nos nós do cluster Kubernetes.

Então, enquanto os nós (nodes) fornecem a infraestrutura física ou virtual para executar os aplicativos, os pods são a menor unidade executável em Kubernetes, contendo um ou mais contêineres e compartilhando recursos em um mesmo contexto.
