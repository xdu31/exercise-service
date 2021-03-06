-- This file will be used to initialize the test database.
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (1,'calves',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (2,'quad_riceps',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (3,'ham_strings',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (4,'gluteus',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (5,'hips_other',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (6,'lower_back',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (7,'lats',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (8,'trapezius',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (9,'abdominals',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (10,'pectorals',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (11,'deltoid',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (12,'triceps',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (13,'biceps',NOW(),NOW());
INSERT INTO muscle_groups (id,name,created_at,updated_at) VALUES (14,'forearms',NOW(),NOW());

INSERT INTO exercises (id,name,created_at,updated_at) VALUES (1,'squat',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (2,'leg_press',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (3,'lunge',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (4,'deadlift',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (5,'leg_extension',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (6,'leg_curl',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (7,'standing_calf_raise',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (8,'seated_calf_raise',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (9,'hip_adductor',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (10,'bench_press',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (11,'chest_fly',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (12,'push_up',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (13,'pull_down',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (14,'pull_up',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (15,'bent_over_row',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (16,'upright_row',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (17,'shoulder_press',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (18,'shoulder_fly',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (19,'lateral_raise',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (20,'shoulder_shrug',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (21,'pushdown',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (22,'triceps_extension',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (23,'biceps_curl',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (24,'crunch',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (25,'russian_twist',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (26,'leg_raise',NOW(),NOW());
INSERT INTO exercises (id,name,created_at,updated_at) VALUES (27,'back_extension',NOW(),NOW());

INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,2);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,6);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,7);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (1,8);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (2,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (2,2);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (2,3);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (2,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (2,5);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (3,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (3,2);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (3,3);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (3,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (3,6);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (4,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (4,2);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (4,3);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (4,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (4,27);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (5,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (5,3);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (5,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (5,9);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (5,26);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (6,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (6,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (6,27);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (7,13);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (7,14);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (7,15);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (8,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (8,16);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (8,20);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (9,1);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (9,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (9,24);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (9,25);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (9,26);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (10,10);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (10,11);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (10,12);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (11,11);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (11,16);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (11,17);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (11,18);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (11,19);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (12,10);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (12,12);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (12,21);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (12,22);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (13,23);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (14,4);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (14,20);
INSERT INTO training_data (muscle_group_id,exercise_id) VALUES (14,22);