<!-- A-BIOMA-DOC-STANDARD -->
<!-- A-BIOMA-DOC-REVISADO -->
<!-- A-BIOMA-DOC-VERSION {"plugin":"a-bioma-doc-standard","version":"2.0.0"} -->
# Documentação — InsulaTestBrowser

**Estado do projeto:** propostas arquiteturais e testes com mocks; serviço não implementado. A arquitetura documental 2.0 está organizada para revisão na MIL-0001; a formalização completa dos comportamentos continua pendente.

## Comece por aqui

| Área | Documento |
| --- | --- |
| Trabalho em andamento e mapa origem/destino | [MIL-0001](milestones/mil-0001.md) |
| Índice central e histórico | [Milestones](milestones/README.md) |
| Módulo e contratos | [Módulos hexagonais](milestones/mil-0001/modulos/README.md) |
| Especificação funcional e aceite | [Catálogo de artefatos](milestones/mil-0001/modulos/automacao/especificacoes/SPEC-001-catalogo-artefatos.md) |
| Exemplo completo de fluxo SDD 2.0 | [Validar session_id](milestones/mil-0001/modulos/automacao/fluxos/FLUX-0003-validar-session-id.md) |
| Arquitetura | [Visão geral proposta](milestones/mil-0001/arquitetura/ARQ-001-visao-geral.md) |
| Histórico de requisitos | [Requisitos](requisitos/README.md) |
| Rastreabilidade e evidências | [Matriz](requisitos/rastreabilidade.md) |
| Validação desta revisão | [Relatório](referencia/REF-002-verificacao-documental.md) |

## Arquitetura da MIL-0001

[ARQ-005 — fronteiras, portas e chamadas](milestones/mil-0001/arquitetura/ARQ-005-arquitetura-hexagonal.md) apresenta a arquitetura hexagonal e a divisão de responsabilidades.

## Organização

`docs/milestones/mil-0001/` espelha os destinos finais em `docs/`: arquitetura, referencia, guias, modulos por responsabilidade, global e requisitos/historias. Após conclusão validada, documentos são publicados nesses destinos e o principal da milestone permanece no histórico. Tipos primitivos compartilhados ficam em global. Não há histórias concluídas nem plano de implementação aprovado.

## Prefixos e catálogo

| Prefixo | Conteúdo |
| --- | --- |
| MIL / US / SPEC | Milestone, história e especificação funcional |
| ARQ / REF | Arquitetura e referência |
| CONTR / FLUX | Contrato e fluxo |
| TIP / CMP / EXC / PAS | Tipo, campo, exceção e passo definidos no SDD |
| PERS / ENT | Relações de persistência e entidade |
| DEC / INTF-TERM | Decisão técnica e interface CLI |
| TEST / EX / PROMPT | Cenário de teste, exemplo e prompt |

O [inventário da milestone](milestones/mil-0001.md) cataloga os documentos e bases. TIP-0001, CONTR-005/006, CMP-0001/0002, EXC-0001 e PAS-0001/0002 são definidos exclusivamente no YAML de FLUX-0003. IDs preservados em propostas têm vínculo explícito com a revisão-base; não são duas definições vigentes.

## Convenções

Implementado exige evidência de código; proposto indica desenho; não verificado indica lacuna. DEVE/NÃO DEVE/RECOMENDA-SE/PODE são convenções internas inspiradas nas RFC 2119/8174. A marca 2.0 não equivale a implementação ou certificação ISO. Os originais fora da milestone são bases legadas preservadas, não a nova arquitetura publicada.
