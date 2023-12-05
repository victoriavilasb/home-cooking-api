curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Ovo caipira mantiqueira",
  "type": "Proteina",
  "purchase_date": "2023-12-01",
  "ingredient": "ovo",
  "due_date": "2023-12-31",
  "is_perishable": true,
  "quantity": {
    "value": 20,
    "type": "un"
  }
}' http://localhost:9251/api/groceries

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Linguiça calabresa",
  "type": "Alimento",
  "purchase_date": "2023-12-01",
  "ingredient": "linguica",
  "due_date": "2023-12-31",
  "is_perishable": false,
  "quantity": {
    "Value": 500,
    "type": "g"
  }
}' http://localhost:9251/api/groceries

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Carne seca",
  "type": "Alimento",
  "purchase_date": "2023-12-01",
  "ingredient": "carne seca",
  "due_date": "2023-12-31",
  "is_perishable": false,
  "quantity": {
    "Value": 500,
    "type": "g"
  }
}' http://localhost:9251/api/groceries

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Feijão preto",
  "type": "Alimento",
  "purchase_date": "2023-12-01",
  "ingredient": "feijão preto",
  "due_date": "2023-12-31",
  "is_perishable": false,
  "quantity": {
    "Value": 1000,
    "type": "g"
  }
}' http://localhost:9251/api/groceries

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Chocolate ao leite",
  "type": "Alimento",
  "purchase_date": "2023-12-01",
  "ingredient": "chocolate",
  "due_date": "2023-12-31",
  "is_perishable": false,
  "quantity": {
    "Value": 1000,
    "type": "g"
  }
}' http://localhost:9251/api/groceries

curl -i -X POST -H "Content-Type: application/json" -d '{
  "name": "Farinha de trigo",
  "type": "Alimento",
  "purchase_date": "2023-12-01",
  "ingredient": "farinha",
  "due_date": "2023-12-31",
  "is_perishable": false,
  "quantity": {
    "Value": 500,
    "type": "g"
  }
}' http://localhost:9251/api/groceries