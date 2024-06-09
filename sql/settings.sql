\c card_db;

DELETE FROM settings;

-- an explanation for each of the bolt card server settings can be found here
-- https://github.com/boltcard/boltcard/blob/main/docs/SETTINGS.md

INSERT INTO settings (name, value) VALUES ('LOG_LEVEL', 'DEBUG');
INSERT INTO settings (name, value) VALUES ('AES_DECRYPT_KEY', 'A7776C2613A6481F486B8A4D84FE62BD');
INSERT INTO settings (name, value) VALUES ('HOST_DOMAIN', 'localhost:9000');
INSERT INTO settings (name, value) VALUES ('MIN_WITHDRAW_SATS', '10');
INSERT INTO settings (name, value) VALUES ('MAX_WITHDRAW_SATS', '10000');
INSERT INTO settings (name, value) VALUES ('LN_HOST', 'localhost');
INSERT INTO settings (name, value) VALUES ('LN_PORT', '10009');
INSERT INTO settings (name, value) VALUES ('LN_TLS_FILE', '/home/xx979xx/boltcard/tls.cert');
INSERT INTO settings (name, value) VALUES ('LN_MACAROON_FILE', '/home/xx979xx/boltcard/admin.macaroon');
INSERT INTO settings (name, value) VALUES ('FEE_LIMIT_SAT', '10');
INSERT INTO settings (name, value) VALUES ('FEE_LIMIT_PERCENT', '0.1');
INSERT INTO settings (name, value) VALUES ('LN_TESTNODE', '');
INSERT INTO settings (name, value) VALUES ('FUNCTION_LNURLW', 'ENABLE');
INSERT INTO settings (name, value) VALUES ('FUNCTION_LNURLP', 'ENABLE');
INSERT INTO settings (name, value) VALUES ('FUNCTION_EMAIL', 'DISABLE');
INSERT INTO settings (name, value) VALUES ('DEFAULT_DESCRIPTION', 'bolt card service');
INSERT INTO settings (name, value) VALUES ('AWS_SES_ID', '');
INSERT INTO settings (name, value) VALUES ('AWS_SES_SECRET', '');
INSERT INTO settings (name, value) VALUES ('AWS_SES_EMAIL_FROM', '');
INSERT INTO settings (name, value) VALUES ('AWS_REGION', 'us-east-1');
INSERT INTO settings (name, value) VALUES ('EMAIL_MAX_TXS', '');
INSERT INTO settings (name, value) VALUES ('FUNCTION_LNDHUB', 'DISABLE');
INSERT INTO settings (name, value) VALUES ('LNDHUB_URL', '');
INSERT INTO settings (name, value) VALUES ('FUNCTION_INTERNAL_API', 'ENABLE');
INSERT INTO settings (name, value) VALUES ('SENDGRID_API_KEY', '');
INSERT INTO settings (name, value) VALUES ('SENDGRID_EMAIL_SENDER', '');
INSERT INTO settings (name, value) VALUES ('LN_INVOICE_EXPIRY_SEC', '3600');
