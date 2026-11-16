<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-121 — Navegar e registrar resultado

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** sessoes.

## Porta pública

Coordenar navegação e gravação do passo; falha no histórico não repete a ação no navegador.

### CONTR-121 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| url | texto | Sim |
| timeout_ms | inteiro | Sim |
| step_id | texto | Sim |
### CONTR-122 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| url | texto | Sim |
| title | texto | Sim |
| document_revision | inteiro | Sim |
### CONTR-123 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão pronta e autorizada no contexto do chamador; step_id emitido pelo serviço e único; URL, aba e timeout dentro da política. Não aceitar contexto de autorização fornecido pelo script.

Aplicação explícita no [fluxo](../fluxos/FLUX-0101-navegar-com-historico.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-121
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
  - id: CONTR-121
    nome: EntradaNavegareregistrarresultado
    campos:
    - id: CMP-1052
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1053
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1054
      nome: url
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1055
      nome: timeout_ms
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1056
      nome: step_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
  - id: CONTR-122
    nome: SaidaNavegareregistrarresultado
    campos:
    - id: CMP-1057
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1058
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1059
      nome: url
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1060
      nome: title
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
    - id: CMP-1061
      nome: document_revision
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
  - id: CONTR-123
    nome: FalhaNavegareregistrarresultado
    campos:
    - id: CMP-1062
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0101
        documento: ../fluxos/FLUX-0101-navegar-com-historico.md#sdd
  excecoes: []
  fluxos: []
```
