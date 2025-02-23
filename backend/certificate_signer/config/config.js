const CERTIFICATE_NAMESPACE = process.env.CERTIFICATE_NAMESPACE || "https://divoc.dev/credentials/vaccination/v1";
const CERTIFICATE_NAMESPACE_V2 = process.env.CERTIFICATE_NAMESPACE_V2 || "https://divoc.dev/credentials/vaccination/v2";
const CERTIFICATE_CONTROLLER_ID = process.env.CERTIFICATE_CONTROLLER_ID || 'https://divoc.dev/';
const CERTIFICATE_PUBKEY_ID = process.env.CERTIFICATE_PUBKEY_ID || 'https://example.com/i/india';
const CERTIFICATE_DID = process.env.CERTIFICATE_DID || 'did:india';
const CERTIFICATE_ISSUER = process.env.CERTIFICATE_ISSUER || "https://divoc.dev/";
const CERTIFICATE_BASE_URL = process.env.CERTIFICATE_BASE_URL || "https://divoc.dev/vaccine/";
const CERTIFICATE_FEEDBACK_BASE_URL = process.env.CERTIFICATE_FEEDBACK_BASE_URL || "https://divoc.dev/?";
const CERTIFICATE_INFO_BASE_URL = process.env.CERTIFICATE_INFO_BASE_URL || "https://divoc.dev/?";
const ENABLE_FEEDBACK_URL = process.env.ENABLE_FEEDBACK_URL || true;

const KAFKA_BOOTSTRAP_SERVER = process.env.KAFKA_BOOTSTRAP_SERVERS || 'localhost:9092';
const KAFKA_CONSUMER_SESSION_TIMEOUT = process.env.KAFKA_CONSUMER_SESSION_TIMEOUT || 300000; // in ms
const CERTIFY_TOPIC = process.env.CERTIFY_TOPIC || 'certify';
const REGISTRY_URL = process.env.REGISTRY_URL || 'http://localhost:8081';
const REGISTRY_CERTIFICATE_SCHEMA = process.env.REGISTRY_CERTIFICATE_SCHEMA || 'VaccinationCertificate';
const CERTIFIED_TOPIC = process.env.CERTIFIED_TOPIC || 'certified';
const ERROR_CERTIFICATE_TOPIC = process.env.ERROR_CERTIFICATE_TOPIC || 'error_certificate';
const DUPLICATE_CERTIFICATE_TOPIC = process.env.DUPLICATE_CERTIFICATE_TOPIC || 'duplicate_certificate';
const CERTIFICATE_ACK_TOPIC = process.env.CERTIFICATE_ACK_TOPIC || 'certify_ack';
const ENABLE_CERTIFY_ACKNOWLEDGEMENT = process.env.ENABLE_CERTIFY_ACKNOWLEDGEMENT || true;
const CERTIFICATE_RETRY_COUNT = process.env.CERTIFICATE_RETRY_COUNT || 5;
const REDIS_ENABLED = process.env.REDIS_ENABLED || true;
const REDIS_URL = process.env.REDIS_URL || 'redis://0.0.0.0:6379';
const REDIS_KEY_EXPIRE = process.env.REDIS_KEY_EXPIRE || 2 * 24 * 60 * 60; // in secs
const ENABLE_PROGRAM_ID_CACHING_KEY = process.env.ENABLE_PROGRAM_ID_CACHING_KEY || true;
const FACILITY_COUNTRY_CODE = process.env.FACILITY_COUNTRY_CODE || '';
const COUNTRY_CODE = process.env.COUNTRY_CODE || '';
const IDENTITY_REJECTION_PATTERN = process.env.IDENTITY_REJECTION_PATTERN || "^NA$|^.?$";

module.exports = {
  CERTIFICATE_NAMESPACE,
  CERTIFICATE_NAMESPACE_V2,
  CERTIFICATE_CONTROLLER_ID,
  CERTIFICATE_DID,
  CERTIFICATE_PUBKEY_ID,
  CERTIFICATE_ISSUER,
  CERTIFICATE_BASE_URL,
  CERTIFICATE_FEEDBACK_BASE_URL,
  CERTIFICATE_INFO_BASE_URL,
  KAFKA_BOOTSTRAP_SERVER,
  CERTIFY_TOPIC,
  REGISTRY_URL,
  CERTIFIED_TOPIC,
  ENABLE_CERTIFY_ACKNOWLEDGEMENT,
  ERROR_CERTIFICATE_TOPIC,
  CERTIFICATE_RETRY_COUNT,
  KAFKA_CONSUMER_SESSION_TIMEOUT,
  REDIS_ENABLED,
  REDIS_URL,
  REDIS_KEY_EXPIRE,
  DUPLICATE_CERTIFICATE_TOPIC,
  REGISTRY_CERTIFICATE_SCHEMA,
  CERTIFICATE_ACK_TOPIC,
  ENABLE_FEEDBACK_URL,
  ENABLE_PROGRAM_ID_CACHING_KEY,
  COUNTRY_CODE,
  FACILITY_COUNTRY_CODE,
  IDENTITY_REJECTION_PATTERN
};
