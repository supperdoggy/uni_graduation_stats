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

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
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

func (db *mongodb) ListCompaniesTopSchools(ctx context.Context, company string) ([]rest.ListCompaniesTopSchools, error) {

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"experiences.company": bson.M{
					"$regex": company,
				},
			},
		},
		bson.M{
			"$unwind": "$education",
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"schoolName": "$education.schoolName",
					"degreeName": "$education.degreeName",
				},
				"count": bson.M{"$sum": 1},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$_id.schoolName",
				"degrees": bson.M{
					"$push": bson.M{
						"degreeName": "$_id.degreeName",
						"count":      "$count",
					},
				},
				"totalCount": bson.M{"$sum": "$count"},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 1,
				"degrees": bson.M{
					"$map": bson.M{
						"input": "$degrees",
						"as":    "degree",
						"in": bson.M{
							"degreeName": bson.M{
								"$cond": bson.M{
									"if":   bson.M{"$eq": bson.A{"$$degree.degreeName", ""}},
									"then": "no data",
									"else": "$$degree.degreeName",
								},
							},
							"count": "$$degree.count",
						},
					},
				},
				"totalCount": 1,
			},
		},
		bson.M{
			"$sort": bson.M{
				"totalCount": -1,
			},
		},
	}

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating companies", zap.Error(err))
		return nil, err
	}

	var schools []rest.ListCompaniesTopSchools
	if err := cur.All(ctx, &schools); err != nil {
		db.log.Error("error getting companies", zap.Error(err))
		return nil, err
	}

	for i, v := range schools {
		for _, degree := range v.Degrees {
			schools[i].Count += degree.Count
		}
	}

	return schools, nil
}
