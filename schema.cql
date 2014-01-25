DROP KEYSPACE twitter_example;

CREATE KEYSPACE twitter_example WITH replication = 
{ 'class' : 'SimpleStrategy', 'replication_factor' : 1 };

USE twitter_example;

CREATE TABLE tweets
(
    id bigint,
    name text,
    PRIMARY KEY (id)
);