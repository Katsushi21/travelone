CREATE TABLE IF NOT public.users (
  id serial PRIMARY KEY,
  email text NOT NULL DEFAULT '' default '',
  password text NOT NULL DEFAULT '',
  type text NOT NULL DEFAULT '',
  session text NOT NULL DEFAULT '',
  name text NOT NULL DEFAULT '',
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
  id serial PRIMARY KEY,
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
  id serial PRIMARY KEY,
  post_id integer NOT NULL DEFAULT '',
  uid integer NOT NULL DEFAULT '',
  body text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.markers (
  id serial PRIMARY KEY,
  post_id integer NOT NULL DEFAULT '',
  title text NOT NULL DEFAULT '',
  lat text NOT NULL DEFAULT '',
  lng text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);

CREATE TABLE IF NOT public.friend_requests (
  id serial PRIMARY KEY,
  request_uid integer NOT NULL DEFAULT '',
  requested_uid integer NOT NULL DEFAULT '',
  status text NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now(),
  deleted_at timestamp
);