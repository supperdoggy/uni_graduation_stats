package storage

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func (db *mongodb) ListSchools(ctx context.Context) ([]rest.ListSchools, error) {
	pipeline := []bson.M{
		{"$unwind": "$education"},
		{"$group": bson.M{
			"_id":   "$education.schoolName",
			"count": bson.M{"$sum": 1},
		}},
		{"$sort": bson.M{"count": -1}},
		{"$project": bson.M{
			"_id":        0,
			"schoolName": "$_id",
			"count":      1,
		}},
	}

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating schools", zap.Error(err))
		return nil, err
	}

	var schools []rest.ListSchools
	if err := cur.All(ctx, &schools); err != nil {
		db.log.Error("error getting schools", zap.Error(err))
		return nil, err
	}

	return schools, nil
}

func (db mongodb) ListSchoolsTopCompanies(ctx context.Context, school string) ([]rest.ListSchoolsTopCompanies, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"education.schoolName": school,
			},
		},
		{
			"$unwind": "$experiences",
		},
		{
			"$group": bson.M{
				"_id": bson.M{
					"$trim": bson.M{
						"input": bson.M{
							"$arrayElemAt": bson.A{
								bson.M{
									"$split": bson.A{
										"$experiences.company",
										"·",
									},
								},
								0,
							},
						},
					},
				},
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":   1,
				"count": 1,
			},
		},
		{
			"$sort": bson.M{
				"count": -1,
			},
		},
	}

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating schools", zap.Error(err))
		return nil, err
	}

	var schools []rest.ListSchoolsTopCompanies
	if err := cur.All(ctx, &schools); err != nil {
		db.log.Error("error getting schools", zap.Error(err))
		return nil, err
	}

	return schools, nil
}
