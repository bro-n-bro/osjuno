CREATE TYPE POOL_SWAP AS
(
    /* todo pool_id as reference */
    pool_id         BIGINT,
    in_coin_denom   TEXT,
    in_coin_amount  TEXT,
    out_coin_denom  TEXT,
    out_coin_amount TEXT
);


CREATE TABLE swap
(
    transaction_hash TEXT        NOT NULL REFERENCES transaction (hash),
    sender           TEXT        NOT NULL REFERENCES account (address),
    routes           POOL_SWAP[] NOT NULL
);

CREATE INDEX swap_transaction_hash_index ON swap (transaction_hash);
CREATE INDEX swap_sender_index ON swap (sender);
/* todo add routes indexes */