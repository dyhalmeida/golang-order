-- Antes de criar a tabela, verifique se a extensão "uuid-ossp" está habilitada.
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Agora, crie a tabela "orders" com o campo "id" como UUID.
CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    price DECIMAL NOT NULL,
    tax DECIMAL NOT NULL,
    final_price DECIMAL NOT NULL
);