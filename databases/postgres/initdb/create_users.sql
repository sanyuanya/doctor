-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	user_id bigserial NOT NULL,
	avatar varchar DEFAULT ''::character varying NOT NULL,
	nickname varchar DEFAULT ''::character varying NOT NULL,
	gender varchar DEFAULT ''::character varying NOT NULL,
	birth_date varchar DEFAULT ''::character varying NOT NULL,
	height_cm varchar DEFAULT ''::character varying NOT NULL,
	weight_kg varchar DEFAULT ''::character varying NOT NULL,
	phone_number varchar DEFAULT ''::character varying NOT NULL,
	open_id varchar DEFAULT ''::character varying NOT NULL,
	session_key varchar DEFAULT ''::character varying NOT NULL,
	union_id varchar DEFAULT ''::character varying NOT NULL,
	emergency_contact_name varchar DEFAULT ''::character varying NOT NULL,
	emergency_contact_relation varchar DEFAULT ''::character varying NOT NULL,
	emergency_contact_phone varchar DEFAULT ''::character varying NOT NULL,
	active_role varchar DEFAULT ''::character varying NOT NULL,
	default_role varchar DEFAULT ''::character varying NOT NULL,
	patient_notification jsonb DEFAULT '[]'::jsonb NOT NULL,
	consultant_notification jsonb DEFAULT '[]'::jsonb NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	deleted_at timestamptz NULL,
	group_type varchar DEFAULT ''::character varying NOT NULL,
	relation_id int8 DEFAULT 0 NOT NULL,
	invite_code varchar DEFAULT ''::character varying NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (user_id)
);