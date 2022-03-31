package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/haguru/martian/mocks"
	"github.com/stretchr/testify/mock"
)

func TestNewRedisDB(t *testing.T) {
	tests := []struct {
		name                    string
		config                  Configuration
		closeReturn             error
		client                  *mocks.RedisClientInterface
		testConnectionReturnErr error
		want                    DBInterface
		wantErr                 bool
	}{
		{
			name:                    "Success client created",
			config:                  Configuration{},
			client:                  &mocks.RedisClientInterface{},
			closeReturn:             nil,
			testConnectionReturnErr: nil,
			want:                    nil,
			wantErr:                 false,
		},
		{
			name:                    "close client failed",
			config:                  Configuration{},
			client:                  &mocks.RedisClientInterface{},
			closeReturn:             fmt.Errorf("failed"),
			testConnectionReturnErr: nil,
			want:                    nil,
			wantErr:                 false,
		},
		{
			name:                    "connection client failed",
			config:                  Configuration{},
			client:                  &mocks.RedisClientInterface{},
			closeReturn:             nil,
			testConnectionReturnErr: fmt.Errorf("failed"),
			want:                    nil,
			wantErr:                 true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn := &mocks.Conn{}
			tt.client.On("TestConnection").Return(conn, tt.testConnectionReturnErr)
			conn.On("Close").Return(tt.closeReturn)
			if tt.testConnectionReturnErr == nil {
				tt.want = &RedisDB{
					client: tt.client,
				}
			}

			got, err := NewRedisDB(tt.client)
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
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		value  string
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doReturnErr error
		wantErr     bool
	}{
		{
			name: "Successfully set key",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "mykey",
				value:  "super sayian",
			},
			doReturnErr: nil,
			wantErr:     false,
		},
		{
			name: "no prefix",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key:    "mykey",
				value:  "super sayian",
			},
			doReturnErr: nil,
			wantErr:     true,
		},
		{
			name: "Do returns error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "mykey",
				value:  "super sayian",
			},
			doReturnErr: fmt.Errorf("failed"),
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "SET", tt.args.prefix+tt.args.key, tt.args.value).Return(nil, tt.doReturnErr)
			if err := db.Set(tt.args.prefix, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_HashSet(t *testing.T) {
	type fields struct {
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix   string
		key      string
		valueMap map[string]interface{}
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doReturnErr error
		wantErr     bool
	}{
		{
			name: "Successful HSET",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix:   "test:",
				key:      "myKey",
				valueMap: map[string]interface{}{"test1": "test2"},
			},
			doReturnErr: nil,
			wantErr:     false,
		},
		{
			name: "no prefix",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix:   "",
				key:      "myKey",
				valueMap: map[string]interface{}{"test1": "test2"},
			},
			doReturnErr: nil,
			wantErr:     true,
		},
		{
			name: "Do fails",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix:   "",
				key:      "myKey",
				valueMap: map[string]interface{}{"test1": "test2"},
			},
			doReturnErr: fmt.Errorf("failed"),
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Send", "MULTI").Return(nil)
			tt.conn.On("Send", "HSET", mock.Anything, mock.Anything, mock.Anything).Return(nil)
			tt.conn.On("Do", "EXEC").Return(nil, tt.doReturnErr)
			if err := db.HashSet(tt.args.prefix, tt.args.key, tt.args.valueMap); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.HashSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_Get(t *testing.T) {
	type fields struct {
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doReturnErr error
		want        interface{}
		wantErr     bool
	}{
		{
			name: "Succesful GET",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "myKey",
			},
			doReturnErr: nil,
			want:        "one above all",
			wantErr:     false,
		},
		{
			name: "no prefix",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key:    "myKey",
			},
			doReturnErr: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name: "Do returned error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "myKey",
			},
			doReturnErr: fmt.Errorf("fail"),
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "GET", tt.args.prefix+tt.args.key).Return(tt.want, tt.doReturnErr)
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
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doReturnErr error
		want        []interface{}
		wantErr     bool
	}{
		{
			name: "Success full GET HASH SET",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "myKey",
			},
			doReturnErr: nil,
			want:        []interface{}{"test1", int64(1), "test2", "12344"},
			wantErr:     false,
		},
		{
			name: "no prefix",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key:    "myKey",
			},
			doReturnErr: nil,
			want:        nil,
			wantErr:     true,
		},
		{
			name: "Do returns error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "myKey",
			},
			doReturnErr: fmt.Errorf("fail"),
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "HGETALL", tt.args.prefix+tt.args.key).Return(tt.want, tt.doReturnErr)
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
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doReturn    int64
		doreturnErr error
		want        bool
		wantErr     bool
	}{
		{
			name: "element exists",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "hadoken",
			},
			doReturn:    1,
			doreturnErr: nil,
			want:        true,
			wantErr:     false,
		},
		{
			name: "element does not exists",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "hadoken",
			},
			doReturn:    0,
			doreturnErr: nil,
			want:        false,
			wantErr:     false,
		},
		{
			name: "Conn Do returns error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key:    "hadoken",
			},
			doReturn:    0,
			doreturnErr: fmt.Errorf("failed"),
			want:        false,
			wantErr:     true,
		},
		{
			name: "no prefix given",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key:    "hadoken",
			},
			doReturn:    0,
			doreturnErr: nil,
			want:        false,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "EXISTS", tt.args.prefix+tt.args.key).Return(tt.doReturn, tt.doreturnErr)
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
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		doreturnErr error
		args        args
		wantErr     bool
	}{
		{
			name: "successfully delete",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			doreturnErr: nil,
			args: args{
				prefix: "test:",
				key:    "hadoken",
			},
			wantErr: false,
		},
		{
			name: "no prefix given",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			doreturnErr: nil,
			args: args{
				prefix: "",
				key:    "hadoken",
			},
			wantErr: true,
		},
		{
			name: "conn Do return failed",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			doreturnErr: fmt.Errorf("failed"),
			args: args{
				prefix: "test:",
				key:    "hadoken",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "DEL", tt.args.prefix+tt.args.key).Return(nil, tt.doreturnErr)
			if err := db.Delete(tt.args.prefix, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_FindHashSets(t *testing.T) {
	type fields struct {
		client *mocks.RedisClientInterface
	}
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name        string
		conn        *mocks.Conn
		fields      fields
		args        args
		doreturn    []interface{}
		doreturnErr error
		want        []string
		wantErr     bool
	}{
		{
			name: "Successfully found hash set keys",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				min: 0,
				max: 0,
			},
			doreturn:    []interface{}{"test1", "test2"},
			doreturnErr: nil,
			want:        []string{"test1", "test2"},
			wantErr:     false,
		},
		{
			name: "Successfully found no hash set keys",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				min: 0,
				max: 0,
			},
			doreturn:    []interface{}{},
			doreturnErr: nil,
			want:        []string{},
			wantErr:     false,
		},
		{
			name: "Con Do return error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				min: 0,
				max: 0,
			},
			doreturn:    []interface{}{},
			doreturnErr: fmt.Errorf("fail"),
			want:        nil,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "ZRANGEBYSCORE", "score", tt.args.min, tt.args.max).Return(tt.doreturn, tt.doreturnErr)
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
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		score  int
	}
	tests := []struct {
		name    string
		conn    *mocks.Conn
		fields  fields
		args    args
		doreturnErr error
		wantErr bool
	}{
		{
			name: "Successfull score set",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: nil,
			wantErr: false,
		},
		{
			name: "Conn Do returns error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: fmt.Errorf("failed"),
			wantErr: true,
		},
		{
			name: "no prefix given",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "ZADD", "score", tt.args.score, tt.args.prefix+tt.args.key).Return(nil, tt.doreturnErr)
			if err := db.SetScore(tt.args.prefix, tt.args.key, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.SetScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRedisDB_UpdateScore(t *testing.T) {
	type fields struct {
		client *mocks.RedisClientInterface
	}
	type args struct {
		prefix string
		key    string
		score  int
	}
	tests := []struct {
		name    string
		conn    *mocks.Conn
		fields  fields
		args    args
		doreturnErr error
		wantErr bool
	}{
		{
			name: "Successfull score set",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: nil,
			wantErr: false,
		},
		{
			name: "Conn Do returns error",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "test:",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: fmt.Errorf("failed"),
			wantErr: true,
		},
		{
			name: "no prefix given",
			conn: &mocks.Conn{},
			fields: fields{
				client: &mocks.RedisClientInterface{},
			},
			args: args{
				prefix: "",
				key: "hadoken",
				score: 0,
			},
			doreturnErr: nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &RedisDB{
				client: tt.fields.client,
			}
			tt.fields.client.On("GetConnection").Return(tt.conn)
			tt.conn.On("Close").Return(nil)
			tt.conn.On("Do", "ZADD", "score", tt.args.score, tt.args.prefix+tt.args.key).Return(nil, tt.doreturnErr)
			if err := db.UpdateScore(tt.args.prefix, tt.args.key, tt.args.score); (err != nil) != tt.wantErr {
				t.Errorf("RedisDB.UpdateScore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
