<!-- A-BIOMA-DOC-STANDARD -->
# CONTR-001 — Operações de sessão e navegação

**Situação:** Proposto. Interface sem implementação. Nomes abaixo são **candidatos lógicos**, não endpoints, métodos Go ou comandos já existentes.

## Objetivo

Permitir que um cliente solicite uma sessão Chromium isolada e execute operações por intermédio do serviço Go, sem acesso direto ao socket Podman ou às portas de depuração.

## Interface conceitual

| Operação lógica | Entrada mínima proposta | Resultado proposto |
| --- | --- | --- |
| `create_session` | Política de imagem, viewport, limites/timeout | `session_id` e estado |
| `get_session` | `session_id` | Estado, versão de imagem e tempos |
| `list_sessions` | Filtros de estado/paginação | Lista resumida de sessões |
| `navigate` | `session_id`, URL, timeout | Estado de navegação e metadados não sensíveis |
| `screenshot` | `session_id`, viewport/formato/opções | Identificador de artefato |
| `interact` | `session_id`, ação tipada/elemento | Resultado ou erro |
| `end_session` | `session_id` | Resultado do encerramento/limpeza |
| `list_images` | Filtros opcionais | Imagens validadas e versões |
| `build_image` | Modo `latest`/`pinned`, versão/origem | Referência e digest de imagem |
| `purge_expired` | Política administrativa | Resumo de remoções e falhas |

As ações de interação podem contemplar localizar elemento, clicar, digitar, preencher formulário e extrair HTML. Os seletores e protocolos serão especificados depois da escolha do canal CDP.

## Regras de acesso e validação

- Interface local por UDS é a preferência arquitetural. Transporte, autenticação e serialização **ainda não foram aprovados**.
- Operações de sessão DEVEM validar autorização e posse da sessão, inclusive quando houver múltiplos clientes.
- `session_id` é opaco; nunca equivale a permissão de acesso por si só.
- A URL DEVE passar por política de destino e restrições de rede; validar redirecionamentos e DNS.
- Screenshots e downloads DEVEM usar `artifact_id`/chave interna; não aceitar caminhos arbitrários de escrita ou leitura no host.
- `end_session` DEVE ser idempotente; repetir a solicitação não pode recriar recursos.
- Não gravar credenciais, cookies ou conteúdo de formulários em logs.

## Classes de erros previstas

| Classe | Situação |
| --- | --- |
| `invalid_request` | Parâmetro incorreto ou incompatível |
| `forbidden` | Operação não autorizada ou destino de rede negado |
| `not_found` | Sessão ou imagem inexistente |
| `invalid_state` | Sessão não está pronta/foi encerrada |
| `resource_exhausted` | Limite de sessões, CPU ou memória |
| `browser_failure` | Falha ao iniciar, controlar ou validar Chromium |
| `runtime_failure` | Falha em Podman, armazenamento ou SQLite |
| `timeout` | Tempo máximo da operação atingido |

## Lacunas intencionais

Permanecem pendentes formato do payload, códigos de erro na interface concreta, modelo de permissão, transporte do CDP e mapeamento final para CLI, biblioteca Go ou MCP. Nenhuma assinatura foi estabelecida.

## Referências

- [Arquitetura](../arquitetura/ARQ-001-visao-geral.md)
- [Fluxo](../fluxos/FLUX-001-ciclo-de-vida.md)
- [Segurança](../arquitetura/ARQ-002-seguranca.md)
