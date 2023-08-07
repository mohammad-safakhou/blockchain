CREATE TABLE IF NOT EXISTS blocks
(
    block_hash          TEXT PRIMARY KEY,
    previous_block_hash TEXT    NOT NULL,
    merkle_root         TEXT    NOT NULL,
    nonce               INTEGER NOT NULL,
    difficulty          INTEGER NOT NULL,
    height              INTEGER NOT NULL,
    timestamp           INTEGER NOT NULL,
    size                INTEGER NOT NULL,
    version             INTEGER NOT NULL,
    bits                INTEGER NOT NULL,
    transactions        JSON    NOT NULL,
    transaction_count   INTEGER NOT NULL,
    transaction_volume  INTEGER NOT NULL,
    transaction_fees    INTEGER NOT NULL,
    created_at          INTEGER NOT NULL,
    updated_at          INTEGER NOT NULL
);