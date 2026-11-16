<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-112 — Persistir sessão provisionada

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** persistencia.

## Porta pública

Transação parametrizada registra vínculo da sessão com contêiner e imagem; retry com mesmos IDs é idempotente; vínculo incompatível falha.

### CONTR-112 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| container_id | texto | Sim |
| image_digest | texto | Sim |
### CONTR-113 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| saved | booleano | Sim |
### CONTR-114 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

IDs e digest não vazios; unicidade de session_id e container_id; imagem já cadastrada por digest. Estado gravado é starting, pois a prontidão Chromium ainda não foi comprovada.

Aplicação explícita no [fluxo](../fluxos/FLUX-0500-salvar-sessao.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-112
  modulo: persistencia
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-112
    nome: EntradaPersistirsessãoprovisionada
    campos:
    - id: CMP-1031
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
    - id: CMP-1032
      nome: container_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
    - id: CMP-1033
      nome: image_digest
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
  - id: CONTR-113
    nome: SaidaPersistirsessãoprovisionada
    campos:
    - id: CMP-1034
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
    - id: CMP-1035
      nome: saved
      tipo:
        id: TIP-0102
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
  - id: CONTR-114
    nome: FalhaPersistirsessãoprovisionada
    campos:
    - id: CMP-1036
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0500
        documento: ../fluxos/FLUX-0500-salvar-sessao.md#sdd
  excecoes: []
  fluxos: []
```
