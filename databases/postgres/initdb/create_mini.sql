-- public.mini definition

-- Drop table

-- DROP TABLE public.mini;

CREATE TABLE public.mini (
	id bigserial NOT NULL,
	access_token varchar DEFAULT ''::character varying NOT NULL,
	expires_in int8 DEFAULT 0 NOT NULL,
	app_id varchar DEFAULT ''::character varying NOT NULL,
	secret varchar DEFAULT ''::character varying NOT NULL,
	CONSTRAINT mini_pk PRIMARY KEY (id)
);

INSERT INTO public.mini
(id, access_token, expires_in, app_id, secret)
VALUES(1, '94_1YZTSooOP_IeeTOEU5QX5WLd5k41_xNyQpHaENCDPaVmxlgLNASjPTQuLsqjsgsLgBDDP745JJdqdfMv3JoPnY7LjM1mxeWfzchPY0zGuzNIkqfIe3Ev8GdV98cQIRiAIABTC', 1753433676, 'wx4b045e0cd2c67eea', '86dee5bc4db693ba2ec4db976caeae94');