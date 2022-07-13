DROP DATABASE name_list;

CREATE DATABASE name_list;

USE name_list;

CREATE table `name_list`(
`sid`   int(11) NOT NULL AUTO_INCREMENT,
`snum` varchar(50) NOT NULL,
`sname` varchar(50) DEFAULT '',
`sage` int(11) DEFAULT '0',
`sex`int(11) DEFAULT '0',
PRIMARY KEY(`sid`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;