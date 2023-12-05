curl -i -X PUT -H "Content-Type: application/json" -d '{
  "ID": "id-do-ovo-mantiqueira-que-comprei-hoje",
  "Name": "Ovo caipira mantiqueira - VERMELHO",
  "Type": "Proteina",
  "PurchaseDate": "2023-12-01",
  "Ingredient": "Ovo",
  "DueDate": "2023-12-31",
  "IsPerishable": true,
  "Quantity": {
    "Value": 12,
    "Type": "un"
  }
}' http://localhost:8080/api/groceries?id=id-do-ovo-mantiqueira-que-comprei-hoje