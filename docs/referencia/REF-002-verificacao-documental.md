<!-- A-BIOMA-DOC-STANDARD -->
<!-- A-BIOMA-DOC-REVISADO -->
<!-- A-BIOMA-DOC-VERSION {"plugin":"a-bioma-doc-standard","version":"2.0.0"} -->
# REF-002 — Verificação documental

Revisão de 2026-11-16, base `7a13bc974cabeffe89a469b1cf106103f834f02e`, branch `docs/estrutura-a-bioma-v2`.

## Resultado e limites

A estrutura de refinamento foi aplicada aos 21 arquivos de conteúdo existentes em docs, preservando as origens. A formalização comportamental completa permanece pendente. O [inventário](../milestones/mil-0001.md) informa origem, documento criado, destino final e lacunas por documento.

| Verificação | Resultado |
| --- | --- |
| Links relativos de arquivos Markdown | Existência dos destinos conferida automaticamente; URLs externas não foram testadas nesta rodada |
| JSON ilustrativo EX-002 | Sintaxe JSON conferida; não representa captura real |
| FLUX-0003 | YAML conferido contra comportamento.schema.json do padrão 2.0; revisão semântica de contratos locais, IDs, campos, sucesso, falha e término |
| Demais contratos/fluxos | Sem marca de validação 2.0; faltam assinaturas completas, campos/regras e exceções tipadas |
| Marcadores da milestone | Uma ocorrência de BACKLOG, MIL-0001 e STATUS REFINING em cada Markdown de trabalho e no principal |
| Preservação das bases | Conteúdo original preservado, com marca LEGADO e link à proposta; sem remoção de arquivos |
| Escopo | Somente README.md e docs; código, testes e CI preservados |
| Diagramas Mermaid | Revisão textual; não houve renderização ou lint Mermaid |
| Testes da aplicação | Não executados: alteração documental; mocks existentes não comprovam catálogo, Chromium, Podman ou UDS |

## Publicação

Documentos de trabalho permanecem propostos na [MIL-0001](../milestones/mil-0001.md). Não houve publicação nos destinos finais, exclusão de legados ou merge. O fluxo tipado é proposta documental; a validação do YAML não comprova implementação nem aprovação funcional.

## Segunda rodada — arquitetura hexagonal

Foram acrescentados cinco módulos, doze fluxos e doze documentos de contratos de porta, além do catálogo de tipos global. Sessões coordena provisionamento com compensação e navegação com histórico. Foram redistribuídas as entidades e decisões propostas, mantendo links e IDs. ARQ-005 registra a exceção local de interfaces públicas autorizada na conversa.

Validação documental: schema 2.0, resolução de referências, IDs novos sem colisão, links relativos, marcadores e chamadas permitidas. A sintaxe dos diagramas foi revisada textualmente, sem renderização. Não foram executados testes da aplicação. Lacunas dos contratos agregados e operações restantes permanecem descritas na milestone.
