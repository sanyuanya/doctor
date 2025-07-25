-- public.feedbacks definition

-- Drop table

-- DROP TABLE public.feedbacks;

CREATE TABLE public.feedbacks (
	feedback_id bigserial NOT NULL,
	"content" varchar DEFAULT ''::character varying NOT NULL,
	file varchar DEFAULT ''::character varying NOT NULL,
	status varchar DEFAULT ''::character varying NOT NULL,
	user_id int8 DEFAULT 0 NOT NULL,
	created_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamptz DEFAULT CURRENT_TIMESTAMP NOT NULL,
	CONSTRAINT feedbacks_pk PRIMARY KEY (feedback_id)
);