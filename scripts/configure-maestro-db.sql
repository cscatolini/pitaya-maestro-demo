-- maestro
-- https://github.com/topfreegames/maestro
--
-- Licensed under the MIT license:
-- http://www.opensource.org/licenses/mit-license
-- Copyright © 2017 Top Free Games <backend@tfgco.com>

CREATE ROLE maestro LOGIN
  SUPERUSER INHERIT CREATEDB CREATEROLE;

CREATE DATABASE maestro
  WITH OWNER = maestro
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       TEMPLATE = template0;

GRANT ALL ON SCHEMA public TO maestro;
