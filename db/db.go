package db

import (
	"context"
	"encoding/json"
	"github.com/jackc/pgx/v5"
	"github.com/mohammad-safakhou/blockchain/app"
	"github.com/mohammad-safakhou/blockchain/block"
)

type DB struct {
	db *pgx.Conn
}

// AddBlock adds a block to the database.
func (d *DB) AddBlock(ctx context.Context, block block.Block) error {
	v := 0.0
	f := 0.0
	for _, t := range block.Transactions {
		v += t.Amount
		f += t.Fee
	}
	t, _ := json.Marshal(block.Transactions)
	_, err := d.db.Exec(ctx, "INSERT INTO blocks (block_hash,previous_block_hash,merkle_root,nonce,difficulty,height,timestamp,size,version,bits,transactions,transaction_count,transaction_volume,transaction_fees,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,NOW(),NOW())",
		block.Header.Hash,
		block.Header.PrevHash,
		block.Header.MerkleRoot,
		block.Header.Nonce,
		block.Header.Difficulty,
		block.Header.Height,
		block.Header.Timestamp,
		block.Header.Size,
		app.Version,
		block.Header.Bits,
		string(t),
		len(block.Transactions),
		v,
		f)
	return err
}

// LastBlock returns the last block in the database.
func (d *DB) LastBlock() block.Block {
	var b block.Block
	_ = d.db.QueryRow(context.Background(), "SELECT block_hash,previous_block_hash,merkle_root,nonce,difficulty,height,timestamp,size,bits,transactions FROM blocks ORDER BY height DESC LIMIT 1").Scan(
		&b.Header.Hash,
		&b.Header.PrevHash,
		&b.Header.MerkleRoot,
		&b.Header.Nonce,
		&b.Header.Difficulty,
		&b.Header.Height,
		&b.Header.Timestamp,
		&b.Header.Size,
		&b.Header.Bits,
		&b.Transactions,
	)
	return b
}

// New returns a new DB.
func New(db *pgx.Conn) *DB {
	return &DB{
		db: db,
	}
}
