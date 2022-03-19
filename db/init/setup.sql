CREATE TABLE IF NOT public.users (
  id integer PRIMARY KEY,
  email text NOT NULL DEFAULT '' default '',
  password text NOT NULL DEFAULT '',
  type text NOT NULL DEFAULT '',
  session text NOT NULL DEFAULT '',
  name text NOT NULL DEFAULT '',
  age integer NOT NULL DEFAULT '',
  gender text NOT NULL DEFAULT '',
  avater text NOT NULL DEFAULT '',
  introduction text NOT NULL DEFAULT '',
  friends integer [] NOT NULL DEFAULT [],
  mute integer [] NOT NULL DEFAULT [],
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.posts (
  id integer PRIMARY KEY,
  uid integer NOT NULL DEFAULT '',
  title text NOT NULL DEFAULT '',
  body text NOT NULL DEFAULT '',
  img text NOT NULL DEFAULT '',
  liked integer [] text NOT NULL DEFAULT [],
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.comments (
  id integer PRIMARY KEY,
  post_id integer NOT NULL DEFAULT '',
  uid integer NOT NULL DEFAULT '',
  body text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.markers (
  id integer PRIMARY KEY,
  post_id integer NOT NULL DEFAULT '',
  title text NOT NULL DEFAULT '',
  lat text NOT NULL DEFAULT '',
  lng text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.requests (
  id integer PRIMARY KEY,
  request integer NOT NULL DEFAULT '',
  requested integer NOT NULL DEFAULT '',
  status text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);