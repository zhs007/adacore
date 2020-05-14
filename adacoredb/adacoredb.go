package adacoredb

import (
	"context"

	"github.com/golang/protobuf/proto"
	adacorebase "github.com/zhs007/adacore/base"
	adacorepb "github.com/zhs007/adacore/adacorepb"
	"github.com/zhs007/ankadb"
	"go.uber.org/zap"
)

// AdaCoreDB -
type AdaCoreDB struct {
	ankaDB ankadb.AnkaDB
}

// NewAdaCoreDB - new AdaCoreDB
func NewAdaCoreDB(dbpath string, httpAddr string, engine string) (*AdaCoreDB, error) {
	cfg := ankadb.NewConfig()

	cfg.AddrHTTP = httpAddr
	cfg.PathDBRoot = dbpath
	cfg.ListDB = append(cfg.ListDB, ankadb.DBConfig{
		Name:   AdaCoreDBName,
		Engine: engine,
		PathDB: AdaCoreDBName,
	})

	ankaDB, err := ankadb.NewAnkaDB(cfg, nil)
	if ankaDB == nil {
		adacorebase.Error("NewAdaCoreDB", zap.Error(err))

		return nil, err
	}

	adacorebase.Info("NewAdaCoreDB", zap.String("dbpath", dbpath),
		zap.String("httpAddr", httpAddr), zap.String("engine", engine))

	db := &AdaCoreDB{
		ankaDB: ankaDB,
	}

	return db, err
}

// AddResource - add a resource
func (db *AdaCoreDB) AddResource(ctx context.Context, hashname string, lst []*adacorepb.ResourceInfo) (int, error) {
	if hashname == "" {
		return 0, adacorebase.ErrAdaCoreDBInvalidHashName
	}

	if len(lst) == 0 || lst[0].HashName != hashname {
		return 0, adacorebase.ErrAdaCoreDBInvalidResList
	}

	nums := 0
	var lastsaveerr error

	for _, cr := range lst {
		curri, _ := db.GetResource(ctx, cr.HashName)
		if curri != nil {
			curri.CitedTimes++

			err := db.SetResource(ctx, curri)
			if err != nil {
				lastsaveerr = err
			}

			nums++
		} else {
			cr.CitedTimes = 1
			cr.CreateTime = adacorebase.GetCurTime()

			if cr.HashName != hashname {
				cr.Type = adacorepb.ResourceType_RT_OTHER
			}

			err := db.SetResource(ctx, cr)
			if err != nil {
				lastsaveerr = err
			}

			nums++
		}
	}

	return nums, lastsaveerr
}

// GetResource - get a resource
func (db *AdaCoreDB) GetResource(ctx context.Context, hashname string) (*adacorepb.ResourceInfo, error) {
	buf, err := db.ankaDB.Get(ctx, AdaCoreDBName, makeKey(hashname))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil, nil
		}

		return nil, err
	}

	ri := &adacorepb.ResourceInfo{}

	err = proto.Unmarshal(buf, ri)
	if err != nil {
		return nil, err
	}

	return ri, nil
}

// DelResource - delete a resource
func (db *AdaCoreDB) DelResource(ctx context.Context, hashname string) error {
	err := db.ankaDB.Delete(ctx, AdaCoreDBName, makeKey(hashname))
	if err != nil {
		if err == ankadb.ErrNotFoundKey {
			return nil
		}

		return err
	}

	return nil
}

// SetResource - set a resource
func (db *AdaCoreDB) SetResource(ctx context.Context, ri *adacorepb.ResourceInfo) error {
	buf, err := proto.Marshal(ri)
	if err != nil {
		return err
	}

	err = db.ankaDB.Set(ctx, AdaCoreDBName, makeKey(ri.HashName), buf)
	if err != nil {
		return err
	}

	return nil
}

// ResetAllCitedTimes - reset all cited times
func (db *AdaCoreDB) ResetAllCitedTimes(ctx context.Context) error {
	err := db.ankaDB.ForEachWithPrefix(ctx, AdaCoreDBName, AdaCoreDBKeyPrefix, func(key string, value []byte) error {
		ri := &adacorepb.ResourceInfo{}

		err := proto.Unmarshal(value, ri)
		if err != nil {
			return err
		}

		if ri.Type == adacorepb.ResourceType_RT_OTHER {
			ri.CitedTimes = int32(db.RecountAllCitedTimes(ctx, ri.HashName))

			if ri.CitedTimes == 0 {
				err = db.ankaDB.Delete(ctx, AdaCoreDBName, key)
				if err != nil {
					return err
				}

				return nil
			}

			db.SetResource(ctx, ri)
		}

		return nil
	})

	return err
}

// RecountAllCitedTimes - recount all cited times
func (db *AdaCoreDB) RecountAllCitedTimes(ctx context.Context, hashname string) int {
	nums := 0

	db.ankaDB.ForEachWithPrefix(ctx, AdaCoreDBName, AdaCoreDBKeyPrefix, func(key string, value []byte) error {
		ri := &adacorepb.ResourceInfo{}

		err := proto.Unmarshal(value, ri)
		if err != nil {
			return err
		}

		if ri.Type == adacorepb.ResourceType_RT_PAGE {
			for _, ch := range ri.Children {
				if ch == hashname {
					nums++
				}
			}
		}

		return nil
	})

	return nums
}
