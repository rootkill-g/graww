
# Graww 🦁

Catalogue your products with ease using **Graww** 🦁.

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

To get the product's information, you must send the `productId` as query to the URI:

```bash
curl -X GET http://localhost:8080/product/info?productId=F5
```

## Add Product to Catalogue 🆕

To add a new product to the catalogue, the payload must have, `name` of the product, the `price` of the product, and the `description` of the product, along with the `productId` in the query:

```bash
curl -X POST -H 'Content-Type: application/json' -d '{"name":"Samsung Galaxy S25 Ultra", "price":"$1299", "description":"Simple. Impactful."}' http://localhost:8080/product/info?productId=SGS25U
```

## Update Product information 📦

Is there a discount on a product? The **Graww** 🦁 app handles this efficiently using the `PATCH` verb:

```bash
curl -X PATCH -H 'Content-Type: application/json' -d '{"price":"$999"}' http://localhost:8080/product/info?productId=SGS25U
```

---

Happy Cataloging 🛸
