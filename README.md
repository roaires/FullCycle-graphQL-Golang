# Sobre

Essa implementação tem como objetivo fixar conhecimento adquirido durante o módulo de Comunicação via graphGL da Formação Full Cycle by Code.Edu.

Dessa forma, será demonstrada a implementação de um Server graphGL em Goland. Quanto ao Client será utilizado o playground fornecido pelo Gqlgen.


# Preparação de ambiente

- https://gqlgen.com/getting-started

# Para executar o servidor GraphQL

```
 go run server.go
```

# Utilização do serviço criado
- Acesse o playground: http://localhost:8080/

## Exemplo inclusões

- Primeira inclusão:
```
mutation createCategory1 {
  createCategory(input : {name: "Arquitetura"}) {
    id
    name
    description
  }
}
```
Output:
```
{
  "data": {
    "createCategory": {
      "id": "T5577006791947779410",
      "name": "Arquitetura",
      "description": null
    }
  }
}
```

- Segunda inclusão:
```
mutation createCategory2 {
  createCategory(input : {name: "Comunição", description: "Comunição em microsserviços"}) {
    id
    name
    description
  }
}
```
Output:
```
{
  "data": {
    "createCategory": {
      "id": "T8674665223082153551",
      "name": "Comunição",
      "description": "Comunição em microsserviços"
    }
  }
}
```

## Exemplo de consulta
```
query findCategories {
  categories {
    id
    description
  }
}
```
Output:
```
{
  "data": {
    "categories": [
      {
        "id": "T5577006791947779410",
        "description": null
      },
      {
        "id": "T8674665223082153551",
        "description": "Comunição em microsserviços"
      }
    ]
  }
}
```