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

## Perfis de criação e imagens personalizadas

O arquivo [config.yaml](../referencia/REF-001-configuracao.md) centraliza a escolha do perfil:

- **Reuso de imagem** (`source: image`): uma imagem OCI validada é a origem para novos contêineres de navegação; não existe clonagem de cookies, cache ou perfil Chromium.
- **Preparação de novo ambiente** (`source: prepare`): iniciar um contêiner separado de preparação, aplicar instruções locais previamente autorizadas e validar o navegador.
- **Salvar imagem preparada** (`save_image.enabled: true`): capturar o filesystem do contêiner de preparação como nova imagem OCI reutilizável, registrando digest e proveniência. A imagem só pode ser disponibilizada depois das verificações de segurança.

A captura de um contêiner iniciado é possível com **`podman commit`**. Por padrão, o Podman pausa o contêiner enquanto faz o commit e **não inclui conteúdo dos volumes montados**. A operação salva alterações de filesystem: não inclui o estado em memória dos processos, nem deve ser apresentada como retomada de uma sessão aberta. Ver [FLUX-002](../fluxos/FLUX-002-preparacao-imagem.md).

O projeto **NÃO DEVE capturar sessões que já navegaram em sites**, porque cookies, credenciais e outros dados privados podem ter sido persistidos em arquivos ou camadas. Evitar introduzir segredos durante a preparação; uma limpeza posterior pode não garantir que dados antigos não apareçam em camadas da imagem.

Imagens também PODEM ser criadas diretamente de um `Containerfile` e base aprovada. A limpeza das imagens DEVE considerar referências de sessões ativas; uma política de expurgo para imagens antigas será definida separadamente da retenção de sessões.

## Dados e auditoria

Guardar no SQLite identificação, digest, versão efetiva, origem, momento de criação e estado de validação da imagem. Não guardar credenciais para repositórios nem manifestos contendo segredos.

## Referências

- [Modelo SQLite](../persistencia/PERS-001-sessoes-sqlite.md)
- [Configuração e perfis](../referencia/REF-001-configuracao.md)
- [Fluxo de preparação](../fluxos/FLUX-002-preparacao-imagem.md)
- [Podman commit](https://docs.podman.io/en/latest/markdown/podman-commit.1.html)
- [Decisão de versões](../decisoes/DEC-003-versoes-chromium.md)
- [Documentação Podman](https://docs.podman.io/en/stable/markdown/podman.1.html)
- [Snapshots Debian](https://snapshot.debian.org/)
