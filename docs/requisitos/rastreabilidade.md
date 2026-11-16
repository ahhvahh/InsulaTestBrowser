<!-- A-BIOMA-DOC-STANDARD -->
<!-- A-BIOMA-DOC-REVISADO -->
<!-- A-BIOMA-DOC-VERSION {"plugin":"a-bioma-doc-standard","version":"2.0.0"} -->
# Rastreabilidade

| História | SPEC/CA | Contrato/fluxo | Milestone | Código | Teste | Situação |
| --- | --- | --- | --- | --- | --- | --- |
| [US-001](../milestones/mil-0001/requisitos/historias/US-001-consultar-artefatos.md) | [SPEC-001 / CA-01 a CA-08](../milestones/mil-0001/modulos/automacao/especificacoes/SPEC-001-catalogo-artefatos.md) | [CONTR-004](../milestones/mil-0001/modulos/automacao/contratos/CONTR-004-artefatos-resultados.md); [FLUX-0003](../milestones/mil-0001/modulos/automacao/fluxos/FLUX-0003-validar-session-id.md) cobre somente restrição lexical | [MIL-0001](../milestones/mil-0001.md) | Ausente | Sem evidência de teste do catálogo | Proposta em refinamento |

Os [testes existentes](../../tests/starlark/runner_test.go) usam mocks e não validam os critérios do catálogo. Não foram executados nesta alteração documental.

A decomposição por camada está em [ARQ-005](../milestones/mil-0001/arquitetura/ARQ-005-arquitetura-hexagonal.md): FLUX-0101 navega e registra passo; FLUX-0401/0502 catalogam artefato. Não representam cobertura integral de CA-01 a CA-08; expurgo, consulta e recuperação continuam pendentes.
