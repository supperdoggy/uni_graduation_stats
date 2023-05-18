package storage

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (db *mongodb) ListSchoolsTopCompanies(ctx context.Context, school string) ([]rest.ListSchoolsTopCompanies, error) {
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

func (db *mongodb) ListJobsBySchool(ctx context.Context, school string) ([]rest.ListJobsBySchool, error) {
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
				"_id":   "$experiences.title",
				"count": bson.M{"$sum": 1},
			},
		},
		{
			"$project": bson.M{
				"_id":   0,
				"title": "$_id",
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

	var jobs []rest.ListJobsBySchool
	if err := cur.All(ctx, &jobs); err != nil {
		db.log.Error("error getting schools", zap.Error(err))
		return nil, err
	}

	return jobs, nil
}

func (db *mongodb) CorrelationBetweenDegreeAndTitle(ctx context.Context, school string) ([]rest.CorrelationDegreeAndTitle, error) {
	pipeline := bson.A{
		bson.D{
			{"$match", bson.D{
				{"education.schoolName", bson.D{
					{"$regex", primitive.Regex{Pattern: school, Options: "i"}},
				}},
			}},
		},
		bson.D{
			{"$unwind", "$education"},
		},
		bson.D{
			{"$match", bson.D{
				{"education.schoolName", bson.D{
					{"$regex", primitive.Regex{Pattern: school, Options: "i"}},
				}},
			}},
		},
		bson.D{
			{"$project", bson.D{
				{"_id", 0},
				{"degreeName", "$education.degreeName"},
				{"experience", bson.D{
					{"$map", bson.D{
						{"input", "$experiences"},
						{"as", "exp"},
						{"in", bson.D{
							{"company", bson.D{
								{"$arrayElemAt", bson.A{
									bson.D{
										{"$split", bson.A{
											"$$exp.company", " · ",
										}},
									},
									0,
								}},
							}},
							{"title", "$$exp.title"},
							{"startDate", "$$exp.startDate"},
							{"endDate", "$$exp.endDate"},
						}},
					}},
				}},
			}},
		},
	}

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating schools", zap.Error(err))
		return nil, err
	}

	var degreeTitles []rest.CorrelationDegreeAndTitle
	if err := cur.All(ctx, &degreeTitles); err != nil {
		db.log.Error("error getting schools", zap.Error(err))
		return nil, err
	}

	return degreeTitles, nil
}

func (db *mongodb) SchoolDegrees(ctx context.Context, school string) ([]rest.SchoolDegrees, error) {
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"education.schoolName": primitive.Regex{Pattern: school, Options: "i"},
			},
		},
		bson.M{"$unwind": "$education"},
		bson.M{
			"$match": bson.M{
				"education.schoolName": primitive.Regex{Pattern: school, Options: "i"},
			},
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
			"$project": bson.M{
				"_id":        0,
				"schoolName": "$_id.schoolName",
				"degreeName": "$_id.degreeName",
				"count":      1,
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
	}

	cur, err := db.students.Aggregate(ctx, pipeline, options.Aggregate().SetAllowDiskUse(true))
	if err != nil {
		db.log.Error("error aggregating degrees", zap.Error(err))
		return nil, err
	}

	var degrees []rest.SchoolDegrees
	if err := cur.All(ctx, &degrees); err != nil {
		db.log.Error("error getting degrees", zap.Error(err))
		return nil, err
	}

	for k, v := range degrees {
		if v.Degree == "" {
			degrees[k].Degree = "no data"
		}
	}

	return degrees, nil
}
