<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-118 — Provisionar sessão

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** sessoes.

## Porta pública

Criar contêiner e persistir associação; compensar criação quando persistência falhar.

### CONTR-118 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| image_digest | texto | Sim |
| profile_name | texto | Sim |
### CONTR-119 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| container_id | texto | Sim |
### CONTR-120 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Entrada vem de caso de uso autenticado; session_id foi emitido pelo serviço, é único, opaco e válido; perfil e digest foram autorizados. Operação ainda não declara navegador pronto.

Aplicação explícita no [fluxo](../fluxos/FLUX-0100-provisionar.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-118
  modulo: sessoes
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-118
    nome: EntradaProvisionarsessão
    campos:
    - id: CMP-1046
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
    - id: CMP-1047
      nome: image_digest
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
    - id: CMP-1048
      nome: profile_name
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
  - id: CONTR-119
    nome: SaidaProvisionarsessão
    campos:
    - id: CMP-1049
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
    - id: CMP-1050
      nome: container_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
  - id: CONTR-120
    nome: FalhaProvisionarsessão
    campos:
    - id: CMP-1051
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0100
        documento: ../fluxos/FLUX-0100-provisionar.md#sdd
  excecoes: []
  fluxos: []
```
