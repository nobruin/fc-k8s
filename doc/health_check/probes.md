# Sondas de Saúde no Kubernetes

As sondas de saúde, também conhecidas como probes de saúde, são mecanismos importantes no Kubernetes para garantir a disponibilidade e a saúde dos pods de uma aplicação. Existem três tipos principais de probes de saúde:

1. **Probe de leiturainess (readiness probe)**: Esta sonda é usada para determinar se um pod está pronto para servir tráfego. Se um pod falhar na sondagem de leiturainess, ele será removido do serviço de balanceamento de carga, garantindo que apenas pods saudáveis recebam tráfego. A sonda de leiturainess é frequentemente usada para verificar se um pod iniciou corretamente e está pronto para receber solicitações.

2. **Probe de verificação de integridade (liveness probe)**: Esta sonda é usada para determinar se um pod está em um estado saudável ou não. Se um pod falhar na sondagem de integridade, o Kubernetes o reiniciará automaticamente. Isso ajuda a garantir que os pods continuem a funcionar corretamente mesmo após falhas internas, como travamentos ou panes. A sonda de integridade é comumente usada para verificar se um aplicativo está funcionando conforme o esperado.

3. **Probe de inicialização (startup probe)**: Esta sonda é usada para determinar se um pod está pronto para receber sondagens de leiturainess e integridade. Ela é usada durante o período de inicialização de um pod para verificar se o aplicativo dentro do contêiner está pronto para iniciar. A sonda de inicialização ajuda a evitar que os pods recebam tráfego antes de estarem prontos para servir solicitações.

Cada sonda de saúde pode ser configurada com base em um determinado tipo de verificação, como HTTP, TCP ou execução de um comando dentro do contêiner. Você pode definir parâmetros específicos, como intervalo de sondagem, tempo limite e número máximo de falhas permitidas antes que um pod seja considerado inoperante.

Usar sondas de saúde no Kubernetes é fundamental para garantir a disponibilidade contínua e a resiliência das aplicações em ambientes de produção. Elas ajudam a identificar e lidar automaticamente com problemas nos pods, garantindo assim uma melhor experiência para os usuários finais.

## Exemplos

- 1 **Probe de leiturainess (readiness probe):**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: myapp-container
    image: nginx
    ports:
    - containerPort: 80
    readinessProbe:
      httpGet:
        path: /healthz
        port: 80
      initialDelaySeconds: 5
      periodSeconds: 10
```

- 2 **Probe de verificação de integridade (liveness probe):**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: myapp-container
    image: nginx
    ports:
    - containerPort: 80
    livenessProbe:
      httpGet:
        path: /healthz
        port: 80
      initialDelaySeconds: 10
      periodSeconds: 15

```

- 3 **Probe de inicialização (startup probe)::**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: myapp-container
    image: nginx
    ports:
    - containerPort: 80
    startupProbe:
      httpGet:
        path: /healthz
        port: 80
      initialDelaySeconds: 10
```

Esses exemplos demonstram como configurar sondas de saúde em um pod no Kubernetes. Cada sonda é configurada com um tipo de verificação (HTTP neste caso), um caminho de verificação específico, uma porta e parâmetros como atraso inicial e período de verificação. Você pode ajustar esses parâmetros conforme necessário para atender aos requisitos da sua aplicação
