package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mc, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://59.110.216.174:27017/?readPreference=primary&ssl=false"))
	if err != nil {
		t.Fatalf("cannot connect mongodb:%v", err)
	}

	m := NewMongo(mc.Database("coolcar"))
	id, err := m.ResoleveAccountID(c, "123")
	if err != nil {
		t.Errorf("faild resoleve account id for 123:%v", err)
	} else {
		want := "60d044a2db7d7a726d4060ae"
		if id != want {
			t.Errorf("resoleve account id: want: %q, got:%q", want, id)
		}
	}

}
