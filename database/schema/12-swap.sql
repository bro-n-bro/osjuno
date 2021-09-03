CREATE TABLE swap
(
    timestamp           TIMESTAMP WITHOUT TIME ZONE NOT NULL REFERENCES block (timestamp),
    hash                TEXT    NOT NULL REFERENCES transaction (hash),
    height              BIGINT  NOT NULL REFERENCES block (height),

    /* Msgs */
    routes              JSONB   NOT NULL DEFAULT '[]'::JSONB,
    sender              VARCHAR(43) NOT NULL,
    token_in_denom      TEXT NOT NULL,
    token_in_amount     NUMBER NOT NULL DEFAULT 0,

    /* Tx response */
    gas_wanted          BIGINT           DEFAULT 0,
    gas_used            BIGINT           DEFAULT 0,
    token_out_denom     TEXT NOT NULL,
    token_out_amount    NUMBER NOT NULL DEFAULT 0,
);
