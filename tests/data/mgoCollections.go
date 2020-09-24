package data

import (
	"errors"
	"github.com/golang/mock/gomock"
	"gopkg.in/mgo.v2"
	mock_interfaces "reddit/tests/mock"
)

func InitCollection(ctrl *gomock.Controller) *mock_interfaces.MockMongoCollection {
	col := mock_interfaces.NewMockMongoCollection(ctrl)
	return col
}

func GetCollection(ctrl *gomock.Controller, status int, data interface{}) *mock_interfaces.MockMongoCollection {
	switch status {
	case 0:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Insert(gomock.Any()).Return(nil)
		return col
	case 1:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Insert(gomock.Any()).Return(errors.New("DB err"))
		return col
	case 2, 3, 4, 5, 7:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, status, data))
		return col
	case 6:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, status, data))
		col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
		return col
	case 8:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, 5, data))
		col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mgo.ErrNotFound)
		return col
	case 9:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Insert(gomock.Any()).Return(nil)
		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, 2, data))
		return col
	case 10:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Remove(gomock.Any()).Return(nil)
		col.EXPECT().Find(gomock.Any()).Return(GetQuery(ctrl, 4, data))
		return col
	case 11:
		col := mock_interfaces.NewMockMongoCollection(ctrl)
		col.EXPECT().Remove(gomock.Any()).Return(mgo.ErrNotFound)
		return col
	}
	return nil
}

func SetExpectInsert(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Insert(gomock.Any()).Return(nil)
	return col
}

func SetExpectInsertWithError(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Insert(gomock.Any()).Return(errors.New("DB err"))
	return col
}

func SetExpectFindReturnQuery(col *mock_interfaces.MockMongoCollection, q *mock_interfaces.MockMongoQuery) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Find(gomock.Any()).Return(q)
	return col
}

func SetExpectUpdate(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)
	return col
}

func SetExpectUpdateWithError(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Update(gomock.Any(), gomock.Any()).Return(mgo.ErrNotFound)
	return col
}

func SetExpectRemove(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Remove(gomock.Any()).Return(nil)
	return col
}

func SetExpectRemoveWithError(col *mock_interfaces.MockMongoCollection) *mock_interfaces.MockMongoCollection {
	col.EXPECT().Remove(gomock.Any()).Return(mgo.ErrNotFound)
	return col
}
