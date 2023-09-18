#!/bin/sh
dockerize -wait tcp://mysql:3306 -timeout 20s mysqldef -h $MYSQL_HOST -u $MYSQL_USER -p $MYSQL_PASSWORD --file=./schema.sql $MYSQL_DATABASE
