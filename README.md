# DORM (Work in Progress)

Dorm is a developer friendly DynamoDB object modeling framework.

## Wishlist of features

- Model Management using Go tags
- Preload validation of models
- Scan/Query/Get support\*
  - Query builder
  - Expression like filter builder
- Put/Upsert/Update/Save support\*
- Delete support
- Reload model support\*\*
- Conditional Expressions support
- Mapping models to GSIs/LSIs
- Single model support for multiple GSIs/LSIs\*\*\*

\* With zero marshal/unsmarshal steps required
\*\* Reloading is similar to get but allows you to update a local record from a remote table
\*\*\* On the fence about this one... (might be an anti-pattern)

## Inspiration

Here are some projects that I used to draw inspiration from for building this package:

- [Gorm (Golang - SQL)](https://github.com/jinzhu/gorm)
- [TypeORM (Typescript - SQL)](https://github.com/typeorm/typeorm)
- [Sequelize (NodeJS - SQL)](https://github.com/sequelize/sequelize)
- [Mongoose (NodeJS - Mongo)](https://github.com/Automattic/mongoose)
