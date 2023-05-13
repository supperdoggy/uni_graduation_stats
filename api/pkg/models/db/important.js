/* global use, db */
// MongoDB Playground
// Use Ctrl+Space inside a snippet or a string literal to trigger completions.

// The current database to use.
use('stud');

// Search for documents in the current collection.
//   db.getCollection('users').aggregate([
//     {
//       $unwind: "$education"
//     },
//     {
//       $group: {
//         _id: "$education.schoolName"
//       }
//     },
//     {
//       $project: {
//         _id: 0,
//         university: "$_id"
//       }
//     }
//   ])

// db.getCollection('users').find({}, { "education.schoolName": 1, _id: 0 })

// db.getCollection('users').aggregate([
//     { $unwind: "$education" },
//     { $group: { _id: "$education.schoolName" } },
//     { $project: { _id: 0, schoolName: "$_id" } }
//   ])

db.getCollection('users').aggregate([
    { $unwind: "$education" },
    { $group: { _id: "$education.schoolName" } }
  ])

// db.getCollection('users').aggregate([
//     {
//       $unwind: "$education"
//     },
//     {
//       $group: {
//         _id: "$education.schoolName"
//       }
//     },
//     {
//       $sort: {
//         _id: 1
//       }
//     }
//   ])

// db.getCollection('users').aggregate([
//     {
//       $unwind: "$education"
//     },
//     {
//       $group: {
//         _id: "$education.schoolName"
//       }
//     },
//     {
//       $sort: {
//         _id: 1
//       }
//     },
//     {
//       $project: {
//         _id: 0,
//         schoolName: "$_id"
//       }
//     }
//   ])


// db.getCollection('users').aggregate([
//     {
//       $unwind: "$education"
//     },
//     {
//       $group: {
//         _id: "$education.schoolName"
//       }
//     },
//     {
//       $sort: {
//         _id: 1
//       }
//     },
//     {
//       $project: {
//         _id: 0,
//         schoolName: { $toLower: "$_id" }
//       }
//     }
//   ])
  