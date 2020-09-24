package data

import (
	"errors"
	"github.com/golang/mock/gomock"
	"gopkg.in/mgo.v2"
	"reddit/tests/mock"
)

func InitQuery(ctrl *gomock.Controller) *mock_interfaces.MockMongoQuery {
	q := mock_interfaces.NewMockMongoQuery(ctrl)
	return q
}

func GetQuery(ctrl *gomock.Controller, status int, data interface{}) *mock_interfaces.MockMongoQuery {
	switch status {
	case 2:
		q := mock_interfaces.NewMockMongoQuery(ctrl)
		q.EXPECT().All(gomock.Any()).SetArg(0, data)
		return q
	case 3:
		q := mock_interfaces.NewMockMongoQuery(ctrl)
		q.EXPECT().All(gomock.Any()).Return(errors.New("DB err"))
		return q
	case 4:
		q := mock_interfaces.NewMockMongoQuery(ctrl)
		q.EXPECT().All(gomock.Any()).Return(mgo.ErrNotFound)
		return q
	case 5, 6:
		q := mock_interfaces.NewMockMongoQuery(ctrl)
		q.EXPECT().One(gomock.Any()).SetArg(0, data)
		return q
	case 7:
		q := mock_interfaces.NewMockMongoQuery(ctrl)
		q.EXPECT().One(gomock.Any()).Return(mgo.ErrNotFound)
		return q
	}
	return nil
}

func SetExpectAllWithArg(q *mock_interfaces.MockMongoQuery, data interface{}) *mock_interfaces.MockMongoQuery {
	q.EXPECT().All(gomock.Any()).SetArg(0, data)
	return q
}

func SetExpectAllWithDbErr(q *mock_interfaces.MockMongoQuery) *mock_interfaces.MockMongoQuery {
	q.EXPECT().All(gomock.Any()).Return(errors.New("DB err"))
	return q
}

func SetExpectAllWithNotFoundErr(q *mock_interfaces.MockMongoQuery) *mock_interfaces.MockMongoQuery {
	q.EXPECT().All(gomock.Any()).Return(mgo.ErrNotFound)
	return q
}

func SetExpectOneWithArg(q *mock_interfaces.MockMongoQuery, data interface{}) *mock_interfaces.MockMongoQuery {
	q.EXPECT().One(gomock.Any()).SetArg(0, data)
	return q
}

func SetExpectOneWithNotFoundErr(q *mock_interfaces.MockMongoQuery) *mock_interfaces.MockMongoQuery {
	q.EXPECT().One(gomock.Any()).Return(mgo.ErrNotFound)
	return q
}

func SetExpectOneWithError(q *mock_interfaces.MockMongoQuery) *mock_interfaces.MockMongoQuery {
	q.EXPECT().One(gomock.Any()).Return(errors.New("DB error"))
	return q
}
