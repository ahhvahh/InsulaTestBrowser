<!-- A-BIOMA-DOC-STANDARD -->
# ARQ-003 — Imagens OCI e atualização do Chromium

**Situação:** Proposto.

## Objetivo

Permitir criar uma imagem OCI reutilizável contendo Chromium e suas dependências, sem instalar Chromium no host e sem reutilizar perfis de sessões.

## Política de seleção

| Modo | Comportamento previsto |
| --- | --- |
| `latest` | Consultar versão estável mais recente disponível nas origens Debian confiáveis configuradas; criar imagem nova após verificar origem/assinaturas |
| `pinned` | Resolver versão exata disponível no repositório configurado ou snapshot validado; falhar se indisponível |
| `cached` | Usar imagem local previamente validada pelo digest, sem consulta externa |

Não confundir o commit mais novo do Chromium com versão estável e atualizada em segurança. A política `latest` depende do canal configurado e seu histórico de atualizações; a versão efetiva e a revisão Debian devem ser registradas.

## Etapas de preparação

1. Identificar modo, distribuição-base, arquitetura e versão exigida.
2. Consultar metadados do pacote e validar origem, versão e assinatura da cadeia de pacotes.
3. Construir imagem OCI incluindo Chromium, bibliotecas necessárias e mecanismo de controle.
4. Testar startup, sandbox, política de rede e screenshot antes de promover a imagem.
5. Registrar imagem pelo digest imutável; manter tag legível apenas como apelido.
6. Criar novas sessões com a imagem selecionada. Não trocar navegador de sessões em andamento.

## Imagens personalizadas

Imagens personalizadas PODEM ser criadas a partir de `Containerfile` e base permitida. NÃO DEVEM ser geradas capturando alterações de contêineres que já navegaram na web: perfis, cookies ou segredos poderiam ser incorporados à camada reutilizável.

A limpeza das imagens DEVE considerar referências de sessões ativas; uma política de expurgo para imagens antigas será definida separadamente da retenção de sessões.

## Dados e auditoria

Guardar no SQLite identificação, digest, versão efetiva, origem, momento de criação e estado de validação da imagem. Não guardar credenciais para repositórios nem manifestos contendo segredos.

## Referências

- [Modelo SQLite](../persistencia/PERS-001-sessoes-sqlite.md)
- [Decisão de versões](../decisoes/DEC-003-versoes-chromium.md)
- [Documentação Podman](https://docs.podman.io/en/stable/markdown/podman.1.html)
- [Snapshots Debian](https://snapshot.debian.org/)
