CREATE TABLE IF NOT EXISTS product
(
  "id" uuid PRIMARY KEY,
  "name" VARCHAR NOT NULL,
  "price" INT NOT NULL,
  "stock" INT NOT NULL,
  "image_url" TEXT NOT NULL,
  "condition" VARCHAR NOT NULL,
  "is_purchasable" boolean NOT NULL,
  "tags" text[] NOT NULL,
  "created_at" timestamptz,
  "updated_at" timestamptz
);
