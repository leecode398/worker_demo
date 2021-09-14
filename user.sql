-- drop table if exist users;
CREATE TABLE users(
	user_id int(11) NOT NULL,
	user_name varchar(20) NOT NULL,
	age int(11) NOT NULL,
	address varchar(30) NOT NULL,
	tel int(15) NOT NULL,
	department varchar(30) NOT NULL,
	primary key(user_id)
);

-- drop table if exist employee;
CREATE TABLE employee(
	id int(11) NOT NULL,
	dept_name varchar(30) NOT NULL,
	position varchar(20) NOT NULL,
	induction_time date NOT NULL,
	foreign key(id) references users(user_id)
);
