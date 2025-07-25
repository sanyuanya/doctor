-- public.blood_glucose_record definition

-- Drop table

-- DROP TABLE public.blood_glucose_record;

CREATE TABLE public.blood_glucose_record (
	blood_glucose_record_id bigserial NOT NULL,
	user_id int8 DEFAULT 0 NOT NULL,
	upload_time int8 DEFAULT 0 NOT NULL,
	notes varchar DEFAULT ''::character varying NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT blood_glucose_record_pk PRIMARY KEY (blood_glucose_record_id)
);