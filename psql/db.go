package psql

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"

	"log"
	"os"
	"sync"
)

var (
	pgxC *PGXContainer
	once sync.Once
)

type PGXContainer struct {
	mu   sync.RWMutex
	conn *pgxpool.Pool
}

func GetInstance() *PGXContainer {
	if pgxC == nil {
		once.Do(func() {
			pgxC = newPGX()
		})
	}

	return pgxC
}

func newPGX() *PGXContainer {
	conn, err := pgxpool.Connect(context.Background(), "postgresql://localhost:5432/postgres")
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return &PGXContainer{
		mu:   sync.RWMutex{},
		conn: conn,
	}
}

//func (cont *PGXContainer) GetTables(uids []string) ([]db.Table, error) {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//
//	uidsStr := "("
//	for i := 0; i < len(uids); i++ {
//		uidsStr += "'" + uids[i] + "', "
//	}
//	uidsStr = uidsStr[:len(uidsStr) - 2] + ")"
//	query := "SELECT * FROM tables WHERE uid in " + uidsStr
//	var tables []db.Table
//	err := pgxscan.Select(context.Background(), cont.conn, &tables, query)
//	if err != nil {
//		return nil, err
//	}
//
//	return tables, nil
//}
//
//func (cont *PGXContainer) InsertTable(tables []db.Table) error {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//
//	ctx := context.Background()
//	var batch pgx.Batch
//	for i := 0; i < len(tables); i++ {
//		batch.Queue("insert into tables(uid, data, chat_id, message_id, channel, type, link) VALUES($1, $2, $3, $4, $5, $6, $7);",
//			tables[i].UID, tables[i].Data, tables[i].ChatId, tables[i].MessageId, tables[i].Channel, tables[i].Type, tables[i].Link)
//	}
//
//	result := cont.conn.SendBatch(ctx, &batch)
//	for i := 0; i < len(tables); i++ {
//		_, err := result.Exec()
//		if err != nil {
//			return err
//		}
//	}
//
//	return result.Close()
//}
//
//func (cont *PGXContainer) AddTags(name string, parentId int) (int, error) {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	row := cont.conn.QueryRow(context.Background(), "INSERT INTO tags(name, parent_id) VALUES ($1, $2) RETURNING id", name, parentId)
//
//	var id int
//
//	err := row.Scan(&id)
//
//	return id, err
//}
//
//func (cont *PGXContainer) GetTag(id int) (*Tags, error) {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	row := cont.conn.QueryRow(context.Background(), "SELECT * FROM tags WHERE id=$1", id)
//	var tag Tags
//	err := row.Scan(&tag.Id, &tag.Name, &tag.ParentId)
//	return &tag, err
//}
//
//
//func (cont *PGXContainer) GetTags() ([]Tags, error) {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	var tags []Tags
//	err := pgxscan.Select(context.Background(), cont.conn, &tags, "SELECT * FROM tags")
//
//	return tags, err
//}
//
//func (cont *PGXContainer) DeleteTag(id int) error {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	_, err := cont.conn.Exec(context.Background(), "DELETE FROM tags WHERE id = $1", id)
//	return err
//}
//
//func (cont *PGXContainer) UpdateTag(id int, name string, parentId int) error {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	_, err := cont.conn.Exec(context.Background(), "UPDATE tags SET name=$1, parentId=$2 WHERE id=$3", name, parentId, id)
//	return err
//}
//
//func (cont *PGXContainer) AssignTag(tagId, messageId, chatId int) error {
//	cont.mu.Lock()
//	defer cont.mu.Unlock()
//	_, err := cont.conn.Exec(context.Background(), "INSERT INTO assign_tags(tag_id, message_id, chat_id) VALUES ($1, $2, $3)", tagId, messageId, chatId)
//	return err
//}
