<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-127 — Baixar arquivo

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** navegacao.

## Porta pública

Solicitar download monitorado e aguardar conclusão; publicar arquivo privado com identificador após checagem de tamanho, tipo real e checksum.

### CONTR-127 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| tab_id | texto | Sim |
| url | texto | Sim |
| timeout_ms | inteiro | Sim |
### CONTR-128 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| artifact_id | texto | Sim |
| storage_key | texto | Sim |
| mime_type | texto | Sim |
| size_bytes | inteiro | Sim |
| sha256 | texto | Sim |
### CONTR-129 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

Sessão/aba autorizadas; URL permitida inclusive redirects; limites de bytes e tempo efetivos; raiz privada configurada pelo host; storage_key é gerada pelo serviço e nunca pelo cliente. O adaptador de download deve aplicar estas regras e limpar parciais antes de relatar falha.

Aplicação explícita no [fluxo](../fluxos/FLUX-0302-baixar-arquivo.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: CONTR-127
  modulo: navegacao
  estado: proposto
  tipos:
  - id: TIP-0100
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0101
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  - id: TIP-0102
    documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
  contratos:
  - id: CONTR-127
    nome: EntradaBaixararquivo
    campos:
    - id: CMP-1072
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1073
      nome: tab_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1074
      nome: url
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1075
      nome: timeout_ms
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
  - id: CONTR-128
    nome: SaidaBaixararquivo
    campos:
    - id: CMP-1076
      nome: session_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1077
      nome: artifact_id
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1078
      nome: storage_key
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1079
      nome: mime_type
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1080
      nome: size_bytes
      tipo:
        id: TIP-0101
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
    - id: CMP-1081
      nome: sha256
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
  - id: CONTR-129
    nome: FalhaBaixararquivo
    campos:
    - id: CMP-1082
      nome: code
      tipo:
        id: TIP-0100
        documento: ../../../global/contratos/CONTR-090-tipos-primitivos.md#sdd
      obrigatorio: true
      regras:
      - id: FLUX-0302
        documento: ../fluxos/FLUX-0302-baixar-arquivo.md#sdd
  excecoes: []
  fluxos: []
```
