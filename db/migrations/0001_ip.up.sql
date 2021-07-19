CREATE TABLE muscle_groups (
  id integer NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL
);

CREATE SEQUENCE muscle_group_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER SEQUENCE muscle_group_id_seq OWNED BY muscle_groups.id;

ALTER TABLE ONLY muscle_groups ALTER COLUMN id SET DEFAULT nextval('muscle_group_id_seq'::regclass);

ALTER TABLE ONLY muscle_groups
  ADD CONSTRAINT pk_muscle_groups PRIMARY KEY (id);

ALTER TABLE ONLY muscle_groups
  ADD CONSTRAINT u_muscle_group_name UNIQUE (name);

CREATE TABLE exercises (
  id integer NOT NULL,
  name varchar(255) NOT NULL,
  created_at timestamptz DEFAULT current_timestamp,
  updated_at timestamptz DEFAULT NULL
);

CREATE SEQUENCE exercise_id_seq
  START WITH 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER SEQUENCE exercise_id_seq OWNED BY exercises.id;

ALTER TABLE ONLY exercises ALTER COLUMN id SET DEFAULT nextval('exercise_id_seq'::regclass);

ALTER TABLE ONLY exercises
  ADD CONSTRAINT pk_exercises PRIMARY KEY (id);

ALTER TABLE ONLY exercises
  ADD CONSTRAINT u_exercise_name UNIQUE (name);

CREATE TABLE training_data (
  muscle_group_id INT REFERENCES muscle_groups(id) NOT NULL,
  exercise_id INT REFERENCES exercises(id) NOT NULL,
  CONSTRAINT pk_training_data
    PRIMARY KEY (muscle_group_id, exercise_id)
);
