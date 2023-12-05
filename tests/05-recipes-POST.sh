curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Bolo de chocolate",
  "ingredients": {
    "farinha": {
      "value": 100,
      "type": "g"
    },
    "chocolate": {
      "value": 50,
      "type": "g"
    }
  },
  "yield": 10,
  "cookTime": 60
}' http://localhost:9251/api/recipes

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Arroz branco",
  "ingredients": {
    "arroz": {
      "value": 100,
      "type": "g"
    }
  },
  "yield": 5,
  "cookTime": 10
}' http://localhost:9251/api/recipes


curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Brigadeiro",
  "ingredients": {
    "chocolate": {
      "value": 100,
      "type": "g"
    },
    "leite condensado": {
      "value": 100,
      "type": "g"
    },
    "manteiga": {
      "value": 50,
      "type": "g"
    }
  },
  "yield": 20,
  "cookTime": 10
}' http://localhost:9251/api/recipes


curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Feijão preto",
  "ingredients": {
    "feijão preto": {
      "value": 100,
      "type": "g"
    },
    "cebola": {
      "value": 50,
      "type": "g"
    },
    "alho": {
      "value": 25,
      "type": "g"
    },
    "sal": {
      "value": 10,
      "type": "g"
    }
  },
  "yield": 500,
  "cookTime": 60
}' http://localhost:9251/api/recipes

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Feijoada",
  "ingredients": {
    "feijão preto": {
      "value": 200,
      "type": "g"
    },
    "carne seca": {
      "value": 200,
      "type": "g"
    },
    "linguiça": {
      "value": 100,
      "type": "g"
    },
    "cebola": {
      "value": 100,
      "type": "g"
    },
    "alho": {
      "value": 50,
      "type": "g"
    },
    "sal": {
      "value": 20,
      "type": "g"
    },
    "louro": {
      "value": 2,
      "type": "unidade"
    }
  },
  "yield": 10,
  "cookTime": 240
}' http://localhost:9251/api/recipes