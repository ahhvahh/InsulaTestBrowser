<!-- DOC-BACKLOG -->
<!-- DOC-MIL-0001 -->
<!-- DOC-STATUS REFINING -->

[Milestone MIL-0001](../../../../mil-0001.md).

# FLUX-0500 — Persistir sessão provisionada

**Estado:** proposto. **Versão aprovada:** nenhuma. **Módulo:** persistencia.

## Objetivo e pré-condições

Transação parametrizada registra vínculo da sessão com contêiner e imagem; retry com mesmos IDs é idempotente; vínculo incompatível falha.

IDs e digest não vazios; unicidade de session_id e container_id; imagem já cadastrada por digest. Estado gravado é starting, pois a prontidão Chromium ainda não foi comprovada.

Contratos: [entrada, saída e falha](../contratos/CONTR-112-salvar-sessao.md). Assinatura pública, consumível somente conforme [matriz de dependências](../../../arquitetura/ARQ-005-arquitetura-hexagonal.md).

## Sequência

```mermaid
sequenceDiagram
    participant U as Chamador
    participant A as persistencia
    participant B as SQLite
    U->>A: Contrato de entrada
    A->>A: Validar campos e contexto
    alt Entrada inválida
        A-->>U: invalid_request
    else Entrada válida
    A->>B: Operação contratada
    alt Dependência confirma
        B-->>A: Resultado tipado
        A->>A: Validar saída
        A-->>U: Resultado ou falha de contrato
    else Dependência falha
        B-->>A: Falha tipada
        A-->>U: Falha sanitizada
    end
    end
```

## Falhas, efeitos e telemetria

Não há retry automático. Cada chamada trata todas as falhas declaradas. Timeouts podem deixar efeito externo incerto; reconciliar antes de repetir. Correlacionar execução, session_id, step_id quando presente, fluxo, PAS e resultado; nunca registrar argumentos completos, tokens ou conteúdo de formulários. Logs técnicos do adaptador não substituem o histórico funcional.

## Critérios documentais

Cobrir sucesso, rejeição antes de efeito, falha da dependência e resposta incompatível. Não existem testes executados nem implementação destas portas. Em provisionamento, testar falha de persistência e falha da compensação; na navegação, testar histórico indisponível após efeito confirmado. As portas de CDP, autorização e migração SQL ainda precisam de implementação.


<a id="sdd"></a>
## SDD — Spec-Driven Development

<!-- A-BIOMA-SPEC:2.0 -->

```yaml
especificacao:
  formato: a-bioma-comportamento
  versao: '2.0'
  documento: FLUX-0500
  modulo: persistencia
  estado: proposto
  tipos: []
  contratos:
  - id: CONTR-112
    documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
  - id: CONTR-113
    documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
  - id: CONTR-114
    documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
  excecoes:
  - id: EXC-0500
    nome: Falha0500
    condicao: Validação, autorização, timeout ou falha da dependência desta porta.
    resultado: Encerrar a operação com código sanitizado, preservando efeitos já confirmados.
    recuperavel: false
    contrato:
      id: CONTR-114
      documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
  fluxos:
  - id: FLUX-0500
    nome: Persistir sessão provisionada
    categoria: metodo
    entradas:
    - nome: solicitacao
      contrato:
        id: CONTR-112
        documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
    saida:
      id: CONTR-113
      documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
    excecoes:
    - id: EXC-0500
    efeitos: &id001
    - Persistir vínculo em sessions
    passos:
    - id: PAS-1027
      acao: validar
      descricao: Validar tipos e regras de entrada antes de acessar dependências.
      condicao:
        descricao: IDs e digest não vazios; unicidade de session_id e container_id;
          imagem já cadastrada por digest. Estado gravado é starting, pois a prontidão
          Chromium ainda não foi comprovada. Todos os campos presentes e com os tipos
          declarados; rejeitar antes de interpretar valores.
        dados:
        - caminho: entradas.solicitacao
      falha:
        id: EXC-0500
      dados_falha:
        literal:
          code: invalid_request
    - id: PAS-1029
      acao: operar
      descricao: Transação parametrizada registra vínculo da sessão com contêiner
        e imagem; retry com mesmos IDs é idempotente; vínculo incompatível falha.
        O adaptador implementa o protocolo e confirma o efeito; invariantes transacionais
        estão descritas nesta porta.
      operacao: sqlite.salvar_sessao
      entradas:
      - nome: pedido
        contrato:
          id: CONTR-112
          documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
      argumentos:
      - nome: pedido
        valor:
          caminho: entradas.solicitacao
      saida:
        id: CONTR-113
        documento: ../contratos/CONTR-112-salvar-sessao.md#sdd
      guardar_em: variaveis.resultado
      excecoes:
      - id: EXC-0500
      erros:
      - origem:
          id: EXC-0500
        guardar_em: falhas.operacao
        passos:
        - id: PAS-1028
          acao: falhar
          descricao: Propagar somente falha tipada e sanitizada do adaptador.
          excecao:
            id: EXC-0500
          dados:
            caminho: falhas.operacao
      efeitos: *id001
    - id: PAS-1030
      acao: validar
      descricao: Verificar contrato de saída antes de retornar.
      condicao:
        descricao: Resultado contém todos os campos obrigatórios nos tipos declarados;
          session_id corresponde à entrada; IDs de recurso não vazios quando presentes;
          revision e sequence não negativos; saved/removed verdadeiros quando presentes.
        dados:
        - caminho: variaveis.resultado
        - caminho: entradas.solicitacao
      falha:
        id: EXC-0500
      dados_falha:
        literal:
          code: invalid_dependency_result
    - id: PAS-1031
      acao: retornar
      descricao: Entregar a única saída normal desta porta.
      valor:
        caminho: variaveis.resultado
```
