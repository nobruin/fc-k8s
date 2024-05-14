# Services no Kubernetes:

Os **pods** no Kubernetes (uma abstração que detém o domínio sobre um container) são efêmeros porque existe um ciclo de vida que determina inclusive a destruição de um pod. Logo, não é sensato que haja preocupação por parte do sistema com endereço IP ou características particulares de um pod. Para isso, os **ReplicaSets** fazem com que cada vez que um pod morra, outro seja criado no seu lugar, partindo do mesmo ponto que o pod destruído.

Já o **Service** no Kubernetes é uma abstração que define um conjunto lógico de Pods e uma política pela qual acessá-los. Serviços permitem um baixo acoplamento entre os Pods dependentes. Um serviço é definido usando YAML ou JSON, como todos os manifestos de objetos Kubernetes. O conjunto de Pods selecionados por um Service é geralmente determinado por um seletor de rótulos que irá definir os alcances daquele serviço, ou seja, quais daqueles pods estarão debaixo das configurações definidas por um serviço.

Embora cada Pod tenha um endereço IP único, esses IPs não são expostos externamente ao **cluster** sem um objeto Service. Objetos Service permitem que suas aplicações recebam tráfego. Services podem ser expostos de formas diferentes especificando um tipo (campo type) na especificação do serviço (campo spec):

- **ClusterIP (padrão)**: Expõe o serviço sob um endereço IP interno no cluster. Este tipo de serviço é acessível somente dentro do cluster.

- **NodePort**: Expõe o serviço sob a mesma porta em cada nó selecionado no cluster usando NAT. Torna o serviço acessível externamente ao cluster usando o endereço <NodeIP>:<NodePort>. É um superconjunto do tipo ClusterIP.

- **LoadBalancer**: Cria um balanceador de carga externo no provedor de nuvem atual (se suportado) e atribui um endereço IP fixo e externo para o serviço. É um superconjunto do tipo NodePort.

- **ExternalName**: Mapeia o Service para o conteúdo do campo externalName (por exemplo, foo.bar.example.com), retornando um registro DNS do tipo CNAME com o seu valor. Nenhum tipo de proxy é configurado. Este tipo requer a versão 1.7 ou mais recente do kube-dns, ou o CoreDNS versão 0.0.8 ou superior.

Profissionalmente, os tipos de serviços mais comuns no Kubernetes são ClusterIP, NodePort e LoadBalancer. Aqui está o porquê:

- **ClusterIP**: Este é o tipo de serviço mais comum para comunicação entre microserviços dentro do cluster. Ele fornece uma maneira confiável para os pods se comunicarem uns com os outros, usando um endereço IP interno que é acessível apenas dentro do cluster. É útil para serviços que precisam ser acessados apenas internamente e não devem ser expostos externamente.

- **NodePort**: Embora não seja recomendado para produção devido a questões de segurança, o NodePort é comumente usado em ambientes de desenvolvimento e teste. Ele fornece uma maneira fácil de acessar serviços diretamente por meio de um endereço IP e porta atribuída a cada nó do cluster. Isso pode ser útil durante o desenvolvimento para acessar rapidamente os serviços sem configurar um balanceador de carga ou DNS.

- **LoadBalancer**: Este é o tipo de serviço mais usado para expor aplicativos para o tráfego externo em ambientes de produção. Ele provisiona automaticamente um balanceador de carga externo (por exemplo, na AWS, GCP, Azure) e roteia o tráfego para os pods do serviço. É escalável, confiável e oferece balanceamento de carga para garantir alta disponibilidade e desempenho para os aplicativos.

Em resumo, o ClusterIP é o mais comum para comunicação interna entre os serviços de um cluster, o NodePort é usado ocasionalmente em ambientes de desenvolvimento e teste, e o LoadBalancer é preferido para exposição de aplicativos para o tráfego externo em ambientes de produção devido à sua escalabilidade e confiabilidade.
