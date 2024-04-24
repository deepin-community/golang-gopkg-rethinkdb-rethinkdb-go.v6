package tests

import (
	"os"
	"time"

	test "gopkg.in/check.v1"
	r "gopkg.in/rethinkdb/rethinkdb-go.v6"
)

func (s *RethinkSuite) TestSessionConnect(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	c.Assert(err, test.IsNil)

	row, err := r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)

	var response string
	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")
}

func (s *RethinkSuite) TestSessionConnectHandshakeV1_0(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address:          url,
		HandshakeVersion: r.HandshakeV1_0,
	})
	c.Assert(err, test.IsNil)

	row, err := r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)

	var response string
	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")
}

func (s *RethinkSuite) TestSessionConnectHandshakeV0_4(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address:          url,
		HandshakeVersion: r.HandshakeV0_4,
	})
	c.Assert(err, test.IsNil)

	row, err := r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)

	var response string
	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")
}

func (s *RethinkSuite) TestSessionReconnect(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	c.Assert(err, test.IsNil)

	row, err := r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)

	var response string
	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World")

	err = session.Reconnect()
	c.Assert(err, test.IsNil)

	row, err = r.Expr("Hello World 2").Run(session)
	c.Assert(err, test.IsNil)

	err = row.One(&response)
	c.Assert(err, test.IsNil)
	c.Assert(response, test.Equals, "Hello World 2")
}

func (s *RethinkSuite) TestSessionConnectError(c *test.C) {
	var err error
	_, err = r.Connect(r.ConnectOpts{
		Address: "nonexistanturl",
		Timeout: time.Second,
	})

	c.Assert(err, test.NotNil)
	c.Assert(err, test.FitsTypeOf, r.RQLConnectionError{})
}

func (s *RethinkSuite) TestSessionClose(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	c.Assert(err, test.IsNil)

	_, err = r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)

	err = session.Close()
	c.Assert(err, test.IsNil)

	_, err = r.Expr("Hello World").Run(session)
	c.Assert(err, test.NotNil)
}

func (s *RethinkSuite) TestSessionServer(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	c.Assert(err, test.IsNil)

	server, err := session.Server()
	c.Assert(err, test.IsNil)

	c.Assert(len(server.ID) > 0, test.Equals, true)
	c.Assert(len(server.Name) > 0, test.Equals, true)
}

func (s *RethinkSuite) TestSessionConnectDatabase(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address:  url,
		AuthKey:  os.Getenv("RETHINKDB_AUTHKEY"),
		Database: "test2",
	})
	c.Assert(err, test.IsNil)
	c.Assert(session.Database(), test.Equals, "test2")

	_, err = r.Table("test2").Run(session)
	c.Assert(err, test.NotNil)
	c.Assert(err.Error(), test.Equals, "rethinkdb: Database `test2` does not exist. in:\nr.Table(\"test2\")")

	session.Use("test3")
	c.Assert(session.Database(), test.Equals, "test3")
}

func (s *RethinkSuite) TestSessionConnectUsername(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address: url,
	})
	c.Assert(err, test.IsNil)

	r.DB("rethinkdb").Table("users").Insert(map[string]string{
		"id":       "rethinkdb_test",
		"password": "password",
	}).Exec(session)

	session, err = r.Connect(r.ConnectOpts{
		Address:  url,
		Username: "rethinkdb_test",
		Password: "password",
	})
	c.Assert(err, test.IsNil)

	_, err = r.Expr("Hello World").Run(session)
	c.Assert(err, test.IsNil)
}

func (s *RethinkSuite) TestSessionIdleConnectionRemainsUsableSmallTimeout(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address:      url,
		NumRetries:   1,
		InitialCap:   1,
		ReadTimeout:  10 * time.Millisecond,
		WriteTimeout: 10 * time.Millisecond,
	})
	c.Assert(err, test.IsNil)

	time.Sleep(20 * time.Millisecond)

	var num int
	err = r.Expr(5).ReadOne(&num, session)
	c.Assert(err, test.IsNil)
	c.Assert(num, test.Equals, 5)

	time.Sleep(20 * time.Millisecond)

	err = r.Expr(6).ReadOne(&num, session)
	c.Assert(err, test.IsNil)
	c.Assert(num, test.Equals, 6)
}

func (s *RethinkSuite) TestSessionIdleConnectionRemainsUsableNoTimeout(c *test.C) {
	session, err := r.Connect(r.ConnectOpts{
		Address:    url,
		NumRetries: 1,
		InitialCap: 1,
	})
	c.Assert(err, test.IsNil)

	time.Sleep(10 * time.Millisecond)

	var num int
	err = r.Expr(5).ReadOne(&num, session)
	c.Assert(err, test.IsNil)
	c.Assert(num, test.Equals, 5)

	time.Sleep(10 * time.Millisecond)

	err = r.Expr(6).ReadOne(&num, session)
	c.Assert(err, test.IsNil)
	c.Assert(num, test.Equals, 6)
}
