package db

import (
	"reflect"
	"testing"
)

func TestNewRedisDB(t *testing.T) {
	type args struct {
		client RedisClientInterface
	}
	tests := []struct {
		name    string
		args    args
		want    DBInterface
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRedisDB(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRedisDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRedisDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_Set(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		value  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			if err := db.Set(tt.args.prefix, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_HashSet(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix   string
		key      string
		valueMap map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			if err := db.HashSet(tt.args.prefix, tt.args.key, tt.args.valueMap); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.HashSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_Get(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			got, err := db.Get(tt.args.prefix, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisDB.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_GetHashSet(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			got, err := db.GetHashSet(tt.args.prefix, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.GetHashSet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisDB.GetHashSet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_Exist(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			got, err := db.Exist(tt.args.prefix, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Exist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RedisDB.Exist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_Delete(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			if err := db.Delete(tt.args.prefix, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_FindHashSets(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			got, err := db.FindHashSets(tt.args.min, tt.args.max)
			if (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.FindHashSets() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RedisDB.FindHashSets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRedisDB_SetScore(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		score  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			if err := db.SetScore(tt.args.prefix, tt.args.key, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.SetScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_UpdateScore(t *testing.T) {
	type fields struct {
		client RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		score  int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			if err := db.UpdateScore(tt.args.prefix, tt.args.key, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.UpdateScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
