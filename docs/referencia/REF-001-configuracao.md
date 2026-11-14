<!-- A-BIOMA-DOC-STANDARD -->
# REF-001 — Arquivo de configuração do serviço

**Situação:** Proposto. Exemplo e campos abaixo são um **contrato de configuração planejado**, não um parser ou arquivo implementado.

## Objetivo e localização

O executável InsulaTestBrowser DEVE carregar um arquivo YAML com nome fixo `config.yaml`. A resolução do arquivo deve ser previsível e independente do diretório corrente do terminal.

1. Sem argumento, procurar `config.yaml` **na pasta do executável** (localização efetiva do binário).
2. Com `--config-dir <diretorio>`, procurar exclusivamente `<diretorio>/config.yaml`.
3. Se o caminho explícito for inválido ou o arquivo estiver ausente, falhar com erro claro; **não recorrer silenciosamente** à pasta do executável.
4. Não criar arquivo com defaults implícitos quando não for encontrado. Fazer validação antes de iniciar Podman, SQLite ou o listener.
5. Resolver caminhos relativos de banco, sockets, arquivos de segredos, scripts de preparo e volumes **a partir da pasta do arquivo de configuração**, nunca do diretório corrente de execução.
6. Ao iniciar, registrar somente o caminho efetivo da configuração e sua validação, sem expor credenciais.

O parâmetro `--config-dir` e o nome `config.yaml` são **propostas de interface** aprovadas para documentação; nenhuma implementação foi verificada. Consulte [INTF-TERM-001](../interfaces/terminal/INTF-TERM-001-inicializacao.md).

## Exemplo ilustrativo de `config.yaml`

Os nomes dos campos são a proposta inicial; versões, digests e IDs abaixo são **placeholders**, não valores reais.

```yaml
config_version: 1

database:
  path: ./data/insula-test-browser.sqlite

access:
  socket: ./run/itb.sock
  allowed_uids: [1001]
  credentials:
    mode: file
    path: ./secrets/service.token

podman:
  socket: /run/user/1001/podman/podman.sock
  rootless_required: true

containers:
  default_profile: chromium-base
  profiles:
    chromium-base:
      source: image
      image: localhost/insula-test-browser:stable
      mounts: []                  # volumes adicionais permitidos, privados por sessão

    chromium-preparado:
      source: prepare
      base_image: "debian:13@sha256:<DIGEST_VALIDADO>"
      setup_file: ./container-setup.yaml
      save_image:
        enabled: true
        name: localhost/insula-test-browser:prepared
      mounts: []

scripts:                       # proposta de configuração do executor; não implementada
  enabled: false               # somente habilitar após validação de sandbox/canal CDP
  api_version: 1
  max_concurrent: 2
  max_steps: 1000000
  max_duration: 30m
  max_pages: 100
  max_downloads: 500
  max_download_bytes: 52428800
  max_dom_nodes: 500
  allow_inline_source: false

sessions:
  max_concurrent: 5
  max_duration: 30m
  retention:
    browser_profile: 0m
    screenshots: 1h
    downloads: 1h
    logs: 1h
    metadata: 24h
```

## Campos e responsabilidades

| Grupo | Responsabilidade |
| --- | --- |
| `database.path` | Local do banco SQLite no host |
| `access.socket` e `access.allowed_uids` | Ponto de acesso local e restrição de clientes |
| `access.credentials` | Referência para credencial de serviço fora do YAML; método definitivo de autenticação ainda depende do protocolo |
| `podman.socket` | Socket Podman Rootless do usuário executor |
| `containers.default_profile` | Perfil escolhido quando não for indicado outro |
| `containers.profiles.*.source: image` | Instanciar nova sessão a partir de imagem OCI previamente validada |
| `containers.profiles.*.source: prepare` | Inicializar **contêiner de preparação**, executar instruções autorizadas e opcionalmente salvar nova imagem reutilizável |
| `containers.profiles.*.mounts` | Política declarativa de montagens; nunca compartilhar perfil Chromium, segredos ou dados de sessão |
| `containers.profiles.*.save_image` | Autorizar snapshot OCI do contêiner de preparação depois de sanitização e validação |
| `sessions` | Concorrência, limite de execução e prazos de retenção |
| `scripts` | Regras do executor Starlark: habilitação, versão da API, limites de passos/tempo/páginas/downloads/DOM e permissão de fonte inline (propostas) |

O formato definitivo de `container-setup.yaml`, regras de montagem, restrições de clientes e extensões de autenticação serão especificados quando a interface estiver pronta. **Não existe execução arbitrária de shell implícita a partir do YAML de exemplo.**

Os limites efetivos do script DEVEM ser os mais restritivos entre `scripts`, perfil, usuário e solicitação externa. Os novos campos são candidatos de esquema, ainda não validados ou implementados; a interface detalhada está em [ARQ-004](../arquitetura/ARQ-004-motor-starlark.md) e [CONTR-003](../contratos/CONTR-003-administracao-scripts.md).

## Segurança e validação

- O arquivo de configuração e a pasta de segredos DEVEM ter permissões restritas ao usuário do serviço. A validação de propriedade/permissões deve impedir exposição de credenciais; recomendação inicial: `0600` para arquivos de segredos.
- Senhas e tokens não devem ser registrados diretamente no `config.yaml`; usar **referências a arquivos privados** ou mecanismo equivalente de segredos.
- Não serializar credenciais em SQLite, logs ou imagens OCI. Arquivos apontados por `access.credentials` **não** serão montados nos contêineres.
- Caminhos indicados para banco, segredos e artefatos devem ser normalizados e validados; impedir escrita/leitura fora das áreas autorizadas.
- Contêineres de preparação e de navegação são papéis diferentes: captura de imagem é autorizada apenas para preparação, não para uma sessão que acessou sites.
- A política de imagens e sua validação ficam em [ARQ-003](../arquitetura/ARQ-003-imagens-chromium.md); os detalhes de isolamento ficam em [ARQ-002](../arquitetura/ARQ-002-seguranca.md).

## Erros propostos

`config_not_found`, `config_invalid`, `permission_denied`, `secret_unavailable`, `invalid_profile` e `podman_unavailable`. Códigos de saída da CLI e formato do erro ainda não foram implementados.

## Referências

- [Interface de inicialização](../interfaces/terminal/INTF-TERM-001-inicializacao.md)
- [Arquitetura](../arquitetura/ARQ-001-visao-geral.md)
- [Preparação de imagens](../fluxos/FLUX-002-preparacao-imagem.md)
- [Decisão de configuração](../decisoes/DEC-004-configuracao-local.md)
