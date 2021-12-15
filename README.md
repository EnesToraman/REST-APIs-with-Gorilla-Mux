# REST API design with Gorilla Mux
Project is deployed to Heroku. Accessible via URL below:

[go-task-v2.herokuapp.com](https://go-task-v1.herokuapp.com/)
## Endpoints

- /users/
- /products/
- /payments/
- /brands/
- /customers/
- /baskets/

### Query Parameters

- /users/ -> id, email, username, isActive
- /products/ -> sku, name, price, stock
- /payments/ -> userID, amount, discount, tax
- /brands/ -> id, name, productQty, totalWorth
- /customers/ -> id, userID, purchaseAmount, orderQty
- /baskets/ -> userID, productID, sku, quantity

#### Example - endpoints with query parameters

- /users/?id=5
- /products/?sku=11004545
- /payments/?userID=2
