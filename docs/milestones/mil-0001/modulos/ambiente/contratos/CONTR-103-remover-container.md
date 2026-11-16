<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# CONTR-103 — Remover contêiner

**Estado:** em refinamento. **Versão aprovada:** nenhuma. **Módulo:** ambiente.

## Porta pública

Parar e remover somente o recurso rotulado com a sessão; ausência prévia é sucesso idempotente.

### CONTR-103 — Entrada

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| container_id | texto | Sim |
### CONTR-104 — Saida

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| session_id | texto | Sim |
| removed | booleano | Sim |
### CONTR-105 — Falha

| Campo | Tipo | Obrigatório |
| --- | --- | --- |
| code | texto | Sim |

## Regras e falhas

IDs não vazios; o container_id foi associado à session_id pelo serviço. Adaptador verifica rótulo real antes de remover.

Aplicação explícita no [fluxo](../fluxos/FLUX-0201-remover-container.md). Campos de saída são montados pelo provedor e verificados pelo chamador. Falha usa code fixo sanitizado; nunca carrega payload da página ou segredo. Tipos externos inválidos devem ser rejeitados na entrada antes de operações dependentes do tipo. Não há chamada recursiva por referência de campo.

## Acesso

Porta interna do serviço Go. O ponto de composição injeta o adaptador e o contexto confiável; IDs não concedem permissão. Não é um endpoint público independente. Veja a matriz em [ARQ-005](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).


## Formato piloto

Este documento usa **Spec Rules 0.2.0**, proposta local da MIL-0001. A versão identifica o novo formato, não uma release do plugin nem compatibilidade com o schema SDD 2.0. As referências abreviadas são resolvidas pelo [catálogo do piloto](../../../referencia/REF-003-spec-rules.md). Estado documental: **em refinamento**; implementação não comprovada.
<a id="sdd"></a>
<a id="spec-rules"></a>
## Spec Rules

<!-- A-BIOMA-SPEC-RULES:0.2.0 -->

```yaml
CONTR-103:
  formato: "a-bioma-spec-rules"
  versao: "0.2.0"
  modulo: "ambiente"
  estado: "em refinamento"
  implementacao: "nao implementado"
  objetivo: "Parar e remover somente o contêiner associado à sessão; ausência prévia é sucesso idempotente."
  acesso:
    tipo: "porta interna"
    contexto: "Injetado pelo serviço; IDs não concedem autorização."
    politica: "Apenas chamadas públicas permitidas pela ARQ-005."
  entrada:
    nome: "solicitacao"
    doc: "CONTR-103"
  saida:
    doc: "CONTR-104"
  contratos:
    CONTR-103:
      nome: "EntradaRemoverConteiner"
      campos:
        - id: "CMP-1007"
          nome: "session_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
        - id: "CMP-1008"
          nome: "container_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
    CONTR-104:
      nome: "SaidaRemoverConteiner"
      campos:
        - id: "CMP-1009"
          nome: "session_id"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
        - id: "CMP-1010"
          nome: "removed"
          tipo: "@TIP-0102"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
    CONTR-105:
      nome: "FalhaRemoverConteiner"
      campos:
        - id: "CMP-1011"
          nome: "code"
          tipo: "@TIP-0100"
          obrigatorio: true
          regras:
            - "@[ambiente](FLUX-0201)"
  regras:
    - "Campos têm aplicação explícita pelo fluxo; referências declarativas não executam chamadas nem criam recursão."
    - "Resposta normal preserva session_id e confirma removed=true."
    - "Código de falha é sanitizado; não inclui dados da página nem segredos."
  fluxo:
    atores:
      - ator: "chamador"
        descricao: "Caso de uso autorizado que consome a porta."
      - ator: "ambiente"
        descricao: "Módulo que valida o pedido e adapta a operação."
      - ator: "podman"
        descricao: "Runtime rootless que administra o recurso."
    root:
      - id: "CHAMADA-103-01"
        acao: "chamar"
        de: "chamador"
        para: "ambiente"
        descricao: "Executar o fluxo encapsulado com o contrato de entrada."
        doc: "@[ambiente](FLUX-0201)"
        argumentos:
          solicitacao:
            caminho: "entrada.solicitacao"
        guardar_em: "variaveis.resultado"
        erros:
          - excecao: "@[ambiente](EXC-0201)"
            guardar_em: "falhas.remocao"
            passos:
              - id: "CHAMADA-103-02"
                acao: "falhar"
                de: "ambiente"
                para: "chamador"
                contrato: "CONTR-105"
                excecao: "@[ambiente](EXC-0201)"
                dados:
                  caminho: "falhas.remocao"
      - id: "CHAMADA-103-03"
        acao: "retornar"
        de: "ambiente"
        para: "chamador"
        contrato: "CONTR-104"
        valor:
          caminho: "variaveis.resultado"
  excecoes:
    - doc: "@[ambiente](EXC-0201)"
      contrato: "@[ambiente](CONTR-105)"
  efeitos:
    - "Parar e remover o contêiner pertencente à sessão, conforme o fluxo encapsulado."
  criterios:
    - "Entrada inválida não aciona remoção."
    - "Ausência prévia é sucesso idempotente."
    - "Vínculo incompatível não permite remover outro contêiner."
    - "Falha do runtime ou resposta incompatível produz CONTR-105."
  pendencias:
    - "Proposta do formato piloto; schema e migração geral ainda pendentes."
    - "Implementação do adaptador e testes de integração pendentes."
```
