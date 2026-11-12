<!-- A-BIOMA-DOC-STANDARD -->
# DEC-001 — Podman Rootless e Chromium por sessão

**Decisão:** adotada para o projeto.  
**Implementação:** não verificada.

## Contexto

O InsulaTestBrowser deve executar navegação automatizada sem instalar um navegador no host e sem permitir que diferentes sessões compartilhem cookies, arquivos ou perfis. Segurança e separação têm prioridade sobre a velocidade de inicialização.

## Decisão

Usar Podman em modo rootless sob um usuário de serviço dedicado. A aplicação Go gerencia um contêiner descartável por sessão, com Chromium e dependências inclusos na imagem OCI. As operações de Podman serão realizadas por socket Unix local.

## Alternativas consideradas

- **Chromium no host com perfis distintos:** menor sobrecarga, mas fronteira de isolamento insuficiente.
- **Docker Rootless:** viável, porém Podman Rootless foi selecionado.
- **VM por sessão:** isolamento do kernel mais forte, com maior custo operacional; permanece opção de endurecimento futuro.

## Consequências

- Precisamos administrar imagens OCI, quotas, rede e limpeza de contêineres.
- O host precisa de Podman, mas não precisa do pacote Chromium instalado.
- Contêineres compartilham o kernel do host; não equivalem a máquinas virtuais.
- A execução só será considerada segura após testes de sandbox, user namespaces e política de rede.
- O socket Podman e o arquivo SQLite ficam acessíveis somente ao serviço, nunca ao Chromium.

## Referências

- [Arquitetura](../arquitetura/ARQ-001-visao-geral.md)
- [Segurança](../arquitetura/ARQ-002-seguranca.md)
- [Podman Rootless](https://docs.podman.io/en/stable/markdown/podman.1.html)
