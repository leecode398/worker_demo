-- drop table if exist users;
CREATE TABLE users(
	user_id int(11) NOT NULL COMMENT '工号',
	user_name varchar(20) NOT NULL COMMENT '姓名',
	age int(11) NOT NULL COMMENT '年龄',
	address varchar(30) NOT NULL COMMENT '住址',
	tel int(15) NOT NULL COMMENT '电话',
	department varchar(30) NOT NULL COMMENT '部门',
	primary key(user_id)
);

-- drop table if exist employee;
CREATE TABLE employee(
	id int(11) NOT NULL COMMENT '工号',
	dept_name varchar(30) NOT NULL COMMENT '部门',
	position varchar(20) NOT NULL COMMENT '职位',
	induction_time date NOT NULL COMMENT '入职时间',
	foreign key(id) references users(user_id)
);
