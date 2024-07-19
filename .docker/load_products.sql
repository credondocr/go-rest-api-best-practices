COPY products(id, name, description, price)
FROM '/docker-entrypoint-initdb.d/products.csv'
DELIMITER ','
CSV HEADER;