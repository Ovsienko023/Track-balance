package sqllite

func (c *Client) InitDb() error {
	if err := c.initUsers(); err != nil {
		return err
	}
	if err := c.initCircle(); err != nil {
		return err
	}
	if err := c.initAreas(); err != nil {
		return err
	}

	return nil
}

func (c *Client) initUsers() error {
	query := `create table if not exists "users" (
		id integer primary key autoincrement,
		login text not null unique,
		hash text not null,
		display_name text not null,
		avatar blob 
	)`

	if _, err := c.Driver.Exec(query); err != nil {
		return err
	}

	queryAdmin := `insert or ignore into users (login, hash, display_name) 
					values ('admin', 'admin', 'SuperDuper Admin')`
	_, err := c.Driver.Exec(queryAdmin)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) initCircle() error {
	queryCreateTable := `create table if not exists circles (
		id integer primary key autoincrement,
		user_id integer not null,
		created_at integer not null,
		description text,
        foreign key (user_id) references users (id) on delete cascade
	)`

	if _, err := c.Driver.Exec(queryCreateTable); err != nil {
		return err
	}

	return nil
}

func (c *Client) initAreas() error {
	queryCreateTable := `create table if not exists areas (
		id integer primary key autoincrement,
		user_id integer not null,
		circle_id integer not null,
		display_name text not null,
		description text,
		grade integer not null,
        foreign key (user_id) references users (id) on delete cascade,
        foreign key (circle_id) references circles (id) on delete cascade
	)`

	if _, err := c.Driver.Exec(queryCreateTable); err != nil {
		return err
	}

	return nil
}
