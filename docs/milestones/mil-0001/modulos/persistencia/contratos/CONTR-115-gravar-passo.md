<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-115 — Persistir passo operacional

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** persistencia.

## Porta pública

Transação curta registra step_id único e aloca sequence monotônica por sessão; repetir mesmo step_id e dados devolve a sequência existente; conflito de dados falha.

### CONTR-115 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| step_id | texto | Sim |
| operation | texto | Sim |
| target | texto | Sim |
| status | texto | Sim |
### CONTR-116 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| step_id | texto | Sim |
| sequence | inteiro | Sim |
### CONTR-117 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão existe; step_id não vazio; operation e status válidos; target já sanitizado. Chave única (session_id, step_id) e (session_id, sequence) propostas; sequência alocada atomicamente.

Aplicação explícita no [fluxo](../fluxos/FLUX-0501-gravar-passo.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-115
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
  - id: CONTR-115
    nome: EntradaPersistirpassooperacional
    campos:
    - id: CMP-1037
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1038
      nome: step_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1039
      nome: operation
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1040
      nome: target
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1041
      nome: status
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
  - id: CONTR-116
    nome: SaidaPersistirpassooperacional
    campos:
    - id: CMP-1042
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1043
      nome: step_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
    - id: CMP-1044
      nome: sequence
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
  - id: CONTR-117
    nome: FalhaPersistirpassooperacional
    campos:
    - id: CMP-1045
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0501
        documento: ../fluxos/FLUX-0501-gravar-passo.md#sdd
  excecoes: []
  fluxos: []
```
