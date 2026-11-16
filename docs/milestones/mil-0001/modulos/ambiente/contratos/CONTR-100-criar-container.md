<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-100 — Criar contêiner

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** ambiente.

## Porta pública

Criar um contêiner rootless rotulado com session_id a partir do digest e perfil autorizados; não iniciar navegação.

### CONTR-100 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| image_digest | texto | Sim |
| profile_name | texto | Sim |
### CONTR-101 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| container_id | texto | Sim |
### CONTR-102 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

session_id é componente opaco não vazio, sem /, .. ou controle; image_digest é digest OCI imutável permitido; profile_name existe na configuração confiável.

Aplicação explícita no [fluxo](../fluxos/FLUX-0200-criar-container.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-100
  modulo: ambiente
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-100
    nome: EntradaCriarcontêiner
    campos:
    - id: CMP-1001
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
    - id: CMP-1002
      nome: image_digest
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
    - id: CMP-1003
      nome: profile_name
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
  - id: CONTR-101
    nome: SaidaCriarcontêiner
    campos:
    - id: CMP-1004
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
    - id: CMP-1005
      nome: container_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
  - id: CONTR-102
    nome: FalhaCriarcontêiner
    campos:
    - id: CMP-1006
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0200
        documento: ../fluxos/FLUX-0200-criar-container.md#sdd
  excecoes: []
  fluxos: []
```
