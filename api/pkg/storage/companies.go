package storage

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (db *mongodb) ListCompanies(ctx context.Context) ([]rest.ListCompanies, error) {

	pipeline := []bson.M{
		bson.M{"$unwind": "$experiences"},
		bson.M{
			"$project": bson.M{
				"company": bson.M{
					"$arrayElemAt": []interface{}{
						bson.M{"$split": []interface{}{"$experiences.company", " · Full-time"}},
						0,
					},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   "$company",
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         0,
				"companyName": "$_id",
				"count":       1,
			},
		},
	}

	cur, err := db.users.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating companies", zap.Error(err))
		return nil, err
	}

	var companies []rest.ListCompanies
	if err := cur.All(ctx, &companies); err != nil {
		db.log.Error("error getting companies", zap.Error(err))
		return nil, err
	}

	return companies, nil
}
