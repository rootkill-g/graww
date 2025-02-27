
# Graww 🦁

Catalogue your products with ease and security using **Graww** 🦁.

  - Uses *JSON Web Token* for **Authentication**

Available **APIs**:

  - **GET** */product/info*
  - **POST** */product/info*
  - **PATCH** */product/info*

## Tech stack 🤖

* **GO** programming language *(golang)*

## Get Started 🚀

To run the cataloger you must build the code:

```bash
go build
```

Then run the **Graww** 🦁 app:

```bash
./graww
```

The cataloger will run on port `8080`.

## Get Product Information 🏷️

To get the product's information, you must send the `JWT`` for `auth` with `Bearer` token (We used a mock value for Bearer Token):

```bash
curl -X -H "Authorization: Bearer 1234567890" GET http://localhost:8080/product/info?productId=F5
```

## Add Product to Catalogue 🆕

To add a new product to the catalogue, the payload must have `id`, `name`, `price`, and `descriptionn` of the product:

```bash
curl -X POST -H "Authorization: Bearer 1234567890" -H 'Content-Type: application/json' -d '{"id":"SGS25U","name":"Samsung Galaxy S25 Ultra", "price":"$1299", "description":"Simple. Impactful."}' http://localhost:8080/product/info
```

## Update Product information 📦

Is there a discount on a product? You may need to edit a product's information. Don't worry, **Graww** 🦁 app handles this efficiently using the `PATCH` verb:

```bash
curl -X PATCH -H "Authorization: Bearer 1234567890" -H 'Content-Type: application/json' -d '{"price":"$999"}' http://localhost:8080/product/info?productId=SGS25U
```

<hr />

Happy Cataloging 🦁
