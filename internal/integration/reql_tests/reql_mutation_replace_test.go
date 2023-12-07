// Code generated by gen_tests.py and process_polyglot.py.
// Do not edit this file directly.
// The template for this file is located at:
// ../template.go.tpl
package reql_tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
	"gopkg.in/rethinkdb/rethinkdb-go.v6/internal/compare"
)

// Tests replacement of selections
func TestMutationReplaceSuite(t *testing.T) {
	suite.Run(t, new(MutationReplaceSuite))
}

type MutationReplaceSuite struct {
	suite.Suite

	session *r.Session
}

func (suite *MutationReplaceSuite) SetupTest() {
	suite.T().Log("Setting up MutationReplaceSuite")
	// Use imports to prevent errors
	_ = time.Time{}
	_ = compare.AnythingIsFine

	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	suite.Require().NoError(err, "Error returned when connecting to server")
	suite.session = session

	r.DBDrop("db_mut_rep").Exec(suite.session)
	err = r.DBCreate("db_mut_rep").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_rep").Wait().Exec(suite.session)
	suite.Require().NoError(err)

	r.DB("db_mut_rep").TableDrop("table_test_mutation_replace").Exec(suite.session)
	err = r.DB("db_mut_rep").TableCreate("table_test_mutation_replace").Exec(suite.session)
	suite.Require().NoError(err)
	err = r.DB("db_mut_rep").Table("table_test_mutation_replace").Wait().Exec(suite.session)
	suite.Require().NoError(err)
}

func (suite *MutationReplaceSuite) TearDownSuite() {
	suite.T().Log("Tearing down MutationReplaceSuite")

	if suite.session != nil {
		r.DB("rethinkdb").Table("_debug_scratch").Delete().Exec(suite.session)
		r.DB("db_mut_rep").TableDrop("table_test_mutation_replace").Exec(suite.session)
		r.DBDrop("db_mut_rep").Exec(suite.session)

		suite.session.Close()
	}
}

func (suite *MutationReplaceSuite) TestCases() {
	suite.T().Log("Running MutationReplaceSuite: Tests replacement of selections")

	table_test_mutation_replace := r.DB("db_mut_rep").Table("table_test_mutation_replace")
	_ = table_test_mutation_replace // Prevent any noused variable errors

	{
		// mutation/replace.yaml line #7
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':100}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 100}
		/* table_test_mutation_replace.insert([{'id':i} for i in xrange(100)]) */

		suite.T().Log("About to run line #7: table_test_mutation_replace.Insert((func() []interface{} {\n    res := []interface{}{}\n    for iterator_ := 0; iterator_ < 100; iterator_++ {\n        i := iterator_\n        res = append(res, map[interface{}]interface{}{'id': i, })\n    }\n    return res\n}()))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Insert((func() []interface{} {
			res := []interface{}{}
			for iterator_ := 0; iterator_ < 100; iterator_++ {
				i := iterator_
				res = append(res, map[interface{}]interface{}{"id": i})
			}
			return res
		}())), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #7")
	}

	{
		// mutation/replace.yaml line #19
		/* 100 */
		var expected_ int = 100
		/* table_test_mutation_replace.count() */

		suite.T().Log("About to run line #19: table_test_mutation_replace.Count()")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Count(), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #19")
	}

	{
		// mutation/replace.yaml line #24
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(12).replace(lambda row:{'id':row['id']}) */

		suite.T().Log("About to run line #24: table_test_mutation_replace.Get(12).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{'id': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(12).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{"id": row.AtIndex("id")} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #24")
	}

	{
		// mutation/replace.yaml line #31
		/* ({'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(12).replace(lambda row:{'id':row['id'], 'a':row['id']}) */

		suite.T().Log("About to run line #31: table_test_mutation_replace.Get(12).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{'id': row.AtIndex('id'), 'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(12).Replace(func(row r.Term) interface{} {
			return map[interface{}]interface{}{"id": row.AtIndex("id"), "a": row.AtIndex("id")}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #31")
	}

	{
		// mutation/replace.yaml line #36
		/* ({'deleted':1,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 1, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(13).replace(lambda row:null) */

		suite.T().Log("About to run line #36: table_test_mutation_replace.Get(13).Replace(func(row r.Term) interface{} { return nil})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(13).Replace(func(row r.Term) interface{} { return nil }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #36")
	}

	{
		// mutation/replace.yaml line #43
		/* ({'first_error':'Inserted object must have primary key `id`:\n{\n\t\"a\":\t1\n}','deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':10,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"first_error": "Inserted object must have primary key `id`:\n{\n\t\"a\":\t1\n}", "deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 10, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.between(10, 20, right_bound='closed').replace(lambda row:{'a':1}) */

		suite.T().Log("About to run line #43: table_test_mutation_replace.Between(10, 20).OptArgs(r.BetweenOpts{RightBound: 'closed', }).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{'a': 1, }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Between(10, 20).OptArgs(r.BetweenOpts{RightBound: "closed"}).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{"a": 1} }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #43")
	}

	{
		// mutation/replace.yaml line #47
		/* ({'deleted':0.0,'replaced':8,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 8, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.filter(lambda row:(row['id'] >= 10) & (row['id'] < 20)).replace(lambda row:{'id':row['id'], 'a':row['id']}) */

		suite.T().Log("About to run line #47: table_test_mutation_replace.Filter(func(row r.Term) interface{} { return row.AtIndex('id').Ge(10).And(row.AtIndex('id').Lt(20))}).Replace(func(row r.Term) interface{} { return map[interface{}]interface{}{'id': row.AtIndex('id'), 'a': row.AtIndex('id'), }})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Filter(func(row r.Term) interface{} { return row.AtIndex("id").Ge(10).And(row.AtIndex("id").Lt(20)) }).Replace(func(row r.Term) interface{} {
			return map[interface{}]interface{}{"id": row.AtIndex("id"), "a": row.AtIndex("id")}
		}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #47")
	}

	{
		// mutation/replace.yaml line #56
		/* ({'first_error':"Primary key `id` cannot be changed (`{\n\t\"id\":\t1\n}` -> `{\n\t\"a\":\t1,\n\t\"id\":\t2\n}`).",'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':1,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"first_error": "Primary key `id` cannot be changed (`{\n\t\"id\":\t1\n}` -> `{\n\t\"a\":\t1,\n\t\"id\":\t2\n}`).", "deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 1, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(1).replace({'id':2,'a':1}) */

		suite.T().Log("About to run line #56: table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{'id': 2, 'a': 1, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{"id": 2, "a": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #56")
	}

	{
		// mutation/replace.yaml line #61
		/* ({'first_error':"Inserted object must have primary key `id`:\n{\n\t\"a\":\t1\n}",'deleted':0.0,'replaced':0.0,'unchanged':0.0,'errors':1,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"first_error": "Inserted object must have primary key `id`:\n{\n\t\"a\":\t1\n}", "deleted": 0.0, "replaced": 0.0, "unchanged": 0.0, "errors": 1, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(1).replace({'a':1}) */

		suite.T().Log("About to run line #61: table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{'a': 1, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{"a": 1}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #61")
	}

	{
		// mutation/replace.yaml line #65
		/* ({'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(1).replace({'id':r.row['id'],'a':'b'}) */

		suite.T().Log("About to run line #65: table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{'id': r.Row.AtIndex('id'), 'a': 'b', })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(map[interface{}]interface{}{"id": r.Row.AtIndex("id"), "a": "b"}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #65")
	}

	{
		// mutation/replace.yaml line #70
		/* ({'deleted':0.0,'replaced':0.0,'unchanged':1,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 0.0, "unchanged": 1, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(1).replace(r.row.merge({'a':'b'})) */

		suite.T().Log("About to run line #70: table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{'a': 'b', }))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{"a": "b"})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #70")
	}

	{
		// mutation/replace.yaml line #75
		/* err('ReqlQueryLogicError', 'Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?")
		/* table_test_mutation_replace.get(1).replace(r.row.merge({'c':r.js('5')})) */

		suite.T().Log("About to run line #75: table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{'c': r.JS('5'), }))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{"c": r.JS("5")})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #75")
	}

	{
		// mutation/replace.yaml line #79
		/* err('ReqlQueryLogicError', 'Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?', [0]) */
		var expected_ Err = err("ReqlQueryLogicError", "Could not prove argument deterministic.  Maybe you want to use the non_atomic flag?")
		/* table_test_mutation_replace.get(1).replace(r.row.merge({'c':table_test_mutation_replace.nth(0)})) */

		suite.T().Log("About to run line #79: table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{'c': table_test_mutation_replace.Nth(0), }))")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{"c": table_test_mutation_replace.Nth(0)})), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #79")
	}

	{
		// mutation/replace.yaml line #83
		/* ({'deleted':0.0,'replaced':1,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 0.0, "replaced": 1, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.get(1).replace(r.row.merge({'c':r.js('5')}), non_atomic=True) */

		suite.T().Log("About to run line #83: table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{'c': r.JS('5'), })).OptArgs(r.ReplaceOpts{NonAtomic: true, })")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get(1).Replace(r.Row.Merge(map[interface{}]interface{}{"c": r.JS("5")})).OptArgs(r.ReplaceOpts{NonAtomic: true}), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #83")
	}

	{
		// mutation/replace.yaml line #99
		/* ({'deleted':99,'replaced':0.0,'unchanged':0.0,'errors':0.0,'skipped':0.0,'inserted':0.0}) */
		var expected_ map[interface{}]interface{} = map[interface{}]interface{}{"deleted": 99, "replaced": 0.0, "unchanged": 0.0, "errors": 0.0, "skipped": 0.0, "inserted": 0.0}
		/* table_test_mutation_replace.replace(lambda row:null) */

		suite.T().Log("About to run line #99: table_test_mutation_replace.Replace(func(row r.Term) interface{} { return nil})")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Replace(func(row r.Term) interface{} { return nil }), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #99")
	}

	{
		// mutation/replace.yaml line #104
		/* 1 */
		var expected_ int = 1
		/* table_test_mutation_replace.get('sdfjk').replace({'id':'sdfjk'})['inserted'] */

		suite.T().Log("About to run line #104: table_test_mutation_replace.Get('sdfjk').Replace(map[interface{}]interface{}{'id': 'sdfjk', }).AtIndex('inserted')")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get("sdfjk").Replace(map[interface{}]interface{}{"id": "sdfjk"}).AtIndex("inserted"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #104")
	}

	{
		// mutation/replace.yaml line #107
		/* 1 */
		var expected_ int = 1
		/* table_test_mutation_replace.get('sdfjki').replace({'id':'sdfjk'})['errors'] */

		suite.T().Log("About to run line #107: table_test_mutation_replace.Get('sdfjki').Replace(map[interface{}]interface{}{'id': 'sdfjk', }).AtIndex('errors')")

		runAndAssert(suite.Suite, expected_, table_test_mutation_replace.Get("sdfjki").Replace(map[interface{}]interface{}{"id": "sdfjk"}).AtIndex("errors"), suite.session, r.RunOpts{
			GeometryFormat: "raw",
			GroupFormat:    "map",
		})
		suite.T().Log("Finished running line #107")
	}
}
